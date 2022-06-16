// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// BTokenErrorCodesABI is the input ABI used to generate the binding from.
const BTokenErrorCodesABI = "[{\"inputs\":[],\"name\":\"BTOKEN_BURN_AMOUNT_EXCEEDS_SUPPLY\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BTOKEN_CONTRACTS_DISALLOWED_DEPOSITS\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BTOKEN_DEPOSIT_AMOUNT_ZERO\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BTOKEN_DEPOSIT_BURN_FAIL\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BTOKEN_INVALID_BALANCE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BTOKEN_INVALID_BURN_AMOUNT\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BTOKEN_INVALID_DEPOSIT_ID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BTOKEN_MARKET_SPREAD_TOO_LOW\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BTOKEN_MINIMUM_BURN_NOT_MET\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BTOKEN_MINIMUM_MINT_NOT_MET\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BTOKEN_MINT_INSUFFICIENT_ETH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BTOKEN_SPLIT_VALUE_SUM_ERROR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// BTokenErrorCodes is an auto generated Go binding around an Ethereum contract.
type BTokenErrorCodes struct {
	BTokenErrorCodesCaller     // Read-only binding to the contract
	BTokenErrorCodesTransactor // Write-only binding to the contract
	BTokenErrorCodesFilterer   // Log filterer for contract events
}

// BTokenErrorCodesCaller is an auto generated read-only Go binding around an Ethereum contract.
type BTokenErrorCodesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BTokenErrorCodesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BTokenErrorCodesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BTokenErrorCodesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BTokenErrorCodesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BTokenErrorCodesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BTokenErrorCodesSession struct {
	Contract     *BTokenErrorCodes // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BTokenErrorCodesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BTokenErrorCodesCallerSession struct {
	Contract *BTokenErrorCodesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// BTokenErrorCodesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BTokenErrorCodesTransactorSession struct {
	Contract     *BTokenErrorCodesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// BTokenErrorCodesRaw is an auto generated low-level Go binding around an Ethereum contract.
type BTokenErrorCodesRaw struct {
	Contract *BTokenErrorCodes // Generic contract binding to access the raw methods on
}

// BTokenErrorCodesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BTokenErrorCodesCallerRaw struct {
	Contract *BTokenErrorCodesCaller // Generic read-only contract binding to access the raw methods on
}

// BTokenErrorCodesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BTokenErrorCodesTransactorRaw struct {
	Contract *BTokenErrorCodesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBTokenErrorCodes creates a new instance of BTokenErrorCodes, bound to a specific deployed contract.
