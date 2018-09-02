// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// MainABI is the input ABI used to generate the binding from.
const MainABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"data\",\"type\":\"bytes32[]\"}],\"name\":\"createUser\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"data\",\"type\":\"bytes32[]\"}],\"name\":\"deleteUser\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"data\",\"type\":\"bytes32[]\"},{\"name\":\"holders\",\"type\":\"address[]\"}],\"name\":\"addOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"data\",\"type\":\"bytes32[]\"},{\"name\":\"holders1\",\"type\":\"address[]\"},{\"name\":\"holders2\",\"type\":\"address[]\"}],\"name\":\"removeOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"data\",\"type\":\"bytes32[]\"}],\"name\":\"resetUserData\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"data\",\"type\":\"bytes32[]\"},{\"name\":\"indexes\",\"type\":\"uint256[]\"},{\"name\":\"items\",\"type\":\"bytes32[]\"}],\"name\":\"updateUserData\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Main is an auto generated Go binding around an Ethereum contract.
type Main struct {
	MainCaller     // Read-only binding to the contract
	MainTransactor // Write-only binding to the contract
	MainFilterer   // Log filterer for contract events
}

// MainCaller is an auto generated read-only Go binding around an Ethereum contract.
type MainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MainFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MainSession struct {
	Contract     *Main             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MainCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MainCallerSession struct {
	Contract *MainCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MainTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MainTransactorSession struct {
	Contract     *MainTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MainRaw is an auto generated low-level Go binding around an Ethereum contract.
type MainRaw struct {
	Contract *Main // Generic contract binding to access the raw methods on
}

// MainCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MainCallerRaw struct {
	Contract *MainCaller // Generic read-only contract binding to access the raw methods on
}

// MainTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MainTransactorRaw struct {
	Contract *MainTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMain creates a new instance of Main, bound to a specific deployed contract.
func NewMain(address common.Address, backend bind.ContractBackend) (*Main, error) {
	contract, err := bindMain(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Main{MainCaller: MainCaller{contract: contract}, MainTransactor: MainTransactor{contract: contract}, MainFilterer: MainFilterer{contract: contract}}, nil
}

// NewMainCaller creates a new read-only instance of Main, bound to a specific deployed contract.
func NewMainCaller(address common.Address, caller bind.ContractCaller) (*MainCaller, error) {
	contract, err := bindMain(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MainCaller{contract: contract}, nil
}

// NewMainTransactor creates a new write-only instance of Main, bound to a specific deployed contract.
func NewMainTransactor(address common.Address, transactor bind.ContractTransactor) (*MainTransactor, error) {
	contract, err := bindMain(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MainTransactor{contract: contract}, nil
}

// NewMainFilterer creates a new log filterer instance of Main, bound to a specific deployed contract.
func NewMainFilterer(address common.Address, filterer bind.ContractFilterer) (*MainFilterer, error) {
	contract, err := bindMain(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MainFilterer{contract: contract}, nil
}

// bindMain binds a generic wrapper to an already deployed contract.
func bindMain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MainABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Main *MainRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Main.Contract.MainCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Main *MainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.Contract.MainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Main *MainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Main.Contract.MainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Main *MainCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Main.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Main *MainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Main *MainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Main.Contract.contract.Transact(opts, method, params...)
}

// AddOwner is a paid mutator transaction binding the contract method 0x662f1478.
//
// Solidity: function addOwner(name string, data bytes32[], holders address[]) returns(bool)
func (_Main *MainTransactor) AddOwner(opts *bind.TransactOpts, name string, data [][32]byte, holders []common.Address) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "addOwner", name, data, holders)
}

// AddOwner is a paid mutator transaction binding the contract method 0x662f1478.
//
// Solidity: function addOwner(name string, data bytes32[], holders address[]) returns(bool)
func (_Main *MainSession) AddOwner(name string, data [][32]byte, holders []common.Address) (*types.Transaction, error) {
	return _Main.Contract.AddOwner(&_Main.TransactOpts, name, data, holders)
}

// AddOwner is a paid mutator transaction binding the contract method 0x662f1478.
//
// Solidity: function addOwner(name string, data bytes32[], holders address[]) returns(bool)
func (_Main *MainTransactorSession) AddOwner(name string, data [][32]byte, holders []common.Address) (*types.Transaction, error) {
	return _Main.Contract.AddOwner(&_Main.TransactOpts, name, data, holders)
}

// CreateUser is a paid mutator transaction binding the contract method 0xd57a5e1a.
//
// Solidity: function createUser(name string, data bytes32[]) returns(bool)
func (_Main *MainTransactor) CreateUser(opts *bind.TransactOpts, name string, data [][32]byte) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "createUser", name, data)
}

// CreateUser is a paid mutator transaction binding the contract method 0xd57a5e1a.
//
// Solidity: function createUser(name string, data bytes32[]) returns(bool)
func (_Main *MainSession) CreateUser(name string, data [][32]byte) (*types.Transaction, error) {
	return _Main.Contract.CreateUser(&_Main.TransactOpts, name, data)
}

// CreateUser is a paid mutator transaction binding the contract method 0xd57a5e1a.
//
// Solidity: function createUser(name string, data bytes32[]) returns(bool)
func (_Main *MainTransactorSession) CreateUser(name string, data [][32]byte) (*types.Transaction, error) {
	return _Main.Contract.CreateUser(&_Main.TransactOpts, name, data)
}

// DeleteUser is a paid mutator transaction binding the contract method 0xdec827c1.
//
// Solidity: function deleteUser(name string, data bytes32[]) returns(bool)
func (_Main *MainTransactor) DeleteUser(opts *bind.TransactOpts, name string, data [][32]byte) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "deleteUser", name, data)
}

