// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package parent

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

// ParentContractPlatform is an auto generated low-level Go binding around an user-defined struct.
type ParentContractPlatform struct {
	Id       *big.Int
	Platform common.Address
	Details  []byte
}

// ParentMetaData contains all meta data concerning the Parent contract.
var ParentMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_details\",\"type\":\"bytes\"}],\"name\":\"createPlatform\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"getPlatform\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"platform\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"details\",\"type\":\"bytes\"}],\"internalType\":\"structParentContract.Platform\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPlatforms\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"platform\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"details\",\"type\":\"bytes\"}],\"internalType\":\"structParentContract.Platform[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUserLatestPlatform\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"platform\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"details\",\"type\":\"bytes\"}],\"internalType\":\"structParentContract.Platform\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUserPlatforms\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"platform\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"details\",\"type\":\"bytes\"}],\"internalType\":\"structParentContract.Platform[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ParentABI is the input ABI used to generate the binding from.
// Deprecated: Use ParentMetaData.ABI instead.
var ParentABI = ParentMetaData.ABI

// Parent is an auto generated Go binding around an Ethereum contract.
type Parent struct {
	ParentCaller     // Read-only binding to the contract
	ParentTransactor // Write-only binding to the contract
	ParentFilterer   // Log filterer for contract events
}

