package gethcontracttesting_test

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"testing"

	"github.com/amidgo/geth-contract-testing/contract"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/eth/ethconfig"
	"github.com/ethereum/go-ethereum/ethclient/simulated"
	"github.com/ethereum/go-ethereum/node"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const chainID int64 = 12345

var Ether = big.NewInt(1000000000000000000)

// middleware для того чтобы у всех сеток был одинаковый chainID
var withChainID = func(chainID int64) func(nodeConf *node.Config, ethConf *ethconfig.Config) {
	return func(nodeConf *node.Config, ethConf *ethconfig.Config) {
		ethConf.Genesis.Config.ChainID = big.NewInt(chainID)
	}
}

func Test_UpdateWhiteList(t *testing.T) {
	ctx := context.Background()
	// создаём случайного пользователя с 10 ether
	user, err := NewAccount(big.NewInt(0).Mul(Ether, big.NewInt(10)))
	require.NoError(t, err, "failed create new account")

	// создаём сетку
	network, err := NewNetwork(ctx, user)
	require.NoError(t, err, "failed create new network")

	// тут мы должны проверить что по умолчанию случайный пользователь не включен в whitelist
	inWhiteList, err := network.Contract.WhiteList(network.CallOpts(ctx, &user), user.Address())
	require.NoError(t, err, "failed call contract white list")

	// ассертим что false
	assert.False(t, inWhiteList, "user in white list")

	// тут мы проверяем что обновить контракт может только owner
	_, err = network.Contract.UpdateWhiteList(user.TransactOpts, user.Address())
	require.Equal(t, "execution reverted: you not owner", err.Error())
	log.Printf("%T", err)

	network.Commit()

	// проверяем что ничего не изменилось
	inWhiteList, err = network.Contract.WhiteList(network.CallOpts(ctx, &user), user.Address())
	require.NoError(t, err, "failed call contract white list")

	assert.False(t, inWhiteList, "user in white list")

	// эта штука нужна чтобы мы могли совершить новую транзакцию, обязатель после деплоя или иных транзакций

	// добавляем пользователя в whitelist от владельца
	_, err = network.Contract.UpdateWhiteList(network.Owner.TransactOpts, user.Address())
	require.NoError(t, err, "failed call contract white list")

	network.Commit()

	// проверяем что всё чикибамбони
	inWhiteList, err = network.Contract.WhiteList(network.CallOpts(ctx, &network.Owner), user.Address())
	require.NoError(t, err, "failed call contract white list")

	assert.True(t, inWhiteList)
}

func Test_CreateUser(t *testing.T) {
	ctx := context.Background()
	// создаём случайного пользователя с 10 ether
	user, err := NewAccount(big.NewInt(0).Mul(Ether, big.NewInt(10)))
	require.NoError(t, err, "failed create new account")

	// создаём сетку
	network, err := NewNetwork(ctx, user)
	require.NoError(t, err, "failed create new network")

	const userLogin = "user"

	// убеждаемся что пользователя пока что не существует
	address, err := network.Contract.Logins(network.CallOpts(ctx, &user), userLogin)
	require.NoError(t, err, "failed get user by login")

	assert.Equal(t, common.Address{}, address, "wrong address")

	// регистрируем пользователя под логином
	_, err = network.Contract.CreateUser(user.TransactOpts, userLogin)
	require.NoError(t, err, "failed create user")

	network.Commit()

	// убеждаемся что пользователь успешно зарегистрирован
	address, err = network.Contract.Logins(network.CallOpts(ctx, &user), userLogin)
	require.NoError(t, err, "failed get user by login")

	assert.Equal(t, user.Address(), address, "wrong address")

	// пытаемся зарегистрировать пользователя с тем же логином
	_, err = network.Contract.CreateUser(user.TransactOpts, userLogin)
	require.Equal(t, err.Error(), "execution reverted: this login exist")

	network.Commit()
}

func Test_SellProduct(t *testing.T) {
	ctx := context.Background()
	// создаём случайного пользователя с 10 ether
	user, err := NewAccount(big.NewInt(0).Mul(Ether, big.NewInt(10)))
	require.NoError(t, err, "failed create new account")

	// создаём сетку
	network, err := NewNetwork(ctx, user)
	require.NoError(t, err, "failed create new network")

	_, err = network.Contract.UpdateWhiteList(network.Owner.TransactOpts, user.Address())
	require.NoError(t, err, "failed update white list")

	network.Commit()

	_, err = network.Contract.CreateProduct(user.TransactOpts, "coca-cola", Ether, "праздник к нам приходит")
	require.NoError(t, err, "failed create product")

	network.Commit()

	_, err = network.Contract.ApproveProduct(network.Owner.TransactOpts, big.NewInt(0))
	require.NoError(t, err, "failed approve product")

	network.Commit()

	_, err = network.Contract.SellProduct(network.Owner.TransactOpts, big.NewInt(0))
	require.Equal(t, err.Error(), "execution reverted: you not a product owner")

	network.Commit()
}

type Account struct {
	Balance      *big.Int
	TransactOpts *bind.TransactOpts
}

func (a *Account) Address() common.Address {
	if a.TransactOpts == nil {
		return common.Address{}
	}

	return a.TransactOpts.From
}

func NewAccount(balance *big.Int) (Account, error) {
	key, err := crypto.GenerateKey()
	if err != nil {
		return Account{}, fmt.Errorf("failed generate key")
	}

	tOpts, err := bind.NewKeyedTransactorWithChainID(key, big.NewInt(chainID))
	if err != nil {
		return Account{}, fmt.Errorf("failed generate transactor")
	}

	return Account{
		TransactOpts: tOpts,
		Balance:      balance,
	}, nil
}

type Network struct {
	Owner    Account
	Contract *contract.Contract
	Backend  *simulated.Backend
}

func (n *Network) Commit() {
	n.Backend.Commit()
}

func (n *Network) CallOpts(ctx context.Context, acc *Account) *bind.CallOpts {
	blockNumber, err := n.Backend.Client().BlockNumber(ctx)
	if err != nil {
		log.Fatalf(err.Error())
	}

	return &bind.CallOpts{
		Pending:     true,
		From:        acc.Address(),
		BlockNumber: big.NewInt(int64(blockNumber)),
		Context:     ctx,
	}
}

func NewNetwork(ctx context.Context, accounts ...Account) (*Network, error) {
	ownerBalance := big.NewInt(0).Mul(big.NewInt(1000000000000000000), big.NewInt(1000000000000000000))

	owner, err := NewAccount(ownerBalance)
	if err != nil {
		return nil, fmt.Errorf("failed create owner account")
	}

	alloc := make(types.GenesisAlloc, len(accounts)+1)
	alloc[owner.Address()] = types.Account{Balance: ownerBalance}

	for _, account := range accounts {
		alloc[account.TransactOpts.From] = types.Account{
			Balance: account.Balance,
		}
	}

	backend := simulated.NewBackend(alloc, withChainID(chainID))

	_, _, cntr, err := contract.DeployContract(owner.TransactOpts, backend.Client())
	if err != nil {
		return nil, fmt.Errorf("failed deploy contract")
	}

	backend.Commit()

	return &Network{
		Owner:    owner,
		Backend:  backend,
		Contract: cntr,
	}, nil
}
