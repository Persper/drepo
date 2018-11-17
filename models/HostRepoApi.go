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

// HostRepoApiABI is the input ABI used to generate the binding from.
const HostRepoApiABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"host\",\"type\":\"bytes32\"},{\"name\":\"indexes\",\"type\":\"uint256[]\"},{\"name\":\"items\",\"type\":\"bytes32[]\"}],\"name\":\"updateHostData\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"host\",\"type\":\"bytes32\"},{\"name\":\"holders\",\"type\":\"address[]\"}],\"name\":\"removeOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"host\",\"type\":\"bytes32\"},{\"name\":\"holders\",\"type\":\"address[]\"}],\"name\":\"addOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"host\",\"type\":\"bytes32\"},{\"name\":\"data\",\"type\":\"bytes32[]\"}],\"name\":\"resetHostData\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"name\",\"type\":\"string\"}],\"name\":\"deleteHost\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"host\",\"type\":\"bytes32\"}],\"name\":\"retrieveOwner\",\"outputs\":[{\"name\":\"owner\",\"type\":\"address[]\"},{\"name\":\"data\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"data\",\"type\":\"bytes32[]\"}],\"name\":\"createHost\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"host\",\"type\":\"bytes32\"},{\"name\":\"repo\",\"type\":\"bytes32\"}],\"name\":\"retrieveAdmin\",\"outputs\":[{\"name\":\"admin\",\"type\":\"address[]\"},{\"name\":\"data\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"host\",\"type\":\"bytes32\"},{\"name\":\"repo\",\"type\":\"bytes32\"},{\"name\":\"holders\",\"type\":\"address[]\"}],\"name\":\"addAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"host\",\"type\":\"bytes32\"},{\"name\":\"repo\",\"type\":\"bytes32\"},{\"name\":\"holders\",\"type\":\"address[]\"}],\"name\":\"removeAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"host\",\"type\":\"bytes32\"},{\"name\":\"repo\",\"type\":\"bytes32\"},{\"name\":\"data\",\"type\":\"bytes32[]\"}],\"name\":\"resetRepoData\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"host\",\"type\":\"bytes32\"},{\"name\":\"repo\",\"type\":\"bytes32\"},{\"name\":\"indexes\",\"type\":\"uint256[]\"},{\"name\":\"items\",\"type\":\"bytes32[]\"}],\"name\":\"updateRepoData\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// HostRepoApi is an auto generated Go binding around an Ethereum contract.
type HostRepoApi struct {
	HostRepoApiCaller     // Read-only binding to the contract
	HostRepoApiTransactor // Write-only binding to the contract
	HostRepoApiFilterer   // Log filterer for contract events
}