// ParentCaller is an auto generated read-only Go binding around an Ethereum contract.
type ParentCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ParentTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ParentTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ParentFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ParentFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ParentSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ParentSession struct {
	Contract     *Parent           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ParentCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ParentCallerSession struct {
	Contract *ParentCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ParentTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ParentTransactorSession struct {
	Contract     *ParentTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ParentRaw is an auto generated low-level Go binding around an Ethereum contract.
type ParentRaw struct {
	Contract *Parent // Generic contract binding to access the raw methods on
}

// ParentCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ParentCallerRaw struct {
	Contract *ParentCaller // Generic read-only contract binding to access the raw methods on
}

// ParentTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ParentTransactorRaw struct {
	Contract *ParentTransactor // Generic write-only contract binding to access the raw methods on
}

// NewParent creates a new instance of Parent, bound to a specific deployed contract.
func NewParent(address common.Address, backend bind.ContractBackend) (*Parent, error) {
	contract, err := bindParent(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Parent{ParentCaller: ParentCaller{contract: contract}, ParentTransactor: ParentTransactor{contract: contract}, ParentFilterer: ParentFilterer{contract: contract}}, nil
}

// NewParentCaller creates a new read-only instance of Parent, bound to a specific deployed contract.
func NewParentCaller(address common.Address, caller bind.ContractCaller) (*ParentCaller, error) {
	contract, err := bindParent(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ParentCaller{contract: contract}, nil
}

// NewParentTransactor creates a new write-only instance of Parent, bound to a specific deployed contract.
func NewParentTransactor(address common.Address, transactor bind.ContractTransactor) (*ParentTransactor, error) {
	contract, err := bindParent(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ParentTransactor{contract: contract}, nil
}

// NewParentFilterer creates a new log filterer instance of Parent, bound to a specific deployed contract.
func NewParentFilterer(address common.Address, filterer bind.ContractFilterer) (*ParentFilterer, error) {
	contract, err := bindParent(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ParentFilterer{contract: contract}, nil
}

// bindParent binds a generic wrapper to an already deployed contract.
func bindParent(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ParentABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Parent *ParentRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Parent.Contract.ParentCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Parent *ParentRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Parent.Contract.ParentTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Parent *ParentRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Parent.Contract.ParentTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Parent *ParentCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Parent.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Parent *ParentTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Parent.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Parent *ParentTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Parent.Contract.contract.Transact(opts, method, params...)
}

// GetPlatform is a free data retrieval call binding the contract method 0x5bd22c98.
//
// Solidity: function getPlatform(uint256 _id) view returns((uint256,address,bytes))
func (_Parent *ParentCaller) GetPlatform(opts *bind.CallOpts, _id *big.Int) (ParentContractPlatform, error) {
	var out []interface{}
	err := _Parent.contract.Call(opts, &out, "getPlatform", _id)

	if err != nil {
		return *new(ParentContractPlatform), err
	}

	out0 := *abi.ConvertType(out[0], new(ParentContractPlatform)).(*ParentContractPlatform)

	return out0, err

}

// GetPlatform is a free data retrieval call binding the contract method 0x5bd22c98.
//
// Solidity: function getPlatform(uint256 _id) view returns((uint256,address,bytes))
func (_Parent *ParentSession) GetPlatform(_id *big.Int) (ParentContractPlatform, error) {
	return _Parent.Contract.GetPlatform(&_Parent.CallOpts, _id)
}

// GetPlatform is a free data retrieval call binding the contract method 0x5bd22c98.
//
// Solidity: function getPlatform(uint256 _id) view returns((uint256,address,bytes))
func (_Parent *ParentCallerSession) GetPlatform(_id *big.Int) (ParentContractPlatform, error) {
	return _Parent.Contract.GetPlatform(&_Parent.CallOpts, _id)
}

// GetPlatforms is a free data retrieval call binding the contract method 0x1fbe552e.
//
// Solidity: function getPlatforms() view returns((uint256,address,bytes)[])
func (_Parent *ParentCaller) GetPlatforms(opts *bind.CallOpts) ([]ParentContractPlatform, error) {
	var out []interface{}
	err := _Parent.contract.Call(opts, &out, "getPlatforms")

	if err != nil {
		return *new([]ParentContractPlatform), err
	}

	out0 := *abi.ConvertType(out[0], new([]ParentContractPlatform)).(*[]ParentContractPlatform)

	return out0, err

}

// GetPlatforms is a free data retrieval call binding the contract method 0x1fbe552e.
//
// Solidity: function getPlatforms() view returns((uint256,address,bytes)[])
func (_Parent *ParentSession) GetPlatforms() ([]ParentContractPlatform, error) {
	return _Parent.Contract.GetPlatforms(&_Parent.CallOpts)
}

// GetPlatforms is a free data retrieval call binding the contract method 0x1fbe552e.
//
// Solidity: function getPlatforms() view returns((uint256,address,bytes)[])
func (_Parent *ParentCallerSession) GetPlatforms() ([]ParentContractPlatform, error) {
	return _Parent.Contract.GetPlatforms(&_Parent.CallOpts)
}

// GetUserLatestPlatform is a free data retrieval call binding the contract method 0x02e8ce0a.
//
// Solidity: function getUserLatestPlatform() view returns((uint256,address,bytes))
func (_Parent *ParentCaller) GetUserLatestPlatform(opts *bind.CallOpts) (ParentContractPlatform, error) {
	var out []interface{}
	err := _Parent.contract.Call(opts, &out, "getUserLatestPlatform")

	if err != nil {
		return *new(ParentContractPlatform), err
	}

	out0 := *abi.ConvertType(out[0], new(ParentContractPlatform)).(*ParentContractPlatform)

	return out0, err

}

// GetUserLatestPlatform is a free data retrieval call binding the contract method 0x02e8ce0a.
//
// Solidity: function getUserLatestPlatform() view returns((uint256,address,bytes))
func (_Parent *ParentSession) GetUserLatestPlatform() (ParentContractPlatform, error) {
	return _Parent.Contract.GetUserLatestPlatform(&_Parent.CallOpts)
}

// GetUserLatestPlatform is a free data retrieval call binding the contract method 0x02e8ce0a.
//
// Solidity: function getUserLatestPlatform() view returns((uint256,address,bytes))
func (_Parent *ParentCallerSession) GetUserLatestPlatform() (ParentContractPlatform, error) {
	return _Parent.Contract.GetUserLatestPlatform(&_Parent.CallOpts)
}

// GetUserPlatforms is a free data retrieval call binding the contract method 0x5eb6d5f0.
//
// Solidity: function getUserPlatforms() view returns((uint256,address,bytes)[])
func (_Parent *ParentCaller) GetUserPlatforms(opts *bind.CallOpts) ([]ParentContractPlatform, error) {
	var out []interface{}
	err := _Parent.contract.Call(opts, &out, "getUserPlatforms")

	if err != nil {
		return *new([]ParentContractPlatform), err
	}

	out0 := *abi.ConvertType(out[0], new([]ParentContractPlatform)).(*[]ParentContractPlatform)

	return out0, err

}

// GetUserPlatforms is a free data retrieval call binding the contract method 0x5eb6d5f0.
//
// Solidity: function getUserPlatforms() view returns((uint256,address,bytes)[])
func (_Parent *ParentSession) GetUserPlatforms() ([]ParentContractPlatform, error) {
	return _Parent.Contract.GetUserPlatforms(&_Parent.CallOpts)
}

// GetUserPlatforms is a free data retrieval call binding the contract method 0x5eb6d5f0.
//
// Solidity: function getUserPlatforms() view returns((uint256,address,bytes)[])
func (_Parent *ParentCallerSession) GetUserPlatforms() ([]ParentContractPlatform, error) {
	return _Parent.Contract.GetUserPlatforms(&_Parent.CallOpts)
}

// CreatePlatform is a paid mutator transaction binding the contract method 0xc6e79b30.
//
// Solidity: function createPlatform(bytes _details) returns()
func (_Parent *ParentTransactor) CreatePlatform(opts *bind.TransactOpts, _details []byte) (*types.Transaction, error) {
	return _Parent.contract.Transact(opts, "createPlatform", _details)
}

// CreatePlatform is a paid mutator transaction binding the contract method 0xc6e79b30.
//
// Solidity: function createPlatform(bytes _details) returns()
func (_Parent *ParentSession) CreatePlatform(_details []byte) (*types.Transaction, error) {
	return _Parent.Contract.CreatePlatform(&_Parent.TransactOpts, _details)
}

// CreatePlatform is a paid mutator transaction binding the contract method 0xc6e79b30.
//
// Solidity: function createPlatform(bytes _details) returns()
func (_Parent *ParentTransactorSession) CreatePlatform(_details []byte) (*types.Transaction, error) {
	return _Parent.Contract.CreatePlatform(&_Parent.TransactOpts, _details)
}
