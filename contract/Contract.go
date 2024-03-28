// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ContractMetaData contains all meta data concerning the Contract contract.
var ContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"product_id\",\"type\":\"uint256\"}],\"name\":\"approveProduct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"product_id\",\"type\":\"uint256\"}],\"name\":\"buyProduct\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"info\",\"type\":\"string\"}],\"name\":\"createProduct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_login\",\"type\":\"string\"}],\"name\":\"createUser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"logins\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"products\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"product_id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"enumOwner.Status\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"info\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"product_id\",\"type\":\"uint256\"}],\"name\":\"sellProduct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"}],\"name\":\"updateWhiteList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"whiteList\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550611af2806100606000396000f3fe6080604052600436106100915760003560e01c806373373d311161005957806373373d31146101775780637acc0b20146101b45780638642269e146101f6578063e941fa7814610212578063ff38ec0f1461022957610091565b806327d5549514610096578063372c12b1146100bf578063492deb15146100fc5780634fcf895114610125578063507ffba51461014e575b600080fd5b3480156100a257600080fd5b506100bd60048036038101906100b891906113eb565b610252565b005b3480156100cb57600080fd5b506100e660048036038101906100e19190611302565b61048c565b6040516100f3919061162c565b60405180910390f35b34801561010857600080fd5b50610123600480360381019061011e91906113eb565b6104ac565b005b34801561013157600080fd5b5061014c60048036038101906101479190611302565b610733565b005b34801561015a57600080fd5b506101756004803603810190610170919061132b565b61081c565b005b34801561018357600080fd5b5061019e6004803603810190610199919061132b565b61092b565b6040516101ab9190611611565b60405180910390f35b3480156101c057600080fd5b506101db60048036038101906101d691906113eb565b610974565b6040516101ed969594939291906116e7565b60405180910390f35b610210600480360381019061020b91906113eb565b610afd565b005b34801561021e57600080fd5b50610227610eaa565b005b34801561023557600080fd5b50610250600480360381019061024b919061136c565b610fa1565b005b80600080600281111561028e577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b600383815481106102c8577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b906000526020600020906006020160040160009054906101000a900460ff16600281111561031f577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b1461035f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161035690611687565b60405180910390fd5b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146103ed576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103e4906116c7565b60405180910390fd5b600160038481548110610429577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b906000526020600020906006020160040160006101000a81548160ff02191690836002811115610482577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b0217905550505050565b60026020528060005260406000206000915054906101000a900460ff1681565b803373ffffffffffffffffffffffffffffffffffffffff16600382815481106104fe577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b906000526020600020906006020160020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610586576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161057d90611647565b60405180910390fd5b8160028060028111156105c2577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b600383815481106105fc577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b906000526020600020906006020160040160009054906101000a900460ff166002811115610653577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b14610693576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161068a90611687565b60405180910390fd5b6001600385815481106106cf577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b906000526020600020906006020160040160006101000a81548160ff02191690836002811115610728577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b021790555050505050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146107c1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107b8906116c7565b60405180910390fd5b6001600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555050565b80600073ffffffffffffffffffffffffffffffffffffffff1660018260405161084591906115fa565b908152602001604051809103902060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146108ca576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108c1906116a7565b60405180910390fd5b336001836040516108db91906115fa565b908152602001604051809103902060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050565b6001818051602081018201805184825260208301602085012081835280955050505050506000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6003818154811061098457600080fd5b90600052602060002090600602016000915090508060000154908060010180546109ad9061194c565b80601f01602080910402602001604051908101604052809291908181526020018280546109d99061194c565b8015610a265780601f106109fb57610100808354040283529160200191610a26565b820191906000526020600020905b815481529060010190602001808311610a0957829003601f168201915b5050505050908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060030154908060040160009054906101000a900460ff1690806005018054610a7a9061194c565b80601f0160208091040260200160405190810160405280929190818152602001828054610aa69061194c565b8015610af35780601f10610ac857610100808354040283529160200191610af3565b820191906000526020600020905b815481529060010190602001808311610ad657829003601f168201915b5050505050905086565b806001806002811115610b39577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60038381548110610b73577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b906000526020600020906006020160040160009054906101000a900460ff166002811115610bca577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b14610c0a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c0190611687565b60405180910390fd5b60006064600160038681548110610c4a577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b906000526020600020906006020160030154610c66919061180f565b610c7091906117de565b905060008160038681548110610caf577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b906000526020600020906006020160030154610ccb9190611869565b905060038581548110610d07577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b906000526020600020906006020160020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc829081150290604051600060405180830381858888f19350505050158015610d7e573d6000803e3d6000fd5b503360038681548110610dba577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b906000526020600020906006020160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600260038681548110610e45577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b906000526020600020906006020160040160006101000a81548160ff02191690836002811115610e9e577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b02179055505050505050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610f38576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f2f906116c7565b60405180910390fd5b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc479081150290604051600060405180830381858888f19350505050158015610f9e573d6000803e3d6000fd5b50565b33600260008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1661102e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161102590611667565b60405180910390fd5b60036040518060c0016040528060038054905081526020018681526020013373ffffffffffffffffffffffffffffffffffffffff168152602001858152602001600060028111156110a8577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b81526020018481525090806001815401808255809150506001900390600052602060002090600602016000909190919091506000820151816000015560208201518160010190805190602001906111009291906111cd565b5060408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506060820151816003015560808201518160040160006101000a81548160ff021916908360028111156111a3577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b021790555060a08201518160050190805190602001906111c49291906111cd565b50505050505050565b8280546111d99061194c565b90600052602060002090601f0160209004810192826111fb5760008555611242565b82601f1061121457805160ff1916838001178555611242565b82800160010185558215611242579182015b82811115611241578251825591602001919060010190611226565b5b50905061124f9190611253565b5090565b5b8082111561126c576000816000905550600101611254565b5090565b600061128361127e84611787565b611756565b90508281526020810184848401111561129b57600080fd5b6112a684828561190a565b509392505050565b6000813590506112bd81611a8e565b92915050565b600082601f8301126112d457600080fd5b81356112e4848260208601611270565b91505092915050565b6000813590506112fc81611aa5565b92915050565b60006020828403121561131457600080fd5b6000611322848285016112ae565b91505092915050565b60006020828403121561133d57600080fd5b600082013567ffffffffffffffff81111561135757600080fd5b611363848285016112c3565b91505092915050565b60008060006060848603121561138157600080fd5b600084013567ffffffffffffffff81111561139b57600080fd5b6113a7868287016112c3565b93505060206113b8868287016112ed565b925050604084013567ffffffffffffffff8111156113d557600080fd5b6113e1868287016112c3565b9150509250925092565b6000602082840312156113fd57600080fd5b600061140b848285016112ed565b91505092915050565b61141d8161189d565b82525050565b61142c816118af565b82525050565b61143b816118f8565b82525050565b600061144c826117b7565b61145681856117c2565b9350611466818560208601611919565b61146f81611a69565b840191505092915050565b6000611485826117b7565b61148f81856117d3565b935061149f818560208601611919565b80840191505092915050565b60006114b86017836117c2565b91507f796f75206e6f7420612070726f64756374206f776e65720000000000000000006000830152602082019050919050565b60006114f8601d836117c2565b91507f746869732061646472657373206e6f7420696e2077686974656c6973740000006000830152602082019050919050565b60006115386016836117c2565b91507f70726f647563652073746174757320696e76616c6964000000000000000000006000830152602082019050919050565b60006115786010836117c2565b91507f74686973206c6f67696e206578697374000000000000000000000000000000006000830152602082019050919050565b60006115b8600d836117c2565b91507f796f75206e6f74206f776e6572000000000000000000000000000000000000006000830152602082019050919050565b6115f4816118ee565b82525050565b6000611606828461147a565b915081905092915050565b60006020820190506116266000830184611414565b92915050565b60006020820190506116416000830184611423565b92915050565b60006020820190508181036000830152611660816114ab565b9050919050565b60006020820190508181036000830152611680816114eb565b9050919050565b600060208201905081810360008301526116a08161152b565b9050919050565b600060208201905081810360008301526116c08161156b565b9050919050565b600060208201905081810360008301526116e0816115ab565b9050919050565b600060c0820190506116fc60008301896115eb565b818103602083015261170e8188611441565b905061171d6040830187611414565b61172a60608301866115eb565b6117376080830185611432565b81810360a08301526117498184611441565b9050979650505050505050565b6000604051905081810181811067ffffffffffffffff8211171561177d5761177c611a3a565b5b8060405250919050565b600067ffffffffffffffff8211156117a2576117a1611a3a565b5b601f19601f8301169050602081019050919050565b600081519050919050565b600082825260208201905092915050565b600081905092915050565b60006117e9826118ee565b91506117f4836118ee565b925082611804576118036119ad565b5b828204905092915050565b600061181a826118ee565b9150611825836118ee565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048311821515161561185e5761185d61197e565b5b828202905092915050565b6000611874826118ee565b915061187f836118ee565b9250828210156118925761189161197e565b5b828203905092915050565b60006118a8826118ce565b9050919050565b60008115159050919050565b60008190506118c982611a7a565b919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b6000611903826118bb565b9050919050565b82818337600083830152505050565b60005b8381101561193757808201518184015260208101905061191c565b83811115611946576000848401525b50505050565b6000600282049050600182168061196457607f821691505b6020821081141561197857611977611a0b565b5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000601f19601f8301169050919050565b60038110611a8b57611a8a6119dc565b5b50565b611a978161189d565b8114611aa257600080fd5b50565b611aae816118ee565b8114611ab957600080fd5b5056fea2646970667358221220c1478549efef75ce8aa45972f97c087b64d7d972a85001df140a438479ac653264736f6c63430008000033",
}

// ContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractMetaData.ABI instead.
var ContractABI = ContractMetaData.ABI

// ContractBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractMetaData.Bin instead.
var ContractBin = ContractMetaData.Bin

// DeployContract deploys a new Ethereum contract, binding an instance of Contract to it.
func DeployContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Contract, error) {
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// Contract is an auto generated Go binding around an Ethereum contract.
type Contract struct {
	ContractCaller     // Read-only binding to the contract
	ContractTransactor // Write-only binding to the contract
	ContractFilterer   // Log filterer for contract events
}

// ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSession struct {
	Contract     *Contract         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCallerSession struct {
	Contract *ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTransactorSession struct {
	Contract     *ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRaw struct {
	Contract *Contract // Generic contract binding to access the raw methods on
}

// ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCallerRaw struct {
	Contract *ContractCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTransactorRaw struct {
	Contract *ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContract creates a new instance of Contract, bound to a specific deployed contract.
func NewContract(address common.Address, backend bind.ContractBackend) (*Contract, error) {
	contract, err := bindContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// NewContractCaller creates a new read-only instance of Contract, bound to a specific deployed contract.
func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
	contract, err := bindContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{contract: contract}, nil
}

// NewContractTransactor creates a new write-only instance of Contract, bound to a specific deployed contract.
func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
	contract, err := bindContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTransactor{contract: contract}, nil
}

// NewContractFilterer creates a new log filterer instance of Contract, bound to a specific deployed contract.
func NewContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFilterer, error) {
	contract, err := bindContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractFilterer{contract: contract}, nil
}

// bindContract binds a generic wrapper to an already deployed contract.
func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transact(opts, method, params...)
}

// Logins is a free data retrieval call binding the contract method 0x73373d31.
//
// Solidity: function logins(string ) view returns(address)
func (_Contract *ContractCaller) Logins(opts *bind.CallOpts, arg0 string) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "logins", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Logins is a free data retrieval call binding the contract method 0x73373d31.
//
// Solidity: function logins(string ) view returns(address)
func (_Contract *ContractSession) Logins(arg0 string) (common.Address, error) {
	return _Contract.Contract.Logins(&_Contract.CallOpts, arg0)
}

// Logins is a free data retrieval call binding the contract method 0x73373d31.
//
// Solidity: function logins(string ) view returns(address)
func (_Contract *ContractCallerSession) Logins(arg0 string) (common.Address, error) {
	return _Contract.Contract.Logins(&_Contract.CallOpts, arg0)
}

// Products is a free data retrieval call binding the contract method 0x7acc0b20.
//
// Solidity: function products(uint256 ) view returns(uint256 product_id, string name, address owner, uint256 value, uint8 status, string info)
func (_Contract *ContractCaller) Products(opts *bind.CallOpts, arg0 *big.Int) (struct {
	ProductId *big.Int
	Name      string
	Owner     common.Address
	Value     *big.Int
	Status    uint8
	Info      string
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "products", arg0)

	outstruct := new(struct {
		ProductId *big.Int
		Name      string
		Owner     common.Address
		Value     *big.Int
		Status    uint8
		Info      string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ProductId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Owner = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Value = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Status = *abi.ConvertType(out[4], new(uint8)).(*uint8)
	outstruct.Info = *abi.ConvertType(out[5], new(string)).(*string)

	return *outstruct, err

}

// Products is a free data retrieval call binding the contract method 0x7acc0b20.
//
// Solidity: function products(uint256 ) view returns(uint256 product_id, string name, address owner, uint256 value, uint8 status, string info)
func (_Contract *ContractSession) Products(arg0 *big.Int) (struct {
	ProductId *big.Int
	Name      string
	Owner     common.Address
	Value     *big.Int
	Status    uint8
	Info      string
}, error) {
	return _Contract.Contract.Products(&_Contract.CallOpts, arg0)
}

// Products is a free data retrieval call binding the contract method 0x7acc0b20.
//
// Solidity: function products(uint256 ) view returns(uint256 product_id, string name, address owner, uint256 value, uint8 status, string info)
func (_Contract *ContractCallerSession) Products(arg0 *big.Int) (struct {
	ProductId *big.Int
	Name      string
	Owner     common.Address
	Value     *big.Int
	Status    uint8
	Info      string
}, error) {
	return _Contract.Contract.Products(&_Contract.CallOpts, arg0)
}

// WhiteList is a free data retrieval call binding the contract method 0x372c12b1.
//
// Solidity: function whiteList(address ) view returns(bool)
func (_Contract *ContractCaller) WhiteList(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "whiteList", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WhiteList is a free data retrieval call binding the contract method 0x372c12b1.
//
// Solidity: function whiteList(address ) view returns(bool)
func (_Contract *ContractSession) WhiteList(arg0 common.Address) (bool, error) {
	return _Contract.Contract.WhiteList(&_Contract.CallOpts, arg0)
}

// WhiteList is a free data retrieval call binding the contract method 0x372c12b1.
//
// Solidity: function whiteList(address ) view returns(bool)
func (_Contract *ContractCallerSession) WhiteList(arg0 common.Address) (bool, error) {
	return _Contract.Contract.WhiteList(&_Contract.CallOpts, arg0)
}

// ApproveProduct is a paid mutator transaction binding the contract method 0x27d55495.
//
// Solidity: function approveProduct(uint256 product_id) returns()
func (_Contract *ContractTransactor) ApproveProduct(opts *bind.TransactOpts, product_id *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "approveProduct", product_id)
}

// ApproveProduct is a paid mutator transaction binding the contract method 0x27d55495.
//
// Solidity: function approveProduct(uint256 product_id) returns()
func (_Contract *ContractSession) ApproveProduct(product_id *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.ApproveProduct(&_Contract.TransactOpts, product_id)
}

// ApproveProduct is a paid mutator transaction binding the contract method 0x27d55495.
//
// Solidity: function approveProduct(uint256 product_id) returns()
func (_Contract *ContractTransactorSession) ApproveProduct(product_id *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.ApproveProduct(&_Contract.TransactOpts, product_id)
}

// BuyProduct is a paid mutator transaction binding the contract method 0x8642269e.
//
// Solidity: function buyProduct(uint256 product_id) payable returns()
func (_Contract *ContractTransactor) BuyProduct(opts *bind.TransactOpts, product_id *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "buyProduct", product_id)
}

// BuyProduct is a paid mutator transaction binding the contract method 0x8642269e.
//
// Solidity: function buyProduct(uint256 product_id) payable returns()
func (_Contract *ContractSession) BuyProduct(product_id *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.BuyProduct(&_Contract.TransactOpts, product_id)
}

// BuyProduct is a paid mutator transaction binding the contract method 0x8642269e.
//
// Solidity: function buyProduct(uint256 product_id) payable returns()
func (_Contract *ContractTransactorSession) BuyProduct(product_id *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.BuyProduct(&_Contract.TransactOpts, product_id)
}

// CreateProduct is a paid mutator transaction binding the contract method 0xff38ec0f.
//
// Solidity: function createProduct(string name, uint256 value, string info) returns()
func (_Contract *ContractTransactor) CreateProduct(opts *bind.TransactOpts, name string, value *big.Int, info string) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "createProduct", name, value, info)
}

