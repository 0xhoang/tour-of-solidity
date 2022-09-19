// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package pkg

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

// CalleeMetaData contains all meta data concerning the Callee contract.
var CalleeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"initialValue\",\"type\":\"uint256\"}],\"name\":\"getValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getValues\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"storeValue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"values\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610288806100206000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c80630ff4c9161461005157806319eb4a90146100815780635e383d211461009f5780637221a26a146100cf575b600080fd5b61006b60048036038101906100669190610198565b6100eb565b60405161007891906101d4565b60405180910390f35b610089610101565b60405161009691906101d4565b60405180910390f35b6100b960048036038101906100b49190610198565b61010d565b6040516100c691906101d4565b60405180910390f35b6100e960048036038101906100e49190610198565b610131565b005b60006096826100fa919061021e565b9050919050565b60008080549050905090565b6000818154811061011d57600080fd5b906000526020600020016000915090505481565b600081908060018154018082558091505060019003906000526020600020016000909190919091505550565b600080fd5b6000819050919050565b61017581610162565b811461018057600080fd5b50565b6000813590506101928161016c565b92915050565b6000602082840312156101ae576101ad61015d565b5b60006101bc84828501610183565b91505092915050565b6101ce81610162565b82525050565b60006020820190506101e960008301846101c5565b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061022982610162565b915061023483610162565b925082820190508082111561024c5761024b6101ef565b5b9291505056fea264697066735822122013d912cbf22e0eaef70d592b049ba8f258ce9fd6eafb95c6327b17b46aa4f71c64736f6c63430008100033",
}

// CalleeABI is the input ABI used to generate the binding from.
// Deprecated: Use CalleeMetaData.ABI instead.
var CalleeABI = CalleeMetaData.ABI

// CalleeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CalleeMetaData.Bin instead.
var CalleeBin = CalleeMetaData.Bin

// DeployCallee deploys a new Ethereum contract, binding an instance of Callee to it.
func DeployCallee(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Callee, error) {
	parsed, err := CalleeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CalleeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Callee{CalleeCaller: CalleeCaller{contract: contract}, CalleeTransactor: CalleeTransactor{contract: contract}, CalleeFilterer: CalleeFilterer{contract: contract}}, nil
}

// Callee is an auto generated Go binding around an Ethereum contract.
type Callee struct {
	CalleeCaller     // Read-only binding to the contract
	CalleeTransactor // Write-only binding to the contract
	CalleeFilterer   // Log filterer for contract events
}

