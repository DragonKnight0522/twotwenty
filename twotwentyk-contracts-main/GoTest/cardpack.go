// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

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

// CardpackMetaData contains all meta data concerning the Cardpack contract.
var CardpackMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_standardAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_premiumAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_eliteAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_totalAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"packNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"typeOf\",\"type\":\"string\"}],\"name\":\"CardPackCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"CardPackOpened\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newLimit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"typeOf\",\"type\":\"string\"}],\"name\":\"LimitChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousMinter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newMinter\",\"type\":\"address\"}],\"name\":\"MinterContractChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newAdmin\",\"type\":\"address\"}],\"name\":\"changeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newMinter\",\"type\":\"address\"}],\"name\":\"changeMinter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newLimit\",\"type\":\"uint256\"}],\"name\":\"changeToEliteLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"changeToOpened\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newLimit\",\"type\":\"uint256\"}],\"name\":\"changeToPremiumLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newLimit\",\"type\":\"uint256\"}],\"name\":\"changeToStandardLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newLimit\",\"type\":\"uint256\"}],\"name\":\"changeToTotalLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_tokenURI\",\"type\":\"string\"}],\"name\":\"createEliteCard\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_tokenURI\",\"type\":\"string\"}],\"name\":\"createPremiumCard\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_tokenURI\",\"type\":\"string\"}],\"name\":\"createStandardCard\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eliteLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"isOpened\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minterContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"opened\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"premiumLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"standardLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenCounterElite\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenCounterPremium\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenCounterStandard\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// CardpackABI is the input ABI used to generate the binding from.
// Deprecated: Use CardpackMetaData.ABI instead.
var CardpackABI = CardpackMetaData.ABI

// Cardpack is an auto generated Go binding around an Ethereum contract.
type Cardpack struct {
	CardpackCaller     // Read-only binding to the contract
	CardpackTransactor // Write-only binding to the contract
	CardpackFilterer   // Log filterer for contract events
}

