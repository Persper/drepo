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

// HostRepoDevApiABI is the input ABI used to generate the binding from.
const HostRepoDevApiABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"host\",\"type\":\"bytes32\"},{\"name\":\"indexes\",\"type\":\"uint256[]\"},{\"name\":\"items\",\"type\":\"bytes32[]\"}],\"name\":\"updateHostData\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"host\",\"type\":\"bytes32\"},{\"name\":\"repo\",\"type\":\"bytes32\"},{\"name\":\"holders\",\"type\":\"address[]\"}],\"name\":\"addAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"host\",\"type\":\"bytes32\"},{\"name\":\"holders\",\"type\":\"address[]\"}],\"name\":\"removeOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"host\",\"type\":\"bytes32\"},{\"name\":\"repo\",\"type\":\"bytes32\"}],\"name\":\"retrieveWrite\",\"outputs\":[{\"name\":\"admin\",\"type\":\"address[]\"},{\"name\":\"data\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"host\",\"type\":\"bytes32\"},{\"name\":\"holders\",\"type\":\"address[]\"}],\"name\":\"addOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"host\",\"type\":\"bytes32\"},{\"name\":\"repo\",\"type\":\"bytes32\"},{\"name\":\"data\",\"type\":\"bytes32[]\"}],\"name\":\"resetRepoData\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"host\",\"type\":\"bytes32\"},{\"name\":\"repo\",\"type\":\"bytes32\"},{\"name\":\"indexes\",\"type\":\"uint256[]\"},{\"name\":\"items\",\"type\":\"bytes32[]\"}],\"name\":\"updateRepoData\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"host\",\"type\":\"bytes32\"},{\"name\":\"data\",\"type\":\"bytes32[]\"}],\"name\":\"resetHostData\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"name\",\"type\":\"string\"}],\"name\":\"deleteHost\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"host\",\"type\":\"bytes32\"}],\"name\":\"retrieveOwner\",\"outputs\":[{\"name\":\"owner\",\"type\":\"address[]\"},{\"name\":\"data\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"data\",\"type\":\"bytes32[]\"}],\"name\":\"createHost\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"host\",\"type\":\"bytes32\"},{\"name\":\"repo\",\"type\":\"bytes32\"}],\"name\":\"retrieveAdmin\",\"outputs\":[{\"name\":\"admin\",\"type\":\"address[]\"},{\"name\":\"data\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"host\",\"type\":\"bytes32\"},{\"name\":\"repo\",\"type\":\"bytes32\"},{\"name\":\"holders\",\"type\":\"address[]\"}],\"name\":\"removeAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"hostName\",\"type\":\"string\"},{\"name\":\"repoName\",\"type\":\"string\"},{\"name\":\"data\",\"type\":\"bytes32[]\"}],\"name\":\"createRepo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"hostName\",\"type\":\"string\"},{\"name\":\"repoName\",\"type\":\"string\"}],\"name\":\"deleteRepo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"host\",\"type\":\"bytes32\"},{\"name\":\"repo\",\"type\":\"bytes32\"},{\"name\":\"holders\",\"type\":\"address[]\"}],\"name\":\"addWrite\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"host\",\"type\":\"bytes32\"},{\"name\":\"repo\",\"type\":\"bytes32\"},{\"name\":\"holders\",\"type\":\"address[]\"}],\"name\":\"removeWrite\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"host\",\"type\":\"bytes32\"},{\"name\":\"repo\",\"type\":\"bytes32\"},{\"name\":\"data\",\"type\":\"bytes32[]\"}],\"name\":\"resetDevData\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"host\",\"type\":\"bytes32\"},{\"name\":\"repo\",\"type\":\"bytes32\"},{\"name\":\"indexes\",\"type\":\"uint256[]\"},{\"name\":\"items\",\"type\":\"bytes32[]\"}],\"name\":\"updateDevData\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// HostRepoDevApi is an auto generated Go binding around an Ethereum contract.
type HostRepoDevApi struct {
	HostRepoDevApiCaller     // Read-only binding to the contract
	HostRepoDevApiTransactor // Write-only binding to the contract
	HostRepoDevApiFilterer   // Log filterer for contract events
}