// CalleeCaller is an auto generated read-only Go binding around an Ethereum contract.
type CalleeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CalleeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CalleeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CalleeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CalleeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CalleeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CalleeSession struct {
	Contract     *Callee           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CalleeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CalleeCallerSession struct {
	Contract *CalleeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// CalleeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CalleeTransactorSession struct {
	Contract     *CalleeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CalleeRaw is an auto generated low-level Go binding around an Ethereum contract.
type CalleeRaw struct {
	Contract *Callee // Generic contract binding to access the raw methods on
}

// CalleeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CalleeCallerRaw struct {
	Contract *CalleeCaller // Generic read-only contract binding to access the raw methods on
}

// CalleeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CalleeTransactorRaw struct {
	Contract *CalleeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCallee creates a new instance of Callee, bound to a specific deployed contract.
func NewCallee(address common.Address, backend bind.ContractBackend) (*Callee, error) {
	contract, err := bindCallee(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Callee{CalleeCaller: CalleeCaller{contract: contract}, CalleeTransactor: CalleeTransactor{contract: contract}, CalleeFilterer: CalleeFilterer{contract: contract}}, nil
}

// NewCalleeCaller creates a new read-only instance of Callee, bound to a specific deployed contract.
func NewCalleeCaller(address common.Address, caller bind.ContractCaller) (*CalleeCaller, error) {
	contract, err := bindCallee(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CalleeCaller{contract: contract}, nil
}

// NewCalleeTransactor creates a new write-only instance of Callee, bound to a specific deployed contract.
func NewCalleeTransactor(address common.Address, transactor bind.ContractTransactor) (*CalleeTransactor, error) {
	contract, err := bindCallee(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CalleeTransactor{contract: contract}, nil
}

// NewCalleeFilterer creates a new log filterer instance of Callee, bound to a specific deployed contract.
func NewCalleeFilterer(address common.Address, filterer bind.ContractFilterer) (*CalleeFilterer, error) {
	contract, err := bindCallee(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CalleeFilterer{contract: contract}, nil
}

// bindCallee binds a generic wrapper to an already deployed contract.
func bindCallee(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CalleeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Callee *CalleeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Callee.Contract.CalleeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Callee *CalleeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Callee.Contract.CalleeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Callee *CalleeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Callee.Contract.CalleeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Callee *CalleeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Callee.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Callee *CalleeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Callee.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Callee *CalleeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Callee.Contract.contract.Transact(opts, method, params...)
}

// GetValue is a free data retrieval call binding the contract method 0x0ff4c916.
//
// Solidity: function getValue(uint256 initialValue) pure returns(uint256)
func (_Callee *CalleeCaller) GetValue(opts *bind.CallOpts, initialValue *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Callee.contract.Call(opts, &out, "getValue", initialValue)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetValue is a free data retrieval call binding the contract method 0x0ff4c916.
//
// Solidity: function getValue(uint256 initialValue) pure returns(uint256)
func (_Callee *CalleeSession) GetValue(initialValue *big.Int) (*big.Int, error) {
	return _Callee.Contract.GetValue(&_Callee.CallOpts, initialValue)
}

// GetValue is a free data retrieval call binding the contract method 0x0ff4c916.
//
// Solidity: function getValue(uint256 initialValue) pure returns(uint256)
func (_Callee *CalleeCallerSession) GetValue(initialValue *big.Int) (*big.Int, error) {
	return _Callee.Contract.GetValue(&_Callee.CallOpts, initialValue)
}

// Values is a free data retrieval call binding the contract method 0x5e383d21.
//
// Solidity: function values(uint256 ) view returns(uint256)
func (_Callee *CalleeCaller) Values(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Callee.contract.Call(opts, &out, "values", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Values is a free data retrieval call binding the contract method 0x5e383d21.
//
// Solidity: function values(uint256 ) view returns(uint256)
func (_Callee *CalleeSession) Values(arg0 *big.Int) (*big.Int, error) {
	return _Callee.Contract.Values(&_Callee.CallOpts, arg0)
}

// Values is a free data retrieval call binding the contract method 0x5e383d21.
//
// Solidity: function values(uint256 ) view returns(uint256)
func (_Callee *CalleeCallerSession) Values(arg0 *big.Int) (*big.Int, error) {
	return _Callee.Contract.Values(&_Callee.CallOpts, arg0)
}

// GetValues is a paid mutator transaction binding the contract method 0x19eb4a90.
//
// Solidity: function getValues() returns(uint256)
func (_Callee *CalleeTransactor) GetValues(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Callee.contract.Transact(opts, "getValues")
}

// GetValues is a paid mutator transaction binding the contract method 0x19eb4a90.
//
// Solidity: function getValues() returns(uint256)
func (_Callee *CalleeSession) GetValues() (*types.Transaction, error) {
	return _Callee.Contract.GetValues(&_Callee.TransactOpts)
}

// GetValues is a paid mutator transaction binding the contract method 0x19eb4a90.
//
// Solidity: function getValues() returns(uint256)
func (_Callee *CalleeTransactorSession) GetValues() (*types.Transaction, error) {
	return _Callee.Contract.GetValues(&_Callee.TransactOpts)
}

// StoreValue is a paid mutator transaction binding the contract method 0x7221a26a.
//
// Solidity: function storeValue(uint256 value) returns()
func (_Callee *CalleeTransactor) StoreValue(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _Callee.contract.Transact(opts, "storeValue", value)
}

// StoreValue is a paid mutator transaction binding the contract method 0x7221a26a.
//
// Solidity: function storeValue(uint256 value) returns()
func (_Callee *CalleeSession) StoreValue(value *big.Int) (*types.Transaction, error) {
	return _Callee.Contract.StoreValue(&_Callee.TransactOpts, value)
}

// StoreValue is a paid mutator transaction binding the contract method 0x7221a26a.
//
// Solidity: function storeValue(uint256 value) returns()
func (_Callee *CalleeTransactorSession) StoreValue(value *big.Int) (*types.Transaction, error) {
	return _Callee.Contract.StoreValue(&_Callee.TransactOpts, value)
}
