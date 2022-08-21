// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package deploy

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

// CallerMetaData contains all meta data concerning the Caller contract.
var CallerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"someAction\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"someUnsafeAction\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"storeAction\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506106d6806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c8063748fa59914610046578063e8a6223514610076578063fba9cea9146100a6575b600080fd5b610060600480360381019061005b91906103f2565b6100d6565b60405161006d919061043a565b60405180910390f35b610090600480360381019061008b91906103f2565b61021b565b60405161009d919061046e565b60405180910390f35b6100c060048036038101906100bb91906103f2565b610306565b6040516100cd919061046e565b60405180910390f35b6000808273ffffffffffffffffffffffffffffffffffffffff16606460405160240161010291906104db565b6040516020818303038152906040527f7221a26a000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff838183161783525050505060405161018c9190610567565b6000604051808303816000865af19150503d80600081146101c9576040519150601f19603f3d011682016040523d82523d6000602084013e6101ce565b606091505b5050905080610212576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610209906105db565b60405180910390fd5b80915050919050565b6000808290508073ffffffffffffffffffffffffffffffffffffffff16637221a26a60646040518263ffffffff1660e01b815260040161025b919061062c565b600060405180830381600087803b15801561027557600080fd5b505af1158015610289573d6000803e3d6000fd5b505050508073ffffffffffffffffffffffffffffffffffffffff166319eb4a906040518163ffffffff1660e01b81526004016020604051808303816000875af11580156102da573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102fe9190610673565b915050919050565b6000808290508073ffffffffffffffffffffffffffffffffffffffff16630ff4c91660646040518263ffffffff1660e01b8152600401610346919061062c565b602060405180830381865afa158015610363573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103879190610673565b915050919050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006103bf82610394565b9050919050565b6103cf816103b4565b81146103da57600080fd5b50565b6000813590506103ec816103c6565b92915050565b6000602082840312156104085761040761038f565b5b6000610416848285016103dd565b91505092915050565b60008115159050919050565b6104348161041f565b82525050565b600060208201905061044f600083018461042b565b92915050565b6000819050919050565b61046881610455565b82525050565b6000602082019050610483600083018461045f565b92915050565b6000819050919050565b600060ff82169050919050565b6000819050919050565b60006104c56104c06104bb84610489565b6104a0565b610493565b9050919050565b6104d5816104aa565b82525050565b60006020820190506104f060008301846104cc565b92915050565b600081519050919050565b600081905092915050565b60005b8381101561052a57808201518184015260208101905061050f565b60008484015250505050565b6000610541826104f6565b61054b8185610501565b935061055b81856020860161050c565b80840191505092915050565b60006105738284610536565b915081905092915050565b600082825260208201905092915050565b7f73746f726556616c7565206661696c65642e0000000000000000000000000000600082015250565b60006105c560128361057e565b91506105d08261058f565b602082019050919050565b600060208201905081810360008301526105f4816105b8565b9050919050565b600061061661061161060c84610489565b6104a0565b610455565b9050919050565b610626816105fb565b82525050565b6000602082019050610641600083018461061d565b92915050565b61065081610455565b811461065b57600080fd5b50565b60008151905061066d81610647565b92915050565b6000602082840312156106895761068861038f565b5b60006106978482850161065e565b9150509291505056fea2646970667358221220e0a7a9bf0f2cfd1547b7f484d20fc87951ab6b57618e634b7377a32a0b6b75e564736f6c63430008100033",
}

// CallerABI is the input ABI used to generate the binding from.
// Deprecated: Use CallerMetaData.ABI instead.
var CallerABI = CallerMetaData.ABI

// CallerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CallerMetaData.Bin instead.
var CallerBin = CallerMetaData.Bin

// DeployCaller deploys a new Ethereum contract, binding an instance of Caller to it.
func DeployCaller(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Caller, error) {
	parsed, err := CallerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CallerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Caller{CallerCaller: CallerCaller{contract: contract}, CallerTransactor: CallerTransactor{contract: contract}, CallerFilterer: CallerFilterer{contract: contract}}, nil
}

// Caller is an auto generated Go binding around an Ethereum contract.
type Caller struct {
	CallerCaller     // Read-only binding to the contract
	CallerTransactor // Write-only binding to the contract
	CallerFilterer   // Log filterer for contract events
}

