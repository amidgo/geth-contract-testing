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

var withChainID = func(chainID int64) func(nodeConf *node.Config, ethConf *ethconfig.Config) {
	return func(nodeConf *node.Config, ethConf *ethconfig.Config) {
		ethConf.Genesis.Config.ChainID = big.NewInt(chainID)
	}
}

func Test_UpdateWhiteList(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	user, err := NewAccount(big.NewInt(0).Mul(Ether, big.NewInt(10)))
	require.NoError(t, err, "failed create new account")

	network, err := NewNetwork(ctx, user)
	require.NoError(t, err, "failed create new network")

	inWhiteList, err := network.Contract.WhiteList(network.CallOpts(ctx, &user), user.Address())
	require.NoError(t, err, "failed call contract white list")

	assert.False(t, inWhiteList, "user in white list")

	_, err = network.Contract.UpdateWhiteList(user.TransactOpts, user.Address())
	require.NoError(t, err)

	inWhiteList, err = network.Contract.WhiteList(network.CallOpts(ctx, &user), user.Address())
	require.NoError(t, err, "failed call contract white list")

	assert.False(t, inWhiteList, "user in white list")

	_, err = network.Contract.UpdateWhiteList(network.Owner.TransactOpts, user.Address())
	require.NoError(t, err)

	inWhiteList, err = network.Contract.WhiteList(network.CallOpts(ctx, &network.Owner), user.Address())
	require.NoError(t, err, "failed call contract white list")

	assert.True(t, inWhiteList, "user not in white list")
}

func Test_CreateUser(t *testing.T) {

}

func Test_SellProduct(t *testing.T) {

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
		alloc[account.TransactOpts.From] = types.Account{Balance: account.Balance}
	}

	backend := simulated.NewBackend(alloc, withChainID(chainID))

	_, _, cntr, err := contract.DeployContract(owner.TransactOpts, backend.Client())
	if err != nil {
		return nil, fmt.Errorf("failed deploy contract")
	}

	return &Network{
		Owner:    owner,
		Backend:  backend,
		Contract: cntr,
	}, nil
}