// CreateProduct is a paid mutator transaction binding the contract method 0xff38ec0f.
//
// Solidity: function createProduct(string name, uint256 value, string info) returns()
func (_Contract *ContractSession) CreateProduct(name string, value *big.Int, info string) (*types.Transaction, error) {
	return _Contract.Contract.CreateProduct(&_Contract.TransactOpts, name, value, info)
}

// CreateProduct is a paid mutator transaction binding the contract method 0xff38ec0f.
//
// Solidity: function createProduct(string name, uint256 value, string info) returns()
func (_Contract *ContractTransactorSession) CreateProduct(name string, value *big.Int, info string) (*types.Transaction, error) {
	return _Contract.Contract.CreateProduct(&_Contract.TransactOpts, name, value, info)
}

// CreateUser is a paid mutator transaction binding the contract method 0x507ffba5.
//
// Solidity: function createUser(string _login) returns()
func (_Contract *ContractTransactor) CreateUser(opts *bind.TransactOpts, _login string) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "createUser", _login)
}

// CreateUser is a paid mutator transaction binding the contract method 0x507ffba5.
//
// Solidity: function createUser(string _login) returns()
func (_Contract *ContractSession) CreateUser(_login string) (*types.Transaction, error) {
	return _Contract.Contract.CreateUser(&_Contract.TransactOpts, _login)
}

