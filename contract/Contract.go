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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"product_id\",\"type\":\"uint256\"}],\"name\":\"approveProduct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"product_id\",\"type\":\"uint256\"}],\"name\":\"buyProduct\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"info\",\"type\":\"string\"}],\"name\":\"createProduct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"logins\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"products\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"product_id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"enumOwner.Status\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"info\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_login\",\"type\":\"string\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"product_id\",\"type\":\"uint256\"}],\"name\":\"sellProduct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"}],\"name\":\"updateWhiteList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"whiteList\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506119b8806100606000396000f3fe6080604052600436106100915760003560e01c80637acc0b20116100595780637acc0b201461018b5780638642269e146101cd578063e941fa78146101e9578063f2c298be14610200578063ff38ec0f1461022957610091565b806327d5549514610096578063372c12b1146100bf578063492deb15146100fc5780634fcf89511461012557806373373d311461014e575b600080fd5b3480156100a257600080fd5b506100bd60048036038101906100b89190611311565b610252565b005b3480156100cb57600080fd5b506100e660048036038101906100e19190611228565b61048c565b6040516100f39190611512565b60405180910390f35b34801561010857600080fd5b50610123600480360381019061011e9190611311565b6104ac565b005b34801561013157600080fd5b5061014c60048036038101906101479190611228565b610658565b005b34801561015a57600080fd5b5061017560048036038101906101709190611251565b610741565b60405161018291906114f7565b60405180910390f35b34801561019757600080fd5b506101b260048036038101906101ad9190611311565b61078a565b6040516101c4969594939291906115ad565b60405180910390f35b6101e760048036038101906101e29190611311565b610913565b005b3480156101f557600080fd5b506101fe610cc0565b005b34801561020c57600080fd5b5061022760048036038101906102229190611251565b610db7565b005b34801561023557600080fd5b50610250600480360381019061024b9190611292565b610ec6565b005b80600080600281111561028e577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b600383815481106102c8577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b906000526020600020906006020160040160009054906101000a900460ff16600281111561031f577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b1461035f576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103569061154d565b60405180910390fd5b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146103ed576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103e49061158d565b60405180910390fd5b600160038481548110610429577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b906000526020600020906006020160040160006101000a81548160ff02191690836002811115610482577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b0217905550505050565b60026020528060005260406000206000915054906101000a900460ff1681565b8060028060028111156104e8577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60038381548110610522577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b906000526020600020906006020160040160009054906101000a900460ff166002811115610579577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b146105b9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105b09061154d565b60405180910390fd5b6001600384815481106105f5577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b906000526020600020906006020160040160006101000a81548160ff0219169083600281111561064e577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b0217905550505050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146106e6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106dd9061158d565b60405180910390fd5b6001600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555050565b6001818051602081018201805184825260208301602085012081835280955050505050506000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6003818154811061079a57600080fd5b90600052602060002090600602016000915090508060000154908060010180546107c390611812565b80601f01602080910402602001604051908101604052809291908181526020018280546107ef90611812565b801561083c5780601f106108115761010080835404028352916020019161083c565b820191906000526020600020905b81548152906001019060200180831161081f57829003601f168201915b5050505050908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060030154908060040160009054906101000a900460ff169080600501805461089090611812565b80601f01602080910402602001604051908101604052809291908181526020018280546108bc90611812565b80156109095780601f106108de57610100808354040283529160200191610909565b820191906000526020600020905b8154815290600101906020018083116108ec57829003601f168201915b5050505050905086565b80600180600281111561094f577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60038381548110610989577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b906000526020600020906006020160040160009054906101000a900460ff1660028111156109e0577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b14610a20576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a179061154d565b60405180910390fd5b60006064600160038681548110610a60577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b906000526020600020906006020160030154610a7c91906116d5565b610a8691906116a4565b905060008160038681548110610ac5577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b906000526020600020906006020160030154610ae1919061172f565b905060038581548110610b1d577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b906000526020600020906006020160020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc829081150290604051600060405180830381858888f19350505050158015610b94573d6000803e3d6000fd5b503360038681548110610bd0577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b906000526020600020906006020160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600260038681548110610c5b577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b906000526020600020906006020160040160006101000a81548160ff02191690836002811115610cb4577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b02179055505050505050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610d4e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d459061158d565b60405180910390fd5b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc479081150290604051600060405180830381858888f19350505050158015610db4573d6000803e3d6000fd5b50565b80600073ffffffffffffffffffffffffffffffffffffffff16600182604051610de091906114e0565b908152602001604051809103902060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610e65576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e5c9061156d565b60405180910390fd5b33600183604051610e7691906114e0565b908152602001604051809103902060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050565b33600260008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615610f54576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f4b9061152d565b60405180910390fd5b60036040518060c0016040528060038054905081526020018681526020013373ffffffffffffffffffffffffffffffffffffffff16815260200185815260200160006002811115610fce577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b81526020018481525090806001815401808255809150506001900390600052602060002090600602016000909190919091506000820151816000015560208201518160010190805190602001906110269291906110f3565b5060408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506060820151816003015560808201518160040160006101000a81548160ff021916908360028111156110c9577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b021790555060a08201518160050190805190602001906110ea9291906110f3565b50505050505050565b8280546110ff90611812565b90600052602060002090601f0160209004810192826111215760008555611168565b82601f1061113a57805160ff1916838001178555611168565b82800160010185558215611168579182015b8281111561116757825182559160200191906001019061114c565b5b5090506111759190611179565b5090565b5b8082111561119257600081600090555060010161117a565b5090565b60006111a96111a48461164d565b61161c565b9050828152602081018484840111156111c157600080fd5b6111cc8482856117d0565b509392505050565b6000813590506111e381611954565b92915050565b600082601f8301126111fa57600080fd5b813561120a848260208601611196565b91505092915050565b6000813590506112228161196b565b92915050565b60006020828403121561123a57600080fd5b6000611248848285016111d4565b91505092915050565b60006020828403121561126357600080fd5b600082013567ffffffffffffffff81111561127d57600080fd5b611289848285016111e9565b91505092915050565b6000806000606084860312156112a757600080fd5b600084013567ffffffffffffffff8111156112c157600080fd5b6112cd868287016111e9565b93505060206112de86828701611213565b925050604084013567ffffffffffffffff8111156112fb57600080fd5b611307868287016111e9565b9150509250925092565b60006020828403121561132357600080fd5b600061133184828501611213565b91505092915050565b61134381611763565b82525050565b61135281611775565b82525050565b611361816117be565b82525050565b60006113728261167d565b61137c8185611688565b935061138c8185602086016117df565b6113958161192f565b840191505092915050565b60006113ab8261167d565b6113b58185611699565b93506113c58185602086016117df565b80840191505092915050565b60006113de601d83611688565b91507f746869732061646472657373206e6f7420696e2077686974656c6973740000006000830152602082019050919050565b600061141e601683611688565b91507f70726f647563652073746174757320696e76616c6964000000000000000000006000830152602082019050919050565b600061145e601083611688565b91507f74686973206c6f67696e206578697374000000000000000000000000000000006000830152602082019050919050565b600061149e600d83611688565b91507f796f75206e6f74206f776e6572000000000000000000000000000000000000006000830152602082019050919050565b6114da816117b4565b82525050565b60006114ec82846113a0565b915081905092915050565b600060208201905061150c600083018461133a565b92915050565b60006020820190506115276000830184611349565b92915050565b60006020820190508181036000830152611546816113d1565b9050919050565b6000602082019050818103600083015261156681611411565b9050919050565b6000602082019050818103600083015261158681611451565b9050919050565b600060208201905081810360008301526115a681611491565b9050919050565b600060c0820190506115c260008301896114d1565b81810360208301526115d48188611367565b90506115e3604083018761133a565b6115f060608301866114d1565b6115fd6080830185611358565b81810360a083015261160f8184611367565b9050979650505050505050565b6000604051905081810181811067ffffffffffffffff8211171561164357611642611900565b5b8060405250919050565b600067ffffffffffffffff82111561166857611667611900565b5b601f19601f8301169050602081019050919050565b600081519050919050565b600082825260208201905092915050565b600081905092915050565b60006116af826117b4565b91506116ba836117b4565b9250826116ca576116c9611873565b5b828204905092915050565b60006116e0826117b4565b91506116eb836117b4565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048311821515161561172457611723611844565b5b828202905092915050565b600061173a826117b4565b9150611745836117b4565b92508282101561175857611757611844565b5b828203905092915050565b600061176e82611794565b9050919050565b60008115159050919050565b600081905061178f82611940565b919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b60006117c982611781565b9050919050565b82818337600083830152505050565b60005b838110156117fd5780820151818401526020810190506117e2565b8381111561180c576000848401525b50505050565b6000600282049050600182168061182a57607f821691505b6020821081141561183e5761183d6118d1565b5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000601f19601f8301169050919050565b60038110611951576119506118a2565b5b50565b61195d81611763565b811461196857600080fd5b50565b611974816117b4565b811461197f57600080fd5b5056fea26469706673582212206d418c3640c5dd768a2aa09b595e38c3c4b8871cdb35edaa44cd537a2b0847ce64736f6c63430008000033",
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

// Register is a paid mutator transaction binding the contract method 0xf2c298be.
//
// Solidity: function register(string _login) returns()
func (_Contract *ContractTransactor) Register(opts *bind.TransactOpts, _login string) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "register", _login)
}

// Register is a paid mutator transaction binding the contract method 0xf2c298be.
//
// Solidity: function register(string _login) returns()
func (_Contract *ContractSession) Register(_login string) (*types.Transaction, error) {
	return _Contract.Contract.Register(&_Contract.TransactOpts, _login)
}

// Register is a paid mutator transaction binding the contract method 0xf2c298be.
//
// Solidity: function register(string _login) returns()
func (_Contract *ContractTransactorSession) Register(_login string) (*types.Transaction, error) {
	return _Contract.Contract.Register(&_Contract.TransactOpts, _login)
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
