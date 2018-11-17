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

// HostApiABI is the input ABI used to generate the binding from.
const HostApiABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"host\",\"type\":\"bytes32\"}],\"name\":\"retrieveOwner\",\"outputs\":[{\"name\":\"owner\",\"type\":\"address[]\"},{\"name\":\"data\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"data\",\"type\":\"bytes32[]\"}],\"name\":\"createHost\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"name\",\"type\":\"string\"}],\"name\":\"deleteHost\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"host\",\"type\":\"bytes32\"},{\"name\":\"holders\",\"type\":\"address[]\"}],\"name\":\"addOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"host\",\"type\":\"bytes32\"},{\"name\":\"holders\",\"type\":\"address[]\"}],\"name\":\"removeOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"host\",\"type\":\"bytes32\"},{\"name\":\"data\",\"type\":\"bytes32[]\"}],\"name\":\"resetHostData\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"host\",\"type\":\"bytes32\"},{\"name\":\"indexes\",\"type\":\"uint256[]\"},{\"name\":\"items\",\"type\":\"bytes32[]\"}],\"name\":\"updateHostData\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// HostApi is an auto generated Go binding around an Ethereum contract.
type HostApi struct {
	HostApiCaller     // Read-only binding to the contract
	HostApiTransactor // Write-only binding to the contract
	HostApiFilterer   // Log filterer for contract events
}

