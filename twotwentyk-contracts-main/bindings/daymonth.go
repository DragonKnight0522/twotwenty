// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package daymonth

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

// DaymonthMetaData contains all meta data concerning the Daymonth contract.
var DaymonthMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_collectionId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousCrafter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newCrafter\",\"type\":\"address\"}],\"name\":\"CraftingContractChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"collectionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"}],\"name\":\"DayMonthMinted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousMinter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newMinter\",\"type\":\"address\"}],\"name\":\"MinterContractChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"PartCrafted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newCrafter\",\"type\":\"address\"}],\"name\":\"changeCrafter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newMinter\",\"type\":\"address\"}],\"name\":\"changeMinter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"changeToCrafted\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"collectionId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"isCrafted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_tokenURI\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"mintDayMonth\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// DaymonthABI is the input ABI used to generate the binding from.
// Deprecated: Use DaymonthMetaData.ABI instead.
var DaymonthABI = DaymonthMetaData.ABI

// Daymonth is an auto generated Go binding around an Ethereum contract.
type Daymonth struct {
	DaymonthCaller     // Read-only binding to the contract
	DaymonthTransactor // Write-only binding to the contract
	DaymonthFilterer   // Log filterer for contract events
}

// DaymonthCaller is an auto generated read-only Go binding around an Ethereum contract.
type DaymonthCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DaymonthTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DaymonthTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DaymonthFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DaymonthFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DaymonthSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DaymonthSession struct {
	Contract     *Daymonth         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DaymonthCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DaymonthCallerSession struct {
	Contract *DaymonthCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// DaymonthTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DaymonthTransactorSession struct {
	Contract     *DaymonthTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// DaymonthRaw is an auto generated low-level Go binding around an Ethereum contract.
type DaymonthRaw struct {
	Contract *Daymonth // Generic contract binding to access the raw methods on
}

// DaymonthCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DaymonthCallerRaw struct {
	Contract *DaymonthCaller // Generic read-only contract binding to access the raw methods on
}

// DaymonthTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DaymonthTransactorRaw struct {
	Contract *DaymonthTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDaymonth creates a new instance of Daymonth, bound to a specific deployed contract.
func NewDaymonth(address common.Address, backend bind.ContractBackend) (*Daymonth, error) {
	contract, err := bindDaymonth(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Daymonth{DaymonthCaller: DaymonthCaller{contract: contract}, DaymonthTransactor: DaymonthTransactor{contract: contract}, DaymonthFilterer: DaymonthFilterer{contract: contract}}, nil
}

// NewDaymonthCaller creates a new read-only instance of Daymonth, bound to a specific deployed contract.
func NewDaymonthCaller(address common.Address, caller bind.ContractCaller) (*DaymonthCaller, error) {
	contract, err := bindDaymonth(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DaymonthCaller{contract: contract}, nil
}

// NewDaymonthTransactor creates a new write-only instance of Daymonth, bound to a specific deployed contract.
func NewDaymonthTransactor(address common.Address, transactor bind.ContractTransactor) (*DaymonthTransactor, error) {
	contract, err := bindDaymonth(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DaymonthTransactor{contract: contract}, nil
}

// NewDaymonthFilterer creates a new log filterer instance of Daymonth, bound to a specific deployed contract.
func NewDaymonthFilterer(address common.Address, filterer bind.ContractFilterer) (*DaymonthFilterer, error) {
	contract, err := bindDaymonth(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DaymonthFilterer{contract: contract}, nil
}

// bindDaymonth binds a generic wrapper to an already deployed contract.
func bindDaymonth(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DaymonthMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Daymonth *DaymonthRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Daymonth.Contract.DaymonthCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Daymonth *DaymonthRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Daymonth.Contract.DaymonthTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Daymonth *DaymonthRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Daymonth.Contract.DaymonthTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Daymonth *DaymonthCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Daymonth.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Daymonth *DaymonthTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Daymonth.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Daymonth *DaymonthTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Daymonth.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Daymonth *DaymonthCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Daymonth.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Daymonth *DaymonthSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Daymonth.Contract.BalanceOf(&_Daymonth.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Daymonth *DaymonthCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Daymonth.Contract.BalanceOf(&_Daymonth.CallOpts, owner)
}

// CollectionId is a free data retrieval call binding the contract method 0x3d26bb67.
//
// Solidity: function collectionId() view returns(uint256)
func (_Daymonth *DaymonthCaller) CollectionId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Daymonth.contract.Call(opts, &out, "collectionId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CollectionId is a free data retrieval call binding the contract method 0x3d26bb67.
//
// Solidity: function collectionId() view returns(uint256)
func (_Daymonth *DaymonthSession) CollectionId() (*big.Int, error) {
	return _Daymonth.Contract.CollectionId(&_Daymonth.CallOpts)
}

// CollectionId is a free data retrieval call binding the contract method 0x3d26bb67.
//
// Solidity: function collectionId() view returns(uint256)
func (_Daymonth *DaymonthCallerSession) CollectionId() (*big.Int, error) {
	return _Daymonth.Contract.CollectionId(&_Daymonth.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Daymonth *DaymonthCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Daymonth.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Daymonth *DaymonthSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Daymonth.Contract.GetApproved(&_Daymonth.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Daymonth *DaymonthCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Daymonth.Contract.GetApproved(&_Daymonth.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Daymonth *DaymonthCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Daymonth.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Daymonth *DaymonthSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Daymonth.Contract.IsApprovedForAll(&_Daymonth.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Daymonth *DaymonthCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Daymonth.Contract.IsApprovedForAll(&_Daymonth.CallOpts, owner, operator)
}

// IsCrafted is a free data retrieval call binding the contract method 0xfd96527d.
//
// Solidity: function isCrafted(uint256 tokenId) view returns(bool)
func (_Daymonth *DaymonthCaller) IsCrafted(opts *bind.CallOpts, tokenId *big.Int) (bool, error) {
	var out []interface{}
	err := _Daymonth.contract.Call(opts, &out, "isCrafted", tokenId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsCrafted is a free data retrieval call binding the contract method 0xfd96527d.
//
// Solidity: function isCrafted(uint256 tokenId) view returns(bool)
func (_Daymonth *DaymonthSession) IsCrafted(tokenId *big.Int) (bool, error) {
	return _Daymonth.Contract.IsCrafted(&_Daymonth.CallOpts, tokenId)
}

// IsCrafted is a free data retrieval call binding the contract method 0xfd96527d.
//
// Solidity: function isCrafted(uint256 tokenId) view returns(bool)
func (_Daymonth *DaymonthCallerSession) IsCrafted(tokenId *big.Int) (bool, error) {
	return _Daymonth.Contract.IsCrafted(&_Daymonth.CallOpts, tokenId)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Daymonth *DaymonthCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Daymonth.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Daymonth *DaymonthSession) Name() (string, error) {
	return _Daymonth.Contract.Name(&_Daymonth.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Daymonth *DaymonthCallerSession) Name() (string, error) {
	return _Daymonth.Contract.Name(&_Daymonth.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Daymonth *DaymonthCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Daymonth.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Daymonth *DaymonthSession) Owner() (common.Address, error) {
	return _Daymonth.Contract.Owner(&_Daymonth.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Daymonth *DaymonthCallerSession) Owner() (common.Address, error) {
	return _Daymonth.Contract.Owner(&_Daymonth.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Daymonth *DaymonthCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Daymonth.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Daymonth *DaymonthSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Daymonth.Contract.OwnerOf(&_Daymonth.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Daymonth *DaymonthCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Daymonth.Contract.OwnerOf(&_Daymonth.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Daymonth *DaymonthCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Daymonth.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Daymonth *DaymonthSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Daymonth.Contract.SupportsInterface(&_Daymonth.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Daymonth *DaymonthCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Daymonth.Contract.SupportsInterface(&_Daymonth.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Daymonth *DaymonthCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Daymonth.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Daymonth *DaymonthSession) Symbol() (string, error) {
	return _Daymonth.Contract.Symbol(&_Daymonth.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Daymonth *DaymonthCallerSession) Symbol() (string, error) {
	return _Daymonth.Contract.Symbol(&_Daymonth.CallOpts)
}

// TokenCounter is a free data retrieval call binding the contract method 0xd082e381.
//
// Solidity: function tokenCounter() view returns(uint256)
func (_Daymonth *DaymonthCaller) TokenCounter(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Daymonth.contract.Call(opts, &out, "tokenCounter")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenCounter is a free data retrieval call binding the contract method 0xd082e381.
//
// Solidity: function tokenCounter() view returns(uint256)
func (_Daymonth *DaymonthSession) TokenCounter() (*big.Int, error) {
	return _Daymonth.Contract.TokenCounter(&_Daymonth.CallOpts)
}

// TokenCounter is a free data retrieval call binding the contract method 0xd082e381.
//
// Solidity: function tokenCounter() view returns(uint256)
func (_Daymonth *DaymonthCallerSession) TokenCounter() (*big.Int, error) {
	return _Daymonth.Contract.TokenCounter(&_Daymonth.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Daymonth *DaymonthCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Daymonth.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Daymonth *DaymonthSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Daymonth.Contract.TokenURI(&_Daymonth.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Daymonth *DaymonthCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Daymonth.Contract.TokenURI(&_Daymonth.CallOpts, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Daymonth *DaymonthTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Daymonth.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Daymonth *DaymonthSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Daymonth.Contract.Approve(&_Daymonth.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Daymonth *DaymonthTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Daymonth.Contract.Approve(&_Daymonth.TransactOpts, to, tokenId)
}

// ChangeCrafter is a paid mutator transaction binding the contract method 0xe7d3b069.
//
// Solidity: function changeCrafter(address _newCrafter) returns()
func (_Daymonth *DaymonthTransactor) ChangeCrafter(opts *bind.TransactOpts, _newCrafter common.Address) (*types.Transaction, error) {
	return _Daymonth.contract.Transact(opts, "changeCrafter", _newCrafter)
}

// ChangeCrafter is a paid mutator transaction binding the contract method 0xe7d3b069.
//
// Solidity: function changeCrafter(address _newCrafter) returns()
func (_Daymonth *DaymonthSession) ChangeCrafter(_newCrafter common.Address) (*types.Transaction, error) {
	return _Daymonth.Contract.ChangeCrafter(&_Daymonth.TransactOpts, _newCrafter)
}

// ChangeCrafter is a paid mutator transaction binding the contract method 0xe7d3b069.
//
// Solidity: function changeCrafter(address _newCrafter) returns()
func (_Daymonth *DaymonthTransactorSession) ChangeCrafter(_newCrafter common.Address) (*types.Transaction, error) {
	return _Daymonth.Contract.ChangeCrafter(&_Daymonth.TransactOpts, _newCrafter)
}

// ChangeMinter is a paid mutator transaction binding the contract method 0x2c4d4d18.
//
// Solidity: function changeMinter(address _newMinter) returns()
func (_Daymonth *DaymonthTransactor) ChangeMinter(opts *bind.TransactOpts, _newMinter common.Address) (*types.Transaction, error) {
	return _Daymonth.contract.Transact(opts, "changeMinter", _newMinter)
}

// ChangeMinter is a paid mutator transaction binding the contract method 0x2c4d4d18.
//
// Solidity: function changeMinter(address _newMinter) returns()
func (_Daymonth *DaymonthSession) ChangeMinter(_newMinter common.Address) (*types.Transaction, error) {
	return _Daymonth.Contract.ChangeMinter(&_Daymonth.TransactOpts, _newMinter)
}

// ChangeMinter is a paid mutator transaction binding the contract method 0x2c4d4d18.
//
// Solidity: function changeMinter(address _newMinter) returns()
func (_Daymonth *DaymonthTransactorSession) ChangeMinter(_newMinter common.Address) (*types.Transaction, error) {
	return _Daymonth.Contract.ChangeMinter(&_Daymonth.TransactOpts, _newMinter)
}

// ChangeToCrafted is a paid mutator transaction binding the contract method 0xe6b60d9b.
//
// Solidity: function changeToCrafted(uint256 tokenId) returns()
func (_Daymonth *DaymonthTransactor) ChangeToCrafted(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _Daymonth.contract.Transact(opts, "changeToCrafted", tokenId)
}

// ChangeToCrafted is a paid mutator transaction binding the contract method 0xe6b60d9b.
//
// Solidity: function changeToCrafted(uint256 tokenId) returns()
func (_Daymonth *DaymonthSession) ChangeToCrafted(tokenId *big.Int) (*types.Transaction, error) {
	return _Daymonth.Contract.ChangeToCrafted(&_Daymonth.TransactOpts, tokenId)
}

// ChangeToCrafted is a paid mutator transaction binding the contract method 0xe6b60d9b.
//
// Solidity: function changeToCrafted(uint256 tokenId) returns()
func (_Daymonth *DaymonthTransactorSession) ChangeToCrafted(tokenId *big.Int) (*types.Transaction, error) {
	return _Daymonth.Contract.ChangeToCrafted(&_Daymonth.TransactOpts, tokenId)
}

// MintDayMonth is a paid mutator transaction binding the contract method 0xdfe2ead4.
//
// Solidity: function mintDayMonth(string _tokenURI, address _owner) returns(uint256)
func (_Daymonth *DaymonthTransactor) MintDayMonth(opts *bind.TransactOpts, _tokenURI string, _owner common.Address) (*types.Transaction, error) {
	return _Daymonth.contract.Transact(opts, "mintDayMonth", _tokenURI, _owner)
}

// MintDayMonth is a paid mutator transaction binding the contract method 0xdfe2ead4.
//
// Solidity: function mintDayMonth(string _tokenURI, address _owner) returns(uint256)
func (_Daymonth *DaymonthSession) MintDayMonth(_tokenURI string, _owner common.Address) (*types.Transaction, error) {
	return _Daymonth.Contract.MintDayMonth(&_Daymonth.TransactOpts, _tokenURI, _owner)
}

// MintDayMonth is a paid mutator transaction binding the contract method 0xdfe2ead4.
//
// Solidity: function mintDayMonth(string _tokenURI, address _owner) returns(uint256)
func (_Daymonth *DaymonthTransactorSession) MintDayMonth(_tokenURI string, _owner common.Address) (*types.Transaction, error) {
	return _Daymonth.Contract.MintDayMonth(&_Daymonth.TransactOpts, _tokenURI, _owner)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Daymonth *DaymonthTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Daymonth.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Daymonth *DaymonthSession) RenounceOwnership() (*types.Transaction, error) {
	return _Daymonth.Contract.RenounceOwnership(&_Daymonth.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Daymonth *DaymonthTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Daymonth.Contract.RenounceOwnership(&_Daymonth.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Daymonth *DaymonthTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Daymonth.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Daymonth *DaymonthSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Daymonth.Contract.SafeTransferFrom(&_Daymonth.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Daymonth *DaymonthTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Daymonth.Contract.SafeTransferFrom(&_Daymonth.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Daymonth *DaymonthTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Daymonth.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Daymonth *DaymonthSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Daymonth.Contract.SafeTransferFrom0(&_Daymonth.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Daymonth *DaymonthTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Daymonth.Contract.SafeTransferFrom0(&_Daymonth.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Daymonth *DaymonthTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Daymonth.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Daymonth *DaymonthSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Daymonth.Contract.SetApprovalForAll(&_Daymonth.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Daymonth *DaymonthTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Daymonth.Contract.SetApprovalForAll(&_Daymonth.TransactOpts, operator, approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Daymonth *DaymonthTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Daymonth.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Daymonth *DaymonthSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Daymonth.Contract.TransferFrom(&_Daymonth.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Daymonth *DaymonthTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Daymonth.Contract.TransferFrom(&_Daymonth.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Daymonth *DaymonthTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Daymonth.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Daymonth *DaymonthSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Daymonth.Contract.TransferOwnership(&_Daymonth.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Daymonth *DaymonthTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Daymonth.Contract.TransferOwnership(&_Daymonth.TransactOpts, newOwner)
}

// DaymonthApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Daymonth contract.
type DaymonthApprovalIterator struct {
	Event *DaymonthApproval // Event containing the contract specifics and raw log

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
func (it *DaymonthApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DaymonthApproval)
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
		it.Event = new(DaymonthApproval)
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
func (it *DaymonthApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DaymonthApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DaymonthApproval represents a Approval event raised by the Daymonth contract.
type DaymonthApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Daymonth *DaymonthFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*DaymonthApprovalIterator, error) {

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

	logs, sub, err := _Daymonth.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &DaymonthApprovalIterator{contract: _Daymonth.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Daymonth *DaymonthFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *DaymonthApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Daymonth.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DaymonthApproval)
				if err := _Daymonth.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Daymonth *DaymonthFilterer) ParseApproval(log types.Log) (*DaymonthApproval, error) {
	event := new(DaymonthApproval)
	if err := _Daymonth.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DaymonthApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Daymonth contract.
type DaymonthApprovalForAllIterator struct {
	Event *DaymonthApprovalForAll // Event containing the contract specifics and raw log

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
func (it *DaymonthApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DaymonthApprovalForAll)
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
		it.Event = new(DaymonthApprovalForAll)
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
func (it *DaymonthApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DaymonthApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DaymonthApprovalForAll represents a ApprovalForAll event raised by the Daymonth contract.
type DaymonthApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Daymonth *DaymonthFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*DaymonthApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Daymonth.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &DaymonthApprovalForAllIterator{contract: _Daymonth.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Daymonth *DaymonthFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *DaymonthApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Daymonth.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DaymonthApprovalForAll)
				if err := _Daymonth.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_Daymonth *DaymonthFilterer) ParseApprovalForAll(log types.Log) (*DaymonthApprovalForAll, error) {
	event := new(DaymonthApprovalForAll)
	if err := _Daymonth.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DaymonthCraftingContractChangedIterator is returned from FilterCraftingContractChanged and is used to iterate over the raw logs and unpacked data for CraftingContractChanged events raised by the Daymonth contract.
type DaymonthCraftingContractChangedIterator struct {
	Event *DaymonthCraftingContractChanged // Event containing the contract specifics and raw log

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
func (it *DaymonthCraftingContractChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DaymonthCraftingContractChanged)
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
		it.Event = new(DaymonthCraftingContractChanged)
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
func (it *DaymonthCraftingContractChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DaymonthCraftingContractChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DaymonthCraftingContractChanged represents a CraftingContractChanged event raised by the Daymonth contract.
type DaymonthCraftingContractChanged struct {
	PreviousCrafter common.Address
	NewCrafter      common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterCraftingContractChanged is a free log retrieval operation binding the contract event 0x83149e90b909cde0e8099cb6998203f8d8dce5fa29151688382650f5637d7d88.
//
// Solidity: event CraftingContractChanged(address indexed previousCrafter, address indexed newCrafter)
func (_Daymonth *DaymonthFilterer) FilterCraftingContractChanged(opts *bind.FilterOpts, previousCrafter []common.Address, newCrafter []common.Address) (*DaymonthCraftingContractChangedIterator, error) {

	var previousCrafterRule []interface{}
	for _, previousCrafterItem := range previousCrafter {
		previousCrafterRule = append(previousCrafterRule, previousCrafterItem)
	}
	var newCrafterRule []interface{}
	for _, newCrafterItem := range newCrafter {
		newCrafterRule = append(newCrafterRule, newCrafterItem)
	}

	logs, sub, err := _Daymonth.contract.FilterLogs(opts, "CraftingContractChanged", previousCrafterRule, newCrafterRule)
	if err != nil {
		return nil, err
	}
	return &DaymonthCraftingContractChangedIterator{contract: _Daymonth.contract, event: "CraftingContractChanged", logs: logs, sub: sub}, nil
}

// WatchCraftingContractChanged is a free log subscription operation binding the contract event 0x83149e90b909cde0e8099cb6998203f8d8dce5fa29151688382650f5637d7d88.
//
// Solidity: event CraftingContractChanged(address indexed previousCrafter, address indexed newCrafter)
func (_Daymonth *DaymonthFilterer) WatchCraftingContractChanged(opts *bind.WatchOpts, sink chan<- *DaymonthCraftingContractChanged, previousCrafter []common.Address, newCrafter []common.Address) (event.Subscription, error) {

	var previousCrafterRule []interface{}
	for _, previousCrafterItem := range previousCrafter {
		previousCrafterRule = append(previousCrafterRule, previousCrafterItem)
	}
	var newCrafterRule []interface{}
	for _, newCrafterItem := range newCrafter {
		newCrafterRule = append(newCrafterRule, newCrafterItem)
	}

	logs, sub, err := _Daymonth.contract.WatchLogs(opts, "CraftingContractChanged", previousCrafterRule, newCrafterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DaymonthCraftingContractChanged)
				if err := _Daymonth.contract.UnpackLog(event, "CraftingContractChanged", log); err != nil {
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

// ParseCraftingContractChanged is a log parse operation binding the contract event 0x83149e90b909cde0e8099cb6998203f8d8dce5fa29151688382650f5637d7d88.
//
// Solidity: event CraftingContractChanged(address indexed previousCrafter, address indexed newCrafter)
func (_Daymonth *DaymonthFilterer) ParseCraftingContractChanged(log types.Log) (*DaymonthCraftingContractChanged, error) {
	event := new(DaymonthCraftingContractChanged)
	if err := _Daymonth.contract.UnpackLog(event, "CraftingContractChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DaymonthDayMonthMintedIterator is returned from FilterDayMonthMinted and is used to iterate over the raw logs and unpacked data for DayMonthMinted events raised by the Daymonth contract.
type DaymonthDayMonthMintedIterator struct {
	Event *DaymonthDayMonthMinted // Event containing the contract specifics and raw log

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
func (it *DaymonthDayMonthMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DaymonthDayMonthMinted)
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
		it.Event = new(DaymonthDayMonthMinted)
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
func (it *DaymonthDayMonthMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DaymonthDayMonthMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DaymonthDayMonthMinted represents a DayMonthMinted event raised by the Daymonth contract.
type DaymonthDayMonthMinted struct {
	CollectionId *big.Int
	TokenId      *big.Int
	TokenURI     string
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDayMonthMinted is a free log retrieval operation binding the contract event 0x68db35241ab1a261832cc79ecbc964485f72ffb226ad6d5e1713d63659757de2.
//
// Solidity: event DayMonthMinted(uint256 collectionId, uint256 indexed tokenId, string tokenURI)
func (_Daymonth *DaymonthFilterer) FilterDayMonthMinted(opts *bind.FilterOpts, tokenId []*big.Int) (*DaymonthDayMonthMintedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Daymonth.contract.FilterLogs(opts, "DayMonthMinted", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &DaymonthDayMonthMintedIterator{contract: _Daymonth.contract, event: "DayMonthMinted", logs: logs, sub: sub}, nil
}

// WatchDayMonthMinted is a free log subscription operation binding the contract event 0x68db35241ab1a261832cc79ecbc964485f72ffb226ad6d5e1713d63659757de2.
//
// Solidity: event DayMonthMinted(uint256 collectionId, uint256 indexed tokenId, string tokenURI)
func (_Daymonth *DaymonthFilterer) WatchDayMonthMinted(opts *bind.WatchOpts, sink chan<- *DaymonthDayMonthMinted, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Daymonth.contract.WatchLogs(opts, "DayMonthMinted", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DaymonthDayMonthMinted)
				if err := _Daymonth.contract.UnpackLog(event, "DayMonthMinted", log); err != nil {
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

// ParseDayMonthMinted is a log parse operation binding the contract event 0x68db35241ab1a261832cc79ecbc964485f72ffb226ad6d5e1713d63659757de2.
//
// Solidity: event DayMonthMinted(uint256 collectionId, uint256 indexed tokenId, string tokenURI)
func (_Daymonth *DaymonthFilterer) ParseDayMonthMinted(log types.Log) (*DaymonthDayMonthMinted, error) {
	event := new(DaymonthDayMonthMinted)
	if err := _Daymonth.contract.UnpackLog(event, "DayMonthMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DaymonthMinterContractChangedIterator is returned from FilterMinterContractChanged and is used to iterate over the raw logs and unpacked data for MinterContractChanged events raised by the Daymonth contract.
type DaymonthMinterContractChangedIterator struct {
	Event *DaymonthMinterContractChanged // Event containing the contract specifics and raw log

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
func (it *DaymonthMinterContractChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DaymonthMinterContractChanged)
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
		it.Event = new(DaymonthMinterContractChanged)
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
func (it *DaymonthMinterContractChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DaymonthMinterContractChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DaymonthMinterContractChanged represents a MinterContractChanged event raised by the Daymonth contract.
type DaymonthMinterContractChanged struct {
	PreviousMinter common.Address
	NewMinter      common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterMinterContractChanged is a free log retrieval operation binding the contract event 0x114d06280373bf8a65ca81a7d5ec748e7747747d23dbaa3e41f3c4ea46223079.
//
// Solidity: event MinterContractChanged(address indexed previousMinter, address indexed newMinter)
func (_Daymonth *DaymonthFilterer) FilterMinterContractChanged(opts *bind.FilterOpts, previousMinter []common.Address, newMinter []common.Address) (*DaymonthMinterContractChangedIterator, error) {

	var previousMinterRule []interface{}
	for _, previousMinterItem := range previousMinter {
		previousMinterRule = append(previousMinterRule, previousMinterItem)
	}
	var newMinterRule []interface{}
	for _, newMinterItem := range newMinter {
		newMinterRule = append(newMinterRule, newMinterItem)
	}

	logs, sub, err := _Daymonth.contract.FilterLogs(opts, "MinterContractChanged", previousMinterRule, newMinterRule)
	if err != nil {
		return nil, err
	}
	return &DaymonthMinterContractChangedIterator{contract: _Daymonth.contract, event: "MinterContractChanged", logs: logs, sub: sub}, nil
}

// WatchMinterContractChanged is a free log subscription operation binding the contract event 0x114d06280373bf8a65ca81a7d5ec748e7747747d23dbaa3e41f3c4ea46223079.
//
// Solidity: event MinterContractChanged(address indexed previousMinter, address indexed newMinter)
func (_Daymonth *DaymonthFilterer) WatchMinterContractChanged(opts *bind.WatchOpts, sink chan<- *DaymonthMinterContractChanged, previousMinter []common.Address, newMinter []common.Address) (event.Subscription, error) {

	var previousMinterRule []interface{}
	for _, previousMinterItem := range previousMinter {
		previousMinterRule = append(previousMinterRule, previousMinterItem)
	}
	var newMinterRule []interface{}
	for _, newMinterItem := range newMinter {
		newMinterRule = append(newMinterRule, newMinterItem)
	}

	logs, sub, err := _Daymonth.contract.WatchLogs(opts, "MinterContractChanged", previousMinterRule, newMinterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DaymonthMinterContractChanged)
				if err := _Daymonth.contract.UnpackLog(event, "MinterContractChanged", log); err != nil {
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
func (_Daymonth *DaymonthFilterer) ParseMinterContractChanged(log types.Log) (*DaymonthMinterContractChanged, error) {
	event := new(DaymonthMinterContractChanged)
	if err := _Daymonth.contract.UnpackLog(event, "MinterContractChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DaymonthOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Daymonth contract.
type DaymonthOwnershipTransferredIterator struct {
	Event *DaymonthOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *DaymonthOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DaymonthOwnershipTransferred)
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
		it.Event = new(DaymonthOwnershipTransferred)
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
func (it *DaymonthOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DaymonthOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DaymonthOwnershipTransferred represents a OwnershipTransferred event raised by the Daymonth contract.
type DaymonthOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Daymonth *DaymonthFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DaymonthOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Daymonth.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DaymonthOwnershipTransferredIterator{contract: _Daymonth.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Daymonth *DaymonthFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DaymonthOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Daymonth.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DaymonthOwnershipTransferred)
				if err := _Daymonth.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Daymonth *DaymonthFilterer) ParseOwnershipTransferred(log types.Log) (*DaymonthOwnershipTransferred, error) {
	event := new(DaymonthOwnershipTransferred)
	if err := _Daymonth.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DaymonthPartCraftedIterator is returned from FilterPartCrafted and is used to iterate over the raw logs and unpacked data for PartCrafted events raised by the Daymonth contract.
type DaymonthPartCraftedIterator struct {
	Event *DaymonthPartCrafted // Event containing the contract specifics and raw log

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
func (it *DaymonthPartCraftedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DaymonthPartCrafted)
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
		it.Event = new(DaymonthPartCrafted)
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
func (it *DaymonthPartCraftedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DaymonthPartCraftedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DaymonthPartCrafted represents a PartCrafted event raised by the Daymonth contract.
type DaymonthPartCrafted struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPartCrafted is a free log retrieval operation binding the contract event 0x71094059908fc1a5d676895c3ab6ec0e8957a09dc7a09914b791a2be4dc7fb15.
//
// Solidity: event PartCrafted(uint256 indexed tokenId)
func (_Daymonth *DaymonthFilterer) FilterPartCrafted(opts *bind.FilterOpts, tokenId []*big.Int) (*DaymonthPartCraftedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Daymonth.contract.FilterLogs(opts, "PartCrafted", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &DaymonthPartCraftedIterator{contract: _Daymonth.contract, event: "PartCrafted", logs: logs, sub: sub}, nil
}

// WatchPartCrafted is a free log subscription operation binding the contract event 0x71094059908fc1a5d676895c3ab6ec0e8957a09dc7a09914b791a2be4dc7fb15.
//
// Solidity: event PartCrafted(uint256 indexed tokenId)
func (_Daymonth *DaymonthFilterer) WatchPartCrafted(opts *bind.WatchOpts, sink chan<- *DaymonthPartCrafted, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Daymonth.contract.WatchLogs(opts, "PartCrafted", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DaymonthPartCrafted)
				if err := _Daymonth.contract.UnpackLog(event, "PartCrafted", log); err != nil {
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

// ParsePartCrafted is a log parse operation binding the contract event 0x71094059908fc1a5d676895c3ab6ec0e8957a09dc7a09914b791a2be4dc7fb15.
//
// Solidity: event PartCrafted(uint256 indexed tokenId)
func (_Daymonth *DaymonthFilterer) ParsePartCrafted(log types.Log) (*DaymonthPartCrafted, error) {
	event := new(DaymonthPartCrafted)
	if err := _Daymonth.contract.UnpackLog(event, "PartCrafted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DaymonthTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Daymonth contract.
type DaymonthTransferIterator struct {
	Event *DaymonthTransfer // Event containing the contract specifics and raw log

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
func (it *DaymonthTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DaymonthTransfer)
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
		it.Event = new(DaymonthTransfer)
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
func (it *DaymonthTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DaymonthTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DaymonthTransfer represents a Transfer event raised by the Daymonth contract.
type DaymonthTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Daymonth *DaymonthFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*DaymonthTransferIterator, error) {

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

	logs, sub, err := _Daymonth.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &DaymonthTransferIterator{contract: _Daymonth.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Daymonth *DaymonthFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *DaymonthTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Daymonth.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DaymonthTransfer)
				if err := _Daymonth.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Daymonth *DaymonthFilterer) ParseTransfer(log types.Log) (*DaymonthTransfer, error) {
	event := new(DaymonthTransfer)
	if err := _Daymonth.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