func NewBTokenErrorCodes(address common.Address, backend bind.ContractBackend) (*BTokenErrorCodes, error) {
	contract, err := bindBTokenErrorCodes(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BTokenErrorCodes{BTokenErrorCodesCaller: BTokenErrorCodesCaller{contract: contract}, BTokenErrorCodesTransactor: BTokenErrorCodesTransactor{contract: contract}, BTokenErrorCodesFilterer: BTokenErrorCodesFilterer{contract: contract}}, nil
}

// NewBTokenErrorCodesCaller creates a new read-only instance of BTokenErrorCodes, bound to a specific deployed contract.
func NewBTokenErrorCodesCaller(address common.Address, caller bind.ContractCaller) (*BTokenErrorCodesCaller, error) {
	contract, err := bindBTokenErrorCodes(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BTokenErrorCodesCaller{contract: contract}, nil
}

// NewBTokenErrorCodesTransactor creates a new write-only instance of BTokenErrorCodes, bound to a specific deployed contract.
func NewBTokenErrorCodesTransactor(address common.Address, transactor bind.ContractTransactor) (*BTokenErrorCodesTransactor, error) {
	contract, err := bindBTokenErrorCodes(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BTokenErrorCodesTransactor{contract: contract}, nil
}

// NewBTokenErrorCodesFilterer creates a new log filterer instance of BTokenErrorCodes, bound to a specific deployed contract.
func NewBTokenErrorCodesFilterer(address common.Address, filterer bind.ContractFilterer) (*BTokenErrorCodesFilterer, error) {
	contract, err := bindBTokenErrorCodes(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BTokenErrorCodesFilterer{contract: contract}, nil
}

// bindBTokenErrorCodes binds a generic wrapper to an already deployed contract.
func bindBTokenErrorCodes(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BTokenErrorCodesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BTokenErrorCodes *BTokenErrorCodesRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BTokenErrorCodes.Contract.BTokenErrorCodesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BTokenErrorCodes *BTokenErrorCodesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BTokenErrorCodes.Contract.BTokenErrorCodesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BTokenErrorCodes *BTokenErrorCodesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BTokenErrorCodes.Contract.BTokenErrorCodesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BTokenErrorCodes *BTokenErrorCodesCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BTokenErrorCodes.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BTokenErrorCodes *BTokenErrorCodesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BTokenErrorCodes.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BTokenErrorCodes *BTokenErrorCodesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BTokenErrorCodes.Contract.contract.Transact(opts, method, params...)
}

// BTOKENBURNAMOUNTEXCEEDSSUPPLY is a paid mutator transaction binding the contract method 0x1e13bfbb.
//
// Solidity: function BTOKEN_BURN_AMOUNT_EXCEEDS_SUPPLY() returns(bytes32)
func (_BTokenErrorCodes *BTokenErrorCodesTransactor) BTOKENBURNAMOUNTEXCEEDSSUPPLY(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BTokenErrorCodes.contract.Transact(opts, "BTOKEN_BURN_AMOUNT_EXCEEDS_SUPPLY")
}

// BTOKENBURNAMOUNTEXCEEDSSUPPLY is a paid mutator transaction binding the contract method 0x1e13bfbb.
//
// Solidity: function BTOKEN_BURN_AMOUNT_EXCEEDS_SUPPLY() returns(bytes32)
func (_BTokenErrorCodes *BTokenErrorCodesSession) BTOKENBURNAMOUNTEXCEEDSSUPPLY() (*types.Transaction, error) {
	return _BTokenErrorCodes.Contract.BTOKENBURNAMOUNTEXCEEDSSUPPLY(&_BTokenErrorCodes.TransactOpts)
}

// BTOKENBURNAMOUNTEXCEEDSSUPPLY is a paid mutator transaction binding the contract method 0x1e13bfbb.
//
// Solidity: function BTOKEN_BURN_AMOUNT_EXCEEDS_SUPPLY() returns(bytes32)
func (_BTokenErrorCodes *BTokenErrorCodesTransactorSession) BTOKENBURNAMOUNTEXCEEDSSUPPLY() (*types.Transaction, error) {
	return _BTokenErrorCodes.Contract.BTOKENBURNAMOUNTEXCEEDSSUPPLY(&_BTokenErrorCodes.TransactOpts)
}

// BTOKENCONTRACTSDISALLOWEDDEPOSITS is a paid mutator transaction binding the contract method 0x0e19d024.
//
// Solidity: function BTOKEN_CONTRACTS_DISALLOWED_DEPOSITS() returns(bytes32)
func (_BTokenErrorCodes *BTokenErrorCodesTransactor) BTOKENCONTRACTSDISALLOWEDDEPOSITS(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BTokenErrorCodes.contract.Transact(opts, "BTOKEN_CONTRACTS_DISALLOWED_DEPOSITS")
}

// BTOKENCONTRACTSDISALLOWEDDEPOSITS is a paid mutator transaction binding the contract method 0x0e19d024.
//
// Solidity: function BTOKEN_CONTRACTS_DISALLOWED_DEPOSITS() returns(bytes32)
func (_BTokenErrorCodes *BTokenErrorCodesSession) BTOKENCONTRACTSDISALLOWEDDEPOSITS() (*types.Transaction, error) {
	return _BTokenErrorCodes.Contract.BTOKENCONTRACTSDISALLOWEDDEPOSITS(&_BTokenErrorCodes.TransactOpts)
}

// BTOKENCONTRACTSDISALLOWEDDEPOSITS is a paid mutator transaction binding the contract method 0x0e19d024.
//
// Solidity: function BTOKEN_CONTRACTS_DISALLOWED_DEPOSITS() returns(bytes32)
func (_BTokenErrorCodes *BTokenErrorCodesTransactorSession) BTOKENCONTRACTSDISALLOWEDDEPOSITS() (*types.Transaction, error) {
	return _BTokenErrorCodes.Contract.BTOKENCONTRACTSDISALLOWEDDEPOSITS(&_BTokenErrorCodes.TransactOpts)
}

// BTOKENDEPOSITAMOUNTZERO is a paid mutator transaction binding the contract method 0xb949d485.
//
// Solidity: function BTOKEN_DEPOSIT_AMOUNT_ZERO() returns(bytes32)
func (_BTokenErrorCodes *BTokenErrorCodesTransactor) BTOKENDEPOSITAMOUNTZERO(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BTokenErrorCodes.contract.Transact(opts, "BTOKEN_DEPOSIT_AMOUNT_ZERO")
}

// BTOKENDEPOSITAMOUNTZERO is a paid mutator transaction binding the contract method 0xb949d485.
//
// Solidity: function BTOKEN_DEPOSIT_AMOUNT_ZERO() returns(bytes32)
func (_BTokenErrorCodes *BTokenErrorCodesSession) BTOKENDEPOSITAMOUNTZERO() (*types.Transaction, error) {
	return _BTokenErrorCodes.Contract.BTOKENDEPOSITAMOUNTZERO(&_BTokenErrorCodes.TransactOpts)
}

// BTOKENDEPOSITAMOUNTZERO is a paid mutator transaction binding the contract method 0xb949d485.
//
// Solidity: function BTOKEN_DEPOSIT_AMOUNT_ZERO() returns(bytes32)
func (_BTokenErrorCodes *BTokenErrorCodesTransactorSession) BTOKENDEPOSITAMOUNTZERO() (*types.Transaction, error) {
	return _BTokenErrorCodes.Contract.BTOKENDEPOSITAMOUNTZERO(&_BTokenErrorCodes.TransactOpts)
}

// BTOKENDEPOSITBURNFAIL is a paid mutator transaction binding the contract method 0xfe4a969a.
//
// Solidity: function BTOKEN_DEPOSIT_BURN_FAIL() returns(bytes32)
func (_BTokenErrorCodes *BTokenErrorCodesTransactor) BTOKENDEPOSITBURNFAIL(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BTokenErrorCodes.contract.Transact(opts, "BTOKEN_DEPOSIT_BURN_FAIL")
}

// BTOKENDEPOSITBURNFAIL is a paid mutator transaction binding the contract method 0xfe4a969a.
//
// Solidity: function BTOKEN_DEPOSIT_BURN_FAIL() returns(bytes32)
func (_BTokenErrorCodes *BTokenErrorCodesSession) BTOKENDEPOSITBURNFAIL() (*types.Transaction, error) {
	return _BTokenErrorCodes.Contract.BTOKENDEPOSITBURNFAIL(&_BTokenErrorCodes.TransactOpts)
}

// BTOKENDEPOSITBURNFAIL is a paid mutator transaction binding the contract method 0xfe4a969a.
//
// Solidity: function BTOKEN_DEPOSIT_BURN_FAIL() returns(bytes32)
func (_BTokenErrorCodes *BTokenErrorCodesTransactorSession) BTOKENDEPOSITBURNFAIL() (*types.Transaction, error) {
	return _BTokenErrorCodes.Contract.BTOKENDEPOSITBURNFAIL(&_BTokenErrorCodes.TransactOpts)
}

// BTOKENINVALIDBALANCE is a paid mutator transaction binding the contract method 0xa3d600f1.
//
// Solidity: function BTOKEN_INVALID_BALANCE() returns(bytes32)
func (_BTokenErrorCodes *BTokenErrorCodesTransactor) BTOKENINVALIDBALANCE(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BTokenErrorCodes.contract.Transact(opts, "BTOKEN_INVALID_BALANCE")
}

// BTOKENINVALIDBALANCE is a paid mutator transaction binding the contract method 0xa3d600f1.
//
// Solidity: function BTOKEN_INVALID_BALANCE() returns(bytes32)
func (_BTokenErrorCodes *BTokenErrorCodesSession) BTOKENINVALIDBALANCE() (*types.Transaction, error) {
	return _BTokenErrorCodes.Contract.BTOKENINVALIDBALANCE(&_BTokenErrorCodes.TransactOpts)
}

// BTOKENINVALIDBALANCE is a paid mutator transaction binding the contract method 0xa3d600f1.
//
// Solidity: function BTOKEN_INVALID_BALANCE() returns(bytes32)
func (_BTokenErrorCodes *BTokenErrorCodesTransactorSession) BTOKENINVALIDBALANCE() (*types.Transaction, error) {
	return _BTokenErrorCodes.Contract.BTOKENINVALIDBALANCE(&_BTokenErrorCodes.TransactOpts)
}

// BTOKENINVALIDBURNAMOUNT is a paid mutator transaction binding the contract method 0xf87e114e.
//
// Solidity: function BTOKEN_INVALID_BURN_AMOUNT() returns(bytes32)
func (_BTokenErrorCodes *BTokenErrorCodesTransactor) BTOKENINVALIDBURNAMOUNT(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BTokenErrorCodes.contract.Transact(opts, "BTOKEN_INVALID_BURN_AMOUNT")
}

// BTOKENINVALIDBURNAMOUNT is a paid mutator transaction binding the contract method 0xf87e114e.
//
// Solidity: function BTOKEN_INVALID_BURN_AMOUNT() returns(bytes32)
func (_BTokenErrorCodes *BTokenErrorCodesSession) BTOKENINVALIDBURNAMOUNT() (*types.Transaction, error) {
	return _BTokenErrorCodes.Contract.BTOKENINVALIDBURNAMOUNT(&_BTokenErrorCodes.TransactOpts)
}

// BTOKENINVALIDBURNAMOUNT is a paid mutator transaction binding the contract method 0xf87e114e.
//
// Solidity: function BTOKEN_INVALID_BURN_AMOUNT() returns(bytes32)
func (_BTokenErrorCodes *BTokenErrorCodesTransactorSession) BTOKENINVALIDBURNAMOUNT() (*types.Transaction, error) {
	return _BTokenErrorCodes.Contract.BTOKENINVALIDBURNAMOUNT(&_BTokenErrorCodes.TransactOpts)
}

// BTOKENINVALIDDEPOSITID is a paid mutator transaction binding the contract method 0x3dd42816.
//
// Solidity: function BTOKEN_INVALID_DEPOSIT_ID() returns(bytes32)
func (_BTokenErrorCodes *BTokenErrorCodesTransactor) BTOKENINVALIDDEPOSITID(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BTokenErrorCodes.contract.Transact(opts, "BTOKEN_INVALID_DEPOSIT_ID")
}

// BTOKENINVALIDDEPOSITID is a paid mutator transaction binding the contract method 0x3dd42816.
//
// Solidity: function BTOKEN_INVALID_DEPOSIT_ID() returns(bytes32)
func (_BTokenErrorCodes *BTokenErrorCodesSession) BTOKENINVALIDDEPOSITID() (*types.Transaction, error) {
	return _BTokenErrorCodes.Contract.BTOKENINVALIDDEPOSITID(&_BTokenErrorCodes.TransactOpts)
}

// BTOKENINVALIDDEPOSITID is a paid mutator transaction binding the contract method 0x3dd42816.
//
// Solidity: function BTOKEN_INVALID_DEPOSIT_ID() returns(bytes32)
func (_BTokenErrorCodes *BTokenErrorCodesTransactorSession) BTOKENINVALIDDEPOSITID() (*types.Transaction, error) {
	return _BTokenErrorCodes.Contract.BTOKENINVALIDDEPOSITID(&_BTokenErrorCodes.TransactOpts)
}

// BTOKENMARKETSPREADTOOLOW is a paid mutator transaction binding the contract method 0x17be6132.
//
// Solidity: function BTOKEN_MARKET_SPREAD_TOO_LOW() returns(bytes32)
func (_BTokenErrorCodes *BTokenErrorCodesTransactor) BTOKENMARKETSPREADTOOLOW(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BTokenErrorCodes.contract.Transact(opts, "BTOKEN_MARKET_SPREAD_TOO_LOW")
}

// BTOKENMARKETSPREADTOOLOW is a paid mutator transaction binding the contract method 0x17be6132.
//
// Solidity: function BTOKEN_MARKET_SPREAD_TOO_LOW() returns(bytes32)
func (_BTokenErrorCodes *BTokenErrorCodesSession) BTOKENMARKETSPREADTOOLOW() (*types.Transaction, error) {
	return _BTokenErrorCodes.Contract.BTOKENMARKETSPREADTOOLOW(&_BTokenErrorCodes.TransactOpts)
}

// BTOKENMARKETSPREADTOOLOW is a paid mutator transaction binding the contract method 0x17be6132.
//
// Solidity: function BTOKEN_MARKET_SPREAD_TOO_LOW() returns(bytes32)
func (_BTokenErrorCodes *BTokenErrorCodesTransactorSession) BTOKENMARKETSPREADTOOLOW() (*types.Transaction, error) {
	return _BTokenErrorCodes.Contract.BTOKENMARKETSPREADTOOLOW(&_BTokenErrorCodes.TransactOpts)
}

// BTOKENMINIMUMBURNNOTMET is a paid mutator transaction binding the contract method 0x42e745e4.
//
// Solidity: function BTOKEN_MINIMUM_BURN_NOT_MET() returns(bytes32)
func (_BTokenErrorCodes *BTokenErrorCodesTransactor) BTOKENMINIMUMBURNNOTMET(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BTokenErrorCodes.contract.Transact(opts, "BTOKEN_MINIMUM_BURN_NOT_MET")
}

// BTOKENMINIMUMBURNNOTMET is a paid mutator transaction binding the contract method 0x42e745e4.
//
// Solidity: function BTOKEN_MINIMUM_BURN_NOT_MET() returns(bytes32)
func (_BTokenErrorCodes *BTokenErrorCodesSession) BTOKENMINIMUMBURNNOTMET() (*types.Transaction, error) {
	return _BTokenErrorCodes.Contract.BTOKENMINIMUMBURNNOTMET(&_BTokenErrorCodes.TransactOpts)
}

// BTOKENMINIMUMBURNNOTMET is a paid mutator transaction binding the contract method 0x42e745e4.
//
// Solidity: function BTOKEN_MINIMUM_BURN_NOT_MET() returns(bytes32)
func (_BTokenErrorCodes *BTokenErrorCodesTransactorSession) BTOKENMINIMUMBURNNOTMET() (*types.Transaction, error) {
	return _BTokenErrorCodes.Contract.BTOKENMINIMUMBURNNOTMET(&_BTokenErrorCodes.TransactOpts)
}

// BTOKENMINIMUMMINTNOTMET is a paid mutator transaction binding the contract method 0xfc45f4cf.
//
// Solidity: function BTOKEN_MINIMUM_MINT_NOT_MET() returns(bytes32)
func (_BTokenErrorCodes *BTokenErrorCodesTransactor) BTOKENMINIMUMMINTNOTMET(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BTokenErrorCodes.contract.Transact(opts, "BTOKEN_MINIMUM_MINT_NOT_MET")
}

// BTOKENMINIMUMMINTNOTMET is a paid mutator transaction binding the contract method 0xfc45f4cf.
//
// Solidity: function BTOKEN_MINIMUM_MINT_NOT_MET() returns(bytes32)
func (_BTokenErrorCodes *BTokenErrorCodesSession) BTOKENMINIMUMMINTNOTMET() (*types.Transaction, error) {
	return _BTokenErrorCodes.Contract.BTOKENMINIMUMMINTNOTMET(&_BTokenErrorCodes.TransactOpts)
}

// BTOKENMINIMUMMINTNOTMET is a paid mutator transaction binding the contract method 0xfc45f4cf.
//
// Solidity: function BTOKEN_MINIMUM_MINT_NOT_MET() returns(bytes32)
func (_BTokenErrorCodes *BTokenErrorCodesTransactorSession) BTOKENMINIMUMMINTNOTMET() (*types.Transaction, error) {
	return _BTokenErrorCodes.Contract.BTOKENMINIMUMMINTNOTMET(&_BTokenErrorCodes.TransactOpts)
}

// BTOKENMINTINSUFFICIENTETH is a paid mutator transaction binding the contract method 0x2810d142.
//
// Solidity: function BTOKEN_MINT_INSUFFICIENT_ETH() returns(bytes32)
func (_BTokenErrorCodes *BTokenErrorCodesTransactor) BTOKENMINTINSUFFICIENTETH(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BTokenErrorCodes.contract.Transact(opts, "BTOKEN_MINT_INSUFFICIENT_ETH")
}

// BTOKENMINTINSUFFICIENTETH is a paid mutator transaction binding the contract method 0x2810d142.
//
// Solidity: function BTOKEN_MINT_INSUFFICIENT_ETH() returns(bytes32)
func (_BTokenErrorCodes *BTokenErrorCodesSession) BTOKENMINTINSUFFICIENTETH() (*types.Transaction, error) {
	return _BTokenErrorCodes.Contract.BTOKENMINTINSUFFICIENTETH(&_BTokenErrorCodes.TransactOpts)
}

// BTOKENMINTINSUFFICIENTETH is a paid mutator transaction binding the contract method 0x2810d142.
//
// Solidity: function BTOKEN_MINT_INSUFFICIENT_ETH() returns(bytes32)
func (_BTokenErrorCodes *BTokenErrorCodesTransactorSession) BTOKENMINTINSUFFICIENTETH() (*types.Transaction, error) {
	return _BTokenErrorCodes.Contract.BTOKENMINTINSUFFICIENTETH(&_BTokenErrorCodes.TransactOpts)
}

// BTOKENSPLITVALUESUMERROR is a paid mutator transaction binding the contract method 0xc7f53c4f.
//
// Solidity: function BTOKEN_SPLIT_VALUE_SUM_ERROR() returns(bytes32)
func (_BTokenErrorCodes *BTokenErrorCodesTransactor) BTOKENSPLITVALUESUMERROR(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BTokenErrorCodes.contract.Transact(opts, "BTOKEN_SPLIT_VALUE_SUM_ERROR")
}

// BTOKENSPLITVALUESUMERROR is a paid mutator transaction binding the contract method 0xc7f53c4f.
//
// Solidity: function BTOKEN_SPLIT_VALUE_SUM_ERROR() returns(bytes32)
func (_BTokenErrorCodes *BTokenErrorCodesSession) BTOKENSPLITVALUESUMERROR() (*types.Transaction, error) {
	return _BTokenErrorCodes.Contract.BTOKENSPLITVALUESUMERROR(&_BTokenErrorCodes.TransactOpts)
}

// BTOKENSPLITVALUESUMERROR is a paid mutator transaction binding the contract method 0xc7f53c4f.
//
// Solidity: function BTOKEN_SPLIT_VALUE_SUM_ERROR() returns(bytes32)
func (_BTokenErrorCodes *BTokenErrorCodesTransactorSession) BTOKENSPLITVALUESUMERROR() (*types.Transaction, error) {
	return _BTokenErrorCodes.Contract.BTOKENSPLITVALUESUMERROR(&_BTokenErrorCodes.TransactOpts)
}