// HostRepoApiCaller is an auto generated read-only Go binding around an Ethereum contract.
type HostRepoApiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HostRepoApiTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HostRepoApiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HostRepoApiFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HostRepoApiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HostRepoApiSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HostRepoApiSession struct {
	Contract     *HostRepoApi      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HostRepoApiCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HostRepoApiCallerSession struct {
	Contract *HostRepoApiCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// HostRepoApiTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HostRepoApiTransactorSession struct {
	Contract     *HostRepoApiTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// HostRepoApiRaw is an auto generated low-level Go binding around an Ethereum contract.
type HostRepoApiRaw struct {
	Contract *HostRepoApi // Generic contract binding to access the raw methods on
}

// HostRepoApiCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HostRepoApiCallerRaw struct {
	Contract *HostRepoApiCaller // Generic read-only contract binding to access the raw methods on
}

// HostRepoApiTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HostRepoApiTransactorRaw struct {
	Contract *HostRepoApiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHostRepoApi creates a new instance of HostRepoApi, bound to a specific deployed contract.
func NewHostRepoApi(address common.Address, backend bind.ContractBackend) (*HostRepoApi, error) {
	contract, err := bindHostRepoApi(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HostRepoApi{HostRepoApiCaller: HostRepoApiCaller{contract: contract}, HostRepoApiTransactor: HostRepoApiTransactor{contract: contract}, HostRepoApiFilterer: HostRepoApiFilterer{contract: contract}}, nil
}

// NewHostRepoApiCaller creates a new read-only instance of HostRepoApi, bound to a specific deployed contract.
func NewHostRepoApiCaller(address common.Address, caller bind.ContractCaller) (*HostRepoApiCaller, error) {
	contract, err := bindHostRepoApi(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HostRepoApiCaller{contract: contract}, nil
}

// NewHostRepoApiTransactor creates a new write-only instance of HostRepoApi, bound to a specific deployed contract.
func NewHostRepoApiTransactor(address common.Address, transactor bind.ContractTransactor) (*HostRepoApiTransactor, error) {
	contract, err := bindHostRepoApi(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HostRepoApiTransactor{contract: contract}, nil
}

// NewHostRepoApiFilterer creates a new log filterer instance of HostRepoApi, bound to a specific deployed contract.
func NewHostRepoApiFilterer(address common.Address, filterer bind.ContractFilterer) (*HostRepoApiFilterer, error) {
	contract, err := bindHostRepoApi(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HostRepoApiFilterer{contract: contract}, nil
}

// bindHostRepoApi binds a generic wrapper to an already deployed contract.
func bindHostRepoApi(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HostRepoApiABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HostRepoApi *HostRepoApiRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _HostRepoApi.Contract.HostRepoApiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HostRepoApi *HostRepoApiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HostRepoApi.Contract.HostRepoApiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HostRepoApi *HostRepoApiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HostRepoApi.Contract.HostRepoApiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HostRepoApi *HostRepoApiCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _HostRepoApi.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HostRepoApi *HostRepoApiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HostRepoApi.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HostRepoApi *HostRepoApiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HostRepoApi.Contract.contract.Transact(opts, method, params...)
}

// RetrieveAdmin is a free data retrieval call binding the contract method 0xe6d118a0.
//
// Solidity: function retrieveAdmin(host bytes32, repo bytes32) constant returns(admin address[], data bytes32[])
func (_HostRepoApi *HostRepoApiCaller) RetrieveAdmin(opts *bind.CallOpts, host [32]byte, repo [32]byte) (struct {
	Admin []common.Address
	Data  [][32]byte
}, error) {
	ret := new(struct {
		Admin []common.Address
		Data  [][32]byte
	})
	out := ret
	err := _HostRepoApi.contract.Call(opts, out, "retrieveAdmin", host, repo)
	return *ret, err
}

// RetrieveAdmin is a free data retrieval call binding the contract method 0xe6d118a0.
//
// Solidity: function retrieveAdmin(host bytes32, repo bytes32) constant returns(admin address[], data bytes32[])
func (_HostRepoApi *HostRepoApiSession) RetrieveAdmin(host [32]byte, repo [32]byte) (struct {
	Admin []common.Address
	Data  [][32]byte
}, error) {
	return _HostRepoApi.Contract.RetrieveAdmin(&_HostRepoApi.CallOpts, host, repo)
}

// RetrieveAdmin is a free data retrieval call binding the contract method 0xe6d118a0.
//
// Solidity: function retrieveAdmin(host bytes32, repo bytes32) constant returns(admin address[], data bytes32[])
func (_HostRepoApi *HostRepoApiCallerSession) RetrieveAdmin(host [32]byte, repo [32]byte) (struct {
	Admin []common.Address
	Data  [][32]byte
}, error) {
	return _HostRepoApi.Contract.RetrieveAdmin(&_HostRepoApi.CallOpts, host, repo)
}

// RetrieveOwner is a free data retrieval call binding the contract method 0xaca9fb46.
//
// Solidity: function retrieveOwner(host bytes32) constant returns(owner address[], data bytes32[])
func (_HostRepoApi *HostRepoApiCaller) RetrieveOwner(opts *bind.CallOpts, host [32]byte) (struct {
	Owner []common.Address
	Data  [][32]byte
}, error) {
	ret := new(struct {
		Owner []common.Address
		Data  [][32]byte
	})
	out := ret
	err := _HostRepoApi.contract.Call(opts, out, "retrieveOwner", host)
	return *ret, err
}

// RetrieveOwner is a free data retrieval call binding the contract method 0xaca9fb46.
//
// Solidity: function retrieveOwner(host bytes32) constant returns(owner address[], data bytes32[])
func (_HostRepoApi *HostRepoApiSession) RetrieveOwner(host [32]byte) (struct {
	Owner []common.Address
	Data  [][32]byte
}, error) {
	return _HostRepoApi.Contract.RetrieveOwner(&_HostRepoApi.CallOpts, host)
}

// RetrieveOwner is a free data retrieval call binding the contract method 0xaca9fb46.
//
// Solidity: function retrieveOwner(host bytes32) constant returns(owner address[], data bytes32[])
func (_HostRepoApi *HostRepoApiCallerSession) RetrieveOwner(host [32]byte) (struct {
	Owner []common.Address
	Data  [][32]byte
}, error) {
	return _HostRepoApi.Contract.RetrieveOwner(&_HostRepoApi.CallOpts, host)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x37c237fb.
//
// Solidity: function addAdmin(host bytes32, repo bytes32, holders address[]) returns()
func (_HostRepoApi *HostRepoApiTransactor) AddAdmin(opts *bind.TransactOpts, host [32]byte, repo [32]byte, holders []common.Address) (*types.Transaction, error) {
	return _HostRepoApi.contract.Transact(opts, "addAdmin", host, repo, holders)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x37c237fb.
//
// Solidity: function addAdmin(host bytes32, repo bytes32, holders address[]) returns()
func (_HostRepoApi *HostRepoApiSession) AddAdmin(host [32]byte, repo [32]byte, holders []common.Address) (*types.Transaction, error) {
	return _HostRepoApi.Contract.AddAdmin(&_HostRepoApi.TransactOpts, host, repo, holders)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x37c237fb.
//
// Solidity: function addAdmin(host bytes32, repo bytes32, holders address[]) returns()
func (_HostRepoApi *HostRepoApiTransactorSession) AddAdmin(host [32]byte, repo [32]byte, holders []common.Address) (*types.Transaction, error) {
	return _HostRepoApi.Contract.AddAdmin(&_HostRepoApi.TransactOpts, host, repo, holders)
}

// AddOwner is a paid mutator transaction binding the contract method 0x5c35d072.
//
// Solidity: function addOwner(host bytes32, holders address[]) returns()
func (_HostRepoApi *HostRepoApiTransactor) AddOwner(opts *bind.TransactOpts, host [32]byte, holders []common.Address) (*types.Transaction, error) {
	return _HostRepoApi.contract.Transact(opts, "addOwner", host, holders)
}

// AddOwner is a paid mutator transaction binding the contract method 0x5c35d072.
//
// Solidity: function addOwner(host bytes32, holders address[]) returns()
func (_HostRepoApi *HostRepoApiSession) AddOwner(host [32]byte, holders []common.Address) (*types.Transaction, error) {
	return _HostRepoApi.Contract.AddOwner(&_HostRepoApi.TransactOpts, host, holders)
}

// AddOwner is a paid mutator transaction binding the contract method 0x5c35d072.
//
// Solidity: function addOwner(host bytes32, holders address[]) returns()
func (_HostRepoApi *HostRepoApiTransactorSession) AddOwner(host [32]byte, holders []common.Address) (*types.Transaction, error) {
	return _HostRepoApi.Contract.AddOwner(&_HostRepoApi.TransactOpts, host, holders)
}

// CreateHost is a paid mutator transaction binding the contract method 0xc0247349.
//
// Solidity: function createHost(name string, data bytes32[]) returns()
func (_HostRepoApi *HostRepoApiTransactor) CreateHost(opts *bind.TransactOpts, name string, data [][32]byte) (*types.Transaction, error) {
	return _HostRepoApi.contract.Transact(opts, "createHost", name, data)
}

// CreateHost is a paid mutator transaction binding the contract method 0xc0247349.
//
// Solidity: function createHost(name string, data bytes32[]) returns()
func (_HostRepoApi *HostRepoApiSession) CreateHost(name string, data [][32]byte) (*types.Transaction, error) {
	return _HostRepoApi.Contract.CreateHost(&_HostRepoApi.TransactOpts, name, data)
}

// CreateHost is a paid mutator transaction binding the contract method 0xc0247349.
//
// Solidity: function createHost(name string, data bytes32[]) returns()
func (_HostRepoApi *HostRepoApiTransactorSession) CreateHost(name string, data [][32]byte) (*types.Transaction, error) {
	return _HostRepoApi.Contract.CreateHost(&_HostRepoApi.TransactOpts, name, data)
}

// DeleteHost is a paid mutator transaction binding the contract method 0xa60864a4.
//
// Solidity: function deleteHost(name string) returns()
func (_HostRepoApi *HostRepoApiTransactor) DeleteHost(opts *bind.TransactOpts, name string) (*types.Transaction, error) {
	return _HostRepoApi.contract.Transact(opts, "deleteHost", name)
}

// DeleteHost is a paid mutator transaction binding the contract method 0xa60864a4.
//
// Solidity: function deleteHost(name string) returns()
func (_HostRepoApi *HostRepoApiSession) DeleteHost(name string) (*types.Transaction, error) {
	return _HostRepoApi.Contract.DeleteHost(&_HostRepoApi.TransactOpts, name)
}

// DeleteHost is a paid mutator transaction binding the contract method 0xa60864a4.
//
// Solidity: function deleteHost(name string) returns()
func (_HostRepoApi *HostRepoApiTransactorSession) DeleteHost(name string) (*types.Transaction, error) {
	return _HostRepoApi.Contract.DeleteHost(&_HostRepoApi.TransactOpts, name)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0xf6f66166.
//
// Solidity: function removeAdmin(host bytes32, repo bytes32, holders address[]) returns()
func (_HostRepoApi *HostRepoApiTransactor) RemoveAdmin(opts *bind.TransactOpts, host [32]byte, repo [32]byte, holders []common.Address) (*types.Transaction, error) {
	return _HostRepoApi.contract.Transact(opts, "removeAdmin", host, repo, holders)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0xf6f66166.
//
// Solidity: function removeAdmin(host bytes32, repo bytes32, holders address[]) returns()
func (_HostRepoApi *HostRepoApiSession) RemoveAdmin(host [32]byte, repo [32]byte, holders []common.Address) (*types.Transaction, error) {
	return _HostRepoApi.Contract.RemoveAdmin(&_HostRepoApi.TransactOpts, host, repo, holders)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0xf6f66166.
//
// Solidity: function removeAdmin(host bytes32, repo bytes32, holders address[]) returns()
func (_HostRepoApi *HostRepoApiTransactorSession) RemoveAdmin(host [32]byte, repo [32]byte, holders []common.Address) (*types.Transaction, error) {
	return _HostRepoApi.Contract.RemoveAdmin(&_HostRepoApi.TransactOpts, host, repo, holders)
}

// RemoveOwner is a paid mutator transaction binding the contract method 0x39cc75aa.
//
// Solidity: function removeOwner(host bytes32, holders address[]) returns()
func (_HostRepoApi *HostRepoApiTransactor) RemoveOwner(opts *bind.TransactOpts, host [32]byte, holders []common.Address) (*types.Transaction, error) {
	return _HostRepoApi.contract.Transact(opts, "removeOwner", host, holders)
}

// RemoveOwner is a paid mutator transaction binding the contract method 0x39cc75aa.
//
// Solidity: function removeOwner(host bytes32, holders address[]) returns()
func (_HostRepoApi *HostRepoApiSession) RemoveOwner(host [32]byte, holders []common.Address) (*types.Transaction, error) {
	return _HostRepoApi.Contract.RemoveOwner(&_HostRepoApi.TransactOpts, host, holders)
}

// RemoveOwner is a paid mutator transaction binding the contract method 0x39cc75aa.
//
// Solidity: function removeOwner(host bytes32, holders address[]) returns()
func (_HostRepoApi *HostRepoApiTransactorSession) RemoveOwner(host [32]byte, holders []common.Address) (*types.Transaction, error) {
	return _HostRepoApi.Contract.RemoveOwner(&_HostRepoApi.TransactOpts, host, holders)
}

// ResetHostData is a paid mutator transaction binding the contract method 0xa51f9275.
//
// Solidity: function resetHostData(host bytes32, data bytes32[]) returns()
func (_HostRepoApi *HostRepoApiTransactor) ResetHostData(opts *bind.TransactOpts, host [32]byte, data [][32]byte) (*types.Transaction, error) {
	return _HostRepoApi.contract.Transact(opts, "resetHostData", host, data)
}

// ResetHostData is a paid mutator transaction binding the contract method 0xa51f9275.
//
// Solidity: function resetHostData(host bytes32, data bytes32[]) returns()
func (_HostRepoApi *HostRepoApiSession) ResetHostData(host [32]byte, data [][32]byte) (*types.Transaction, error) {
	return _HostRepoApi.Contract.ResetHostData(&_HostRepoApi.TransactOpts, host, data)
}

// ResetHostData is a paid mutator transaction binding the contract method 0xa51f9275.
//
// Solidity: function resetHostData(host bytes32, data bytes32[]) returns()
func (_HostRepoApi *HostRepoApiTransactorSession) ResetHostData(host [32]byte, data [][32]byte) (*types.Transaction, error) {
	return _HostRepoApi.Contract.ResetHostData(&_HostRepoApi.TransactOpts, host, data)
}

// ResetRepoData is a paid mutator transaction binding the contract method 0x5f260715.
//
// Solidity: function resetRepoData(host bytes32, repo bytes32, data bytes32[]) returns()
func (_HostRepoApi *HostRepoApiTransactor) ResetRepoData(opts *bind.TransactOpts, host [32]byte, repo [32]byte, data [][32]byte) (*types.Transaction, error) {
	return _HostRepoApi.contract.Transact(opts, "resetRepoData", host, repo, data)
}

// ResetRepoData is a paid mutator transaction binding the contract method 0x5f260715.
//
// Solidity: function resetRepoData(host bytes32, repo bytes32, data bytes32[]) returns()
func (_HostRepoApi *HostRepoApiSession) ResetRepoData(host [32]byte, repo [32]byte, data [][32]byte) (*types.Transaction, error) {
	return _HostRepoApi.Contract.ResetRepoData(&_HostRepoApi.TransactOpts, host, repo, data)
}

// ResetRepoData is a paid mutator transaction binding the contract method 0x5f260715.
//
// Solidity: function resetRepoData(host bytes32, repo bytes32, data bytes32[]) returns()
func (_HostRepoApi *HostRepoApiTransactorSession) ResetRepoData(host [32]byte, repo [32]byte, data [][32]byte) (*types.Transaction, error) {
	return _HostRepoApi.Contract.ResetRepoData(&_HostRepoApi.TransactOpts, host, repo, data)
}

// UpdateHostData is a paid mutator transaction binding the contract method 0x202f58d9.
//
// Solidity: function updateHostData(host bytes32, indexes uint256[], items bytes32[]) returns()
func (_HostRepoApi *HostRepoApiTransactor) UpdateHostData(opts *bind.TransactOpts, host [32]byte, indexes []*big.Int, items [][32]byte) (*types.Transaction, error) {
	return _HostRepoApi.contract.Transact(opts, "updateHostData", host, indexes, items)
}

// UpdateHostData is a paid mutator transaction binding the contract method 0x202f58d9.
//
// Solidity: function updateHostData(host bytes32, indexes uint256[], items bytes32[]) returns()
func (_HostRepoApi *HostRepoApiSession) UpdateHostData(host [32]byte, indexes []*big.Int, items [][32]byte) (*types.Transaction, error) {
	return _HostRepoApi.Contract.UpdateHostData(&_HostRepoApi.TransactOpts, host, indexes, items)
}

// UpdateHostData is a paid mutator transaction binding the contract method 0x202f58d9.
//
// Solidity: function updateHostData(host bytes32, indexes uint256[], items bytes32[]) returns()
func (_HostRepoApi *HostRepoApiTransactorSession) UpdateHostData(host [32]byte, indexes []*big.Int, items [][32]byte) (*types.Transaction, error) {
	return _HostRepoApi.Contract.UpdateHostData(&_HostRepoApi.TransactOpts, host, indexes, items)
}

// UpdateRepoData is a paid mutator transaction binding the contract method 0x88c5d860.
//
// Solidity: function updateRepoData(host bytes32, repo bytes32, indexes uint256[], items bytes32[]) returns()
func (_HostRepoApi *HostRepoApiTransactor) UpdateRepoData(opts *bind.TransactOpts, host [32]byte, repo [32]byte, indexes []*big.Int, items [][32]byte) (*types.Transaction, error) {
	return _HostRepoApi.contract.Transact(opts, "updateRepoData", host, repo, indexes, items)
}

// UpdateRepoData is a paid mutator transaction binding the contract method 0x88c5d860.
//
// Solidity: function updateRepoData(host bytes32, repo bytes32, indexes uint256[], items bytes32[]) returns()
func (_HostRepoApi *HostRepoApiSession) UpdateRepoData(host [32]byte, repo [32]byte, indexes []*big.Int, items [][32]byte) (*types.Transaction, error) {
	return _HostRepoApi.Contract.UpdateRepoData(&_HostRepoApi.TransactOpts, host, repo, indexes, items)
}

// UpdateRepoData is a paid mutator transaction binding the contract method 0x88c5d860.
//
// Solidity: function updateRepoData(host bytes32, repo bytes32, indexes uint256[], items bytes32[]) returns()
func (_HostRepoApi *HostRepoApiTransactorSession) UpdateRepoData(host [32]byte, repo [32]byte, indexes []*big.Int, items [][32]byte) (*types.Transaction, error) {
	return _HostRepoApi.Contract.UpdateRepoData(&_HostRepoApi.TransactOpts, host, repo, indexes, items)
}