// HostRepoDevApiCaller is an auto generated read-only Go binding around an Ethereum contract.
type HostRepoDevApiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HostRepoDevApiTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HostRepoDevApiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HostRepoDevApiFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HostRepoDevApiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HostRepoDevApiSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HostRepoDevApiSession struct {
	Contract     *HostRepoDevApi   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HostRepoDevApiCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HostRepoDevApiCallerSession struct {
	Contract *HostRepoDevApiCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// HostRepoDevApiTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HostRepoDevApiTransactorSession struct {
	Contract     *HostRepoDevApiTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// HostRepoDevApiRaw is an auto generated low-level Go binding around an Ethereum contract.
type HostRepoDevApiRaw struct {
	Contract *HostRepoDevApi // Generic contract binding to access the raw methods on
}

// HostRepoDevApiCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HostRepoDevApiCallerRaw struct {
	Contract *HostRepoDevApiCaller // Generic read-only contract binding to access the raw methods on
}

// HostRepoDevApiTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HostRepoDevApiTransactorRaw struct {
	Contract *HostRepoDevApiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHostRepoDevApi creates a new instance of HostRepoDevApi, bound to a specific deployed contract.
func NewHostRepoDevApi(address common.Address, backend bind.ContractBackend) (*HostRepoDevApi, error) {
	contract, err := bindHostRepoDevApi(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HostRepoDevApi{HostRepoDevApiCaller: HostRepoDevApiCaller{contract: contract}, HostRepoDevApiTransactor: HostRepoDevApiTransactor{contract: contract}, HostRepoDevApiFilterer: HostRepoDevApiFilterer{contract: contract}}, nil
}

// NewHostRepoDevApiCaller creates a new read-only instance of HostRepoDevApi, bound to a specific deployed contract.
func NewHostRepoDevApiCaller(address common.Address, caller bind.ContractCaller) (*HostRepoDevApiCaller, error) {
	contract, err := bindHostRepoDevApi(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HostRepoDevApiCaller{contract: contract}, nil
}

// NewHostRepoDevApiTransactor creates a new write-only instance of HostRepoDevApi, bound to a specific deployed contract.
func NewHostRepoDevApiTransactor(address common.Address, transactor bind.ContractTransactor) (*HostRepoDevApiTransactor, error) {
	contract, err := bindHostRepoDevApi(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HostRepoDevApiTransactor{contract: contract}, nil
}

// NewHostRepoDevApiFilterer creates a new log filterer instance of HostRepoDevApi, bound to a specific deployed contract.
func NewHostRepoDevApiFilterer(address common.Address, filterer bind.ContractFilterer) (*HostRepoDevApiFilterer, error) {
	contract, err := bindHostRepoDevApi(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HostRepoDevApiFilterer{contract: contract}, nil
}

// bindHostRepoDevApi binds a generic wrapper to an already deployed contract.
func bindHostRepoDevApi(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HostRepoDevApiABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HostRepoDevApi *HostRepoDevApiRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _HostRepoDevApi.Contract.HostRepoDevApiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HostRepoDevApi *HostRepoDevApiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HostRepoDevApi.Contract.HostRepoDevApiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HostRepoDevApi *HostRepoDevApiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HostRepoDevApi.Contract.HostRepoDevApiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HostRepoDevApi *HostRepoDevApiCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _HostRepoDevApi.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HostRepoDevApi *HostRepoDevApiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HostRepoDevApi.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HostRepoDevApi *HostRepoDevApiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HostRepoDevApi.Contract.contract.Transact(opts, method, params...)
}

// RetrieveAdmin is a free data retrieval call binding the contract method 0xe6d118a0.
//
// Solidity: function retrieveAdmin(host bytes32, repo bytes32) constant returns(admin address[], data bytes32[])
func (_HostRepoDevApi *HostRepoDevApiCaller) RetrieveAdmin(opts *bind.CallOpts, host [32]byte, repo [32]byte) (struct {
	Admin []common.Address
	Data  [][32]byte
}, error) {
	ret := new(struct {
		Admin []common.Address
		Data  [][32]byte
	})
	out := ret
	err := _HostRepoDevApi.contract.Call(opts, out, "retrieveAdmin", host, repo)
	return *ret, err
}

// RetrieveAdmin is a free data retrieval call binding the contract method 0xe6d118a0.
//
// Solidity: function retrieveAdmin(host bytes32, repo bytes32) constant returns(admin address[], data bytes32[])
func (_HostRepoDevApi *HostRepoDevApiSession) RetrieveAdmin(host [32]byte, repo [32]byte) (struct {
	Admin []common.Address
	Data  [][32]byte
}, error) {
	return _HostRepoDevApi.Contract.RetrieveAdmin(&_HostRepoDevApi.CallOpts, host, repo)
}

// RetrieveAdmin is a free data retrieval call binding the contract method 0xe6d118a0.
//
// Solidity: function retrieveAdmin(host bytes32, repo bytes32) constant returns(admin address[], data bytes32[])
func (_HostRepoDevApi *HostRepoDevApiCallerSession) RetrieveAdmin(host [32]byte, repo [32]byte) (struct {
	Admin []common.Address
	Data  [][32]byte
}, error) {
	return _HostRepoDevApi.Contract.RetrieveAdmin(&_HostRepoDevApi.CallOpts, host, repo)
}

// RetrieveOwner is a free data retrieval call binding the contract method 0xaca9fb46.
//
// Solidity: function retrieveOwner(host bytes32) constant returns(owner address[], data bytes32[])
func (_HostRepoDevApi *HostRepoDevApiCaller) RetrieveOwner(opts *bind.CallOpts, host [32]byte) (struct {
	Owner []common.Address
	Data  [][32]byte
}, error) {
	ret := new(struct {
		Owner []common.Address
		Data  [][32]byte
	})
	out := ret
	err := _HostRepoDevApi.contract.Call(opts, out, "retrieveOwner", host)
	return *ret, err
}

// RetrieveOwner is a free data retrieval call binding the contract method 0xaca9fb46.
//
// Solidity: function retrieveOwner(host bytes32) constant returns(owner address[], data bytes32[])
func (_HostRepoDevApi *HostRepoDevApiSession) RetrieveOwner(host [32]byte) (struct {
	Owner []common.Address
	Data  [][32]byte
}, error) {
	return _HostRepoDevApi.Contract.RetrieveOwner(&_HostRepoDevApi.CallOpts, host)
}

// RetrieveOwner is a free data retrieval call binding the contract method 0xaca9fb46.
//
// Solidity: function retrieveOwner(host bytes32) constant returns(owner address[], data bytes32[])
func (_HostRepoDevApi *HostRepoDevApiCallerSession) RetrieveOwner(host [32]byte) (struct {
	Owner []common.Address
	Data  [][32]byte
}, error) {
	return _HostRepoDevApi.Contract.RetrieveOwner(&_HostRepoDevApi.CallOpts, host)
}

// RetrieveWrite is a free data retrieval call binding the contract method 0x4000e5e7.
//
// Solidity: function retrieveWrite(host bytes32, repo bytes32) constant returns(admin address[], data bytes32[])
func (_HostRepoDevApi *HostRepoDevApiCaller) RetrieveWrite(opts *bind.CallOpts, host [32]byte, repo [32]byte) (struct {
	Admin []common.Address
	Data  [][32]byte
}, error) {
	ret := new(struct {
		Admin []common.Address
		Data  [][32]byte
	})
	out := ret
	err := _HostRepoDevApi.contract.Call(opts, out, "retrieveWrite", host, repo)
	return *ret, err
}

// RetrieveWrite is a free data retrieval call binding the contract method 0x4000e5e7.
//
// Solidity: function retrieveWrite(host bytes32, repo bytes32) constant returns(admin address[], data bytes32[])
func (_HostRepoDevApi *HostRepoDevApiSession) RetrieveWrite(host [32]byte, repo [32]byte) (struct {
	Admin []common.Address
	Data  [][32]byte
}, error) {
	return _HostRepoDevApi.Contract.RetrieveWrite(&_HostRepoDevApi.CallOpts, host, repo)
}

// RetrieveWrite is a free data retrieval call binding the contract method 0x4000e5e7.
//
// Solidity: function retrieveWrite(host bytes32, repo bytes32) constant returns(admin address[], data bytes32[])
func (_HostRepoDevApi *HostRepoDevApiCallerSession) RetrieveWrite(host [32]byte, repo [32]byte) (struct {
	Admin []common.Address
	Data  [][32]byte
}, error) {
	return _HostRepoDevApi.Contract.RetrieveWrite(&_HostRepoDevApi.CallOpts, host, repo)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x37c237fb.
//
// Solidity: function addAdmin(host bytes32, repo bytes32, holders address[]) returns()
func (_HostRepoDevApi *HostRepoDevApiTransactor) AddAdmin(opts *bind.TransactOpts, host [32]byte, repo [32]byte, holders []common.Address) (*types.Transaction, error) {
	return _HostRepoDevApi.contract.Transact(opts, "addAdmin", host, repo, holders)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x37c237fb.
//
// Solidity: function addAdmin(host bytes32, repo bytes32, holders address[]) returns()
func (_HostRepoDevApi *HostRepoDevApiSession) AddAdmin(host [32]byte, repo [32]byte, holders []common.Address) (*types.Transaction, error) {
	return _HostRepoDevApi.Contract.AddAdmin(&_HostRepoDevApi.TransactOpts, host, repo, holders)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x37c237fb.
//
// Solidity: function addAdmin(host bytes32, repo bytes32, holders address[]) returns()
func (_HostRepoDevApi *HostRepoDevApiTransactorSession) AddAdmin(host [32]byte, repo [32]byte, holders []common.Address) (*types.Transaction, error) {
	return _HostRepoDevApi.Contract.AddAdmin(&_HostRepoDevApi.TransactOpts, host, repo, holders)
}

// AddOwner is a paid mutator transaction binding the contract method 0x5c35d072.
//
// Solidity: function addOwner(host bytes32, holders address[]) returns()
func (_HostRepoDevApi *HostRepoDevApiTransactor) AddOwner(opts *bind.TransactOpts, host [32]byte, holders []common.Address) (*types.Transaction, error) {
	return _HostRepoDevApi.contract.Transact(opts, "addOwner", host, holders)
}

// AddOwner is a paid mutator transaction binding the contract method 0x5c35d072.
//
// Solidity: function addOwner(host bytes32, holders address[]) returns()
func (_HostRepoDevApi *HostRepoDevApiSession) AddOwner(host [32]byte, holders []common.Address) (*types.Transaction, error) {
	return _HostRepoDevApi.Contract.AddOwner(&_HostRepoDevApi.TransactOpts, host, holders)
}

// AddOwner is a paid mutator transaction binding the contract method 0x5c35d072.
//
// Solidity: function addOwner(host bytes32, holders address[]) returns()
func (_HostRepoDevApi *HostRepoDevApiTransactorSession) AddOwner(host [32]byte, holders []common.Address) (*types.Transaction, error) {
	return _HostRepoDevApi.Contract.AddOwner(&_HostRepoDevApi.TransactOpts, host, holders)
}

// AddWrite is a paid mutator transaction binding the contract method 0xd219d6f8.
//
// Solidity: function addWrite(host bytes32, repo bytes32, holders address[]) returns()
func (_HostRepoDevApi *HostRepoDevApiTransactor) AddWrite(opts *bind.TransactOpts, host [32]byte, repo [32]byte, holders []common.Address) (*types.Transaction, error) {
	return _HostRepoDevApi.contract.Transact(opts, "addWrite", host, repo, holders)
}

// AddWrite is a paid mutator transaction binding the contract method 0xd219d6f8.
//
// Solidity: function addWrite(host bytes32, repo bytes32, holders address[]) returns()
func (_HostRepoDevApi *HostRepoDevApiSession) AddWrite(host [32]byte, repo [32]byte, holders []common.Address) (*types.Transaction, error) {
	return _HostRepoDevApi.Contract.AddWrite(&_HostRepoDevApi.TransactOpts, host, repo, holders)
}

// AddWrite is a paid mutator transaction binding the contract method 0xd219d6f8.
//
// Solidity: function addWrite(host bytes32, repo bytes32, holders address[]) returns()
func (_HostRepoDevApi *HostRepoDevApiTransactorSession) AddWrite(host [32]byte, repo [32]byte, holders []common.Address) (*types.Transaction, error) {
	return _HostRepoDevApi.Contract.AddWrite(&_HostRepoDevApi.TransactOpts, host, repo, holders)
}

// CreateHost is a paid mutator transaction binding the contract method 0xc0247349.
//
// Solidity: function createHost(name string, data bytes32[]) returns()
func (_HostRepoDevApi *HostRepoDevApiTransactor) CreateHost(opts *bind.TransactOpts, name string, data [][32]byte) (*types.Transaction, error) {
	return _HostRepoDevApi.contract.Transact(opts, "createHost", name, data)
}

// CreateHost is a paid mutator transaction binding the contract method 0xc0247349.
//
// Solidity: function createHost(name string, data bytes32[]) returns()
func (_HostRepoDevApi *HostRepoDevApiSession) CreateHost(name string, data [][32]byte) (*types.Transaction, error) {
	return _HostRepoDevApi.Contract.CreateHost(&_HostRepoDevApi.TransactOpts, name, data)
}

// CreateHost is a paid mutator transaction binding the contract method 0xc0247349.
//
// Solidity: function createHost(name string, data bytes32[]) returns()
func (_HostRepoDevApi *HostRepoDevApiTransactorSession) CreateHost(name string, data [][32]byte) (*types.Transaction, error) {
	return _HostRepoDevApi.Contract.CreateHost(&_HostRepoDevApi.TransactOpts, name, data)
}

// CreateRepo is a paid mutator transaction binding the contract method 0x16cea0ea.
//
// Solidity: function createRepo(hostName string, repoName string, data bytes32[]) returns()
func (_HostRepoDevApi *HostRepoDevApiTransactor) CreateRepo(opts *bind.TransactOpts, hostName string, repoName string, data [][32]byte) (*types.Transaction, error) {
	return _HostRepoDevApi.contract.Transact(opts, "createRepo", hostName, repoName, data)
}

// CreateRepo is a paid mutator transaction binding the contract method 0x16cea0ea.
//
// Solidity: function createRepo(hostName string, repoName string, data bytes32[]) returns()
func (_HostRepoDevApi *HostRepoDevApiSession) CreateRepo(hostName string, repoName string, data [][32]byte) (*types.Transaction, error) {
	return _HostRepoDevApi.Contract.CreateRepo(&_HostRepoDevApi.TransactOpts, hostName, repoName, data)
}

// CreateRepo is a paid mutator transaction binding the contract method 0x16cea0ea.
//
// Solidity: function createRepo(hostName string, repoName string, data bytes32[]) returns()
func (_HostRepoDevApi *HostRepoDevApiTransactorSession) CreateRepo(hostName string, repoName string, data [][32]byte) (*types.Transaction, error) {
	return _HostRepoDevApi.Contract.CreateRepo(&_HostRepoDevApi.TransactOpts, hostName, repoName, data)
}

// DeleteHost is a paid mutator transaction binding the contract method 0xa60864a4.
//
// Solidity: function deleteHost(name string) returns()
func (_HostRepoDevApi *HostRepoDevApiTransactor) DeleteHost(opts *bind.TransactOpts, name string) (*types.Transaction, error) {
	return _HostRepoDevApi.contract.Transact(opts, "deleteHost", name)
}

// DeleteHost is a paid mutator transaction binding the contract method 0xa60864a4.
//
// Solidity: function deleteHost(name string) returns()
func (_HostRepoDevApi *HostRepoDevApiSession) DeleteHost(name string) (*types.Transaction, error) {
	return _HostRepoDevApi.Contract.DeleteHost(&_HostRepoDevApi.TransactOpts, name)
}

// DeleteHost is a paid mutator transaction binding the contract method 0xa60864a4.
//
// Solidity: function deleteHost(name string) returns()
func (_HostRepoDevApi *HostRepoDevApiTransactorSession) DeleteHost(name string) (*types.Transaction, error) {
	return _HostRepoDevApi.Contract.DeleteHost(&_HostRepoDevApi.TransactOpts, name)
}

// DeleteRepo is a paid mutator transaction binding the contract method 0x79289b21.
//
// Solidity: function deleteRepo(hostName string, repoName string) returns()
func (_HostRepoDevApi *HostRepoDevApiTransactor) DeleteRepo(opts *bind.TransactOpts, hostName string, repoName string) (*types.Transaction, error) {
	return _HostRepoDevApi.contract.Transact(opts, "deleteRepo", hostName, repoName)
}

// DeleteRepo is a paid mutator transaction binding the contract method 0x79289b21.
//
// Solidity: function deleteRepo(hostName string, repoName string) returns()
func (_HostRepoDevApi *HostRepoDevApiSession) DeleteRepo(hostName string, repoName string) (*types.Transaction, error) {
	return _HostRepoDevApi.Contract.DeleteRepo(&_HostRepoDevApi.TransactOpts, hostName, repoName)
}

// DeleteRepo is a paid mutator transaction binding the contract method 0x79289b21.
//
// Solidity: function deleteRepo(hostName string, repoName string) returns()
func (_HostRepoDevApi *HostRepoDevApiTransactorSession) DeleteRepo(hostName string, repoName string) (*types.Transaction, error) {
	return _HostRepoDevApi.Contract.DeleteRepo(&_HostRepoDevApi.TransactOpts, hostName, repoName)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0xf6f66166.
//
// Solidity: function removeAdmin(host bytes32, repo bytes32, holders address[]) returns()
func (_HostRepoDevApi *HostRepoDevApiTransactor) RemoveAdmin(opts *bind.TransactOpts, host [32]byte, repo [32]byte, holders []common.Address) (*types.Transaction, error) {
	return _HostRepoDevApi.contract.Transact(opts, "removeAdmin", host, repo, holders)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0xf6f66166.
//
// Solidity: function removeAdmin(host bytes32, repo bytes32, holders address[]) returns()
func (_HostRepoDevApi *HostRepoDevApiSession) RemoveAdmin(host [32]byte, repo [32]byte, holders []common.Address) (*types.Transaction, error) {
	return _HostRepoDevApi.Contract.RemoveAdmin(&_HostRepoDevApi.TransactOpts, host, repo, holders)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0xf6f66166.
//
// Solidity: function removeAdmin(host bytes32, repo bytes32, holders address[]) returns()
func (_HostRepoDevApi *HostRepoDevApiTransactorSession) RemoveAdmin(host [32]byte, repo [32]byte, holders []common.Address) (*types.Transaction, error) {
	return _HostRepoDevApi.Contract.RemoveAdmin(&_HostRepoDevApi.TransactOpts, host, repo, holders)
}

// RemoveOwner is a paid mutator transaction binding the contract method 0x39cc75aa.
//
// Solidity: function removeOwner(host bytes32, holders address[]) returns()
func (_HostRepoDevApi *HostRepoDevApiTransactor) RemoveOwner(opts *bind.TransactOpts, host [32]byte, holders []common.Address) (*types.Transaction, error) {
	return _HostRepoDevApi.contract.Transact(opts, "removeOwner", host, holders)
}

// RemoveOwner is a paid mutator transaction binding the contract method 0x39cc75aa.
//
// Solidity: function removeOwner(host bytes32, holders address[]) returns()
func (_HostRepoDevApi *HostRepoDevApiSession) RemoveOwner(host [32]byte, holders []common.Address) (*types.Transaction, error) {
	return _HostRepoDevApi.Contract.RemoveOwner(&_HostRepoDevApi.TransactOpts, host, holders)
}

// RemoveOwner is a paid mutator transaction binding the contract method 0x39cc75aa.
//
// Solidity: function removeOwner(host bytes32, holders address[]) returns()
func (_HostRepoDevApi *HostRepoDevApiTransactorSession) RemoveOwner(host [32]byte, holders []common.Address) (*types.Transaction, error) {
	return _HostRepoDevApi.Contract.RemoveOwner(&_HostRepoDevApi.TransactOpts, host, holders)
}

// RemoveWrite is a paid mutator transaction binding the contract method 0xb92f915a.
//
// Solidity: function removeWrite(host bytes32, repo bytes32, holders address[]) returns()
func (_HostRepoDevApi *HostRepoDevApiTransactor) RemoveWrite(opts *bind.TransactOpts, host [32]byte, repo [32]byte, holders []common.Address) (*types.Transaction, error) {
	return _HostRepoDevApi.contract.Transact(opts, "removeWrite", host, repo, holders)
}

// RemoveWrite is a paid mutator transaction binding the contract method 0xb92f915a.
//
// Solidity: function removeWrite(host bytes32, repo bytes32, holders address[]) returns()
func (_HostRepoDevApi *HostRepoDevApiSession) RemoveWrite(host [32]byte, repo [32]byte, holders []common.Address) (*types.Transaction, error) {
	return _HostRepoDevApi.Contract.RemoveWrite(&_HostRepoDevApi.TransactOpts, host, repo, holders)
}

// RemoveWrite is a paid mutator transaction binding the contract method 0xb92f915a.
//
// Solidity: function removeWrite(host bytes32, repo bytes32, holders address[]) returns()
func (_HostRepoDevApi *HostRepoDevApiTransactorSession) RemoveWrite(host [32]byte, repo [32]byte, holders []common.Address) (*types.Transaction, error) {
	return _HostRepoDevApi.Contract.RemoveWrite(&_HostRepoDevApi.TransactOpts, host, repo, holders)
}

// ResetDevData is a paid mutator transaction binding the contract method 0x96cff04e.
//
// Solidity: function resetDevData(host bytes32, repo bytes32, data bytes32[]) returns()
func (_HostRepoDevApi *HostRepoDevApiTransactor) ResetDevData(opts *bind.TransactOpts, host [32]byte, repo [32]byte, data [][32]byte) (*types.Transaction, error) {
	return _HostRepoDevApi.contract.Transact(opts, "resetDevData", host, repo, data)
}

// ResetDevData is a paid mutator transaction binding the contract method 0x96cff04e.
//
// Solidity: function resetDevData(host bytes32, repo bytes32, data bytes32[]) returns()
func (_HostRepoDevApi *HostRepoDevApiSession) ResetDevData(host [32]byte, repo [32]byte, data [][32]byte) (*types.Transaction, error) {
	return _HostRepoDevApi.Contract.ResetDevData(&_HostRepoDevApi.TransactOpts, host, repo, data)
}

// ResetDevData is a paid mutator transaction binding the contract method 0x96cff04e.
//
// Solidity: function resetDevData(host bytes32, repo bytes32, data bytes32[]) returns()
func (_HostRepoDevApi *HostRepoDevApiTransactorSession) ResetDevData(host [32]byte, repo [32]byte, data [][32]byte) (*types.Transaction, error) {
	return _HostRepoDevApi.Contract.ResetDevData(&_HostRepoDevApi.TransactOpts, host, repo, data)
}

// ResetHostData is a paid mutator transaction binding the contract method 0xa51f9275.
//
// Solidity: function resetHostData(host bytes32, data bytes32[]) returns()
func (_HostRepoDevApi *HostRepoDevApiTransactor) ResetHostData(opts *bind.TransactOpts, host [32]byte, data [][32]byte) (*types.Transaction, error) {
	return _HostRepoDevApi.contract.Transact(opts, "resetHostData", host, data)
}

// ResetHostData is a paid mutator transaction binding the contract method 0xa51f9275.
//
// Solidity: function resetHostData(host bytes32, data bytes32[]) returns()
func (_HostRepoDevApi *HostRepoDevApiSession) ResetHostData(host [32]byte, data [][32]byte) (*types.Transaction, error) {
	return _HostRepoDevApi.Contract.ResetHostData(&_HostRepoDevApi.TransactOpts, host, data)
}

// ResetHostData is a paid mutator transaction binding the contract method 0xa51f9275.
//
// Solidity: function resetHostData(host bytes32, data bytes32[]) returns()
func (_HostRepoDevApi *HostRepoDevApiTransactorSession) ResetHostData(host [32]byte, data [][32]byte) (*types.Transaction, error) {
	return _HostRepoDevApi.Contract.ResetHostData(&_HostRepoDevApi.TransactOpts, host, data)
}

// ResetRepoData is a paid mutator transaction binding the contract method 0x5f260715.
//
// Solidity: function resetRepoData(host bytes32, repo bytes32, data bytes32[]) returns()
func (_HostRepoDevApi *HostRepoDevApiTransactor) ResetRepoData(opts *bind.TransactOpts, host [32]byte, repo [32]byte, data [][32]byte) (*types.Transaction, error) {
	return _HostRepoDevApi.contract.Transact(opts, "resetRepoData", host, repo, data)
}

// ResetRepoData is a paid mutator transaction binding the contract method 0x5f260715.
//
// Solidity: function resetRepoData(host bytes32, repo bytes32, data bytes32[]) returns()
func (_HostRepoDevApi *HostRepoDevApiSession) ResetRepoData(host [32]byte, repo [32]byte, data [][32]byte) (*types.Transaction, error) {
	return _HostRepoDevApi.Contract.ResetRepoData(&_HostRepoDevApi.TransactOpts, host, repo, data)
}

// ResetRepoData is a paid mutator transaction binding the contract method 0x5f260715.
//
// Solidity: function resetRepoData(host bytes32, repo bytes32, data bytes32[]) returns()
func (_HostRepoDevApi *HostRepoDevApiTransactorSession) ResetRepoData(host [32]byte, repo [32]byte, data [][32]byte) (*types.Transaction, error) {
	return _HostRepoDevApi.Contract.ResetRepoData(&_HostRepoDevApi.TransactOpts, host, repo, data)
}

// UpdateDevData is a paid mutator transaction binding the contract method 0xee811423.
//
// Solidity: function updateDevData(host bytes32, repo bytes32, indexes uint256[], items bytes32[]) returns()
func (_HostRepoDevApi *HostRepoDevApiTransactor) UpdateDevData(opts *bind.TransactOpts, host [32]byte, repo [32]byte, indexes []*big.Int, items [][32]byte) (*types.Transaction, error) {
	return _HostRepoDevApi.contract.Transact(opts, "updateDevData", host, repo, indexes, items)
}

// UpdateDevData is a paid mutator transaction binding the contract method 0xee811423.
//
// Solidity: function updateDevData(host bytes32, repo bytes32, indexes uint256[], items bytes32[]) returns()
func (_HostRepoDevApi *HostRepoDevApiSession) UpdateDevData(host [32]byte, repo [32]byte, indexes []*big.Int, items [][32]byte) (*types.Transaction, error) {
	return _HostRepoDevApi.Contract.UpdateDevData(&_HostRepoDevApi.TransactOpts, host, repo, indexes, items)
}

// UpdateDevData is a paid mutator transaction binding the contract method 0xee811423.
//
// Solidity: function updateDevData(host bytes32, repo bytes32, indexes uint256[], items bytes32[]) returns()
func (_HostRepoDevApi *HostRepoDevApiTransactorSession) UpdateDevData(host [32]byte, repo [32]byte, indexes []*big.Int, items [][32]byte) (*types.Transaction, error) {
	return _HostRepoDevApi.Contract.UpdateDevData(&_HostRepoDevApi.TransactOpts, host, repo, indexes, items)
}

// UpdateHostData is a paid mutator transaction binding the contract method 0x202f58d9.
//
// Solidity: function updateHostData(host bytes32, indexes uint256[], items bytes32[]) returns()
func (_HostRepoDevApi *HostRepoDevApiTransactor) UpdateHostData(opts *bind.TransactOpts, host [32]byte, indexes []*big.Int, items [][32]byte) (*types.Transaction, error) {
	return _HostRepoDevApi.contract.Transact(opts, "updateHostData", host, indexes, items)
}

// UpdateHostData is a paid mutator transaction binding the contract method 0x202f58d9.
//
// Solidity: function updateHostData(host bytes32, indexes uint256[], items bytes32[]) returns()
func (_HostRepoDevApi *HostRepoDevApiSession) UpdateHostData(host [32]byte, indexes []*big.Int, items [][32]byte) (*types.Transaction, error) {
	return _HostRepoDevApi.Contract.UpdateHostData(&_HostRepoDevApi.TransactOpts, host, indexes, items)
}

// UpdateHostData is a paid mutator transaction binding the contract method 0x202f58d9.
//
// Solidity: function updateHostData(host bytes32, indexes uint256[], items bytes32[]) returns()
func (_HostRepoDevApi *HostRepoDevApiTransactorSession) UpdateHostData(host [32]byte, indexes []*big.Int, items [][32]byte) (*types.Transaction, error) {
	return _HostRepoDevApi.Contract.UpdateHostData(&_HostRepoDevApi.TransactOpts, host, indexes, items)
}

// UpdateRepoData is a paid mutator transaction binding the contract method 0x88c5d860.
//
// Solidity: function updateRepoData(host bytes32, repo bytes32, indexes uint256[], items bytes32[]) returns()
func (_HostRepoDevApi *HostRepoDevApiTransactor) UpdateRepoData(opts *bind.TransactOpts, host [32]byte, repo [32]byte, indexes []*big.Int, items [][32]byte) (*types.Transaction, error) {
	return _HostRepoDevApi.contract.Transact(opts, "updateRepoData", host, repo, indexes, items)
}

// UpdateRepoData is a paid mutator transaction binding the contract method 0x88c5d860.
//
// Solidity: function updateRepoData(host bytes32, repo bytes32, indexes uint256[], items bytes32[]) returns()
func (_HostRepoDevApi *HostRepoDevApiSession) UpdateRepoData(host [32]byte, repo [32]byte, indexes []*big.Int, items [][32]byte) (*types.Transaction, error) {
	return _HostRepoDevApi.Contract.UpdateRepoData(&_HostRepoDevApi.TransactOpts, host, repo, indexes, items)
}

// UpdateRepoData is a paid mutator transaction binding the contract method 0x88c5d860.
//
// Solidity: function updateRepoData(host bytes32, repo bytes32, indexes uint256[], items bytes32[]) returns()
func (_HostRepoDevApi *HostRepoDevApiTransactorSession) UpdateRepoData(host [32]byte, repo [32]byte, indexes []*big.Int, items [][32]byte) (*types.Transaction, error) {
	return _HostRepoDevApi.Contract.UpdateRepoData(&_HostRepoDevApi.TransactOpts, host, repo, indexes, items)
}