// HostApiCaller is an auto generated read-only Go binding around an Ethereum contract.
type HostApiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HostApiTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HostApiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HostApiFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HostApiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HostApiSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HostApiSession struct {
	Contract     *HostApi          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HostApiCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HostApiCallerSession struct {
	Contract *HostApiCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// HostApiTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HostApiTransactorSession struct {
	Contract     *HostApiTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// HostApiRaw is an auto generated low-level Go binding around an Ethereum contract.
type HostApiRaw struct {
	Contract *HostApi // Generic contract binding to access the raw methods on
}

// HostApiCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HostApiCallerRaw struct {
	Contract *HostApiCaller // Generic read-only contract binding to access the raw methods on
}

// HostApiTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HostApiTransactorRaw struct {
	Contract *HostApiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHostApi creates a new instance of HostApi, bound to a specific deployed contract.
func NewHostApi(address common.Address, backend bind.ContractBackend) (*HostApi, error) {
	contract, err := bindHostApi(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HostApi{HostApiCaller: HostApiCaller{contract: contract}, HostApiTransactor: HostApiTransactor{contract: contract}, HostApiFilterer: HostApiFilterer{contract: contract}}, nil
}

// NewHostApiCaller creates a new read-only instance of HostApi, bound to a specific deployed contract.
func NewHostApiCaller(address common.Address, caller bind.ContractCaller) (*HostApiCaller, error) {
	contract, err := bindHostApi(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HostApiCaller{contract: contract}, nil
}

// NewHostApiTransactor creates a new write-only instance of HostApi, bound to a specific deployed contract.
func NewHostApiTransactor(address common.Address, transactor bind.ContractTransactor) (*HostApiTransactor, error) {
	contract, err := bindHostApi(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HostApiTransactor{contract: contract}, nil
}

// NewHostApiFilterer creates a new log filterer instance of HostApi, bound to a specific deployed contract.
func NewHostApiFilterer(address common.Address, filterer bind.ContractFilterer) (*HostApiFilterer, error) {
	contract, err := bindHostApi(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HostApiFilterer{contract: contract}, nil
}

// bindHostApi binds a generic wrapper to an already deployed contract.
func bindHostApi(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HostApiABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HostApi *HostApiRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _HostApi.Contract.HostApiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HostApi *HostApiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HostApi.Contract.HostApiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HostApi *HostApiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HostApi.Contract.HostApiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HostApi *HostApiCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _HostApi.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HostApi *HostApiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HostApi.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HostApi *HostApiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HostApi.Contract.contract.Transact(opts, method, params...)
}

// RetrieveOwner is a free data retrieval call binding the contract method 0xaca9fb46.
//
// Solidity: function retrieveOwner(host bytes32) constant returns(owner address[], data bytes32[])
func (_HostApi *HostApiCaller) RetrieveOwner(opts *bind.CallOpts, host [32]byte) (struct {
	Owner []common.Address
	Data  [][32]byte
}, error) {
	ret := new(struct {
		Owner []common.Address
		Data  [][32]byte
	})
	out := ret
	err := _HostApi.contract.Call(opts, out, "retrieveOwner", host)
	return *ret, err
}

// RetrieveOwner is a free data retrieval call binding the contract method 0xaca9fb46.
//
// Solidity: function retrieveOwner(host bytes32) constant returns(owner address[], data bytes32[])
func (_HostApi *HostApiSession) RetrieveOwner(host [32]byte) (struct {
	Owner []common.Address
	Data  [][32]byte
}, error) {
	return _HostApi.Contract.RetrieveOwner(&_HostApi.CallOpts, host)
}

// RetrieveOwner is a free data retrieval call binding the contract method 0xaca9fb46.
//
// Solidity: function retrieveOwner(host bytes32) constant returns(owner address[], data bytes32[])
func (_HostApi *HostApiCallerSession) RetrieveOwner(host [32]byte) (struct {
	Owner []common.Address
	Data  [][32]byte
}, error) {
	return _HostApi.Contract.RetrieveOwner(&_HostApi.CallOpts, host)
}

// AddOwner is a paid mutator transaction binding the contract method 0x5c35d072.
//
// Solidity: function addOwner(host bytes32, holders address[]) returns()
func (_HostApi *HostApiTransactor) AddOwner(opts *bind.TransactOpts, host [32]byte, holders []common.Address) (*types.Transaction, error) {
	return _HostApi.contract.Transact(opts, "addOwner", host, holders)
}

// AddOwner is a paid mutator transaction binding the contract method 0x5c35d072.
//
// Solidity: function addOwner(host bytes32, holders address[]) returns()
func (_HostApi *HostApiSession) AddOwner(host [32]byte, holders []common.Address) (*types.Transaction, error) {
	return _HostApi.Contract.AddOwner(&_HostApi.TransactOpts, host, holders)
}

// AddOwner is a paid mutator transaction binding the contract method 0x5c35d072.
//
// Solidity: function addOwner(host bytes32, holders address[]) returns()
func (_HostApi *HostApiTransactorSession) AddOwner(host [32]byte, holders []common.Address) (*types.Transaction, error) {
	return _HostApi.Contract.AddOwner(&_HostApi.TransactOpts, host, holders)
}

// CreateHost is a paid mutator transaction binding the contract method 0xc0247349.
//
// Solidity: function createHost(name string, data bytes32[]) returns()
func (_HostApi *HostApiTransactor) CreateHost(opts *bind.TransactOpts, name string, data [][32]byte) (*types.Transaction, error) {
	return _HostApi.contract.Transact(opts, "createHost", name, data)
}

// CreateHost is a paid mutator transaction binding the contract method 0xc0247349.
//
// Solidity: function createHost(name string, data bytes32[]) returns()
func (_HostApi *HostApiSession) CreateHost(name string, data [][32]byte) (*types.Transaction, error) {
	return _HostApi.Contract.CreateHost(&_HostApi.TransactOpts, name, data)
}

// CreateHost is a paid mutator transaction binding the contract method 0xc0247349.
//
// Solidity: function createHost(name string, data bytes32[]) returns()
func (_HostApi *HostApiTransactorSession) CreateHost(name string, data [][32]byte) (*types.Transaction, error) {
	return _HostApi.Contract.CreateHost(&_HostApi.TransactOpts, name, data)
}

// DeleteHost is a paid mutator transaction binding the contract method 0xa60864a4.
//
// Solidity: function deleteHost(name string) returns()
func (_HostApi *HostApiTransactor) DeleteHost(opts *bind.TransactOpts, name string) (*types.Transaction, error) {
	return _HostApi.contract.Transact(opts, "deleteHost", name)
}

// DeleteHost is a paid mutator transaction binding the contract method 0xa60864a4.
//
// Solidity: function deleteHost(name string) returns()
func (_HostApi *HostApiSession) DeleteHost(name string) (*types.Transaction, error) {
	return _HostApi.Contract.DeleteHost(&_HostApi.TransactOpts, name)
}

// DeleteHost is a paid mutator transaction binding the contract method 0xa60864a4.
//
// Solidity: function deleteHost(name string) returns()
func (_HostApi *HostApiTransactorSession) DeleteHost(name string) (*types.Transaction, error) {
	return _HostApi.Contract.DeleteHost(&_HostApi.TransactOpts, name)
}

// RemoveOwner is a paid mutator transaction binding the contract method 0x39cc75aa.
//
// Solidity: function removeOwner(host bytes32, holders address[]) returns()
func (_HostApi *HostApiTransactor) RemoveOwner(opts *bind.TransactOpts, host [32]byte, holders []common.Address) (*types.Transaction, error) {
	return _HostApi.contract.Transact(opts, "removeOwner", host, holders)
}

// RemoveOwner is a paid mutator transaction binding the contract method 0x39cc75aa.
//
// Solidity: function removeOwner(host bytes32, holders address[]) returns()
func (_HostApi *HostApiSession) RemoveOwner(host [32]byte, holders []common.Address) (*types.Transaction, error) {
	return _HostApi.Contract.RemoveOwner(&_HostApi.TransactOpts, host, holders)
}

// RemoveOwner is a paid mutator transaction binding the contract method 0x39cc75aa.
//
// Solidity: function removeOwner(host bytes32, holders address[]) returns()
func (_HostApi *HostApiTransactorSession) RemoveOwner(host [32]byte, holders []common.Address) (*types.Transaction, error) {
	return _HostApi.Contract.RemoveOwner(&_HostApi.TransactOpts, host, holders)
}

// ResetHostData is a paid mutator transaction binding the contract method 0xa51f9275.
//
// Solidity: function resetHostData(host bytes32, data bytes32[]) returns()
func (_HostApi *HostApiTransactor) ResetHostData(opts *bind.TransactOpts, host [32]byte, data [][32]byte) (*types.Transaction, error) {
	return _HostApi.contract.Transact(opts, "resetHostData", host, data)
}

// ResetHostData is a paid mutator transaction binding the contract method 0xa51f9275.
//
// Solidity: function resetHostData(host bytes32, data bytes32[]) returns()
func (_HostApi *HostApiSession) ResetHostData(host [32]byte, data [][32]byte) (*types.Transaction, error) {
	return _HostApi.Contract.ResetHostData(&_HostApi.TransactOpts, host, data)
}

// ResetHostData is a paid mutator transaction binding the contract method 0xa51f9275.
//
// Solidity: function resetHostData(host bytes32, data bytes32[]) returns()
func (_HostApi *HostApiTransactorSession) ResetHostData(host [32]byte, data [][32]byte) (*types.Transaction, error) {
	return _HostApi.Contract.ResetHostData(&_HostApi.TransactOpts, host, data)
}

// UpdateHostData is a paid mutator transaction binding the contract method 0x202f58d9.
//
// Solidity: function updateHostData(host bytes32, indexes uint256[], items bytes32[]) returns()
func (_HostApi *HostApiTransactor) UpdateHostData(opts *bind.TransactOpts, host [32]byte, indexes []*big.Int, items [][32]byte) (*types.Transaction, error) {
	return _HostApi.contract.Transact(opts, "updateHostData", host, indexes, items)
}

// UpdateHostData is a paid mutator transaction binding the contract method 0x202f58d9.
//
// Solidity: function updateHostData(host bytes32, indexes uint256[], items bytes32[]) returns()
func (_HostApi *HostApiSession) UpdateHostData(host [32]byte, indexes []*big.Int, items [][32]byte) (*types.Transaction, error) {
	return _HostApi.Contract.UpdateHostData(&_HostApi.TransactOpts, host, indexes, items)
}

// UpdateHostData is a paid mutator transaction binding the contract method 0x202f58d9.
//
// Solidity: function updateHostData(host bytes32, indexes uint256[], items bytes32[]) returns()
func (_HostApi *HostApiTransactorSession) UpdateHostData(host [32]byte, indexes []*big.Int, items [][32]byte) (*types.Transaction, error) {
	return _HostApi.Contract.UpdateHostData(&_HostApi.TransactOpts, host, indexes, items)
}