// CallerCaller is an auto generated read-only Go binding around an Ethereum contract.
type CallerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CallerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CallerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CallerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CallerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CallerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CallerSession struct {
	Contract     *Caller           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CallerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CallerCallerSession struct {
	Contract *CallerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// CallerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CallerTransactorSession struct {
	Contract     *CallerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CallerRaw is an auto generated low-level Go binding around an Ethereum contract.
type CallerRaw struct {
	Contract *Caller // Generic contract binding to access the raw methods on
}

// CallerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CallerCallerRaw struct {
	Contract *CallerCaller // Generic read-only contract binding to access the raw methods on
}

// CallerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CallerTransactorRaw struct {
	Contract *CallerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCaller creates a new instance of Caller, bound to a specific deployed contract.
func NewCaller(address common.Address, backend bind.ContractBackend) (*Caller, error) {
	contract, err := bindCaller(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Caller{CallerCaller: CallerCaller{contract: contract}, CallerTransactor: CallerTransactor{contract: contract}, CallerFilterer: CallerFilterer{contract: contract}}, nil
}

// NewCallerCaller creates a new read-only instance of Caller, bound to a specific deployed contract.
func NewCallerCaller(address common.Address, caller bind.ContractCaller) (*CallerCaller, error) {
	contract, err := bindCaller(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CallerCaller{contract: contract}, nil
}

// NewCallerTransactor creates a new write-only instance of Caller, bound to a specific deployed contract.
func NewCallerTransactor(address common.Address, transactor bind.ContractTransactor) (*CallerTransactor, error) {
	contract, err := bindCaller(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CallerTransactor{contract: contract}, nil
}

// NewCallerFilterer creates a new log filterer instance of Caller, bound to a specific deployed contract.
func NewCallerFilterer(address common.Address, filterer bind.ContractFilterer) (*CallerFilterer, error) {
	contract, err := bindCaller(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CallerFilterer{contract: contract}, nil
}

// bindCaller binds a generic wrapper to an already deployed contract.
func bindCaller(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CallerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Caller *CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Caller.Contract.CallerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Caller *CallerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Caller.Contract.CallerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Caller *CallerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Caller.Contract.CallerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Caller *CallerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Caller.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Caller *CallerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Caller.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Caller *CallerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Caller.Contract.contract.Transact(opts, method, params...)
}

// SomeAction is a free data retrieval call binding the contract method 0xfba9cea9.
//
// Solidity: function someAction(address addr) pure returns(uint256)
func (_Caller *CallerCaller) SomeAction(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Caller.contract.Call(opts, &out, "someAction", addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SomeAction is a free data retrieval call binding the contract method 0xfba9cea9.
//
// Solidity: function someAction(address addr) pure returns(uint256)
func (_Caller *CallerSession) SomeAction(addr common.Address) (*big.Int, error) {
	return _Caller.Contract.SomeAction(&_Caller.CallOpts, addr)
}

// SomeAction is a free data retrieval call binding the contract method 0xfba9cea9.
//
// Solidity: function someAction(address addr) pure returns(uint256)
func (_Caller *CallerCallerSession) SomeAction(addr common.Address) (*big.Int, error) {
	return _Caller.Contract.SomeAction(&_Caller.CallOpts, addr)
}

// SomeUnsafeAction is a paid mutator transaction binding the contract method 0x748fa599.
//
// Solidity: function someUnsafeAction(address addr) returns(bool)
func (_Caller *CallerTransactor) SomeUnsafeAction(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Caller.contract.Transact(opts, "someUnsafeAction", addr)
}

// SomeUnsafeAction is a paid mutator transaction binding the contract method 0x748fa599.
//
// Solidity: function someUnsafeAction(address addr) returns(bool)
func (_Caller *CallerSession) SomeUnsafeAction(addr common.Address) (*types.Transaction, error) {
	return _Caller.Contract.SomeUnsafeAction(&_Caller.TransactOpts, addr)
}

// SomeUnsafeAction is a paid mutator transaction binding the contract method 0x748fa599.
//
// Solidity: function someUnsafeAction(address addr) returns(bool)
func (_Caller *CallerTransactorSession) SomeUnsafeAction(addr common.Address) (*types.Transaction, error) {
	return _Caller.Contract.SomeUnsafeAction(&_Caller.TransactOpts, addr)
}

// StoreAction is a paid mutator transaction binding the contract method 0xe8a62235.
//
// Solidity: function storeAction(address addr) returns(uint256)
func (_Caller *CallerTransactor) StoreAction(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Caller.contract.Transact(opts, "storeAction", addr)
}

// StoreAction is a paid mutator transaction binding the contract method 0xe8a62235.
//
// Solidity: function storeAction(address addr) returns(uint256)
func (_Caller *CallerSession) StoreAction(addr common.Address) (*types.Transaction, error) {
	return _Caller.Contract.StoreAction(&_Caller.TransactOpts, addr)
}

// StoreAction is a paid mutator transaction binding the contract method 0xe8a62235.
//
// Solidity: function storeAction(address addr) returns(uint256)
func (_Caller *CallerTransactorSession) StoreAction(addr common.Address) (*types.Transaction, error) {
	return _Caller.Contract.StoreAction(&_Caller.TransactOpts, addr)
}