// CreateUser is a paid mutator transaction binding the contract method 0x507ffba5.
//
// Solidity: function createUser(string _login) returns()
func (_Contract *ContractTransactorSession) CreateUser(_login string) (*types.Transaction, error) {
	return _Contract.Contract.CreateUser(&_Contract.TransactOpts, _login)
}

// SellProduct is a paid mutator transaction binding the contract method 0x492deb15.
//
// Solidity: function sellProduct(uint256 product_id) returns()
func (_Contract *ContractTransactor) SellProduct(opts *bind.TransactOpts, product_id *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "sellProduct", product_id)
}

// SellProduct is a paid mutator transaction binding the contract method 0x492deb15.
//
// Solidity: function sellProduct(uint256 product_id) returns()
func (_Contract *ContractSession) SellProduct(product_id *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SellProduct(&_Contract.TransactOpts, product_id)
}

// SellProduct is a paid mutator transaction binding the contract method 0x492deb15.
//
// Solidity: function sellProduct(uint256 product_id) returns()
func (_Contract *ContractTransactorSession) SellProduct(product_id *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SellProduct(&_Contract.TransactOpts, product_id)
}

// UpdateWhiteList is a paid mutator transaction binding the contract method 0x4fcf8951.
//
// Solidity: function updateWhiteList(address wallet) returns()
func (_Contract *ContractTransactor) UpdateWhiteList(opts *bind.TransactOpts, wallet common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "updateWhiteList", wallet)
}