// CardpackCaller is an auto generated read-only Go binding around an Ethereum contract.
type CardpackCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CardpackTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CardpackTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CardpackFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CardpackFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CardpackSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CardpackSession struct {
	Contract     *Cardpack         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CardpackCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CardpackCallerSession struct {
	Contract *CardpackCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// CardpackTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CardpackTransactorSession struct {
	Contract     *CardpackTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// CardpackRaw is an auto generated low-level Go binding around an Ethereum contract.
type CardpackRaw struct {
	Contract *Cardpack // Generic contract binding to access the raw methods on
}

// CardpackCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CardpackCallerRaw struct {
	Contract *CardpackCaller // Generic read-only contract binding to access the raw methods on
}

// CardpackTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CardpackTransactorRaw struct {
	Contract *CardpackTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCardpack creates a new instance of Cardpack, bound to a specific deployed contract.
func NewCardpack(address common.Address, backend bind.ContractBackend) (*Cardpack, error) {
	contract, err := bindCardpack(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Cardpack{CardpackCaller: CardpackCaller{contract: contract}, CardpackTransactor: CardpackTransactor{contract: contract}, CardpackFilterer: CardpackFilterer{contract: contract}}, nil
}

// NewCardpackCaller creates a new read-only instance of Cardpack, bound to a specific deployed contract.
func NewCardpackCaller(address common.Address, caller bind.ContractCaller) (*CardpackCaller, error) {
	contract, err := bindCardpack(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CardpackCaller{contract: contract}, nil
}

// NewCardpackTransactor creates a new write-only instance of Cardpack, bound to a specific deployed contract.
func NewCardpackTransactor(address common.Address, transactor bind.ContractTransactor) (*CardpackTransactor, error) {
	contract, err := bindCardpack(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CardpackTransactor{contract: contract}, nil
}

// NewCardpackFilterer creates a new log filterer instance of Cardpack, bound to a specific deployed contract.
func NewCardpackFilterer(address common.Address, filterer bind.ContractFilterer) (*CardpackFilterer, error) {
	contract, err := bindCardpack(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CardpackFilterer{contract: contract}, nil
}

// bindCardpack binds a generic wrapper to an already deployed contract.
func bindCardpack(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CardpackMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Cardpack *CardpackRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Cardpack.Contract.CardpackCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Cardpack *CardpackRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cardpack.Contract.CardpackTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Cardpack *CardpackRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Cardpack.Contract.CardpackTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Cardpack *CardpackCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Cardpack.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Cardpack *CardpackTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cardpack.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Cardpack *CardpackTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Cardpack.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Cardpack *CardpackCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Cardpack.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Cardpack *CardpackSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Cardpack.Contract.BalanceOf(&_Cardpack.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Cardpack *CardpackCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Cardpack.Contract.BalanceOf(&_Cardpack.CallOpts, owner)
}

// EliteLimit is a free data retrieval call binding the contract method 0x9b229845.
//
// Solidity: function eliteLimit() view returns(uint256)
func (_Cardpack *CardpackCaller) EliteLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Cardpack.contract.Call(opts, &out, "eliteLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EliteLimit is a free data retrieval call binding the contract method 0x9b229845.
//
// Solidity: function eliteLimit() view returns(uint256)
func (_Cardpack *CardpackSession) EliteLimit() (*big.Int, error) {
	return _Cardpack.Contract.EliteLimit(&_Cardpack.CallOpts)
}

// EliteLimit is a free data retrieval call binding the contract method 0x9b229845.
//
// Solidity: function eliteLimit() view returns(uint256)
func (_Cardpack *CardpackCallerSession) EliteLimit() (*big.Int, error) {
	return _Cardpack.Contract.EliteLimit(&_Cardpack.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Cardpack *CardpackCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Cardpack.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Cardpack *CardpackSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Cardpack.Contract.GetApproved(&_Cardpack.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Cardpack *CardpackCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Cardpack.Contract.GetApproved(&_Cardpack.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Cardpack *CardpackCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Cardpack.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Cardpack *CardpackSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Cardpack.Contract.IsApprovedForAll(&_Cardpack.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Cardpack *CardpackCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Cardpack.Contract.IsApprovedForAll(&_Cardpack.CallOpts, owner, operator)
}

// IsOpened is a free data retrieval call binding the contract method 0x5faf46bb.
//
// Solidity: function isOpened(uint256 tokenId) view returns(bool)
func (_Cardpack *CardpackCaller) IsOpened(opts *bind.CallOpts, tokenId *big.Int) (bool, error) {
	var out []interface{}
	err := _Cardpack.contract.Call(opts, &out, "isOpened", tokenId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOpened is a free data retrieval call binding the contract method 0x5faf46bb.
//
// Solidity: function isOpened(uint256 tokenId) view returns(bool)
func (_Cardpack *CardpackSession) IsOpened(tokenId *big.Int) (bool, error) {
	return _Cardpack.Contract.IsOpened(&_Cardpack.CallOpts, tokenId)
}

// IsOpened is a free data retrieval call binding the contract method 0x5faf46bb.
//
// Solidity: function isOpened(uint256 tokenId) view returns(bool)
func (_Cardpack *CardpackCallerSession) IsOpened(tokenId *big.Int) (bool, error) {
	return _Cardpack.Contract.IsOpened(&_Cardpack.CallOpts, tokenId)
}

// MinterContract is a free data retrieval call binding the contract method 0x92f00233.
//
// Solidity: function minterContract() view returns(address)
func (_Cardpack *CardpackCaller) MinterContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Cardpack.contract.Call(opts, &out, "minterContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MinterContract is a free data retrieval call binding the contract method 0x92f00233.
//
// Solidity: function minterContract() view returns(address)
func (_Cardpack *CardpackSession) MinterContract() (common.Address, error) {
	return _Cardpack.Contract.MinterContract(&_Cardpack.CallOpts)
}

// MinterContract is a free data retrieval call binding the contract method 0x92f00233.
//
// Solidity: function minterContract() view returns(address)
func (_Cardpack *CardpackCallerSession) MinterContract() (common.Address, error) {
	return _Cardpack.Contract.MinterContract(&_Cardpack.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Cardpack *CardpackCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Cardpack.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Cardpack *CardpackSession) Name() (string, error) {
	return _Cardpack.Contract.Name(&_Cardpack.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Cardpack *CardpackCallerSession) Name() (string, error) {
	return _Cardpack.Contract.Name(&_Cardpack.CallOpts)
}

// Opened is a free data retrieval call binding the contract method 0xf1ea5cd3.
//
// Solidity: function opened(uint256 ) view returns(bool)
func (_Cardpack *CardpackCaller) Opened(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _Cardpack.contract.Call(opts, &out, "opened", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Opened is a free data retrieval call binding the contract method 0xf1ea5cd3.
//
// Solidity: function opened(uint256 ) view returns(bool)
func (_Cardpack *CardpackSession) Opened(arg0 *big.Int) (bool, error) {
	return _Cardpack.Contract.Opened(&_Cardpack.CallOpts, arg0)
}

// Opened is a free data retrieval call binding the contract method 0xf1ea5cd3.
//
// Solidity: function opened(uint256 ) view returns(bool)
func (_Cardpack *CardpackCallerSession) Opened(arg0 *big.Int) (bool, error) {
	return _Cardpack.Contract.Opened(&_Cardpack.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Cardpack *CardpackCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Cardpack.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Cardpack *CardpackSession) Owner() (common.Address, error) {
	return _Cardpack.Contract.Owner(&_Cardpack.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Cardpack *CardpackCallerSession) Owner() (common.Address, error) {
	return _Cardpack.Contract.Owner(&_Cardpack.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Cardpack *CardpackCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Cardpack.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Cardpack *CardpackSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Cardpack.Contract.OwnerOf(&_Cardpack.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Cardpack *CardpackCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Cardpack.Contract.OwnerOf(&_Cardpack.CallOpts, tokenId)
}

// PremiumLimit is a free data retrieval call binding the contract method 0xc359f946.
//
// Solidity: function premiumLimit() view returns(uint256)
func (_Cardpack *CardpackCaller) PremiumLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Cardpack.contract.Call(opts, &out, "premiumLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PremiumLimit is a free data retrieval call binding the contract method 0xc359f946.
//
// Solidity: function premiumLimit() view returns(uint256)
func (_Cardpack *CardpackSession) PremiumLimit() (*big.Int, error) {
	return _Cardpack.Contract.PremiumLimit(&_Cardpack.CallOpts)
}

// PremiumLimit is a free data retrieval call binding the contract method 0xc359f946.
//
// Solidity: function premiumLimit() view returns(uint256)
func (_Cardpack *CardpackCallerSession) PremiumLimit() (*big.Int, error) {
	return _Cardpack.Contract.PremiumLimit(&_Cardpack.CallOpts)
}

// StandardLimit is a free data retrieval call binding the contract method 0x487f2039.
//
// Solidity: function standardLimit() view returns(uint256)
func (_Cardpack *CardpackCaller) StandardLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Cardpack.contract.Call(opts, &out, "standardLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StandardLimit is a free data retrieval call binding the contract method 0x487f2039.
//
// Solidity: function standardLimit() view returns(uint256)
func (_Cardpack *CardpackSession) StandardLimit() (*big.Int, error) {
	return _Cardpack.Contract.StandardLimit(&_Cardpack.CallOpts)
}

// StandardLimit is a free data retrieval call binding the contract method 0x487f2039.
//
// Solidity: function standardLimit() view returns(uint256)
func (_Cardpack *CardpackCallerSession) StandardLimit() (*big.Int, error) {
	return _Cardpack.Contract.StandardLimit(&_Cardpack.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Cardpack *CardpackCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Cardpack.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Cardpack *CardpackSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Cardpack.Contract.SupportsInterface(&_Cardpack.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Cardpack *CardpackCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Cardpack.Contract.SupportsInterface(&_Cardpack.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Cardpack *CardpackCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Cardpack.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Cardpack *CardpackSession) Symbol() (string, error) {
	return _Cardpack.Contract.Symbol(&_Cardpack.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Cardpack *CardpackCallerSession) Symbol() (string, error) {
	return _Cardpack.Contract.Symbol(&_Cardpack.CallOpts)
}

// TokenCounterElite is a free data retrieval call binding the contract method 0xbf5e7328.
//
// Solidity: function tokenCounterElite() view returns(uint256)
func (_Cardpack *CardpackCaller) TokenCounterElite(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Cardpack.contract.Call(opts, &out, "tokenCounterElite")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenCounterElite is a free data retrieval call binding the contract method 0xbf5e7328.
//
// Solidity: function tokenCounterElite() view returns(uint256)
func (_Cardpack *CardpackSession) TokenCounterElite() (*big.Int, error) {
	return _Cardpack.Contract.TokenCounterElite(&_Cardpack.CallOpts)
}

// TokenCounterElite is a free data retrieval call binding the contract method 0xbf5e7328.
//
// Solidity: function tokenCounterElite() view returns(uint256)
func (_Cardpack *CardpackCallerSession) TokenCounterElite() (*big.Int, error) {
	return _Cardpack.Contract.TokenCounterElite(&_Cardpack.CallOpts)
}

// TokenCounterPremium is a free data retrieval call binding the contract method 0xae49228e.
//
// Solidity: function tokenCounterPremium() view returns(uint256)
func (_Cardpack *CardpackCaller) TokenCounterPremium(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Cardpack.contract.Call(opts, &out, "tokenCounterPremium")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenCounterPremium is a free data retrieval call binding the contract method 0xae49228e.
//
// Solidity: function tokenCounterPremium() view returns(uint256)
func (_Cardpack *CardpackSession) TokenCounterPremium() (*big.Int, error) {
	return _Cardpack.Contract.TokenCounterPremium(&_Cardpack.CallOpts)
}

// TokenCounterPremium is a free data retrieval call binding the contract method 0xae49228e.
//
// Solidity: function tokenCounterPremium() view returns(uint256)
func (_Cardpack *CardpackCallerSession) TokenCounterPremium() (*big.Int, error) {
	return _Cardpack.Contract.TokenCounterPremium(&_Cardpack.CallOpts)
}

// TokenCounterStandard is a free data retrieval call binding the contract method 0xf568e90a.
//
// Solidity: function tokenCounterStandard() view returns(uint256)
func (_Cardpack *CardpackCaller) TokenCounterStandard(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Cardpack.contract.Call(opts, &out, "tokenCounterStandard")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenCounterStandard is a free data retrieval call binding the contract method 0xf568e90a.
//
// Solidity: function tokenCounterStandard() view returns(uint256)
func (_Cardpack *CardpackSession) TokenCounterStandard() (*big.Int, error) {
	return _Cardpack.Contract.TokenCounterStandard(&_Cardpack.CallOpts)
}

// TokenCounterStandard is a free data retrieval call binding the contract method 0xf568e90a.
//
// Solidity: function tokenCounterStandard() view returns(uint256)
func (_Cardpack *CardpackCallerSession) TokenCounterStandard() (*big.Int, error) {
	return _Cardpack.Contract.TokenCounterStandard(&_Cardpack.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Cardpack *CardpackCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Cardpack.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Cardpack *CardpackSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Cardpack.Contract.TokenURI(&_Cardpack.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Cardpack *CardpackCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Cardpack.Contract.TokenURI(&_Cardpack.CallOpts, tokenId)
}

// TotalLimit is a free data retrieval call binding the contract method 0xa36298c7.
//
// Solidity: function totalLimit() view returns(uint256)
func (_Cardpack *CardpackCaller) TotalLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Cardpack.contract.Call(opts, &out, "totalLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalLimit is a free data retrieval call binding the contract method 0xa36298c7.
//
// Solidity: function totalLimit() view returns(uint256)
func (_Cardpack *CardpackSession) TotalLimit() (*big.Int, error) {
	return _Cardpack.Contract.TotalLimit(&_Cardpack.CallOpts)
}

// TotalLimit is a free data retrieval call binding the contract method 0xa36298c7.
//
// Solidity: function totalLimit() view returns(uint256)
func (_Cardpack *CardpackCallerSession) TotalLimit() (*big.Int, error) {
	return _Cardpack.Contract.TotalLimit(&_Cardpack.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Cardpack *CardpackTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Cardpack.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Cardpack *CardpackSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Cardpack.Contract.Approve(&_Cardpack.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Cardpack *CardpackTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Cardpack.Contract.Approve(&_Cardpack.TransactOpts, to, tokenId)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address _newAdmin) returns()
func (_Cardpack *CardpackTransactor) ChangeAdmin(opts *bind.TransactOpts, _newAdmin common.Address) (*types.Transaction, error) {
	return _Cardpack.contract.Transact(opts, "changeAdmin", _newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address _newAdmin) returns()
func (_Cardpack *CardpackSession) ChangeAdmin(_newAdmin common.Address) (*types.Transaction, error) {
	return _Cardpack.Contract.ChangeAdmin(&_Cardpack.TransactOpts, _newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address _newAdmin) returns()
func (_Cardpack *CardpackTransactorSession) ChangeAdmin(_newAdmin common.Address) (*types.Transaction, error) {
	return _Cardpack.Contract.ChangeAdmin(&_Cardpack.TransactOpts, _newAdmin)
}

// ChangeMinter is a paid mutator transaction binding the contract method 0x2c4d4d18.
//
// Solidity: function changeMinter(address _newMinter) returns()
func (_Cardpack *CardpackTransactor) ChangeMinter(opts *bind.TransactOpts, _newMinter common.Address) (*types.Transaction, error) {
	return _Cardpack.contract.Transact(opts, "changeMinter", _newMinter)
}

// ChangeMinter is a paid mutator transaction binding the contract method 0x2c4d4d18.
//
// Solidity: function changeMinter(address _newMinter) returns()
func (_Cardpack *CardpackSession) ChangeMinter(_newMinter common.Address) (*types.Transaction, error) {
	return _Cardpack.Contract.ChangeMinter(&_Cardpack.TransactOpts, _newMinter)
}

// ChangeMinter is a paid mutator transaction binding the contract method 0x2c4d4d18.
//
// Solidity: function changeMinter(address _newMinter) returns()
func (_Cardpack *CardpackTransactorSession) ChangeMinter(_newMinter common.Address) (*types.Transaction, error) {
	return _Cardpack.Contract.ChangeMinter(&_Cardpack.TransactOpts, _newMinter)
}

// ChangeToEliteLimit is a paid mutator transaction binding the contract method 0x457c5503.
//
// Solidity: function changeToEliteLimit(uint256 newLimit) returns()
func (_Cardpack *CardpackTransactor) ChangeToEliteLimit(opts *bind.TransactOpts, newLimit *big.Int) (*types.Transaction, error) {
	return _Cardpack.contract.Transact(opts, "changeToEliteLimit", newLimit)
}

// ChangeToEliteLimit is a paid mutator transaction binding the contract method 0x457c5503.
//
// Solidity: function changeToEliteLimit(uint256 newLimit) returns()
func (_Cardpack *CardpackSession) ChangeToEliteLimit(newLimit *big.Int) (*types.Transaction, error) {
	return _Cardpack.Contract.ChangeToEliteLimit(&_Cardpack.TransactOpts, newLimit)
}

// ChangeToEliteLimit is a paid mutator transaction binding the contract method 0x457c5503.
//
// Solidity: function changeToEliteLimit(uint256 newLimit) returns()
func (_Cardpack *CardpackTransactorSession) ChangeToEliteLimit(newLimit *big.Int) (*types.Transaction, error) {
	return _Cardpack.Contract.ChangeToEliteLimit(&_Cardpack.TransactOpts, newLimit)
}

// ChangeToOpened is a paid mutator transaction binding the contract method 0xd9a99ac5.
//
// Solidity: function changeToOpened(uint256 tokenId) returns()
func (_Cardpack *CardpackTransactor) ChangeToOpened(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _Cardpack.contract.Transact(opts, "changeToOpened", tokenId)
}

// ChangeToOpened is a paid mutator transaction binding the contract method 0xd9a99ac5.
//
// Solidity: function changeToOpened(uint256 tokenId) returns()
func (_Cardpack *CardpackSession) ChangeToOpened(tokenId *big.Int) (*types.Transaction, error) {
	return _Cardpack.Contract.ChangeToOpened(&_Cardpack.TransactOpts, tokenId)
}

// ChangeToOpened is a paid mutator transaction binding the contract method 0xd9a99ac5.
//
// Solidity: function changeToOpened(uint256 tokenId) returns()
func (_Cardpack *CardpackTransactorSession) ChangeToOpened(tokenId *big.Int) (*types.Transaction, error) {
	return _Cardpack.Contract.ChangeToOpened(&_Cardpack.TransactOpts, tokenId)
}

// ChangeToPremiumLimit is a paid mutator transaction binding the contract method 0x915a4edc.
//
// Solidity: function changeToPremiumLimit(uint256 newLimit) returns()
func (_Cardpack *CardpackTransactor) ChangeToPremiumLimit(opts *bind.TransactOpts, newLimit *big.Int) (*types.Transaction, error) {
	return _Cardpack.contract.Transact(opts, "changeToPremiumLimit", newLimit)
}

// ChangeToPremiumLimit is a paid mutator transaction binding the contract method 0x915a4edc.
//
// Solidity: function changeToPremiumLimit(uint256 newLimit) returns()
func (_Cardpack *CardpackSession) ChangeToPremiumLimit(newLimit *big.Int) (*types.Transaction, error) {
	return _Cardpack.Contract.ChangeToPremiumLimit(&_Cardpack.TransactOpts, newLimit)
}

// ChangeToPremiumLimit is a paid mutator transaction binding the contract method 0x915a4edc.
//
// Solidity: function changeToPremiumLimit(uint256 newLimit) returns()
func (_Cardpack *CardpackTransactorSession) ChangeToPremiumLimit(newLimit *big.Int) (*types.Transaction, error) {
	return _Cardpack.Contract.ChangeToPremiumLimit(&_Cardpack.TransactOpts, newLimit)
}

// ChangeToStandardLimit is a paid mutator transaction binding the contract method 0x06fe5ebf.
//
// Solidity: function changeToStandardLimit(uint256 newLimit) returns()
func (_Cardpack *CardpackTransactor) ChangeToStandardLimit(opts *bind.TransactOpts, newLimit *big.Int) (*types.Transaction, error) {
	return _Cardpack.contract.Transact(opts, "changeToStandardLimit", newLimit)
}

// ChangeToStandardLimit is a paid mutator transaction binding the contract method 0x06fe5ebf.
//
// Solidity: function changeToStandardLimit(uint256 newLimit) returns()
func (_Cardpack *CardpackSession) ChangeToStandardLimit(newLimit *big.Int) (*types.Transaction, error) {
	return _Cardpack.Contract.ChangeToStandardLimit(&_Cardpack.TransactOpts, newLimit)
}

// ChangeToStandardLimit is a paid mutator transaction binding the contract method 0x06fe5ebf.
//
// Solidity: function changeToStandardLimit(uint256 newLimit) returns()
func (_Cardpack *CardpackTransactorSession) ChangeToStandardLimit(newLimit *big.Int) (*types.Transaction, error) {
	return _Cardpack.Contract.ChangeToStandardLimit(&_Cardpack.TransactOpts, newLimit)
}

// ChangeToTotalLimit is a paid mutator transaction binding the contract method 0x4bc8355a.
//
// Solidity: function changeToTotalLimit(uint256 newLimit) returns()
func (_Cardpack *CardpackTransactor) ChangeToTotalLimit(opts *bind.TransactOpts, newLimit *big.Int) (*types.Transaction, error) {
	return _Cardpack.contract.Transact(opts, "changeToTotalLimit", newLimit)
}

// ChangeToTotalLimit is a paid mutator transaction binding the contract method 0x4bc8355a.
//
// Solidity: function changeToTotalLimit(uint256 newLimit) returns()
func (_Cardpack *CardpackSession) ChangeToTotalLimit(newLimit *big.Int) (*types.Transaction, error) {
	return _Cardpack.Contract.ChangeToTotalLimit(&_Cardpack.TransactOpts, newLimit)
}

// ChangeToTotalLimit is a paid mutator transaction binding the contract method 0x4bc8355a.
//
// Solidity: function changeToTotalLimit(uint256 newLimit) returns()
func (_Cardpack *CardpackTransactorSession) ChangeToTotalLimit(newLimit *big.Int) (*types.Transaction, error) {
	return _Cardpack.Contract.ChangeToTotalLimit(&_Cardpack.TransactOpts, newLimit)
}

// CreateEliteCard is a paid mutator transaction binding the contract method 0x7fb5772d.
//
// Solidity: function createEliteCard(string _tokenURI) returns(uint256)
func (_Cardpack *CardpackTransactor) CreateEliteCard(opts *bind.TransactOpts, _tokenURI string) (*types.Transaction, error) {
	return _Cardpack.contract.Transact(opts, "createEliteCard", _tokenURI)
}

// CreateEliteCard is a paid mutator transaction binding the contract method 0x7fb5772d.
//
// Solidity: function createEliteCard(string _tokenURI) returns(uint256)
func (_Cardpack *CardpackSession) CreateEliteCard(_tokenURI string) (*types.Transaction, error) {
	return _Cardpack.Contract.CreateEliteCard(&_Cardpack.TransactOpts, _tokenURI)
}

// CreateEliteCard is a paid mutator transaction binding the contract method 0x7fb5772d.
//
// Solidity: function createEliteCard(string _tokenURI) returns(uint256)
func (_Cardpack *CardpackTransactorSession) CreateEliteCard(_tokenURI string) (*types.Transaction, error) {
	return _Cardpack.Contract.CreateEliteCard(&_Cardpack.TransactOpts, _tokenURI)
}

// CreatePremiumCard is a paid mutator transaction binding the contract method 0xffdcb746.
//
// Solidity: function createPremiumCard(string _tokenURI) returns(uint256)
func (_Cardpack *CardpackTransactor) CreatePremiumCard(opts *bind.TransactOpts, _tokenURI string) (*types.Transaction, error) {
	return _Cardpack.contract.Transact(opts, "createPremiumCard", _tokenURI)
}

// CreatePremiumCard is a paid mutator transaction binding the contract method 0xffdcb746.
//
// Solidity: function createPremiumCard(string _tokenURI) returns(uint256)
func (_Cardpack *CardpackSession) CreatePremiumCard(_tokenURI string) (*types.Transaction, error) {
	return _Cardpack.Contract.CreatePremiumCard(&_Cardpack.TransactOpts, _tokenURI)
}

// CreatePremiumCard is a paid mutator transaction binding the contract method 0xffdcb746.
//
// Solidity: function createPremiumCard(string _tokenURI) returns(uint256)
func (_Cardpack *CardpackTransactorSession) CreatePremiumCard(_tokenURI string) (*types.Transaction, error) {
	return _Cardpack.Contract.CreatePremiumCard(&_Cardpack.TransactOpts, _tokenURI)
}

// CreateStandardCard is a paid mutator transaction binding the contract method 0x9777d2a1.
//
// Solidity: function createStandardCard(string _tokenURI) returns(uint256)
func (_Cardpack *CardpackTransactor) CreateStandardCard(opts *bind.TransactOpts, _tokenURI string) (*types.Transaction, error) {
	return _Cardpack.contract.Transact(opts, "createStandardCard", _tokenURI)
}

// CreateStandardCard is a paid mutator transaction binding the contract method 0x9777d2a1.
//
// Solidity: function createStandardCard(string _tokenURI) returns(uint256)
func (_Cardpack *CardpackSession) CreateStandardCard(_tokenURI string) (*types.Transaction, error) {
	return _Cardpack.Contract.CreateStandardCard(&_Cardpack.TransactOpts, _tokenURI)
}

// CreateStandardCard is a paid mutator transaction binding the contract method 0x9777d2a1.
//
// Solidity: function createStandardCard(string _tokenURI) returns(uint256)
func (_Cardpack *CardpackTransactorSession) CreateStandardCard(_tokenURI string) (*types.Transaction, error) {
	return _Cardpack.Contract.CreateStandardCard(&_Cardpack.TransactOpts, _tokenURI)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Cardpack *CardpackTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cardpack.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Cardpack *CardpackSession) RenounceOwnership() (*types.Transaction, error) {
	return _Cardpack.Contract.RenounceOwnership(&_Cardpack.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Cardpack *CardpackTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Cardpack.Contract.RenounceOwnership(&_Cardpack.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Cardpack *CardpackTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Cardpack.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Cardpack *CardpackSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Cardpack.Contract.SafeTransferFrom(&_Cardpack.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Cardpack *CardpackTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Cardpack.Contract.SafeTransferFrom(&_Cardpack.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Cardpack *CardpackTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Cardpack.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Cardpack *CardpackSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Cardpack.Contract.SafeTransferFrom0(&_Cardpack.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Cardpack *CardpackTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Cardpack.Contract.SafeTransferFrom0(&_Cardpack.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Cardpack *CardpackTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Cardpack.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Cardpack *CardpackSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Cardpack.Contract.SetApprovalForAll(&_Cardpack.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Cardpack *CardpackTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Cardpack.Contract.SetApprovalForAll(&_Cardpack.TransactOpts, operator, approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Cardpack *CardpackTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Cardpack.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Cardpack *CardpackSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Cardpack.Contract.TransferFrom(&_Cardpack.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Cardpack *CardpackTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Cardpack.Contract.TransferFrom(&_Cardpack.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Cardpack *CardpackTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Cardpack.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Cardpack *CardpackSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Cardpack.Contract.TransferOwnership(&_Cardpack.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Cardpack *CardpackTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Cardpack.Contract.TransferOwnership(&_Cardpack.TransactOpts, newOwner)
}

// CardpackApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Cardpack contract.
type CardpackApprovalIterator struct {
	Event *CardpackApproval // Event containing the contract specifics and raw log

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
func (it *CardpackApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CardpackApproval)
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
		it.Event = new(CardpackApproval)
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
func (it *CardpackApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CardpackApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CardpackApproval represents a Approval event raised by the Cardpack contract.
type CardpackApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Cardpack *CardpackFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*CardpackApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Cardpack.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &CardpackApprovalIterator{contract: _Cardpack.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Cardpack *CardpackFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *CardpackApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Cardpack.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CardpackApproval)
				if err := _Cardpack.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Cardpack *CardpackFilterer) ParseApproval(log types.Log) (*CardpackApproval, error) {
	event := new(CardpackApproval)
	if err := _Cardpack.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CardpackApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Cardpack contract.
type CardpackApprovalForAllIterator struct {
	Event *CardpackApprovalForAll // Event containing the contract specifics and raw log

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
func (it *CardpackApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CardpackApprovalForAll)
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
		it.Event = new(CardpackApprovalForAll)
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
func (it *CardpackApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CardpackApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CardpackApprovalForAll represents a ApprovalForAll event raised by the Cardpack contract.
type CardpackApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Cardpack *CardpackFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*CardpackApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Cardpack.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &CardpackApprovalForAllIterator{contract: _Cardpack.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Cardpack *CardpackFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *CardpackApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Cardpack.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CardpackApprovalForAll)
				if err := _Cardpack.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Cardpack *CardpackFilterer) ParseApprovalForAll(log types.Log) (*CardpackApprovalForAll, error) {
	event := new(CardpackApprovalForAll)
	if err := _Cardpack.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CardpackCardPackCreatedIterator is returned from FilterCardPackCreated and is used to iterate over the raw logs and unpacked data for CardPackCreated events raised by the Cardpack contract.
type CardpackCardPackCreatedIterator struct {
	Event *CardpackCardPackCreated // Event containing the contract specifics and raw log

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
func (it *CardpackCardPackCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CardpackCardPackCreated)
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
		it.Event = new(CardpackCardPackCreated)
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
func (it *CardpackCardPackCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CardpackCardPackCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CardpackCardPackCreated represents a CardPackCreated event raised by the Cardpack contract.
type CardpackCardPackCreated struct {
	TokenId    *big.Int
	PackNumber *big.Int
	TokenURI   string
	TypeOf     common.Hash
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCardPackCreated is a free log retrieval operation binding the contract event 0xe63a3c01b5d7b0b39869d441ac33b298f5da66c069d89f99ef6da3ee1ed2193d.
//
// Solidity: event CardPackCreated(uint256 indexed tokenId, uint256 packNumber, string tokenURI, string indexed typeOf)
func (_Cardpack *CardpackFilterer) FilterCardPackCreated(opts *bind.FilterOpts, tokenId []*big.Int, typeOf []string) (*CardpackCardPackCreatedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	var typeOfRule []interface{}
	for _, typeOfItem := range typeOf {
		typeOfRule = append(typeOfRule, typeOfItem)
	}

	logs, sub, err := _Cardpack.contract.FilterLogs(opts, "CardPackCreated", tokenIdRule, typeOfRule)
	if err != nil {
		return nil, err
	}
	return &CardpackCardPackCreatedIterator{contract: _Cardpack.contract, event: "CardPackCreated", logs: logs, sub: sub}, nil
}

// WatchCardPackCreated is a free log subscription operation binding the contract event 0xe63a3c01b5d7b0b39869d441ac33b298f5da66c069d89f99ef6da3ee1ed2193d.
//
// Solidity: event CardPackCreated(uint256 indexed tokenId, uint256 packNumber, string tokenURI, string indexed typeOf)
func (_Cardpack *CardpackFilterer) WatchCardPackCreated(opts *bind.WatchOpts, sink chan<- *CardpackCardPackCreated, tokenId []*big.Int, typeOf []string) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	var typeOfRule []interface{}
	for _, typeOfItem := range typeOf {
		typeOfRule = append(typeOfRule, typeOfItem)
	}

	logs, sub, err := _Cardpack.contract.WatchLogs(opts, "CardPackCreated", tokenIdRule, typeOfRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CardpackCardPackCreated)
				if err := _Cardpack.contract.UnpackLog(event, "CardPackCreated", log); err != nil {
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

// ParseCardPackCreated is a log parse operation binding the contract event 0xe63a3c01b5d7b0b39869d441ac33b298f5da66c069d89f99ef6da3ee1ed2193d.
//
// Solidity: event CardPackCreated(uint256 indexed tokenId, uint256 packNumber, string tokenURI, string indexed typeOf)
func (_Cardpack *CardpackFilterer) ParseCardPackCreated(log types.Log) (*CardpackCardPackCreated, error) {
	event := new(CardpackCardPackCreated)
	if err := _Cardpack.contract.UnpackLog(event, "CardPackCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CardpackCardPackOpenedIterator is returned from FilterCardPackOpened and is used to iterate over the raw logs and unpacked data for CardPackOpened events raised by the Cardpack contract.
type CardpackCardPackOpenedIterator struct {
	Event *CardpackCardPackOpened // Event containing the contract specifics and raw log

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
func (it *CardpackCardPackOpenedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CardpackCardPackOpened)
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
		it.Event = new(CardpackCardPackOpened)
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
func (it *CardpackCardPackOpenedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CardpackCardPackOpenedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CardpackCardPackOpened represents a CardPackOpened event raised by the Cardpack contract.
type CardpackCardPackOpened struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterCardPackOpened is a free log retrieval operation binding the contract event 0xf3e8995f49b67b9b40d19d0ae915aff4ef00ecd661b45eaacdc47b3684279a34.
//
// Solidity: event CardPackOpened(uint256 indexed tokenId)
func (_Cardpack *CardpackFilterer) FilterCardPackOpened(opts *bind.FilterOpts, tokenId []*big.Int) (*CardpackCardPackOpenedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Cardpack.contract.FilterLogs(opts, "CardPackOpened", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &CardpackCardPackOpenedIterator{contract: _Cardpack.contract, event: "CardPackOpened", logs: logs, sub: sub}, nil
}

// WatchCardPackOpened is a free log subscription operation binding the contract event 0xf3e8995f49b67b9b40d19d0ae915aff4ef00ecd661b45eaacdc47b3684279a34.
//
// Solidity: event CardPackOpened(uint256 indexed tokenId)
func (_Cardpack *CardpackFilterer) WatchCardPackOpened(opts *bind.WatchOpts, sink chan<- *CardpackCardPackOpened, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Cardpack.contract.WatchLogs(opts, "CardPackOpened", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CardpackCardPackOpened)
				if err := _Cardpack.contract.UnpackLog(event, "CardPackOpened", log); err != nil {
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

// ParseCardPackOpened is a log parse operation binding the contract event 0xf3e8995f49b67b9b40d19d0ae915aff4ef00ecd661b45eaacdc47b3684279a34.
//
// Solidity: event CardPackOpened(uint256 indexed tokenId)
func (_Cardpack *CardpackFilterer) ParseCardPackOpened(log types.Log) (*CardpackCardPackOpened, error) {
	event := new(CardpackCardPackOpened)
	if err := _Cardpack.contract.UnpackLog(event, "CardPackOpened", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CardpackLimitChangedIterator is returned from FilterLimitChanged and is used to iterate over the raw logs and unpacked data for LimitChanged events raised by the Cardpack contract.
type CardpackLimitChangedIterator struct {
	Event *CardpackLimitChanged // Event containing the contract specifics and raw log

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
func (it *CardpackLimitChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CardpackLimitChanged)
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
		it.Event = new(CardpackLimitChanged)
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
func (it *CardpackLimitChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CardpackLimitChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CardpackLimitChanged represents a LimitChanged event raised by the Cardpack contract.
type CardpackLimitChanged struct {
	NewLimit *big.Int
	TypeOf   string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLimitChanged is a free log retrieval operation binding the contract event 0x9fc072c98b6eb02f368bc5c6765f963c699b663e23b0bbf06712102ce81d0b65.
//
// Solidity: event LimitChanged(uint256 newLimit, string typeOf)
func (_Cardpack *CardpackFilterer) FilterLimitChanged(opts *bind.FilterOpts) (*CardpackLimitChangedIterator, error) {

	logs, sub, err := _Cardpack.contract.FilterLogs(opts, "LimitChanged")
	if err != nil {
		return nil, err
	}
	return &CardpackLimitChangedIterator{contract: _Cardpack.contract, event: "LimitChanged", logs: logs, sub: sub}, nil
}

// WatchLimitChanged is a free log subscription operation binding the contract event 0x9fc072c98b6eb02f368bc5c6765f963c699b663e23b0bbf06712102ce81d0b65.
//
// Solidity: event LimitChanged(uint256 newLimit, string typeOf)
func (_Cardpack *CardpackFilterer) WatchLimitChanged(opts *bind.WatchOpts, sink chan<- *CardpackLimitChanged) (event.Subscription, error) {

	logs, sub, err := _Cardpack.contract.WatchLogs(opts, "LimitChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CardpackLimitChanged)
				if err := _Cardpack.contract.UnpackLog(event, "LimitChanged", log); err != nil {
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

// ParseLimitChanged is a log parse operation binding the contract event 0x9fc072c98b6eb02f368bc5c6765f963c699b663e23b0bbf06712102ce81d0b65.
//
// Solidity: event LimitChanged(uint256 newLimit, string typeOf)
func (_Cardpack *CardpackFilterer) ParseLimitChanged(log types.Log) (*CardpackLimitChanged, error) {
	event := new(CardpackLimitChanged)
	if err := _Cardpack.contract.UnpackLog(event, "LimitChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CardpackMinterContractChangedIterator is returned from FilterMinterContractChanged and is used to iterate over the raw logs and unpacked data for MinterContractChanged events raised by the Cardpack contract.
type CardpackMinterContractChangedIterator struct {
	Event *CardpackMinterContractChanged // Event containing the contract specifics and raw log

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
func (it *CardpackMinterContractChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CardpackMinterContractChanged)
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
		it.Event = new(CardpackMinterContractChanged)
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
func (it *CardpackMinterContractChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CardpackMinterContractChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CardpackMinterContractChanged represents a MinterContractChanged event raised by the Cardpack contract.
type CardpackMinterContractChanged struct {
	PreviousMinter common.Address
	NewMinter      common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterMinterContractChanged is a free log retrieval operation binding the contract event 0x114d06280373bf8a65ca81a7d5ec748e7747747d23dbaa3e41f3c4ea46223079.
//
// Solidity: event MinterContractChanged(address indexed previousMinter, address indexed newMinter)
func (_Cardpack *CardpackFilterer) FilterMinterContractChanged(opts *bind.FilterOpts, previousMinter []common.Address, newMinter []common.Address) (*CardpackMinterContractChangedIterator, error) {

	var previousMinterRule []interface{}
	for _, previousMinterItem := range previousMinter {
		previousMinterRule = append(previousMinterRule, previousMinterItem)
	}
	var newMinterRule []interface{}
	for _, newMinterItem := range newMinter {
		newMinterRule = append(newMinterRule, newMinterItem)
	}

	logs, sub, err := _Cardpack.contract.FilterLogs(opts, "MinterContractChanged", previousMinterRule, newMinterRule)
	if err != nil {
		return nil, err
	}
	return &CardpackMinterContractChangedIterator{contract: _Cardpack.contract, event: "MinterContractChanged", logs: logs, sub: sub}, nil
}

// WatchMinterContractChanged is a free log subscription operation binding the contract event 0x114d06280373bf8a65ca81a7d5ec748e7747747d23dbaa3e41f3c4ea46223079.
//
// Solidity: event MinterContractChanged(address indexed previousMinter, address indexed newMinter)
func (_Cardpack *CardpackFilterer) WatchMinterContractChanged(opts *bind.WatchOpts, sink chan<- *CardpackMinterContractChanged, previousMinter []common.Address, newMinter []common.Address) (event.Subscription, error) {

	var previousMinterRule []interface{}
	for _, previousMinterItem := range previousMinter {
		previousMinterRule = append(previousMinterRule, previousMinterItem)
	}
	var newMinterRule []interface{}
	for _, newMinterItem := range newMinter {
		newMinterRule = append(newMinterRule, newMinterItem)
	}

	logs, sub, err := _Cardpack.contract.WatchLogs(opts, "MinterContractChanged", previousMinterRule, newMinterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CardpackMinterContractChanged)
				if err := _Cardpack.contract.UnpackLog(event, "MinterContractChanged", log); err != nil {
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

// ParseMinterContractChanged is a log parse operation binding the contract event 0x114d06280373bf8a65ca81a7d5ec748e7747747d23dbaa3e41f3c4ea46223079.
//
// Solidity: event MinterContractChanged(address indexed previousMinter, address indexed newMinter)
func (_Cardpack *CardpackFilterer) ParseMinterContractChanged(log types.Log) (*CardpackMinterContractChanged, error) {
	event := new(CardpackMinterContractChanged)
	if err := _Cardpack.contract.UnpackLog(event, "MinterContractChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CardpackOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Cardpack contract.
type CardpackOwnershipTransferredIterator struct {
	Event *CardpackOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *CardpackOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CardpackOwnershipTransferred)
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
		it.Event = new(CardpackOwnershipTransferred)
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
func (it *CardpackOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CardpackOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CardpackOwnershipTransferred represents a OwnershipTransferred event raised by the Cardpack contract.
type CardpackOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Cardpack *CardpackFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*CardpackOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Cardpack.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CardpackOwnershipTransferredIterator{contract: _Cardpack.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Cardpack *CardpackFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CardpackOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Cardpack.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CardpackOwnershipTransferred)
				if err := _Cardpack.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Cardpack *CardpackFilterer) ParseOwnershipTransferred(log types.Log) (*CardpackOwnershipTransferred, error) {
	event := new(CardpackOwnershipTransferred)
	if err := _Cardpack.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CardpackTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Cardpack contract.
type CardpackTransferIterator struct {
	Event *CardpackTransfer // Event containing the contract specifics and raw log

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
func (it *CardpackTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CardpackTransfer)
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
		it.Event = new(CardpackTransfer)
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
func (it *CardpackTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CardpackTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CardpackTransfer represents a Transfer event raised by the Cardpack contract.
type CardpackTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Cardpack *CardpackFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*CardpackTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Cardpack.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &CardpackTransferIterator{contract: _Cardpack.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Cardpack *CardpackFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *CardpackTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Cardpack.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CardpackTransfer)
				if err := _Cardpack.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Cardpack *CardpackFilterer) ParseTransfer(log types.Log) (*CardpackTransfer, error) {
	event := new(CardpackTransfer)
	if err := _Cardpack.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
