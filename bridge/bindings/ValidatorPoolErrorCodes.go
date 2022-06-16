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

// ValidatorPoolErrorCodesABI is the input ABI used to generate the binding from.
const ValidatorPoolErrorCodesABI = "[{\"inputs\":[],\"name\":\"VALIDATORPOOL_ADDRESS_ALREADY_VALIDATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VALIDATORPOOL_ADDRESS_NOT_ACCUSABLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VALIDATORPOOL_ADDRESS_NOT_VALIDATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VALIDATORPOOL_CALLER_NOT_VALIDATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VALIDATORPOOL_CONSENSUS_RUNNING\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VALIDATORPOOL_DISHONEST_VALIDATOR_NOT_ACCUSABLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VALIDATORPOOL_ETHDKG_ROUND_RUNNING\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VALIDATORPOOL_ETH_BALANCE_CHANGED\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VALIDATORPOOL_FACTORY_SHOULD_OWN_POSITION\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VALIDATORPOOL_INSUFFICIENT_FUNDS_IN_STAKE_POSITION\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VALIDATORPOOL_INVALID_INDEX\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VALIDATORPOOL_MAX_VALIDATORS_MET\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VALIDATORPOOL_MINIMUM_STAKE_NOT_MET\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VALIDATORPOOL_MIN_BLOCK_INTERVAL_NOT_MET\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VALIDATORPOOL_ONLY_CONTRACTS_ALLOWED\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VALIDATORPOOL_PAYOUT_TOO_LOW\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VALIDATORPOOL_PROFITS_ONLY_CLAIMABLE_DURING_CONSENSUS\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VALIDATORPOOL_REGISTRATION_PARAMETER_LENGTH_MISMATCH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VALIDATORPOOL_SENDER_NOT_IN_EXITING_QUEUE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VALIDATORPOOL_TOKEN_BALANCE_CHANGED\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VALIDATORPOOL_VALIDATORS_GREATER_THAN_AVAILABLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VALIDATORPOOL_WAITING_PERIOD_NOT_MET\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ValidatorPoolErrorCodes is an auto generated Go binding around an Ethereum contract.
type ValidatorPoolErrorCodes struct {
	ValidatorPoolErrorCodesCaller     // Read-only binding to the contract
	ValidatorPoolErrorCodesTransactor // Write-only binding to the contract
	ValidatorPoolErrorCodesFilterer   // Log filterer for contract events
}

// ValidatorPoolErrorCodesCaller is an auto generated read-only Go binding around an Ethereum contract.
type ValidatorPoolErrorCodesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorPoolErrorCodesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ValidatorPoolErrorCodesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorPoolErrorCodesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ValidatorPoolErrorCodesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorPoolErrorCodesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ValidatorPoolErrorCodesSession struct {
	Contract     *ValidatorPoolErrorCodes // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ValidatorPoolErrorCodesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ValidatorPoolErrorCodesCallerSession struct {
	Contract *ValidatorPoolErrorCodesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// ValidatorPoolErrorCodesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ValidatorPoolErrorCodesTransactorSession struct {
	Contract     *ValidatorPoolErrorCodesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// ValidatorPoolErrorCodesRaw is an auto generated low-level Go binding around an Ethereum contract.
type ValidatorPoolErrorCodesRaw struct {
	Contract *ValidatorPoolErrorCodes // Generic contract binding to access the raw methods on
}

// ValidatorPoolErrorCodesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ValidatorPoolErrorCodesCallerRaw struct {
	Contract *ValidatorPoolErrorCodesCaller // Generic read-only contract binding to access the raw methods on
}

// ValidatorPoolErrorCodesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ValidatorPoolErrorCodesTransactorRaw struct {
	Contract *ValidatorPoolErrorCodesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewValidatorPoolErrorCodes creates a new instance of ValidatorPoolErrorCodes, bound to a specific deployed contract.
