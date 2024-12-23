// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package utils

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
	_ = abi.ConvertType
)

// CdsAbiMetaData contains all meta data concerning the CdsAbi contract.
var CdsAbiMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"mainAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"_decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_destoryAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_finalTotalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_mintTotalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// CdsAbiABI is the input ABI used to generate the binding from.
// Deprecated: Use CdsAbiMetaData.ABI instead.
var CdsAbiABI = CdsAbiMetaData.ABI

// CdsAbi is an auto generated Go binding around an Ethereum contract.
type CdsAbi struct {
	CdsAbiCaller     // Read-only binding to the contract
	CdsAbiTransactor // Write-only binding to the contract
	CdsAbiFilterer   // Log filterer for contract events
}

// CdsAbiCaller is an auto generated read-only Go binding around an Ethereum contract.
type CdsAbiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CdsAbiTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CdsAbiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CdsAbiFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CdsAbiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CdsAbiSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CdsAbiSession struct {
	Contract     *CdsAbi           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CdsAbiCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CdsAbiCallerSession struct {
	Contract *CdsAbiCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// CdsAbiTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CdsAbiTransactorSession struct {
	Contract     *CdsAbiTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CdsAbiRaw is an auto generated low-level Go binding around an Ethereum contract.
type CdsAbiRaw struct {
	Contract *CdsAbi // Generic contract binding to access the raw methods on
}

// CdsAbiCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CdsAbiCallerRaw struct {
	Contract *CdsAbiCaller // Generic read-only contract binding to access the raw methods on
}

// CdsAbiTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CdsAbiTransactorRaw struct {
	Contract *CdsAbiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCdsAbi creates a new instance of CdsAbi, bound to a specific deployed contract.
func NewCdsAbi(address common.Address, backend bind.ContractBackend) (*CdsAbi, error) {
	contract, err := bindCdsAbi(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CdsAbi{CdsAbiCaller: CdsAbiCaller{contract: contract}, CdsAbiTransactor: CdsAbiTransactor{contract: contract}, CdsAbiFilterer: CdsAbiFilterer{contract: contract}}, nil
}

// NewCdsAbiCaller creates a new read-only instance of CdsAbi, bound to a specific deployed contract.
func NewCdsAbiCaller(address common.Address, caller bind.ContractCaller) (*CdsAbiCaller, error) {
	contract, err := bindCdsAbi(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CdsAbiCaller{contract: contract}, nil
}

// NewCdsAbiTransactor creates a new write-only instance of CdsAbi, bound to a specific deployed contract.
func NewCdsAbiTransactor(address common.Address, transactor bind.ContractTransactor) (*CdsAbiTransactor, error) {
	contract, err := bindCdsAbi(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CdsAbiTransactor{contract: contract}, nil
}

// NewCdsAbiFilterer creates a new log filterer instance of CdsAbi, bound to a specific deployed contract.
func NewCdsAbiFilterer(address common.Address, filterer bind.ContractFilterer) (*CdsAbiFilterer, error) {
	contract, err := bindCdsAbi(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CdsAbiFilterer{contract: contract}, nil
}

// bindCdsAbi binds a generic wrapper to an already deployed contract.
func bindCdsAbi(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CdsAbiMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CdsAbi *CdsAbiRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CdsAbi.Contract.CdsAbiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CdsAbi *CdsAbiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CdsAbi.Contract.CdsAbiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CdsAbi *CdsAbiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CdsAbi.Contract.CdsAbiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CdsAbi *CdsAbiCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CdsAbi.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CdsAbi *CdsAbiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CdsAbi.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CdsAbi *CdsAbiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CdsAbi.Contract.contract.Transact(opts, method, params...)
}

// PrivateDecimals is a free data retrieval call binding the contract method 0x32424aa3.
//
// Solidity: function _decimals() view returns(uint8)
func (_CdsAbi *CdsAbiCaller) PrivateDecimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _CdsAbi.contract.Call(opts, &out, "_decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// PrivateDecimals is a free data retrieval call binding the contract method 0x32424aa3.
//
// Solidity: function _decimals() view returns(uint8)
func (_CdsAbi *CdsAbiSession) PrivateDecimals() (uint8, error) {
	return _CdsAbi.Contract.PrivateDecimals(&_CdsAbi.CallOpts)
}

// PrivateDecimals is a free data retrieval call binding the contract method 0x32424aa3.
//
// Solidity: function _decimals() view returns(uint8)
func (_CdsAbi *CdsAbiCallerSession) PrivateDecimals() (uint8, error) {
	return _CdsAbi.Contract.PrivateDecimals(&_CdsAbi.CallOpts)
}

// DestoryAddress is a free data retrieval call binding the contract method 0x71b70e0b.
//
// Solidity: function _destoryAddress() view returns(address)
func (_CdsAbi *CdsAbiCaller) DestoryAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CdsAbi.contract.Call(opts, &out, "_destoryAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DestoryAddress is a free data retrieval call binding the contract method 0x71b70e0b.
//
// Solidity: function _destoryAddress() view returns(address)
func (_CdsAbi *CdsAbiSession) DestoryAddress() (common.Address, error) {
	return _CdsAbi.Contract.DestoryAddress(&_CdsAbi.CallOpts)
}

// DestoryAddress is a free data retrieval call binding the contract method 0x71b70e0b.
//
// Solidity: function _destoryAddress() view returns(address)
func (_CdsAbi *CdsAbiCallerSession) DestoryAddress() (common.Address, error) {
	return _CdsAbi.Contract.DestoryAddress(&_CdsAbi.CallOpts)
}

// FinalTotalSupply is a free data retrieval call binding the contract method 0xe0ade9f5.
//
// Solidity: function _finalTotalSupply() view returns(uint256)
func (_CdsAbi *CdsAbiCaller) FinalTotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CdsAbi.contract.Call(opts, &out, "_finalTotalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FinalTotalSupply is a free data retrieval call binding the contract method 0xe0ade9f5.
//
// Solidity: function _finalTotalSupply() view returns(uint256)
func (_CdsAbi *CdsAbiSession) FinalTotalSupply() (*big.Int, error) {
	return _CdsAbi.Contract.FinalTotalSupply(&_CdsAbi.CallOpts)
}

// FinalTotalSupply is a free data retrieval call binding the contract method 0xe0ade9f5.
//
// Solidity: function _finalTotalSupply() view returns(uint256)
func (_CdsAbi *CdsAbiCallerSession) FinalTotalSupply() (*big.Int, error) {
	return _CdsAbi.Contract.FinalTotalSupply(&_CdsAbi.CallOpts)
}

// MintTotalSupply is a free data retrieval call binding the contract method 0x13d247f8.
//
// Solidity: function _mintTotalSupply() view returns(uint256)
func (_CdsAbi *CdsAbiCaller) MintTotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CdsAbi.contract.Call(opts, &out, "_mintTotalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MintTotalSupply is a free data retrieval call binding the contract method 0x13d247f8.
//
// Solidity: function _mintTotalSupply() view returns(uint256)
func (_CdsAbi *CdsAbiSession) MintTotalSupply() (*big.Int, error) {
	return _CdsAbi.Contract.MintTotalSupply(&_CdsAbi.CallOpts)
}

// MintTotalSupply is a free data retrieval call binding the contract method 0x13d247f8.
//
// Solidity: function _mintTotalSupply() view returns(uint256)
func (_CdsAbi *CdsAbiCallerSession) MintTotalSupply() (*big.Int, error) {
	return _CdsAbi.Contract.MintTotalSupply(&_CdsAbi.CallOpts)
}

// PrivateName is a free data retrieval call binding the contract method 0xd28d8852.
//
// Solidity: function _name() view returns(string)
func (_CdsAbi *CdsAbiCaller) PrivateName(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CdsAbi.contract.Call(opts, &out, "_name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// PrivateName is a free data retrieval call binding the contract method 0xd28d8852.
//
// Solidity: function _name() view returns(string)
func (_CdsAbi *CdsAbiSession) PrivateName() (string, error) {
	return _CdsAbi.Contract.PrivateName(&_CdsAbi.CallOpts)
}

// PrivateName is a free data retrieval call binding the contract method 0xd28d8852.
//
// Solidity: function _name() view returns(string)
func (_CdsAbi *CdsAbiCallerSession) PrivateName() (string, error) {
	return _CdsAbi.Contract.PrivateName(&_CdsAbi.CallOpts)
}

// PrivateSymbol is a free data retrieval call binding the contract method 0xb09f1266.
//
// Solidity: function _symbol() view returns(string)
func (_CdsAbi *CdsAbiCaller) PrivateSymbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CdsAbi.contract.Call(opts, &out, "_symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// PrivateSymbol is a free data retrieval call binding the contract method 0xb09f1266.
//
// Solidity: function _symbol() view returns(string)
func (_CdsAbi *CdsAbiSession) PrivateSymbol() (string, error) {
	return _CdsAbi.Contract.PrivateSymbol(&_CdsAbi.CallOpts)
}

// PrivateSymbol is a free data retrieval call binding the contract method 0xb09f1266.
//
// Solidity: function _symbol() view returns(string)
func (_CdsAbi *CdsAbiCallerSession) PrivateSymbol() (string, error) {
	return _CdsAbi.Contract.PrivateSymbol(&_CdsAbi.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_CdsAbi *CdsAbiCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CdsAbi.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_CdsAbi *CdsAbiSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _CdsAbi.Contract.Allowance(&_CdsAbi.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_CdsAbi *CdsAbiCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _CdsAbi.Contract.Allowance(&_CdsAbi.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_CdsAbi *CdsAbiCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CdsAbi.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_CdsAbi *CdsAbiSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _CdsAbi.Contract.BalanceOf(&_CdsAbi.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_CdsAbi *CdsAbiCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _CdsAbi.Contract.BalanceOf(&_CdsAbi.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_CdsAbi *CdsAbiCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _CdsAbi.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_CdsAbi *CdsAbiSession) Decimals() (uint8, error) {
	return _CdsAbi.Contract.Decimals(&_CdsAbi.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_CdsAbi *CdsAbiCallerSession) Decimals() (uint8, error) {
	return _CdsAbi.Contract.Decimals(&_CdsAbi.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_CdsAbi *CdsAbiCaller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CdsAbi.contract.Call(opts, &out, "getOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_CdsAbi *CdsAbiSession) GetOwner() (common.Address, error) {
	return _CdsAbi.Contract.GetOwner(&_CdsAbi.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_CdsAbi *CdsAbiCallerSession) GetOwner() (common.Address, error) {
	return _CdsAbi.Contract.GetOwner(&_CdsAbi.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CdsAbi *CdsAbiCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CdsAbi.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CdsAbi *CdsAbiSession) Name() (string, error) {
	return _CdsAbi.Contract.Name(&_CdsAbi.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CdsAbi *CdsAbiCallerSession) Name() (string, error) {
	return _CdsAbi.Contract.Name(&_CdsAbi.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CdsAbi *CdsAbiCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CdsAbi.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CdsAbi *CdsAbiSession) Owner() (common.Address, error) {
	return _CdsAbi.Contract.Owner(&_CdsAbi.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CdsAbi *CdsAbiCallerSession) Owner() (common.Address, error) {
	return _CdsAbi.Contract.Owner(&_CdsAbi.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CdsAbi *CdsAbiCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CdsAbi.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CdsAbi *CdsAbiSession) Symbol() (string, error) {
	return _CdsAbi.Contract.Symbol(&_CdsAbi.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CdsAbi *CdsAbiCallerSession) Symbol() (string, error) {
	return _CdsAbi.Contract.Symbol(&_CdsAbi.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CdsAbi *CdsAbiCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CdsAbi.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CdsAbi *CdsAbiSession) TotalSupply() (*big.Int, error) {
	return _CdsAbi.Contract.TotalSupply(&_CdsAbi.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CdsAbi *CdsAbiCallerSession) TotalSupply() (*big.Int, error) {
	return _CdsAbi.Contract.TotalSupply(&_CdsAbi.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_CdsAbi *CdsAbiTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CdsAbi.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_CdsAbi *CdsAbiSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CdsAbi.Contract.Approve(&_CdsAbi.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_CdsAbi *CdsAbiTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CdsAbi.Contract.Approve(&_CdsAbi.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns(bool)
func (_CdsAbi *CdsAbiTransactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _CdsAbi.contract.Transact(opts, "burn", amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns(bool)
func (_CdsAbi *CdsAbiSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _CdsAbi.Contract.Burn(&_CdsAbi.TransactOpts, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns(bool)
func (_CdsAbi *CdsAbiTransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _CdsAbi.Contract.Burn(&_CdsAbi.TransactOpts, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_CdsAbi *CdsAbiTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _CdsAbi.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_CdsAbi *CdsAbiSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _CdsAbi.Contract.DecreaseAllowance(&_CdsAbi.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_CdsAbi *CdsAbiTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _CdsAbi.Contract.DecreaseAllowance(&_CdsAbi.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_CdsAbi *CdsAbiTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _CdsAbi.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_CdsAbi *CdsAbiSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _CdsAbi.Contract.IncreaseAllowance(&_CdsAbi.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_CdsAbi *CdsAbiTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _CdsAbi.Contract.IncreaseAllowance(&_CdsAbi.TransactOpts, spender, addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address addr, uint256 amount) returns(uint256)
func (_CdsAbi *CdsAbiTransactor) Mint(opts *bind.TransactOpts, addr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CdsAbi.contract.Transact(opts, "mint", addr, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address addr, uint256 amount) returns(uint256)
func (_CdsAbi *CdsAbiSession) Mint(addr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CdsAbi.Contract.Mint(&_CdsAbi.TransactOpts, addr, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address addr, uint256 amount) returns(uint256)
func (_CdsAbi *CdsAbiTransactorSession) Mint(addr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CdsAbi.Contract.Mint(&_CdsAbi.TransactOpts, addr, amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CdsAbi *CdsAbiTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CdsAbi.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CdsAbi *CdsAbiSession) RenounceOwnership() (*types.Transaction, error) {
	return _CdsAbi.Contract.RenounceOwnership(&_CdsAbi.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CdsAbi *CdsAbiTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _CdsAbi.Contract.RenounceOwnership(&_CdsAbi.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_CdsAbi *CdsAbiTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CdsAbi.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_CdsAbi *CdsAbiSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CdsAbi.Contract.Transfer(&_CdsAbi.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_CdsAbi *CdsAbiTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CdsAbi.Contract.Transfer(&_CdsAbi.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_CdsAbi *CdsAbiTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CdsAbi.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_CdsAbi *CdsAbiSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CdsAbi.Contract.TransferFrom(&_CdsAbi.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_CdsAbi *CdsAbiTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CdsAbi.Contract.TransferFrom(&_CdsAbi.TransactOpts, sender, recipient, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CdsAbi *CdsAbiTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _CdsAbi.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CdsAbi *CdsAbiSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CdsAbi.Contract.TransferOwnership(&_CdsAbi.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CdsAbi *CdsAbiTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CdsAbi.Contract.TransferOwnership(&_CdsAbi.TransactOpts, newOwner)
}

// CdsAbiApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the CdsAbi contract.
type CdsAbiApprovalIterator struct {
	Event *CdsAbiApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CdsAbiApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CdsAbiApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CdsAbiApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CdsAbiApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CdsAbiApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CdsAbiApproval represents a Approval event raised by the CdsAbi contract.
type CdsAbiApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_CdsAbi *CdsAbiFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*CdsAbiApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _CdsAbi.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &CdsAbiApprovalIterator{contract: _CdsAbi.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_CdsAbi *CdsAbiFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *CdsAbiApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _CdsAbi.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CdsAbiApproval)
				if err := _CdsAbi.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_CdsAbi *CdsAbiFilterer) ParseApproval(log types.Log) (*CdsAbiApproval, error) {
	event := new(CdsAbiApproval)
	if err := _CdsAbi.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CdsAbiOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the CdsAbi contract.
type CdsAbiOwnershipTransferredIterator struct {
	Event *CdsAbiOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CdsAbiOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CdsAbiOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CdsAbiOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CdsAbiOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CdsAbiOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CdsAbiOwnershipTransferred represents a OwnershipTransferred event raised by the CdsAbi contract.
type CdsAbiOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CdsAbi *CdsAbiFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*CdsAbiOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CdsAbi.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CdsAbiOwnershipTransferredIterator{contract: _CdsAbi.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CdsAbi *CdsAbiFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CdsAbiOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CdsAbi.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CdsAbiOwnershipTransferred)
				if err := _CdsAbi.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CdsAbi *CdsAbiFilterer) ParseOwnershipTransferred(log types.Log) (*CdsAbiOwnershipTransferred, error) {
	event := new(CdsAbiOwnershipTransferred)
	if err := _CdsAbi.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CdsAbiTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the CdsAbi contract.
type CdsAbiTransferIterator struct {
	Event *CdsAbiTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CdsAbiTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CdsAbiTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CdsAbiTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CdsAbiTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CdsAbiTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CdsAbiTransfer represents a Transfer event raised by the CdsAbi contract.
type CdsAbiTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_CdsAbi *CdsAbiFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CdsAbiTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CdsAbi.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CdsAbiTransferIterator{contract: _CdsAbi.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_CdsAbi *CdsAbiFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *CdsAbiTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CdsAbi.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CdsAbiTransfer)
				if err := _CdsAbi.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_CdsAbi *CdsAbiFilterer) ParseTransfer(log types.Log) (*CdsAbiTransfer, error) {
	event := new(CdsAbiTransfer)
	if err := _CdsAbi.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