// UpdateWhiteList is a paid mutator transaction binding the contract method 0x4fcf8951.
//
// Solidity: function updateWhiteList(address wallet) returns()
func (_Contract *ContractSession) UpdateWhiteList(wallet common.Address) (*types.Transaction, error) {
	return _Contract.Contract.UpdateWhiteList(&_Contract.TransactOpts, wallet)
}

// UpdateWhiteList is a paid mutator transaction binding the contract method 0x4fcf8951.
//
// Solidity: function updateWhiteList(address wallet) returns()
func (_Contract *ContractTransactorSession) UpdateWhiteList(wallet common.Address) (*types.Transaction, error) {
	return _Contract.Contract.UpdateWhiteList(&_Contract.TransactOpts, wallet)
}

// WithdrawFee is a paid mutator transaction binding the contract method 0xe941fa78.
//
// Solidity: function withdrawFee() returns()
func (_Contract *ContractTransactor) WithdrawFee(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "withdrawFee")
}

// WithdrawFee is a paid mutator transaction binding the contract method 0xe941fa78.
//
// Solidity: function withdrawFee() returns()
func (_Contract *ContractSession) WithdrawFee() (*types.Transaction, error) {
	return _Contract.Contract.WithdrawFee(&_Contract.TransactOpts)
}

// WithdrawFee is a paid mutator transaction binding the contract method 0xe941fa78.
//
// Solidity: function withdrawFee() returns()
func (_Contract *ContractTransactorSession) WithdrawFee() (*types.Transaction, error) {
	return _Contract.Contract.WithdrawFee(&_Contract.TransactOpts)
}