func NewValidatorPoolErrorCodes(address common.Address, backend bind.ContractBackend) (*ValidatorPoolErrorCodes, error) {
	contract, err := bindValidatorPoolErrorCodes(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ValidatorPoolErrorCodes{ValidatorPoolErrorCodesCaller: ValidatorPoolErrorCodesCaller{contract: contract}, ValidatorPoolErrorCodesTransactor: ValidatorPoolErrorCodesTransactor{contract: contract}, ValidatorPoolErrorCodesFilterer: ValidatorPoolErrorCodesFilterer{contract: contract}}, nil
}

// NewValidatorPoolErrorCodesCaller creates a new read-only instance of ValidatorPoolErrorCodes, bound to a specific deployed contract.
func NewValidatorPoolErrorCodesCaller(address common.Address, caller bind.ContractCaller) (*ValidatorPoolErrorCodesCaller, error) {
	contract, err := bindValidatorPoolErrorCodes(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorPoolErrorCodesCaller{contract: contract}, nil
}

// NewValidatorPoolErrorCodesTransactor creates a new write-only instance of ValidatorPoolErrorCodes, bound to a specific deployed contract.
func NewValidatorPoolErrorCodesTransactor(address common.Address, transactor bind.ContractTransactor) (*ValidatorPoolErrorCodesTransactor, error) {
	contract, err := bindValidatorPoolErrorCodes(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorPoolErrorCodesTransactor{contract: contract}, nil
}

// NewValidatorPoolErrorCodesFilterer creates a new log filterer instance of ValidatorPoolErrorCodes, bound to a specific deployed contract.
func NewValidatorPoolErrorCodesFilterer(address common.Address, filterer bind.ContractFilterer) (*ValidatorPoolErrorCodesFilterer, error) {
	contract, err := bindValidatorPoolErrorCodes(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ValidatorPoolErrorCodesFilterer{contract: contract}, nil
}

// bindValidatorPoolErrorCodes binds a generic wrapper to an already deployed contract.
func bindValidatorPoolErrorCodes(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ValidatorPoolErrorCodesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ValidatorPoolErrorCodes *ValidatorPoolErrorCodesRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ValidatorPoolErrorCodes.Contract.ValidatorPoolErrorCodesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ValidatorPoolErrorCodes *ValidatorPoolErrorCodesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorPoolErrorCodes.Contract.ValidatorPoolErrorCodesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ValidatorPoolErrorCodes *ValidatorPoolErrorCodesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ValidatorPoolErrorCodes.Contract.ValidatorPoolErrorCodesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ValidatorPoolErrorCodes *ValidatorPoolErrorCodesCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ValidatorPoolErrorCodes.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ValidatorPoolErrorCodes *ValidatorPoolErrorCodesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorPoolErrorCodes.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ValidatorPoolErrorCodes *ValidatorPoolErrorCodesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ValidatorPoolErrorCodes.Contract.contract.Transact(opts, method, params...)
}

// VALIDATORPOOLADDRESSALREADYVALIDATOR is a paid mutator transaction binding the contract method 0xb8e6dc4d.
//
// Solidity: function VALIDATORPOOL_ADDRESS_ALREADY_VALIDATOR() returns(bytes32)
func (_ValidatorPoolErrorCodes *ValidatorPoolErrorCodesTransactor) VALIDATORPOOLADDRESSALREADYVALIDATOR(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorPoolErrorCodes.contract.Transact(opts, "VALIDATORPOOL_ADDRESS_ALREADY_VALIDATOR")
}

// VALIDATORPOOLADDRESSALREADYVALIDATOR is a paid mutator transaction binding the contract method 0xb8e6dc4d.
//
// Solidity: function VALIDATORPOOL_ADDRESS_ALREADY_VALIDATOR() returns(bytes32)
func (_ValidatorPoolErrorCodes *ValidatorPoolErrorCodesSession) VALIDATORPOOLADDRESSALREADYVALIDATOR() (*types.Transaction, error) {
	return _ValidatorPoolErrorCodes.Contract.VALIDATORPOOLADDRESSALREADYVALIDATOR(&_ValidatorPoolErrorCodes.TransactOpts)
}

// VALIDATORPOOLADDRESSALREADYVALIDATOR is a paid mutator transaction binding the contract method 0xb8e6dc4d.
//
// Solidity: function VALIDATORPOOL_ADDRESS_ALREADY_VALIDATOR() returns(bytes32)
func (_ValidatorPoolErrorCodes *ValidatorPoolErrorCodesTransactorSession) VALIDATORPOOLADDRESSALREADYVALIDATOR() (*types.Transaction, error) {
	return _ValidatorPoolErrorCodes.Contract.VALIDATORPOOLADDRESSALREADYVALIDATOR(&_ValidatorPoolErrorCodes.TransactOpts)
}

// VALIDATORPOOLADDRESSNOTACCUSABLE is a paid mutator transaction binding the contract method 0x0e144d69.
//
// Solidity: function VALIDATORPOOL_ADDRESS_NOT_ACCUSABLE() returns(bytes32)
func (_ValidatorPoolErrorCodes *ValidatorPoolErrorCodesTransactor) VALIDATORPOOLADDRESSNOTACCUSABLE(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorPoolErrorCodes.contract.Transact(opts, "VALIDATORPOOL_ADDRESS_NOT_ACCUSABLE")
}

// VALIDATORPOOLADDRESSNOTACCUSABLE is a paid mutator transaction binding the contract method 0x0e144d69.
//
// Solidity: function VALIDATORPOOL_ADDRESS_NOT_ACCUSABLE() returns(bytes32)
func (_ValidatorPoolErrorCodes *ValidatorPoolErrorCodesSession) VALIDATORPOOLADDRESSNOTACCUSABLE() (*types.Transaction, error) {
	return _ValidatorPoolErrorCodes.Contract.VALIDATORPOOLADDRESSNOTACCUSABLE(&_ValidatorPoolErrorCodes.TransactOpts)
}

// VALIDATORPOOLADDRESSNOTACCUSABLE is a paid mutator transaction binding the contract method 0x0e144d69.
//
// Solidity: function VALIDATORPOOL_ADDRESS_NOT_ACCUSABLE() returns(bytes32)
func (_ValidatorPoolErrorCodes *ValidatorPoolErrorCodesTransactorSession) VALIDATORPOOLADDRESSNOTACCUSABLE() (*types.Transaction, error) {
	return _ValidatorPoolErrorCodes.Contract.VALIDATORPOOLADDRESSNOTACCUSABLE(&_ValidatorPoolErrorCodes.TransactOpts)
}

// VALIDATORPOOLADDRESSNOTVALIDATOR is a paid mutator transaction binding the contract method 0x4ce78d00.
//
// Solidity: function VALIDATORPOOL_ADDRESS_NOT_VALIDATOR() returns(bytes32)
func (_ValidatorPoolErrorCodes *ValidatorPoolErrorCodesTransactor) VALIDATORPOOLADDRESSNOTVALIDATOR(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorPoolErrorCodes.contract.Transact(opts, "VALIDATORPOOL_ADDRESS_NOT_VALIDATOR")
}

// VALIDATORPOOLADDRESSNOTVALIDATOR is a paid mutator transaction binding the contract method 0x4ce78d00.
//
// Solidity: function VALIDATORPOOL_ADDRESS_NOT_VALIDATOR() returns(bytes32)
func (_ValidatorPoolErrorCodes *ValidatorPoolErrorCodesSession) VALIDATORPOOLADDRESSNOTVALIDATOR() (*types.Transaction, error) {
	return _ValidatorPoolErrorCodes.Contract.VALIDATORPOOLADDRESSNOTVALIDATOR(&_ValidatorPoolErrorCodes.TransactOpts)
}

// VALIDATORPOOLADDRESSNOTVALIDATOR is a paid mutator transaction binding the contract method 0x4ce78d00.
//
// Solidity: function VALIDATORPOOL_ADDRESS_NOT_VALIDATOR() returns(bytes32)
func (_ValidatorPoolErrorCodes *ValidatorPoolErrorCodesTransactorSession) VALIDATORPOOLADDRESSNOTVALIDATOR() (*types.Transaction, error) {
	return _ValidatorPoolErrorCodes.Contract.VALIDATORPOOLADDRESSNOTVALIDATOR(&_ValidatorPoolErrorCodes.TransactOpts)
}

// VALIDATORPOOLCALLERNOTVALIDATOR is a paid mutator transaction binding the contract method 0xdf6f3c26.
//
// Solidity: function VALIDATORPOOL_CALLER_NOT_VALIDATOR() returns(bytes32)
func (_ValidatorPoolErrorCodes *ValidatorPoolErrorCodesTransactor) VALIDATORPOOLCALLERNOTVALIDATOR(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorPoolErrorCodes.contract.Transact(opts, "VALIDATORPOOL_CALLER_NOT_VALIDATOR")
}

// VALIDATORPOOLCALLERNOTVALIDATOR is a paid mutator transaction binding the contract method 0xdf6f3c26.
//
// Solidity: function VALIDATORPOOL_CALLER_NOT_VALIDATOR() returns(bytes32)
func (_ValidatorPoolErrorCodes *ValidatorPoolErrorCodesSession) VALIDATORPOOLCALLERNOTVALIDATOR() (*types.Transaction, error) {
	return _ValidatorPoolErrorCodes.Contract.VALIDATORPOOLCALLERNOTVALIDATOR(&_ValidatorPoolErrorCodes.TransactOpts)
}

// VALIDATORPOOLCALLERNOTVALIDATOR is a paid mutator transaction binding the contract method 0xdf6f3c26.
//
// Solidity: function VALIDATORPOOL_CALLER_NOT_VALIDATOR() returns(bytes32)
func (_ValidatorPoolErrorCodes *ValidatorPoolErrorCodesTransactorSession) VALIDATORPOOLCALLERNOTVALIDATOR() (*types.Transaction, error) {
	return _ValidatorPoolErrorCodes.Contract.VALIDATORPOOLCALLERNOTVALIDATOR(&_ValidatorPoolErrorCodes.TransactOpts)
}

// VALIDATORPOOLCONSENSUSRUNNING is a paid mutator transaction binding the contract method 0x8ee837c4.
//
// Solidity: function VALIDATORPOOL_CONSENSUS_RUNNING() returns(bytes32)
func (_ValidatorPoolErrorCodes *ValidatorPoolErrorCodesTransactor) VALIDATORPOOLCONSENSUSRUNNING(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorPoolErrorCodes.contract.Transact(opts, "VALIDATORPOOL_CONSENSUS_RUNNING")
}

// VALIDATORPOOLCONSENSUSRUNNING is a paid mutator transaction binding the contract method 0x8ee837c4.
//
// Solidity: function VALIDATORPOOL_CONSENSUS_RUNNING() returns(bytes32)
func (_ValidatorPoolErrorCodes *ValidatorPoolErrorCodesSession) VALIDATORPOOLCONSENSUSRUNNING() (*types.Transaction, error) {
	return _ValidatorPoolErrorCodes.Contract.VALIDATORPOOLCONSENSUSRUNNING(&_ValidatorPoolErrorCodes.TransactOpts)
}

// VALIDATORPOOLCONSENSUSRUNNING is a paid mutator transaction binding the contract method 0x8ee837c4.
//
// Solidity: function VALIDATORPOOL_CONSENSUS_RUNNING() returns(bytes32)
func (_ValidatorPoolErrorCodes *ValidatorPoolErrorCodesTransactorSession) VALIDATORPOOLCONSENSUSRUNNING() (*types.Transaction, error) {
	return _ValidatorPoolErrorCodes.Contract.VALIDATORPOOLCONSENSUSRUNNING(&_ValidatorPoolErrorCodes.TransactOpts)
}

// VALIDATORPOOLDISHONESTVALIDATORNOTACCUSABLE is a paid mutator transaction binding the contract method 0x4c08af9b.
//
// Solidity: function VALIDATORPOOL_DISHONEST_VALIDATOR_NOT_ACCUSABLE() returns(bytes32)
func (_ValidatorPoolErrorCodes *ValidatorPoolErrorCodesTransactor) VALIDATORPOOLDISHONESTVALIDATORNOTACCUSABLE(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorPoolErrorCodes.contract.Transact(opts, "VALIDATORPOOL_DISHONEST_VALIDATOR_NOT_ACCUSABLE")
}

// VALIDATORPOOLDISHONESTVALIDATORNOTACCUSABLE is a paid mutator transaction binding the contract method 0x4c08af9b.
//
// Solidity: function VALIDATORPOOL_DISHONEST_VALIDATOR_NOT_ACCUSABLE() returns(bytes32)
func (_ValidatorPoolErrorCodes *ValidatorPoolErrorCodesSession) VALIDATORPOOLDISHONESTVALIDATORNOTACCUSABLE() (*types.Transaction, error) {
	return _ValidatorPoolErrorCodes.Contract.VALIDATORPOOLDISHONESTVALIDATORNOTACCUSABLE(&_ValidatorPoolErrorCodes.TransactOpts)
}

// VALIDATORPOOLDISHONESTVALIDATORNOTACCUSABLE is a paid mutator transaction binding the contract method 0x4c08af9b.
//
// Solidity: function VALIDATORPOOL_DISHONEST_VALIDATOR_NOT_ACCUSABLE() returns(bytes32)
func (_ValidatorPoolErrorCodes *ValidatorPoolErrorCodesTransactorSession) VALIDATORPOOLDISHONESTVALIDATORNOTACCUSABLE() (*types.Transaction, error) {
	return _ValidatorPoolErrorCodes.Contract.VALIDATORPOOLDISHONESTVALIDATORNOTACCUSABLE(&_ValidatorPoolErrorCodes.TransactOpts)
}

// VALIDATORPOOLETHDKGROUNDRUNNING is a paid mutator transaction binding the contract method 0x9126af77.
//
// Solidity: function VALIDATORPOOL_ETHDKG_ROUND_RUNNING() returns(bytes32)
func (_ValidatorPoolErrorCodes *ValidatorPoolErrorCodesTransactor) VALIDATORPOOLETHDKGROUNDRUNNING(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorPoolErrorCodes.contract.Transact(opts, "VALIDATORPOOL_ETHDKG_ROUND_RUNNING")
}

// VALIDATORPOOLETHDKGROUNDRUNNING is a paid mutator transaction binding the contract method 0x9126af77.
//
// Solidity: function VALIDATORPOOL_ETHDKG_ROUND_RUNNING() returns(bytes32)
func (_ValidatorPoolErrorCodes *ValidatorPoolErrorCodesSession) VALIDATORPOOLETHDKGROUNDRUNNING() (*types.Transaction, error) {
	return _ValidatorPoolErrorCodes.Contract.VALIDATORPOOLETHDKGROUNDRUNNING(&_ValidatorPoolErrorCodes.TransactOpts)
}

// VALIDATORPOOLETHDKGROUNDRUNNING is a paid mutator transaction binding the contract method 0x9126af77.
//
// Solidity: function VALIDATORPOOL_ETHDKG_ROUND_RUNNING() returns(bytes32)
func (_ValidatorPoolErrorCodes *ValidatorPoolErrorCodesTransactorSession) VALIDATORPOOLETHDKGROUNDRUNNING() (*types.Transaction, error) {
	return _ValidatorPoolErrorCodes.Contract.VALIDATORPOOLETHDKGROUNDRUNNING(&_ValidatorPoolErrorCodes.TransactOpts)
}

// VALIDATORPOOLETHBALANCECHANGED is a paid mutator transaction binding the contract method 0x826b711b.
//
// Solidity: function VALIDATORPOOL_ETH_BALANCE_CHANGED() returns(bytes32)
func (_ValidatorPoolErrorCodes *ValidatorPoolErrorCodesTransactor) VALIDATORPOOLETHBALANCECHANGED(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorPoolErrorCodes.contract.Transact(opts, "VALIDATORPOOL_ETH_BALANCE_CHANGED")
}

// VALIDATORPOOLETHBALANCECHANGED is a paid mutator transaction binding the contract method 0x826b711b.
//
// Solidity: function VALIDATORPOOL_ETH_BALANCE_CHANGED() returns(bytes32)
func (_ValidatorPoolErrorCodes *ValidatorPoolErrorCodesSession) VALIDATORPOOLETHBALANCECHANGED() (*types.Transaction, error) {
	return _ValidatorPoolErrorCodes.Contract.VALIDATORPOOLETHBALANCECHANGED(&_ValidatorPoolErrorCodes.TransactOpts)
}

// VALIDATORPOOLETHBALANCECHANGED is a paid mutator transaction binding the contract method 0x826b711b.
//
// Solidity: function VALIDATORPOOL_ETH_BALANCE_CHANGED() returns(bytes32)
func (_ValidatorPoolErrorCodes *ValidatorPoolErrorCodesTransactorSession) VALIDATORPOOLETHBALANCECHANGED() (*types.Transaction, error) {
	return _ValidatorPoolErrorCodes.Contract.VALIDATORPOOLETHBALANCECHANGED(&_ValidatorPoolErrorCodes.TransactOpts)
}

// VALIDATORPOOLFACTORYSHOULDOWNPOSITION is a paid mutator transaction binding the contract method 0x527991a3.
//
// Solidity: function VALIDATORPOOL_FACTORY_SHOULD_OWN_POSITION() returns(bytes32)
func (_ValidatorPoolErrorCodes *ValidatorPoolErrorCodesTransactor) VALIDATORPOOLFACTORYSHOULDOWNPOSITION(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorPoolErrorCodes.contract.Transact(opts, "VALIDATORPOOL_FACTORY_SHOULD_OWN_POSITION")
}

// VALIDATORPOOLFACTORYSHOULDOWNPOSITION is a paid mutator transaction binding the contract method 0x527991a3.
//
// Solidity: function VALIDATORPOOL_FACTORY_SHOULD_OWN_POSITION() returns(bytes32)
func (_ValidatorPoolErrorCodes *ValidatorPoolErrorCodesSession) VALIDATORPOOLFACTORYSHOULDOWNPOSITION() (*types.Transaction, error) {
	return _ValidatorPoolErrorCodes.Contract.VALIDATORPOOLFACTORYSHOULDOWNPOSITION(&_ValidatorPoolErrorCodes.TransactOpts)
}

// VALIDATORPOOLFACTORYSHOULDOWNPOSITION is a paid mutator transaction binding the contract method 0x527991a3.
//
// Solidity: function VALIDATORPOOL_FACTORY_SHOULD_OWN_POSITION() returns(bytes32)
func (_ValidatorPoolErrorCodes *ValidatorPoolErrorCodesTransactorSession) VALIDATORPOOLFACTORYSHOULDOWNPOSITION() (*types.Transaction, error) {
	return _ValidatorPoolErrorCodes.Contract.VALIDATORPOOLFACTORYSHOULDOWNPOSITION(&_ValidatorPoolErrorCodes.TransactOpts)
}

// VALIDATORPOOLINSUFFICIENTFUNDSINSTAKEPOSITION is a paid mutator transaction binding the contract method 0x2c20647a.
//
// Solidity: function VALIDATORPOOL_INSUFFICIENT_FUNDS_IN_STAKE_POSITION() returns(bytes32)
func (_ValidatorPoolErrorCodes *ValidatorPoolErrorCodesTransactor) VALIDATORPOOLINSUFFICIENTFUNDSINSTAKEPOSITION(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorPoolErrorCodes.contract.Transact(opts, "VALIDATORPOOL_INSUFFICIENT_FUNDS_IN_STAKE_POSITION")
}

// VALIDATORPOOLINSUFFICIENTFUNDSINSTAKEPOSITION is a paid mutator transaction binding the contract method 0x2c20647a.
//
// Solidity: function VALIDATORPOOL_INSUFFICIENT_FUNDS_IN_STAKE_POSITION() returns(bytes32)
func (_ValidatorPoolErrorCodes *ValidatorPoolErrorCodesSession) VALIDATORPOOLINSUFFICIENTFUNDSINSTAKEPOSITION() (*types.Transaction, error) {
	return _ValidatorPoolErrorCodes.Contract.VALIDATORPOOLINSUFFICIENTFUNDSINSTAKEPOSITION(&_ValidatorPoolErrorCodes.TransactOpts)
}

// VALIDATORPOOLINSUFFICIENTFUNDSINSTAKEPOSITION is a paid mutator transaction binding the contract method 0x2c20647a.
//
// Solidity: function VALIDATORPOOL_INSUFFICIENT_FUNDS_IN_STAKE_POSITION() returns(bytes32)
func (_ValidatorPoolErrorCodes *ValidatorPoolErrorCodesTransactorSession) VALIDATORPOOLINSUFFICIENTFUNDSINSTAKEPOSITION() (*types.Transaction, error) {
	return _ValidatorPoolErrorCodes.Contract.VALIDATORPOOLINSUFFICIENTFUNDSINSTAKEPOSITION(&_ValidatorPoolErrorCodes.TransactOpts)
}

// VALIDATORPOOLINVALIDINDEX is a paid mutator transaction binding the contract method 0x38dce6e0.
//
// Solidity: function VALIDATORPOOL_INVALID_INDEX() returns(bytes32)
func (_ValidatorPoolErrorCodes *ValidatorPoolErrorCodesTransactor) VALIDATORPOOLINVALIDINDEX(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorPoolErrorCodes.contract.Transact(opts, "VALIDATORPOOL_INVALID_INDEX")
}

// VALIDATORPOOLINVALIDINDEX is a paid mutator transaction binding the contract method 0x38dce6e0.
//
// Solidity: function VALIDATORPOOL_INVALID_INDEX() returns(bytes32)
func (_ValidatorPoolErrorCodes *ValidatorPoolErrorCodesSession) VALIDATORPOOLINVALIDINDEX() (*types.Transaction, error) {
	return _ValidatorPoolErrorCodes.Contract.VALIDATORPOOLINVALIDINDEX(&_ValidatorPoolErrorCodes.TransactOpts)
}

// VALIDATORPOOLINVALIDINDEX is a paid mutator transaction binding the contract method 0x38dce6e0.
//
// Solidity: function VALIDATORPOOL_INVALID_INDEX() returns(bytes32)
func (_ValidatorPoolErrorCodes *ValidatorPoolErrorCodesTransactorSession) VALIDATORPOOLINVALIDINDEX() (*types.Transaction, error) {
	return _ValidatorPoolErrorCodes.Contract.VALIDATORPOOLINVALIDINDEX(&_ValidatorPoolErrorCodes.TransactOpts)
}

// VALIDATORPOOLMAXVALIDATORSMET is a paid mutator transaction binding the contract method 0xfb5f6a21.
//
// Solidity: function VALIDATORPOOL_MAX_VALIDATORS_MET() returns(bytes32)
func (_ValidatorPoolErrorCodes *ValidatorPoolErrorCodesTransactor) VALIDATORPOOLMAXVALIDATORSMET(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorPoolErrorCodes.contract.Transact(opts, "VALIDATORPOOL_MAX_VALIDATORS_MET")
}

// VALIDATORPOOLMAXVALIDATORSMET is a paid mutator transaction binding the contract method 0xfb5f6a21.
//
// Solidity: function VALIDATORPOOL_MAX_VALIDATORS_MET() returns(bytes32)
func (_ValidatorPoolErrorCodes *ValidatorPoolErrorCodesSession) VALIDATORPOOLMAXVALIDATORSMET() (*types.Transaction, error) {
	return _ValidatorPoolErrorCodes.Contract.VALIDATORPOOLMAXVALIDATORSMET(&_ValidatorPoolErrorCodes.TransactOpts)
}

// VALIDATORPOOLMAXVALIDATORSMET is a paid mutator transaction binding the contract method 0xfb5f6a21.
//
// Solidity: function VALIDATORPOOL_MAX_VALIDATORS_MET() returns(bytes32)
func (_ValidatorPoolErrorCodes *ValidatorPoolErrorCodesTransactorSession) VALIDATORPOOLMAXVALIDATORSMET() (*types.Transaction, error) {
	return _ValidatorPoolErrorCodes.Contract.VALIDATORPOOLMAXVALIDATORSMET(&_ValidatorPoolErrorCodes.TransactOpts)
}

// VALIDATORPOOLMINIMUMSTAKENOTMET is a paid mutator transaction binding the contract method 0x057f2b29.
//
// Solidity: function VALIDATORPOOL_MINIMUM_STAKE_NOT_MET() returns(bytes32)
func (_ValidatorPoolErrorCodes *ValidatorPoolErrorCodesTransactor) VALIDATORPOOLMINIMUMSTAKENOTMET(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorPoolErrorCodes.contract.Transact(opts, "VALIDATORPOOL_MINIMUM_STAKE_NOT_MET")
}

// VALIDATORPOOLMINIMUMSTAKENOTMET is a paid mutator transaction binding the contract method 0x057f2b29.
//
// Solidity: function VALIDATORPOOL_MINIMUM_STAKE_NOT_MET() returns(bytes32)
func (_ValidatorPoolErrorCodes *ValidatorPoolErrorCodesSession) VALIDATORPOOLMINIMUMSTAKENOTMET() (*types.Transaction, error) {
	return _ValidatorPoolErrorCodes.Contract.VALIDATORPOOLMINIMUMSTAKENOTMET(&_ValidatorPoolErrorCodes.TransactOpts)
}

// VALIDATORPOOLMINIMUMSTAKENOTMET is a paid mutator transaction binding the contract method 0x057f2b29.
//
// Solidity: function VALIDATORPOOL_MINIMUM_STAKE_NOT_MET() returns(bytes32)
func (_ValidatorPoolErrorCodes *ValidatorPoolErrorCodesTransactorSession) VALIDATORPOOLMINIMUMSTAKENOTMET() (*types.Transaction, error) {
	return _ValidatorPoolErrorCodes.Contract.VALIDATORPOOLMINIMUMSTAKENOTMET(&_ValidatorPoolErrorCodes.TransactOpts)
}

// VALIDATORPOOLMINBLOCKINTERVALNOTMET is a paid mutator transaction binding the contract method 0x08abeefe.
//
// Solidity: function VALIDATORPOOL_MIN_BLOCK_INTERVAL_NOT_MET() returns(bytes32)
func (_ValidatorPoolErrorCodes *ValidatorPoolErrorCodesTransactor) VALIDATORPOOLMINBLOCKINTERVALNOTMET(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorPoolErrorCodes.contract.Transact(opts, "VALIDATORPOOL_MIN_BLOCK_INTERVAL_NOT_MET")
}

// VALIDATORPOOLMINBLOCKINTERVALNOTMET is a paid mutator transaction binding the contract method 0x08abeefe.
//
// Solidity: function VALIDATORPOOL_MIN_BLOCK_INTERVAL_NOT_MET() returns(bytes32)
func (_ValidatorPoolErrorCodes *ValidatorPoolErrorCodesSession) VALIDATORPOOLMINBLOCKINTERVALNOTMET() (*types.Transaction, error) {
	return _ValidatorPoolErrorCodes.Contract.VALIDATORPOOLMINBLOCKINTERVALNOTMET(&_ValidatorPoolErrorCodes.TransactOpts)
}

// VALIDATORPOOLMINBLOCKINTERVALNOTMET is a paid mutator transaction binding the contract method 0x08abeefe.
//
// Solidity: function VALIDATORPOOL_MIN_BLOCK_INTERVAL_NOT_MET() returns(bytes32)
func (_ValidatorPoolErrorCodes *ValidatorPoolErrorCodesTransactorSession) VALIDATORPOOLMINBLOCKINTERVALNOTMET() (*types.Transaction, error) {
	return _ValidatorPoolErrorCodes.Contract.VALIDATORPOOLMINBLOCKINTERVALNOTMET(&_ValidatorPoolErrorCodes.TransactOpts)
}

// VALIDATORPOOLONLYCONTRACTSALLOWED is a paid mutator transaction binding the contract method 0x7527a98f.
//
// Solidity: function VALIDATORPOOL_ONLY_CONTRACTS_ALLOWED() returns(bytes32)
func (_ValidatorPoolErrorCodes *ValidatorPoolErrorCodesTransactor) VALIDATORPOOLONLYCONTRACTSALLOWED(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorPoolErrorCodes.contract.Transact(opts, "VALIDATORPOOL_ONLY_CONTRACTS_ALLOWED")
}

// VALIDATORPOOLONLYCONTRACTSALLOWED is a paid mutator transaction binding the contract method 0x7527a98f.
//
// Solidity: function VALIDATORPOOL_ONLY_CONTRACTS_ALLOWED() returns(bytes32)
func (_ValidatorPoolErrorCodes *ValidatorPoolErrorCodesSession) VALIDATORPOOLONLYCONTRACTSALLOWED() (*types.Transaction, error) {
	return _ValidatorPoolErrorCodes.Contract.VALIDATORPOOLONLYCONTRACTSALLOWED(&_ValidatorPoolErrorCodes.TransactOpts)
}

// VALIDATORPOOLONLYCONTRACTSALLOWED is a paid mutator transaction binding the contract method 0x7527a98f.
//
// Solidity: function VALIDATORPOOL_ONLY_CONTRACTS_ALLOWED() returns(bytes32)
func (_ValidatorPoolErrorCodes *ValidatorPoolErrorCodesTransactorSession) VALIDATORPOOLONLYCONTRACTSALLOWED() (*types.Transaction, error) {
	return _ValidatorPoolErrorCodes.Contract.VALIDATORPOOLONLYCONTRACTSALLOWED(&_ValidatorPoolErrorCodes.TransactOpts)
}

// VALIDATORPOOLPAYOUTTOOLOW is a paid mutator transaction binding the contract method 0x0ce7d41e.
//
// Solidity: function VALIDATORPOOL_PAYOUT_TOO_LOW() returns(bytes32)
func (_ValidatorPoolErrorCodes *ValidatorPoolErrorCodesTransactor) VALIDATORPOOLPAYOUTTOOLOW(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorPoolErrorCodes.contract.Transact(opts, "VALIDATORPOOL_PAYOUT_TOO_LOW")
}

// VALIDATORPOOLPAYOUTTOOLOW is a paid mutator transaction binding the contract method 0x0ce7d41e.
//
// Solidity: function VALIDATORPOOL_PAYOUT_TOO_LOW() returns(bytes32)
func (_ValidatorPoolErrorCodes *ValidatorPoolErrorCodesSession) VALIDATORPOOLPAYOUTTOOLOW() (*types.Transaction, error) {
	return _ValidatorPoolErrorCodes.Contract.VALIDATORPOOLPAYOUTTOOLOW(&_ValidatorPoolErrorCodes.TransactOpts)
}

// VALIDATORPOOLPAYOUTTOOLOW is a paid mutator transaction binding the contract method 0x0ce7d41e.
//
// Solidity: function VALIDATORPOOL_PAYOUT_TOO_LOW() returns(bytes32)
func (_ValidatorPoolErrorCodes *ValidatorPoolErrorCodesTransactorSession) VALIDATORPOOLPAYOUTTOOLOW() (*types.Transaction, error) {
	return _ValidatorPoolErrorCodes.Contract.VALIDATORPOOLPAYOUTTOOLOW(&_ValidatorPoolErrorCodes.TransactOpts)
}

// VALIDATORPOOLPROFITSONLYCLAIMABLEDURINGCONSENSUS is a paid mutator transaction binding the contract method 0x4c454a97.
//
// Solidity: function VALIDATORPOOL_PROFITS_ONLY_CLAIMABLE_DURING_CONSENSUS() returns(bytes32)
func (_ValidatorPoolErrorCodes *ValidatorPoolErrorCodesTransactor) VALIDATORPOOLPROFITSONLYCLAIMABLEDURINGCONSENSUS(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorPoolErrorCodes.contract.Transact(opts, "VALIDATORPOOL_PROFITS_ONLY_CLAIMABLE_DURING_CONSENSUS")
}

// VALIDATORPOOLPROFITSONLYCLAIMABLEDURINGCONSENSUS is a paid mutator transaction binding the contract method 0x4c454a97.
//
// Solidity: function VALIDATORPOOL_PROFITS_ONLY_CLAIMABLE_DURING_CONSENSUS() returns(bytes32)
func (_ValidatorPoolErrorCodes *ValidatorPoolErrorCodesSession) VALIDATORPOOLPROFITSONLYCLAIMABLEDURINGCONSENSUS() (*types.Transaction, error) {
	return _ValidatorPoolErrorCodes.Contract.VALIDATORPOOLPROFITSONLYCLAIMABLEDURINGCONSENSUS(&_ValidatorPoolErrorCodes.TransactOpts)
}

// VALIDATORPOOLPROFITSONLYCLAIMABLEDURINGCONSENSUS is a paid mutator transaction binding the contract method 0x4c454a97.
//
// Solidity: function VALIDATORPOOL_PROFITS_ONLY_CLAIMABLE_DURING_CONSENSUS() returns(bytes32)
func (_ValidatorPoolErrorCodes *ValidatorPoolErrorCodesTransactorSession) VALIDATORPOOLPROFITSONLYCLAIMABLEDURINGCONSENSUS() (*types.Transaction, error) {
	return _ValidatorPoolErrorCodes.Contract.VALIDATORPOOLPROFITSONLYCLAIMABLEDURINGCONSENSUS(&_ValidatorPoolErrorCodes.TransactOpts)
}

// VALIDATORPOOLREGISTRATIONPARAMETERLENGTHMISMATCH is a paid mutator transaction binding the contract method 0xe5b50fe6.
//
// Solidity: function VALIDATORPOOL_REGISTRATION_PARAMETER_LENGTH_MISMATCH() returns(bytes32)
func (_ValidatorPoolErrorCodes *ValidatorPoolErrorCodesTransactor) VALIDATORPOOLREGISTRATIONPARAMETERLENGTHMISMATCH(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorPoolErrorCodes.contract.Transact(opts, "VALIDATORPOOL_REGISTRATION_PARAMETER_LENGTH_MISMATCH")
}

// VALIDATORPOOLREGISTRATIONPARAMETERLENGTHMISMATCH is a paid mutator transaction binding the contract method 0xe5b50fe6.
//
// Solidity: function VALIDATORPOOL_REGISTRATION_PARAMETER_LENGTH_MISMATCH() returns(bytes32)
func (_ValidatorPoolErrorCodes *ValidatorPoolErrorCodesSession) VALIDATORPOOLREGISTRATIONPARAMETERLENGTHMISMATCH() (*types.Transaction, error) {
	return _ValidatorPoolErrorCodes.Contract.VALIDATORPOOLREGISTRATIONPARAMETERLENGTHMISMATCH(&_ValidatorPoolErrorCodes.TransactOpts)
}

// VALIDATORPOOLREGISTRATIONPARAMETERLENGTHMISMATCH is a paid mutator transaction binding the contract method 0xe5b50fe6.
//
// Solidity: function VALIDATORPOOL_REGISTRATION_PARAMETER_LENGTH_MISMATCH() returns(bytes32)
func (_ValidatorPoolErrorCodes *ValidatorPoolErrorCodesTransactorSession) VALIDATORPOOLREGISTRATIONPARAMETERLENGTHMISMATCH() (*types.Transaction, error) {
	return _ValidatorPoolErrorCodes.Contract.VALIDATORPOOLREGISTRATIONPARAMETERLENGTHMISMATCH(&_ValidatorPoolErrorCodes.TransactOpts)
}

// VALIDATORPOOLSENDERNOTINEXITINGQUEUE is a paid mutator transaction binding the contract method 0x55efc658.
//
// Solidity: function VALIDATORPOOL_SENDER_NOT_IN_EXITING_QUEUE() returns(bytes32)
func (_ValidatorPoolErrorCodes *ValidatorPoolErrorCodesTransactor) VALIDATORPOOLSENDERNOTINEXITINGQUEUE(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorPoolErrorCodes.contract.Transact(opts, "VALIDATORPOOL_SENDER_NOT_IN_EXITING_QUEUE")
}

// VALIDATORPOOLSENDERNOTINEXITINGQUEUE is a paid mutator transaction binding the contract method 0x55efc658.
//
// Solidity: function VALIDATORPOOL_SENDER_NOT_IN_EXITING_QUEUE() returns(bytes32)
func (_ValidatorPoolErrorCodes *ValidatorPoolErrorCodesSession) VALIDATORPOOLSENDERNOTINEXITINGQUEUE() (*types.Transaction, error) {
	return _ValidatorPoolErrorCodes.Contract.VALIDATORPOOLSENDERNOTINEXITINGQUEUE(&_ValidatorPoolErrorCodes.TransactOpts)
}

// VALIDATORPOOLSENDERNOTINEXITINGQUEUE is a paid mutator transaction binding the contract method 0x55efc658.
//
// Solidity: function VALIDATORPOOL_SENDER_NOT_IN_EXITING_QUEUE() returns(bytes32)
func (_ValidatorPoolErrorCodes *ValidatorPoolErrorCodesTransactorSession) VALIDATORPOOLSENDERNOTINEXITINGQUEUE() (*types.Transaction, error) {
	return _ValidatorPoolErrorCodes.Contract.VALIDATORPOOLSENDERNOTINEXITINGQUEUE(&_ValidatorPoolErrorCodes.TransactOpts)
}

// VALIDATORPOOLTOKENBALANCECHANGED is a paid mutator transaction binding the contract method 0x329fa341.
//
// Solidity: function VALIDATORPOOL_TOKEN_BALANCE_CHANGED() returns(bytes32)
func (_ValidatorPoolErrorCodes *ValidatorPoolErrorCodesTransactor) VALIDATORPOOLTOKENBALANCECHANGED(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorPoolErrorCodes.contract.Transact(opts, "VALIDATORPOOL_TOKEN_BALANCE_CHANGED")
}

// VALIDATORPOOLTOKENBALANCECHANGED is a paid mutator transaction binding the contract method 0x329fa341.
//
// Solidity: function VALIDATORPOOL_TOKEN_BALANCE_CHANGED() returns(bytes32)
func (_ValidatorPoolErrorCodes *ValidatorPoolErrorCodesSession) VALIDATORPOOLTOKENBALANCECHANGED() (*types.Transaction, error) {
	return _ValidatorPoolErrorCodes.Contract.VALIDATORPOOLTOKENBALANCECHANGED(&_ValidatorPoolErrorCodes.TransactOpts)
}

// VALIDATORPOOLTOKENBALANCECHANGED is a paid mutator transaction binding the contract method 0x329fa341.
//
// Solidity: function VALIDATORPOOL_TOKEN_BALANCE_CHANGED() returns(bytes32)
func (_ValidatorPoolErrorCodes *ValidatorPoolErrorCodesTransactorSession) VALIDATORPOOLTOKENBALANCECHANGED() (*types.Transaction, error) {
	return _ValidatorPoolErrorCodes.Contract.VALIDATORPOOLTOKENBALANCECHANGED(&_ValidatorPoolErrorCodes.TransactOpts)
}

// VALIDATORPOOLVALIDATORSGREATERTHANAVAILABLE is a paid mutator transaction binding the contract method 0x7cd62d16.
//
// Solidity: function VALIDATORPOOL_VALIDATORS_GREATER_THAN_AVAILABLE() returns(bytes32)
func (_ValidatorPoolErrorCodes *ValidatorPoolErrorCodesTransactor) VALIDATORPOOLVALIDATORSGREATERTHANAVAILABLE(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorPoolErrorCodes.contract.Transact(opts, "VALIDATORPOOL_VALIDATORS_GREATER_THAN_AVAILABLE")
}

// VALIDATORPOOLVALIDATORSGREATERTHANAVAILABLE is a paid mutator transaction binding the contract method 0x7cd62d16.
//
// Solidity: function VALIDATORPOOL_VALIDATORS_GREATER_THAN_AVAILABLE() returns(bytes32)
func (_ValidatorPoolErrorCodes *ValidatorPoolErrorCodesSession) VALIDATORPOOLVALIDATORSGREATERTHANAVAILABLE() (*types.Transaction, error) {
	return _ValidatorPoolErrorCodes.Contract.VALIDATORPOOLVALIDATORSGREATERTHANAVAILABLE(&_ValidatorPoolErrorCodes.TransactOpts)
}

// VALIDATORPOOLVALIDATORSGREATERTHANAVAILABLE is a paid mutator transaction binding the contract method 0x7cd62d16.
//
// Solidity: function VALIDATORPOOL_VALIDATORS_GREATER_THAN_AVAILABLE() returns(bytes32)
func (_ValidatorPoolErrorCodes *ValidatorPoolErrorCodesTransactorSession) VALIDATORPOOLVALIDATORSGREATERTHANAVAILABLE() (*types.Transaction, error) {
	return _ValidatorPoolErrorCodes.Contract.VALIDATORPOOLVALIDATORSGREATERTHANAVAILABLE(&_ValidatorPoolErrorCodes.TransactOpts)
}

// VALIDATORPOOLWAITINGPERIODNOTMET is a paid mutator transaction binding the contract method 0x0e2d4d55.
//
// Solidity: function VALIDATORPOOL_WAITING_PERIOD_NOT_MET() returns(bytes32)
func (_ValidatorPoolErrorCodes *ValidatorPoolErrorCodesTransactor) VALIDATORPOOLWAITINGPERIODNOTMET(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorPoolErrorCodes.contract.Transact(opts, "VALIDATORPOOL_WAITING_PERIOD_NOT_MET")
}

// VALIDATORPOOLWAITINGPERIODNOTMET is a paid mutator transaction binding the contract method 0x0e2d4d55.
//
// Solidity: function VALIDATORPOOL_WAITING_PERIOD_NOT_MET() returns(bytes32)
func (_ValidatorPoolErrorCodes *ValidatorPoolErrorCodesSession) VALIDATORPOOLWAITINGPERIODNOTMET() (*types.Transaction, error) {
	return _ValidatorPoolErrorCodes.Contract.VALIDATORPOOLWAITINGPERIODNOTMET(&_ValidatorPoolErrorCodes.TransactOpts)
}

// VALIDATORPOOLWAITINGPERIODNOTMET is a paid mutator transaction binding the contract method 0x0e2d4d55.
//
// Solidity: function VALIDATORPOOL_WAITING_PERIOD_NOT_MET() returns(bytes32)
func (_ValidatorPoolErrorCodes *ValidatorPoolErrorCodesTransactorSession) VALIDATORPOOLWAITINGPERIODNOTMET() (*types.Transaction, error) {
	return _ValidatorPoolErrorCodes.Contract.VALIDATORPOOLWAITINGPERIODNOTMET(&_ValidatorPoolErrorCodes.TransactOpts)
}
