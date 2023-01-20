// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package platform

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

// PlatformContractPlan is an auto generated low-level Go binding around an user-defined struct.
type PlatformContractPlan struct {
	Price    *big.Int
	Duration *big.Int
}

// PlatformMetaData contains all meta data concerning the Platform contract.
var PlatformMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_duration\",\"type\":\"uint256\"}],\"name\":\"addPlan\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_planId\",\"type\":\"uint8\"}],\"name\":\"getPlan\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"internalType\":\"structPlatformContract.Plan\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPlans\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"internalType\":\"structPlatformContract.Plan[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getUserSubDated\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_id\",\"type\":\"uint8\"}],\"name\":\"removePlan\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_planId\",\"type\":\"uint8\"}],\"name\":\"subscribe\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"userStatus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// PlatformABI is the input ABI used to generate the binding from.
// Deprecated: Use PlatformMetaData.ABI instead.
var PlatformABI = PlatformMetaData.ABI

// Platform is an auto generated Go binding around an Ethereum contract.
type Platform struct {
	PlatformCaller     // Read-only binding to the contract
	PlatformTransactor // Write-only binding to the contract
	PlatformFilterer   // Log filterer for contract events
}

// PlatformCaller is an auto generated read-only Go binding around an Ethereum contract.
type PlatformCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PlatformTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PlatformTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PlatformFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PlatformFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PlatformSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PlatformSession struct {
	Contract     *Platform         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PlatformCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PlatformCallerSession struct {
	Contract *PlatformCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// PlatformTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PlatformTransactorSession struct {
	Contract     *PlatformTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// PlatformRaw is an auto generated low-level Go binding around an Ethereum contract.
type PlatformRaw struct {
	Contract *Platform // Generic contract binding to access the raw methods on
}

// PlatformCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PlatformCallerRaw struct {
	Contract *PlatformCaller // Generic read-only contract binding to access the raw methods on
}

// PlatformTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PlatformTransactorRaw struct {
	Contract *PlatformTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPlatform creates a new instance of Platform, bound to a specific deployed contract.
func NewPlatform(address common.Address, backend bind.ContractBackend) (*Platform, error) {
	contract, err := bindPlatform(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Platform{PlatformCaller: PlatformCaller{contract: contract}, PlatformTransactor: PlatformTransactor{contract: contract}, PlatformFilterer: PlatformFilterer{contract: contract}}, nil
}

// NewPlatformCaller creates a new read-only instance of Platform, bound to a specific deployed contract.
func NewPlatformCaller(address common.Address, caller bind.ContractCaller) (*PlatformCaller, error) {
	contract, err := bindPlatform(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PlatformCaller{contract: contract}, nil
}

// NewPlatformTransactor creates a new write-only instance of Platform, bound to a specific deployed contract.
func NewPlatformTransactor(address common.Address, transactor bind.ContractTransactor) (*PlatformTransactor, error) {
	contract, err := bindPlatform(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PlatformTransactor{contract: contract}, nil
}

// NewPlatformFilterer creates a new log filterer instance of Platform, bound to a specific deployed contract.
func NewPlatformFilterer(address common.Address, filterer bind.ContractFilterer) (*PlatformFilterer, error) {
	contract, err := bindPlatform(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PlatformFilterer{contract: contract}, nil
}

// bindPlatform binds a generic wrapper to an already deployed contract.
func bindPlatform(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PlatformABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Platform *PlatformRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Platform.Contract.PlatformCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Platform *PlatformRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Platform.Contract.PlatformTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Platform *PlatformRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Platform.Contract.PlatformTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Platform *PlatformCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Platform.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Platform *PlatformTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Platform.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Platform *PlatformTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Platform.Contract.contract.Transact(opts, method, params...)
}

// GetPlan is a free data retrieval call binding the contract method 0xa12b2121.
//
// Solidity: function getPlan(uint8 _planId) view returns((uint256,uint256))
func (_Platform *PlatformCaller) GetPlan(opts *bind.CallOpts, _planId uint8) (PlatformContractPlan, error) {
	var out []interface{}
	err := _Platform.contract.Call(opts, &out, "getPlan", _planId)

	if err != nil {
		return *new(PlatformContractPlan), err
	}

	out0 := *abi.ConvertType(out[0], new(PlatformContractPlan)).(*PlatformContractPlan)

	return out0, err

}

// GetPlan is a free data retrieval call binding the contract method 0xa12b2121.
//
// Solidity: function getPlan(uint8 _planId) view returns((uint256,uint256))
func (_Platform *PlatformSession) GetPlan(_planId uint8) (PlatformContractPlan, error) {
	return _Platform.Contract.GetPlan(&_Platform.CallOpts, _planId)
}

// GetPlan is a free data retrieval call binding the contract method 0xa12b2121.
//
// Solidity: function getPlan(uint8 _planId) view returns((uint256,uint256))
func (_Platform *PlatformCallerSession) GetPlan(_planId uint8) (PlatformContractPlan, error) {
	return _Platform.Contract.GetPlan(&_Platform.CallOpts, _planId)
}

// GetPlans is a free data retrieval call binding the contract method 0xd94a862b.
//
// Solidity: function getPlans() view returns((uint256,uint256)[])
func (_Platform *PlatformCaller) GetPlans(opts *bind.CallOpts) ([]PlatformContractPlan, error) {
	var out []interface{}
	err := _Platform.contract.Call(opts, &out, "getPlans")

	if err != nil {
		return *new([]PlatformContractPlan), err
	}

	out0 := *abi.ConvertType(out[0], new([]PlatformContractPlan)).(*[]PlatformContractPlan)

	return out0, err

}

// GetPlans is a free data retrieval call binding the contract method 0xd94a862b.
//
// Solidity: function getPlans() view returns((uint256,uint256)[])
func (_Platform *PlatformSession) GetPlans() ([]PlatformContractPlan, error) {
	return _Platform.Contract.GetPlans(&_Platform.CallOpts)
}

// GetPlans is a free data retrieval call binding the contract method 0xd94a862b.
//
// Solidity: function getPlans() view returns((uint256,uint256)[])
func (_Platform *PlatformCallerSession) GetPlans() ([]PlatformContractPlan, error) {
	return _Platform.Contract.GetPlans(&_Platform.CallOpts)
}

// GetUserSubDated is a free data retrieval call binding the contract method 0x244a1985.
//
// Solidity: function getUserSubDated(address _user) view returns(uint256)
func (_Platform *PlatformCaller) GetUserSubDated(opts *bind.CallOpts, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Platform.contract.Call(opts, &out, "getUserSubDated", _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserSubDated is a free data retrieval call binding the contract method 0x244a1985.
//
// Solidity: function getUserSubDated(address _user) view returns(uint256)
func (_Platform *PlatformSession) GetUserSubDated(_user common.Address) (*big.Int, error) {
	return _Platform.Contract.GetUserSubDated(&_Platform.CallOpts, _user)
}

// GetUserSubDated is a free data retrieval call binding the contract method 0x244a1985.
//
// Solidity: function getUserSubDated(address _user) view returns(uint256)
func (_Platform *PlatformCallerSession) GetUserSubDated(_user common.Address) (*big.Int, error) {
	return _Platform.Contract.GetUserSubDated(&_Platform.CallOpts, _user)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Platform *PlatformCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Platform.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Platform *PlatformSession) Owner() (common.Address, error) {
	return _Platform.Contract.Owner(&_Platform.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Platform *PlatformCallerSession) Owner() (common.Address, error) {
	return _Platform.Contract.Owner(&_Platform.CallOpts)
}

// UserStatus is a free data retrieval call binding the contract method 0x225d29a1.
//
// Solidity: function userStatus(address _user) view returns(bool status)
func (_Platform *PlatformCaller) UserStatus(opts *bind.CallOpts, _user common.Address) (bool, error) {
	var out []interface{}
	err := _Platform.contract.Call(opts, &out, "userStatus", _user)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UserStatus is a free data retrieval call binding the contract method 0x225d29a1.
//
// Solidity: function userStatus(address _user) view returns(bool status)
func (_Platform *PlatformSession) UserStatus(_user common.Address) (bool, error) {
	return _Platform.Contract.UserStatus(&_Platform.CallOpts, _user)
}

// UserStatus is a free data retrieval call binding the contract method 0x225d29a1.
//
// Solidity: function userStatus(address _user) view returns(bool status)
func (_Platform *PlatformCallerSession) UserStatus(_user common.Address) (bool, error) {
	return _Platform.Contract.UserStatus(&_Platform.CallOpts, _user)
}

// AddPlan is a paid mutator transaction binding the contract method 0x6bd872ce.
//
// Solidity: function addPlan(uint256 _price, uint256 _duration) returns()
func (_Platform *PlatformTransactor) AddPlan(opts *bind.TransactOpts, _price *big.Int, _duration *big.Int) (*types.Transaction, error) {
	return _Platform.contract.Transact(opts, "addPlan", _price, _duration)
}

// AddPlan is a paid mutator transaction binding the contract method 0x6bd872ce.
//
// Solidity: function addPlan(uint256 _price, uint256 _duration) returns()
func (_Platform *PlatformSession) AddPlan(_price *big.Int, _duration *big.Int) (*types.Transaction, error) {
	return _Platform.Contract.AddPlan(&_Platform.TransactOpts, _price, _duration)
}

// AddPlan is a paid mutator transaction binding the contract method 0x6bd872ce.
//
// Solidity: function addPlan(uint256 _price, uint256 _duration) returns()
func (_Platform *PlatformTransactorSession) AddPlan(_price *big.Int, _duration *big.Int) (*types.Transaction, error) {
	return _Platform.Contract.AddPlan(&_Platform.TransactOpts, _price, _duration)
}

// RemovePlan is a paid mutator transaction binding the contract method 0xe63c7e0d.
//
// Solidity: function removePlan(uint8 _id) returns()
func (_Platform *PlatformTransactor) RemovePlan(opts *bind.TransactOpts, _id uint8) (*types.Transaction, error) {
	return _Platform.contract.Transact(opts, "removePlan", _id)
}

// RemovePlan is a paid mutator transaction binding the contract method 0xe63c7e0d.
//
// Solidity: function removePlan(uint8 _id) returns()
func (_Platform *PlatformSession) RemovePlan(_id uint8) (*types.Transaction, error) {
	return _Platform.Contract.RemovePlan(&_Platform.TransactOpts, _id)
}

// RemovePlan is a paid mutator transaction binding the contract method 0xe63c7e0d.
//
// Solidity: function removePlan(uint8 _id) returns()
func (_Platform *PlatformTransactorSession) RemovePlan(_id uint8) (*types.Transaction, error) {
	return _Platform.Contract.RemovePlan(&_Platform.TransactOpts, _id)
}

// Subscribe is a paid mutator transaction binding the contract method 0x49c7e639.
//
// Solidity: function subscribe(uint8 _planId) payable returns()
func (_Platform *PlatformTransactor) Subscribe(opts *bind.TransactOpts, _planId uint8) (*types.Transaction, error) {
	return _Platform.contract.Transact(opts, "subscribe", _planId)
}

// Subscribe is a paid mutator transaction binding the contract method 0x49c7e639.
//
// Solidity: function subscribe(uint8 _planId) payable returns()
func (_Platform *PlatformSession) Subscribe(_planId uint8) (*types.Transaction, error) {
	return _Platform.Contract.Subscribe(&_Platform.TransactOpts, _planId)
}

// Subscribe is a paid mutator transaction binding the contract method 0x49c7e639.
//
// Solidity: function subscribe(uint8 _planId) payable returns()
func (_Platform *PlatformTransactorSession) Subscribe(_planId uint8) (*types.Transaction, error) {
	return _Platform.Contract.Subscribe(&_Platform.TransactOpts, _planId)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Platform *PlatformTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Platform.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Platform *PlatformSession) Receive() (*types.Transaction, error) {
	return _Platform.Contract.Receive(&_Platform.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Platform *PlatformTransactorSession) Receive() (*types.Transaction, error) {
	return _Platform.Contract.Receive(&_Platform.TransactOpts)
}