// DeleteUser is a paid mutator transaction binding the contract method 0xdec827c1.
//
// Solidity: function deleteUser(name string, data bytes32[]) returns(bool)
func (_Main *MainSession) DeleteUser(name string, data [][32]byte) (*types.Transaction, error) {
	return _Main.Contract.DeleteUser(&_Main.TransactOpts, name, data)
}

// DeleteUser is a paid mutator transaction binding the contract method 0xdec827c1.
//
// Solidity: function deleteUser(name string, data bytes32[]) returns(bool)
func (_Main *MainTransactorSession) DeleteUser(name string, data [][32]byte) (*types.Transaction, error) {
	return _Main.Contract.DeleteUser(&_Main.TransactOpts, name, data)
}

// RemoveOwner is a paid mutator transaction binding the contract method 0x19a46927.
//
// Solidity: function removeOwner(name string, data bytes32[], holders1 address[], holders2 address[]) returns(bool)
func (_Main *MainTransactor) RemoveOwner(opts *bind.TransactOpts, name string, data [][32]byte, holders1 []common.Address, holders2 []common.Address) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "removeOwner", name, data, holders1, holders2)
}

// RemoveOwner is a paid mutator transaction binding the contract method 0x19a46927.
//
// Solidity: function removeOwner(name string, data bytes32[], holders1 address[], holders2 address[]) returns(bool)
func (_Main *MainSession) RemoveOwner(name string, data [][32]byte, holders1 []common.Address, holders2 []common.Address) (*types.Transaction, error) {
	return _Main.Contract.RemoveOwner(&_Main.TransactOpts, name, data, holders1, holders2)
}

// RemoveOwner is a paid mutator transaction binding the contract method 0x19a46927.
//
// Solidity: function removeOwner(name string, data bytes32[], holders1 address[], holders2 address[]) returns(bool)
func (_Main *MainTransactorSession) RemoveOwner(name string, data [][32]byte, holders1 []common.Address, holders2 []common.Address) (*types.Transaction, error) {
	return _Main.Contract.RemoveOwner(&_Main.TransactOpts, name, data, holders1, holders2)
}

// ResetUserData is a paid mutator transaction binding the contract method 0x66f46c36.
//
// Solidity: function resetUserData(name string, data bytes32[]) returns(bool)
func (_Main *MainTransactor) ResetUserData(opts *bind.TransactOpts, name string, data [][32]byte) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "resetUserData", name, data)
}

// ResetUserData is a paid mutator transaction binding the contract method 0x66f46c36.
//
// Solidity: function resetUserData(name string, data bytes32[]) returns(bool)
func (_Main *MainSession) ResetUserData(name string, data [][32]byte) (*types.Transaction, error) {
	return _Main.Contract.ResetUserData(&_Main.TransactOpts, name, data)
}

// ResetUserData is a paid mutator transaction binding the contract method 0x66f46c36.
//
// Solidity: function resetUserData(name string, data bytes32[]) returns(bool)
func (_Main *MainTransactorSession) ResetUserData(name string, data [][32]byte) (*types.Transaction, error) {
	return _Main.Contract.ResetUserData(&_Main.TransactOpts, name, data)
}

// UpdateUserData is a paid mutator transaction binding the contract method 0x93177b06.
//
// Solidity: function updateUserData(name string, data bytes32[], indexes uint256[], items bytes32[]) returns(bool)
func (_Main *MainTransactor) UpdateUserData(opts *bind.TransactOpts, name string, data [][32]byte, indexes []*big.Int, items [][32]byte) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "updateUserData", name, data, indexes, items)
}

// UpdateUserData is a paid mutator transaction binding the contract method 0x93177b06.
//
// Solidity: function updateUserData(name string, data bytes32[], indexes uint256[], items bytes32[]) returns(bool)
func (_Main *MainSession) UpdateUserData(name string, data [][32]byte, indexes []*big.Int, items [][32]byte) (*types.Transaction, error) {
	return _Main.Contract.UpdateUserData(&_Main.TransactOpts, name, data, indexes, items)
}

// UpdateUserData is a paid mutator transaction binding the contract method 0x93177b06.
//
// Solidity: function updateUserData(name string, data bytes32[], indexes uint256[], items bytes32[]) returns(bool)
func (_Main *MainTransactorSession) UpdateUserData(name string, data [][32]byte, indexes []*big.Int, items [][32]byte) (*types.Transaction, error) {
	return _Main.Contract.UpdateUserData(&_Main.TransactOpts, name, data, indexes, items)
}
