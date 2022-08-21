// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package caller

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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"someAction\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"someUnsafeAction\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"storeAction\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506104fd806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c8063748fa59914610046578063e8a6223514610062578063fba9cea914610092575b600080fd5b610060600480360381019061005b91906102c7565b6100c2565b005b61007c600480360381019061007791906102c7565b6100f0565b604051610089919061030d565b60405180910390f35b6100ac60048036038101906100a791906102c7565b6101db565b6040516100b9919061030d565b60405180910390f35b6064816040516020016100d69291906103e6565b604051602081830303815290604052805190602001205050565b6000808290508073ffffffffffffffffffffffffffffffffffffffff16637221a26a60646040518263ffffffff1660e01b81526004016101309190610453565b600060405180830381600087803b15801561014a57600080fd5b505af115801561015e573d6000803e3d6000fd5b505050508073ffffffffffffffffffffffffffffffffffffffff166319eb4a906040518163ffffffff1660e01b81526004016020604051808303816000875af11580156101af573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906101d3919061049a565b915050919050565b6000808290508073ffffffffffffffffffffffffffffffffffffffff16630ff4c91660646040518263ffffffff1660e01b815260040161021b9190610453565b602060405180830381865afa158015610238573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061025c919061049a565b915050919050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061029482610269565b9050919050565b6102a481610289565b81146102af57600080fd5b50565b6000813590506102c18161029b565b92915050565b6000602082840312156102dd576102dc610264565b5b60006102eb848285016102b2565b91505092915050565b6000819050919050565b610307816102f4565b82525050565b600060208201905061032260008301846102fe565b92915050565b600082825260208201905092915050565b7f73746f726556616c75652875696e743235362900000000000000000000000000600082015250565b600061036f601383610328565b915061037a82610339565b602082019050919050565b6000819050919050565b600060ff82169050919050565b6000819050919050565b60006103c16103bc6103b784610385565b61039c565b61038f565b9050919050565b6103d1816103a6565b82525050565b6103e081610289565b82525050565b600060608201905081810360008301526103ff81610362565b905061040e60208301856103c8565b61041b60408301846103d7565b9392505050565b600061043d61043861043384610385565b61039c565b6102f4565b9050919050565b61044d81610422565b82525050565b60006020820190506104686000830184610444565b92915050565b610477816102f4565b811461048257600080fd5b50565b6000815190506104948161046e565b92915050565b6000602082840312156104b0576104af610264565b5b60006104be84828501610485565b9150509291505056fea2646970667358221220bdc928941ec43076bc8f47762400030c059f7d616964de557aa2da5c7b34049964736f6c63430008100033",
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

// SomeUnsafeAction is a free data retrieval call binding the contract method 0x748fa599.
//
// Solidity: function someUnsafeAction(address addr) pure returns()
func (_Caller *CallerCaller) SomeUnsafeAction(opts *bind.CallOpts, addr common.Address) error {
	var out []interface{}
	err := _Caller.contract.Call(opts, &out, "someUnsafeAction", addr)

	if err != nil {
		return err
	}

	return err

}

// SomeUnsafeAction is a free data retrieval call binding the contract method 0x748fa599.
//
// Solidity: function someUnsafeAction(address addr) pure returns()
func (_Caller *CallerSession) SomeUnsafeAction(addr common.Address) error {
	return _Caller.Contract.SomeUnsafeAction(&_Caller.CallOpts, addr)
}

// SomeUnsafeAction is a free data retrieval call binding the contract method 0x748fa599.
//
// Solidity: function someUnsafeAction(address addr) pure returns()
func (_Caller *CallerCallerSession) SomeUnsafeAction(addr common.Address) error {
	return _Caller.Contract.SomeUnsafeAction(&_Caller.CallOpts, addr)
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
