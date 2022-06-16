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

// ETHDKGErrorCodesABI is the input ABI used to generate the binding from.
const ETHDKGErrorCodesABI = "[{\"inputs\":[],\"name\":\"ETHDKG_ACCUSED_DID_NOT_DISTRIBUTE_SHARES_IN_ROUND\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ETHDKG_ACCUSED_DID_NOT_PARTICIPATE_IN_GPKJ_SUBMISSION\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ETHDKG_ACCUSED_DID_NOT_SUBMIT_GPKJ_IN_ROUND\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ETHDKG_ACCUSED_DISTRIBUTED_GPKJ\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ETHDKG_ACCUSED_DISTRIBUTED_SHARES_IN_ROUND\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ETHDKG_ACCUSED_HAS_COMMITMENTS\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ETHDKG_ACCUSED_NOT_PARTICIPATING_IN_ROUND\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ETHDKG_ACCUSED_NOT_VALIDATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ETHDKG_ACCUSED_PARTICIPATING_IN_ROUND\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ETHDKG_ACCUSED_SUBMITTED_SHARES_IN_ROUND\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ETHDKG_COMMITMENT_NOT_ON_CURVE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ETHDKG_COMMITMENT_ZERO\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ETHDKG_DISPUTER_DID_NOT_DISTRIBUTE_SHARES_IN_ROUND\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ETHDKG_DISPUTER_DID_NOT_SUBMIT_GPKJ_IN_ROUND\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ETHDKG_DISPUTER_NOT_PARTICIPATING_IN_ROUND\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ETHDKG_DISTRIBUTED_SHARE_HASH_ZERO\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ETHDKG_GPKJ_ZERO\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ETHDKG_INVALID_ARGS\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ETHDKG_INVALID_COMMITMENTS\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ETHDKG_INVALID_KEYSHARE_G1\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ETHDKG_INVALID_KEYSHARE_G2\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ETHDKG_INVALID_KEY_OR_PROOF\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ETHDKG_INVALID_NONCE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ETHDKG_INVALID_NUM_COMMITMENTS\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ETHDKG_INVALID_NUM_ENCRYPTED_SHARES\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ETHDKG_INVALID_OR_DUPLICATED_PARTICIPANT\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ETHDKG_INVALID_SHARES_OR_COMMITMENTS\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ETHDKG_KEYSHARE_PHASE_INVALID_NONCE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ETHDKG_MASTER_PUBLIC_KEY_PAIRING_CHECK_FAILURE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ETHDKG_MIGRATION_INPUT_DATA_MISMATCH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ETHDKG_MIGRATION_INVALID_NONCE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ETHDKG_MIN_VALIDATORS_NOT_MET\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ETHDKG_NOT_IN_DISPUTE_PHASE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ETHDKG_NOT_IN_GPKJ_SUBMISSION_PHASE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ETHDKG_NOT_IN_KEYSHARE_SUBMISSION_PHASE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ETHDKG_NOT_IN_MASTER_PUBLIC_KEY_SUBMISSION_PHASE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ETHDKG_NOT_IN_POST_GPKJ_DISPUTE_PHASE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ETHDKG_NOT_IN_POST_GPKJ_SUBMISSION_PHASE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ETHDKG_NOT_IN_POST_KEYSHARE_SUBMISSION_PHASE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ETHDKG_NOT_IN_POST_REGISTRATION_PHASE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ETHDKG_NOT_IN_POST_SHARED_DISTRIBUTION_PHASE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ETHDKG_NOT_IN_REGISTRATION_PHASE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ETHDKG_NOT_IN_SHARED_DISTRIBUTION_PHASE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ETHDKG_ONLY_VALIDATORS_ALLOWED\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ETHDKG_PARTICIPANT_DISTRIBUTED_SHARES_IN_ROUND\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ETHDKG_PARTICIPANT_PARTICIPATING_IN_ROUND\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ETHDKG_PARTICIPANT_SUBMITTED_GPKJ_IN_ROUND\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ETHDKG_PARTICIPANT_SUBMITTED_KEYSHARES_IN_ROUND\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ETHDKG_PUBLIC_KEY_NOT_ON_CURVE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ETHDKG_PUBLIC_KEY_ZERO\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ETHDKG_REQUISITES_INCOMPLETE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ETHDKG_SHARES_AND_COMMITMENTS_MISMATCH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ETHDKG_VARIABLE_CANNOT_BE_SET_WHILE_RUNNING\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ETHDKGErrorCodes is an auto generated Go binding around an Ethereum contract.
type ETHDKGErrorCodes struct {
	ETHDKGErrorCodesCaller     // Read-only binding to the contract
	ETHDKGErrorCodesTransactor // Write-only binding to the contract
	ETHDKGErrorCodesFilterer   // Log filterer for contract events
}

// ETHDKGErrorCodesCaller is an auto generated read-only Go binding around an Ethereum contract.
type ETHDKGErrorCodesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ETHDKGErrorCodesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ETHDKGErrorCodesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ETHDKGErrorCodesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ETHDKGErrorCodesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ETHDKGErrorCodesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ETHDKGErrorCodesSession struct {
	Contract     *ETHDKGErrorCodes // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ETHDKGErrorCodesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ETHDKGErrorCodesCallerSession struct {
	Contract *ETHDKGErrorCodesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ETHDKGErrorCodesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ETHDKGErrorCodesTransactorSession struct {
	Contract     *ETHDKGErrorCodesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ETHDKGErrorCodesRaw is an auto generated low-level Go binding around an Ethereum contract.
type ETHDKGErrorCodesRaw struct {
	Contract *ETHDKGErrorCodes // Generic contract binding to access the raw methods on
}

// ETHDKGErrorCodesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ETHDKGErrorCodesCallerRaw struct {
	Contract *ETHDKGErrorCodesCaller // Generic read-only contract binding to access the raw methods on
}

// ETHDKGErrorCodesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ETHDKGErrorCodesTransactorRaw struct {
	Contract *ETHDKGErrorCodesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewETHDKGErrorCodes creates a new instance of ETHDKGErrorCodes, bound to a specific deployed contract.
func NewETHDKGErrorCodes(address common.Address, backend bind.ContractBackend) (*ETHDKGErrorCodes, error) {
	contract, err := bindETHDKGErrorCodes(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ETHDKGErrorCodes{ETHDKGErrorCodesCaller: ETHDKGErrorCodesCaller{contract: contract}, ETHDKGErrorCodesTransactor: ETHDKGErrorCodesTransactor{contract: contract}, ETHDKGErrorCodesFilterer: ETHDKGErrorCodesFilterer{contract: contract}}, nil
}

// NewETHDKGErrorCodesCaller creates a new read-only instance of ETHDKGErrorCodes, bound to a specific deployed contract.
func NewETHDKGErrorCodesCaller(address common.Address, caller bind.ContractCaller) (*ETHDKGErrorCodesCaller, error) {
	contract, err := bindETHDKGErrorCodes(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ETHDKGErrorCodesCaller{contract: contract}, nil
}

// NewETHDKGErrorCodesTransactor creates a new write-only instance of ETHDKGErrorCodes, bound to a specific deployed contract.
func NewETHDKGErrorCodesTransactor(address common.Address, transactor bind.ContractTransactor) (*ETHDKGErrorCodesTransactor, error) {
	contract, err := bindETHDKGErrorCodes(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ETHDKGErrorCodesTransactor{contract: contract}, nil
}

// NewETHDKGErrorCodesFilterer creates a new log filterer instance of ETHDKGErrorCodes, bound to a specific deployed contract.
func NewETHDKGErrorCodesFilterer(address common.Address, filterer bind.ContractFilterer) (*ETHDKGErrorCodesFilterer, error) {
	contract, err := bindETHDKGErrorCodes(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ETHDKGErrorCodesFilterer{contract: contract}, nil
}

// bindETHDKGErrorCodes binds a generic wrapper to an already deployed contract.
func bindETHDKGErrorCodes(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ETHDKGErrorCodesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ETHDKGErrorCodes *ETHDKGErrorCodesRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ETHDKGErrorCodes.Contract.ETHDKGErrorCodesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ETHDKGErrorCodes *ETHDKGErrorCodesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ETHDKGErrorCodes.Contract.ETHDKGErrorCodesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ETHDKGErrorCodes *ETHDKGErrorCodesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ETHDKGErrorCodes.Contract.ETHDKGErrorCodesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ETHDKGErrorCodes *ETHDKGErrorCodesCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ETHDKGErrorCodes.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ETHDKGErrorCodes *ETHDKGErrorCodesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ETHDKGErrorCodes.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ETHDKGErrorCodes *ETHDKGErrorCodesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ETHDKGErrorCodes.Contract.contract.Transact(opts, method, params...)
}

// ETHDKGACCUSEDDIDNOTDISTRIBUTESHARESINROUND is a paid mutator transaction binding the contract method 0xfa0f33b9.
//
// Solidity: function ETHDKG_ACCUSED_DID_NOT_DISTRIBUTE_SHARES_IN_ROUND() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesTransactor) ETHDKGACCUSEDDIDNOTDISTRIBUTESHARESINROUND(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ETHDKGErrorCodes.contract.Transact(opts, "ETHDKG_ACCUSED_DID_NOT_DISTRIBUTE_SHARES_IN_ROUND")
}

// ETHDKGACCUSEDDIDNOTDISTRIBUTESHARESINROUND is a paid mutator transaction binding the contract method 0xfa0f33b9.
//
// Solidity: function ETHDKG_ACCUSED_DID_NOT_DISTRIBUTE_SHARES_IN_ROUND() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesSession) ETHDKGACCUSEDDIDNOTDISTRIBUTESHARESINROUND() (*types.Transaction, error) {
	return _ETHDKGErrorCodes.Contract.ETHDKGACCUSEDDIDNOTDISTRIBUTESHARESINROUND(&_ETHDKGErrorCodes.TransactOpts)
}

// ETHDKGACCUSEDDIDNOTDISTRIBUTESHARESINROUND is a paid mutator transaction binding the contract method 0xfa0f33b9.
//
// Solidity: function ETHDKG_ACCUSED_DID_NOT_DISTRIBUTE_SHARES_IN_ROUND() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesTransactorSession) ETHDKGACCUSEDDIDNOTDISTRIBUTESHARESINROUND() (*types.Transaction, error) {
	return _ETHDKGErrorCodes.Contract.ETHDKGACCUSEDDIDNOTDISTRIBUTESHARESINROUND(&_ETHDKGErrorCodes.TransactOpts)
}

// ETHDKGACCUSEDDIDNOTPARTICIPATEINGPKJSUBMISSION is a paid mutator transaction binding the contract method 0x0348f5cc.
//
// Solidity: function ETHDKG_ACCUSED_DID_NOT_PARTICIPATE_IN_GPKJ_SUBMISSION() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesTransactor) ETHDKGACCUSEDDIDNOTPARTICIPATEINGPKJSUBMISSION(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ETHDKGErrorCodes.contract.Transact(opts, "ETHDKG_ACCUSED_DID_NOT_PARTICIPATE_IN_GPKJ_SUBMISSION")
}

// ETHDKGACCUSEDDIDNOTPARTICIPATEINGPKJSUBMISSION is a paid mutator transaction binding the contract method 0x0348f5cc.
//
// Solidity: function ETHDKG_ACCUSED_DID_NOT_PARTICIPATE_IN_GPKJ_SUBMISSION() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesSession) ETHDKGACCUSEDDIDNOTPARTICIPATEINGPKJSUBMISSION() (*types.Transaction, error) {
	return _ETHDKGErrorCodes.Contract.ETHDKGACCUSEDDIDNOTPARTICIPATEINGPKJSUBMISSION(&_ETHDKGErrorCodes.TransactOpts)
}

// ETHDKGACCUSEDDIDNOTPARTICIPATEINGPKJSUBMISSION is a paid mutator transaction binding the contract method 0x0348f5cc.
//
// Solidity: function ETHDKG_ACCUSED_DID_NOT_PARTICIPATE_IN_GPKJ_SUBMISSION() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesTransactorSession) ETHDKGACCUSEDDIDNOTPARTICIPATEINGPKJSUBMISSION() (*types.Transaction, error) {
	return _ETHDKGErrorCodes.Contract.ETHDKGACCUSEDDIDNOTPARTICIPATEINGPKJSUBMISSION(&_ETHDKGErrorCodes.TransactOpts)
}

// ETHDKGACCUSEDDIDNOTSUBMITGPKJINROUND is a paid mutator transaction binding the contract method 0xac4683be.
//
// Solidity: function ETHDKG_ACCUSED_DID_NOT_SUBMIT_GPKJ_IN_ROUND() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesTransactor) ETHDKGACCUSEDDIDNOTSUBMITGPKJINROUND(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ETHDKGErrorCodes.contract.Transact(opts, "ETHDKG_ACCUSED_DID_NOT_SUBMIT_GPKJ_IN_ROUND")
}

// ETHDKGACCUSEDDIDNOTSUBMITGPKJINROUND is a paid mutator transaction binding the contract method 0xac4683be.
//
// Solidity: function ETHDKG_ACCUSED_DID_NOT_SUBMIT_GPKJ_IN_ROUND() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesSession) ETHDKGACCUSEDDIDNOTSUBMITGPKJINROUND() (*types.Transaction, error) {
	return _ETHDKGErrorCodes.Contract.ETHDKGACCUSEDDIDNOTSUBMITGPKJINROUND(&_ETHDKGErrorCodes.TransactOpts)
}

// ETHDKGACCUSEDDIDNOTSUBMITGPKJINROUND is a paid mutator transaction binding the contract method 0xac4683be.
//
// Solidity: function ETHDKG_ACCUSED_DID_NOT_SUBMIT_GPKJ_IN_ROUND() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesTransactorSession) ETHDKGACCUSEDDIDNOTSUBMITGPKJINROUND() (*types.Transaction, error) {
	return _ETHDKGErrorCodes.Contract.ETHDKGACCUSEDDIDNOTSUBMITGPKJINROUND(&_ETHDKGErrorCodes.TransactOpts)
}

// ETHDKGACCUSEDDISTRIBUTEDGPKJ is a paid mutator transaction binding the contract method 0xaf1c8f58.
//
// Solidity: function ETHDKG_ACCUSED_DISTRIBUTED_GPKJ() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesTransactor) ETHDKGACCUSEDDISTRIBUTEDGPKJ(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ETHDKGErrorCodes.contract.Transact(opts, "ETHDKG_ACCUSED_DISTRIBUTED_GPKJ")
}

// ETHDKGACCUSEDDISTRIBUTEDGPKJ is a paid mutator transaction binding the contract method 0xaf1c8f58.
//
// Solidity: function ETHDKG_ACCUSED_DISTRIBUTED_GPKJ() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesSession) ETHDKGACCUSEDDISTRIBUTEDGPKJ() (*types.Transaction, error) {
	return _ETHDKGErrorCodes.Contract.ETHDKGACCUSEDDISTRIBUTEDGPKJ(&_ETHDKGErrorCodes.TransactOpts)
}

// ETHDKGACCUSEDDISTRIBUTEDGPKJ is a paid mutator transaction binding the contract method 0xaf1c8f58.
//
// Solidity: function ETHDKG_ACCUSED_DISTRIBUTED_GPKJ() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesTransactorSession) ETHDKGACCUSEDDISTRIBUTEDGPKJ() (*types.Transaction, error) {
	return _ETHDKGErrorCodes.Contract.ETHDKGACCUSEDDISTRIBUTEDGPKJ(&_ETHDKGErrorCodes.TransactOpts)
}

// ETHDKGACCUSEDDISTRIBUTEDSHARESINROUND is a paid mutator transaction binding the contract method 0x2838edae.
//
// Solidity: function ETHDKG_ACCUSED_DISTRIBUTED_SHARES_IN_ROUND() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesTransactor) ETHDKGACCUSEDDISTRIBUTEDSHARESINROUND(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ETHDKGErrorCodes.contract.Transact(opts, "ETHDKG_ACCUSED_DISTRIBUTED_SHARES_IN_ROUND")
}

// ETHDKGACCUSEDDISTRIBUTEDSHARESINROUND is a paid mutator transaction binding the contract method 0x2838edae.
//
// Solidity: function ETHDKG_ACCUSED_DISTRIBUTED_SHARES_IN_ROUND() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesSession) ETHDKGACCUSEDDISTRIBUTEDSHARESINROUND() (*types.Transaction, error) {
	return _ETHDKGErrorCodes.Contract.ETHDKGACCUSEDDISTRIBUTEDSHARESINROUND(&_ETHDKGErrorCodes.TransactOpts)
}

// ETHDKGACCUSEDDISTRIBUTEDSHARESINROUND is a paid mutator transaction binding the contract method 0x2838edae.
//
// Solidity: function ETHDKG_ACCUSED_DISTRIBUTED_SHARES_IN_ROUND() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesTransactorSession) ETHDKGACCUSEDDISTRIBUTEDSHARESINROUND() (*types.Transaction, error) {
	return _ETHDKGErrorCodes.Contract.ETHDKGACCUSEDDISTRIBUTEDSHARESINROUND(&_ETHDKGErrorCodes.TransactOpts)
}

// ETHDKGACCUSEDHASCOMMITMENTS is a paid mutator transaction binding the contract method 0x4cd291bf.
//
// Solidity: function ETHDKG_ACCUSED_HAS_COMMITMENTS() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesTransactor) ETHDKGACCUSEDHASCOMMITMENTS(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ETHDKGErrorCodes.contract.Transact(opts, "ETHDKG_ACCUSED_HAS_COMMITMENTS")
}

// ETHDKGACCUSEDHASCOMMITMENTS is a paid mutator transaction binding the contract method 0x4cd291bf.
//
// Solidity: function ETHDKG_ACCUSED_HAS_COMMITMENTS() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesSession) ETHDKGACCUSEDHASCOMMITMENTS() (*types.Transaction, error) {
	return _ETHDKGErrorCodes.Contract.ETHDKGACCUSEDHASCOMMITMENTS(&_ETHDKGErrorCodes.TransactOpts)
}

// ETHDKGACCUSEDHASCOMMITMENTS is a paid mutator transaction binding the contract method 0x4cd291bf.
//
// Solidity: function ETHDKG_ACCUSED_HAS_COMMITMENTS() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesTransactorSession) ETHDKGACCUSEDHASCOMMITMENTS() (*types.Transaction, error) {
	return _ETHDKGErrorCodes.Contract.ETHDKGACCUSEDHASCOMMITMENTS(&_ETHDKGErrorCodes.TransactOpts)
}

// ETHDKGACCUSEDNOTPARTICIPATINGINROUND is a paid mutator transaction binding the contract method 0xe11879cc.
//
// Solidity: function ETHDKG_ACCUSED_NOT_PARTICIPATING_IN_ROUND() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesTransactor) ETHDKGACCUSEDNOTPARTICIPATINGINROUND(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ETHDKGErrorCodes.contract.Transact(opts, "ETHDKG_ACCUSED_NOT_PARTICIPATING_IN_ROUND")
}

// ETHDKGACCUSEDNOTPARTICIPATINGINROUND is a paid mutator transaction binding the contract method 0xe11879cc.
//
// Solidity: function ETHDKG_ACCUSED_NOT_PARTICIPATING_IN_ROUND() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesSession) ETHDKGACCUSEDNOTPARTICIPATINGINROUND() (*types.Transaction, error) {
	return _ETHDKGErrorCodes.Contract.ETHDKGACCUSEDNOTPARTICIPATINGINROUND(&_ETHDKGErrorCodes.TransactOpts)
}

// ETHDKGACCUSEDNOTPARTICIPATINGINROUND is a paid mutator transaction binding the contract method 0xe11879cc.
//
// Solidity: function ETHDKG_ACCUSED_NOT_PARTICIPATING_IN_ROUND() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesTransactorSession) ETHDKGACCUSEDNOTPARTICIPATINGINROUND() (*types.Transaction, error) {
	return _ETHDKGErrorCodes.Contract.ETHDKGACCUSEDNOTPARTICIPATINGINROUND(&_ETHDKGErrorCodes.TransactOpts)
}

// ETHDKGACCUSEDNOTVALIDATOR is a paid mutator transaction binding the contract method 0xf5f46e73.
//
// Solidity: function ETHDKG_ACCUSED_NOT_VALIDATOR() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesTransactor) ETHDKGACCUSEDNOTVALIDATOR(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ETHDKGErrorCodes.contract.Transact(opts, "ETHDKG_ACCUSED_NOT_VALIDATOR")
}

// ETHDKGACCUSEDNOTVALIDATOR is a paid mutator transaction binding the contract method 0xf5f46e73.
//
// Solidity: function ETHDKG_ACCUSED_NOT_VALIDATOR() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesSession) ETHDKGACCUSEDNOTVALIDATOR() (*types.Transaction, error) {
	return _ETHDKGErrorCodes.Contract.ETHDKGACCUSEDNOTVALIDATOR(&_ETHDKGErrorCodes.TransactOpts)
}

// ETHDKGACCUSEDNOTVALIDATOR is a paid mutator transaction binding the contract method 0xf5f46e73.
//
// Solidity: function ETHDKG_ACCUSED_NOT_VALIDATOR() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesTransactorSession) ETHDKGACCUSEDNOTVALIDATOR() (*types.Transaction, error) {
	return _ETHDKGErrorCodes.Contract.ETHDKGACCUSEDNOTVALIDATOR(&_ETHDKGErrorCodes.TransactOpts)
}

// ETHDKGACCUSEDPARTICIPATINGINROUND is a paid mutator transaction binding the contract method 0x55b83c56.
//
// Solidity: function ETHDKG_ACCUSED_PARTICIPATING_IN_ROUND() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesTransactor) ETHDKGACCUSEDPARTICIPATINGINROUND(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ETHDKGErrorCodes.contract.Transact(opts, "ETHDKG_ACCUSED_PARTICIPATING_IN_ROUND")
}

// ETHDKGACCUSEDPARTICIPATINGINROUND is a paid mutator transaction binding the contract method 0x55b83c56.
//
// Solidity: function ETHDKG_ACCUSED_PARTICIPATING_IN_ROUND() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesSession) ETHDKGACCUSEDPARTICIPATINGINROUND() (*types.Transaction, error) {
	return _ETHDKGErrorCodes.Contract.ETHDKGACCUSEDPARTICIPATINGINROUND(&_ETHDKGErrorCodes.TransactOpts)
}

// ETHDKGACCUSEDPARTICIPATINGINROUND is a paid mutator transaction binding the contract method 0x55b83c56.
//
// Solidity: function ETHDKG_ACCUSED_PARTICIPATING_IN_ROUND() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesTransactorSession) ETHDKGACCUSEDPARTICIPATINGINROUND() (*types.Transaction, error) {
	return _ETHDKGErrorCodes.Contract.ETHDKGACCUSEDPARTICIPATINGINROUND(&_ETHDKGErrorCodes.TransactOpts)
}

// ETHDKGACCUSEDSUBMITTEDSHARESINROUND is a paid mutator transaction binding the contract method 0xb23b8358.
//
// Solidity: function ETHDKG_ACCUSED_SUBMITTED_SHARES_IN_ROUND() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesTransactor) ETHDKGACCUSEDSUBMITTEDSHARESINROUND(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ETHDKGErrorCodes.contract.Transact(opts, "ETHDKG_ACCUSED_SUBMITTED_SHARES_IN_ROUND")
}

// ETHDKGACCUSEDSUBMITTEDSHARESINROUND is a paid mutator transaction binding the contract method 0xb23b8358.
//
// Solidity: function ETHDKG_ACCUSED_SUBMITTED_SHARES_IN_ROUND() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesSession) ETHDKGACCUSEDSUBMITTEDSHARESINROUND() (*types.Transaction, error) {
	return _ETHDKGErrorCodes.Contract.ETHDKGACCUSEDSUBMITTEDSHARESINROUND(&_ETHDKGErrorCodes.TransactOpts)
}

// ETHDKGACCUSEDSUBMITTEDSHARESINROUND is a paid mutator transaction binding the contract method 0xb23b8358.
//
// Solidity: function ETHDKG_ACCUSED_SUBMITTED_SHARES_IN_ROUND() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesTransactorSession) ETHDKGACCUSEDSUBMITTEDSHARESINROUND() (*types.Transaction, error) {
	return _ETHDKGErrorCodes.Contract.ETHDKGACCUSEDSUBMITTEDSHARESINROUND(&_ETHDKGErrorCodes.TransactOpts)
}

// ETHDKGCOMMITMENTNOTONCURVE is a paid mutator transaction binding the contract method 0xe58f04ed.
//
// Solidity: function ETHDKG_COMMITMENT_NOT_ON_CURVE() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesTransactor) ETHDKGCOMMITMENTNOTONCURVE(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ETHDKGErrorCodes.contract.Transact(opts, "ETHDKG_COMMITMENT_NOT_ON_CURVE")
}

// ETHDKGCOMMITMENTNOTONCURVE is a paid mutator transaction binding the contract method 0xe58f04ed.
//
// Solidity: function ETHDKG_COMMITMENT_NOT_ON_CURVE() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesSession) ETHDKGCOMMITMENTNOTONCURVE() (*types.Transaction, error) {
	return _ETHDKGErrorCodes.Contract.ETHDKGCOMMITMENTNOTONCURVE(&_ETHDKGErrorCodes.TransactOpts)
}

// ETHDKGCOMMITMENTNOTONCURVE is a paid mutator transaction binding the contract method 0xe58f04ed.
//
// Solidity: function ETHDKG_COMMITMENT_NOT_ON_CURVE() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesTransactorSession) ETHDKGCOMMITMENTNOTONCURVE() (*types.Transaction, error) {
	return _ETHDKGErrorCodes.Contract.ETHDKGCOMMITMENTNOTONCURVE(&_ETHDKGErrorCodes.TransactOpts)
}

// ETHDKGCOMMITMENTZERO is a paid mutator transaction binding the contract method 0x81687f80.
//
// Solidity: function ETHDKG_COMMITMENT_ZERO() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesTransactor) ETHDKGCOMMITMENTZERO(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ETHDKGErrorCodes.contract.Transact(opts, "ETHDKG_COMMITMENT_ZERO")
}

// ETHDKGCOMMITMENTZERO is a paid mutator transaction binding the contract method 0x81687f80.
//
// Solidity: function ETHDKG_COMMITMENT_ZERO() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesSession) ETHDKGCOMMITMENTZERO() (*types.Transaction, error) {
	return _ETHDKGErrorCodes.Contract.ETHDKGCOMMITMENTZERO(&_ETHDKGErrorCodes.TransactOpts)
}

// ETHDKGCOMMITMENTZERO is a paid mutator transaction binding the contract method 0x81687f80.
//
// Solidity: function ETHDKG_COMMITMENT_ZERO() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesTransactorSession) ETHDKGCOMMITMENTZERO() (*types.Transaction, error) {
	return _ETHDKGErrorCodes.Contract.ETHDKGCOMMITMENTZERO(&_ETHDKGErrorCodes.TransactOpts)
}

// ETHDKGDISPUTERDIDNOTDISTRIBUTESHARESINROUND is a paid mutator transaction binding the contract method 0xd65915e2.
//
// Solidity: function ETHDKG_DISPUTER_DID_NOT_DISTRIBUTE_SHARES_IN_ROUND() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesTransactor) ETHDKGDISPUTERDIDNOTDISTRIBUTESHARESINROUND(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ETHDKGErrorCodes.contract.Transact(opts, "ETHDKG_DISPUTER_DID_NOT_DISTRIBUTE_SHARES_IN_ROUND")
}

// ETHDKGDISPUTERDIDNOTDISTRIBUTESHARESINROUND is a paid mutator transaction binding the contract method 0xd65915e2.
//
// Solidity: function ETHDKG_DISPUTER_DID_NOT_DISTRIBUTE_SHARES_IN_ROUND() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesSession) ETHDKGDISPUTERDIDNOTDISTRIBUTESHARESINROUND() (*types.Transaction, error) {
	return _ETHDKGErrorCodes.Contract.ETHDKGDISPUTERDIDNOTDISTRIBUTESHARESINROUND(&_ETHDKGErrorCodes.TransactOpts)
}

// ETHDKGDISPUTERDIDNOTDISTRIBUTESHARESINROUND is a paid mutator transaction binding the contract method 0xd65915e2.
//
// Solidity: function ETHDKG_DISPUTER_DID_NOT_DISTRIBUTE_SHARES_IN_ROUND() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesTransactorSession) ETHDKGDISPUTERDIDNOTDISTRIBUTESHARESINROUND() (*types.Transaction, error) {
	return _ETHDKGErrorCodes.Contract.ETHDKGDISPUTERDIDNOTDISTRIBUTESHARESINROUND(&_ETHDKGErrorCodes.TransactOpts)
}

// ETHDKGDISPUTERDIDNOTSUBMITGPKJINROUND is a paid mutator transaction binding the contract method 0x3b2b8245.
//
// Solidity: function ETHDKG_DISPUTER_DID_NOT_SUBMIT_GPKJ_IN_ROUND() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesTransactor) ETHDKGDISPUTERDIDNOTSUBMITGPKJINROUND(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ETHDKGErrorCodes.contract.Transact(opts, "ETHDKG_DISPUTER_DID_NOT_SUBMIT_GPKJ_IN_ROUND")
}

// ETHDKGDISPUTERDIDNOTSUBMITGPKJINROUND is a paid mutator transaction binding the contract method 0x3b2b8245.
//
// Solidity: function ETHDKG_DISPUTER_DID_NOT_SUBMIT_GPKJ_IN_ROUND() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesSession) ETHDKGDISPUTERDIDNOTSUBMITGPKJINROUND() (*types.Transaction, error) {
	return _ETHDKGErrorCodes.Contract.ETHDKGDISPUTERDIDNOTSUBMITGPKJINROUND(&_ETHDKGErrorCodes.TransactOpts)
}

// ETHDKGDISPUTERDIDNOTSUBMITGPKJINROUND is a paid mutator transaction binding the contract method 0x3b2b8245.
//
// Solidity: function ETHDKG_DISPUTER_DID_NOT_SUBMIT_GPKJ_IN_ROUND() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesTransactorSession) ETHDKGDISPUTERDIDNOTSUBMITGPKJINROUND() (*types.Transaction, error) {
	return _ETHDKGErrorCodes.Contract.ETHDKGDISPUTERDIDNOTSUBMITGPKJINROUND(&_ETHDKGErrorCodes.TransactOpts)
}

// ETHDKGDISPUTERNOTPARTICIPATINGINROUND is a paid mutator transaction binding the contract method 0x763df93d.
//
// Solidity: function ETHDKG_DISPUTER_NOT_PARTICIPATING_IN_ROUND() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesTransactor) ETHDKGDISPUTERNOTPARTICIPATINGINROUND(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ETHDKGErrorCodes.contract.Transact(opts, "ETHDKG_DISPUTER_NOT_PARTICIPATING_IN_ROUND")
}

// ETHDKGDISPUTERNOTPARTICIPATINGINROUND is a paid mutator transaction binding the contract method 0x763df93d.
//
// Solidity: function ETHDKG_DISPUTER_NOT_PARTICIPATING_IN_ROUND() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesSession) ETHDKGDISPUTERNOTPARTICIPATINGINROUND() (*types.Transaction, error) {
	return _ETHDKGErrorCodes.Contract.ETHDKGDISPUTERNOTPARTICIPATINGINROUND(&_ETHDKGErrorCodes.TransactOpts)
}

// ETHDKGDISPUTERNOTPARTICIPATINGINROUND is a paid mutator transaction binding the contract method 0x763df93d.
//
// Solidity: function ETHDKG_DISPUTER_NOT_PARTICIPATING_IN_ROUND() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesTransactorSession) ETHDKGDISPUTERNOTPARTICIPATINGINROUND() (*types.Transaction, error) {
	return _ETHDKGErrorCodes.Contract.ETHDKGDISPUTERNOTPARTICIPATINGINROUND(&_ETHDKGErrorCodes.TransactOpts)
}

// ETHDKGDISTRIBUTEDSHAREHASHZERO is a paid mutator transaction binding the contract method 0xf54980c7.
//
// Solidity: function ETHDKG_DISTRIBUTED_SHARE_HASH_ZERO() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesTransactor) ETHDKGDISTRIBUTEDSHAREHASHZERO(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ETHDKGErrorCodes.contract.Transact(opts, "ETHDKG_DISTRIBUTED_SHARE_HASH_ZERO")
}

// ETHDKGDISTRIBUTEDSHAREHASHZERO is a paid mutator transaction binding the contract method 0xf54980c7.
//
// Solidity: function ETHDKG_DISTRIBUTED_SHARE_HASH_ZERO() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesSession) ETHDKGDISTRIBUTEDSHAREHASHZERO() (*types.Transaction, error) {
	return _ETHDKGErrorCodes.Contract.ETHDKGDISTRIBUTEDSHAREHASHZERO(&_ETHDKGErrorCodes.TransactOpts)
}

// ETHDKGDISTRIBUTEDSHAREHASHZERO is a paid mutator transaction binding the contract method 0xf54980c7.
//
// Solidity: function ETHDKG_DISTRIBUTED_SHARE_HASH_ZERO() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesTransactorSession) ETHDKGDISTRIBUTEDSHAREHASHZERO() (*types.Transaction, error) {
	return _ETHDKGErrorCodes.Contract.ETHDKGDISTRIBUTEDSHAREHASHZERO(&_ETHDKGErrorCodes.TransactOpts)
}

// ETHDKGGPKJZERO is a paid mutator transaction binding the contract method 0xc4e9cbe3.
//
// Solidity: function ETHDKG_GPKJ_ZERO() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesTransactor) ETHDKGGPKJZERO(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ETHDKGErrorCodes.contract.Transact(opts, "ETHDKG_GPKJ_ZERO")
}

// ETHDKGGPKJZERO is a paid mutator transaction binding the contract method 0xc4e9cbe3.
//
// Solidity: function ETHDKG_GPKJ_ZERO() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesSession) ETHDKGGPKJZERO() (*types.Transaction, error) {
	return _ETHDKGErrorCodes.Contract.ETHDKGGPKJZERO(&_ETHDKGErrorCodes.TransactOpts)
}

// ETHDKGGPKJZERO is a paid mutator transaction binding the contract method 0xc4e9cbe3.
//
// Solidity: function ETHDKG_GPKJ_ZERO() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesTransactorSession) ETHDKGGPKJZERO() (*types.Transaction, error) {
	return _ETHDKGErrorCodes.Contract.ETHDKGGPKJZERO(&_ETHDKGErrorCodes.TransactOpts)
}

// ETHDKGINVALIDARGS is a paid mutator transaction binding the contract method 0x4d76291d.
//
// Solidity: function ETHDKG_INVALID_ARGS() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesTransactor) ETHDKGINVALIDARGS(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ETHDKGErrorCodes.contract.Transact(opts, "ETHDKG_INVALID_ARGS")
}

// ETHDKGINVALIDARGS is a paid mutator transaction binding the contract method 0x4d76291d.
//
// Solidity: function ETHDKG_INVALID_ARGS() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesSession) ETHDKGINVALIDARGS() (*types.Transaction, error) {
	return _ETHDKGErrorCodes.Contract.ETHDKGINVALIDARGS(&_ETHDKGErrorCodes.TransactOpts)
}

// ETHDKGINVALIDARGS is a paid mutator transaction binding the contract method 0x4d76291d.
//
// Solidity: function ETHDKG_INVALID_ARGS() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesTransactorSession) ETHDKGINVALIDARGS() (*types.Transaction, error) {
	return _ETHDKGErrorCodes.Contract.ETHDKGINVALIDARGS(&_ETHDKGErrorCodes.TransactOpts)
}

// ETHDKGINVALIDCOMMITMENTS is a paid mutator transaction binding the contract method 0xf8fd7944.
//
// Solidity: function ETHDKG_INVALID_COMMITMENTS() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesTransactor) ETHDKGINVALIDCOMMITMENTS(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ETHDKGErrorCodes.contract.Transact(opts, "ETHDKG_INVALID_COMMITMENTS")
}

// ETHDKGINVALIDCOMMITMENTS is a paid mutator transaction binding the contract method 0xf8fd7944.
//
// Solidity: function ETHDKG_INVALID_COMMITMENTS() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesSession) ETHDKGINVALIDCOMMITMENTS() (*types.Transaction, error) {
	return _ETHDKGErrorCodes.Contract.ETHDKGINVALIDCOMMITMENTS(&_ETHDKGErrorCodes.TransactOpts)
}

// ETHDKGINVALIDCOMMITMENTS is a paid mutator transaction binding the contract method 0xf8fd7944.
//
// Solidity: function ETHDKG_INVALID_COMMITMENTS() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesTransactorSession) ETHDKGINVALIDCOMMITMENTS() (*types.Transaction, error) {
	return _ETHDKGErrorCodes.Contract.ETHDKGINVALIDCOMMITMENTS(&_ETHDKGErrorCodes.TransactOpts)
}

// ETHDKGINVALIDKEYSHAREG1 is a paid mutator transaction binding the contract method 0xdd35d7da.
//
// Solidity: function ETHDKG_INVALID_KEYSHARE_G1() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesTransactor) ETHDKGINVALIDKEYSHAREG1(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ETHDKGErrorCodes.contract.Transact(opts, "ETHDKG_INVALID_KEYSHARE_G1")
}

// ETHDKGINVALIDKEYSHAREG1 is a paid mutator transaction binding the contract method 0xdd35d7da.
//
// Solidity: function ETHDKG_INVALID_KEYSHARE_G1() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesSession) ETHDKGINVALIDKEYSHAREG1() (*types.Transaction, error) {
	return _ETHDKGErrorCodes.Contract.ETHDKGINVALIDKEYSHAREG1(&_ETHDKGErrorCodes.TransactOpts)
}

// ETHDKGINVALIDKEYSHAREG1 is a paid mutator transaction binding the contract method 0xdd35d7da.
//
// Solidity: function ETHDKG_INVALID_KEYSHARE_G1() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesTransactorSession) ETHDKGINVALIDKEYSHAREG1() (*types.Transaction, error) {
	return _ETHDKGErrorCodes.Contract.ETHDKGINVALIDKEYSHAREG1(&_ETHDKGErrorCodes.TransactOpts)
}

// ETHDKGINVALIDKEYSHAREG2 is a paid mutator transaction binding the contract method 0x8468c092.
//
// Solidity: function ETHDKG_INVALID_KEYSHARE_G2() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesTransactor) ETHDKGINVALIDKEYSHAREG2(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ETHDKGErrorCodes.contract.Transact(opts, "ETHDKG_INVALID_KEYSHARE_G2")
}

// ETHDKGINVALIDKEYSHAREG2 is a paid mutator transaction binding the contract method 0x8468c092.
//
// Solidity: function ETHDKG_INVALID_KEYSHARE_G2() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesSession) ETHDKGINVALIDKEYSHAREG2() (*types.Transaction, error) {
	return _ETHDKGErrorCodes.Contract.ETHDKGINVALIDKEYSHAREG2(&_ETHDKGErrorCodes.TransactOpts)
}

// ETHDKGINVALIDKEYSHAREG2 is a paid mutator transaction binding the contract method 0x8468c092.
//
// Solidity: function ETHDKG_INVALID_KEYSHARE_G2() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesTransactorSession) ETHDKGINVALIDKEYSHAREG2() (*types.Transaction, error) {
	return _ETHDKGErrorCodes.Contract.ETHDKGINVALIDKEYSHAREG2(&_ETHDKGErrorCodes.TransactOpts)
}

// ETHDKGINVALIDKEYORPROOF is a paid mutator transaction binding the contract method 0xa852713f.
//
// Solidity: function ETHDKG_INVALID_KEY_OR_PROOF() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesTransactor) ETHDKGINVALIDKEYORPROOF(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ETHDKGErrorCodes.contract.Transact(opts, "ETHDKG_INVALID_KEY_OR_PROOF")
}

// ETHDKGINVALIDKEYORPROOF is a paid mutator transaction binding the contract method 0xa852713f.
//
// Solidity: function ETHDKG_INVALID_KEY_OR_PROOF() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesSession) ETHDKGINVALIDKEYORPROOF() (*types.Transaction, error) {
	return _ETHDKGErrorCodes.Contract.ETHDKGINVALIDKEYORPROOF(&_ETHDKGErrorCodes.TransactOpts)
}

// ETHDKGINVALIDKEYORPROOF is a paid mutator transaction binding the contract method 0xa852713f.
//
// Solidity: function ETHDKG_INVALID_KEY_OR_PROOF() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesTransactorSession) ETHDKGINVALIDKEYORPROOF() (*types.Transaction, error) {
	return _ETHDKGErrorCodes.Contract.ETHDKGINVALIDKEYORPROOF(&_ETHDKGErrorCodes.TransactOpts)
}

// ETHDKGINVALIDNONCE is a paid mutator transaction binding the contract method 0x3341bcdf.
//
// Solidity: function ETHDKG_INVALID_NONCE() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesTransactor) ETHDKGINVALIDNONCE(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ETHDKGErrorCodes.contract.Transact(opts, "ETHDKG_INVALID_NONCE")
}

// ETHDKGINVALIDNONCE is a paid mutator transaction binding the contract method 0x3341bcdf.
//
// Solidity: function ETHDKG_INVALID_NONCE() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesSession) ETHDKGINVALIDNONCE() (*types.Transaction, error) {
	return _ETHDKGErrorCodes.Contract.ETHDKGINVALIDNONCE(&_ETHDKGErrorCodes.TransactOpts)
}

// ETHDKGINVALIDNONCE is a paid mutator transaction binding the contract method 0x3341bcdf.
//
// Solidity: function ETHDKG_INVALID_NONCE() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesTransactorSession) ETHDKGINVALIDNONCE() (*types.Transaction, error) {
	return _ETHDKGErrorCodes.Contract.ETHDKGINVALIDNONCE(&_ETHDKGErrorCodes.TransactOpts)
}

// ETHDKGINVALIDNUMCOMMITMENTS is a paid mutator transaction binding the contract method 0x0bca7264.
//
// Solidity: function ETHDKG_INVALID_NUM_COMMITMENTS() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesTransactor) ETHDKGINVALIDNUMCOMMITMENTS(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ETHDKGErrorCodes.contract.Transact(opts, "ETHDKG_INVALID_NUM_COMMITMENTS")
}

// ETHDKGINVALIDNUMCOMMITMENTS is a paid mutator transaction binding the contract method 0x0bca7264.
//
// Solidity: function ETHDKG_INVALID_NUM_COMMITMENTS() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesSession) ETHDKGINVALIDNUMCOMMITMENTS() (*types.Transaction, error) {
	return _ETHDKGErrorCodes.Contract.ETHDKGINVALIDNUMCOMMITMENTS(&_ETHDKGErrorCodes.TransactOpts)
}

// ETHDKGINVALIDNUMCOMMITMENTS is a paid mutator transaction binding the contract method 0x0bca7264.
//
// Solidity: function ETHDKG_INVALID_NUM_COMMITMENTS() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesTransactorSession) ETHDKGINVALIDNUMCOMMITMENTS() (*types.Transaction, error) {
	return _ETHDKGErrorCodes.Contract.ETHDKGINVALIDNUMCOMMITMENTS(&_ETHDKGErrorCodes.TransactOpts)
}

// ETHDKGINVALIDNUMENCRYPTEDSHARES is a paid mutator transaction binding the contract method 0xe8dcd67a.
//
// Solidity: function ETHDKG_INVALID_NUM_ENCRYPTED_SHARES() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesTransactor) ETHDKGINVALIDNUMENCRYPTEDSHARES(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ETHDKGErrorCodes.contract.Transact(opts, "ETHDKG_INVALID_NUM_ENCRYPTED_SHARES")
}

// ETHDKGINVALIDNUMENCRYPTEDSHARES is a paid mutator transaction binding the contract method 0xe8dcd67a.
//
// Solidity: function ETHDKG_INVALID_NUM_ENCRYPTED_SHARES() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesSession) ETHDKGINVALIDNUMENCRYPTEDSHARES() (*types.Transaction, error) {
	return _ETHDKGErrorCodes.Contract.ETHDKGINVALIDNUMENCRYPTEDSHARES(&_ETHDKGErrorCodes.TransactOpts)
}

// ETHDKGINVALIDNUMENCRYPTEDSHARES is a paid mutator transaction binding the contract method 0xe8dcd67a.
//
// Solidity: function ETHDKG_INVALID_NUM_ENCRYPTED_SHARES() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesTransactorSession) ETHDKGINVALIDNUMENCRYPTEDSHARES() (*types.Transaction, error) {
	return _ETHDKGErrorCodes.Contract.ETHDKGINVALIDNUMENCRYPTEDSHARES(&_ETHDKGErrorCodes.TransactOpts)
}

// ETHDKGINVALIDORDUPLICATEDPARTICIPANT is a paid mutator transaction binding the contract method 0x2f5e98de.
//
// Solidity: function ETHDKG_INVALID_OR_DUPLICATED_PARTICIPANT() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesTransactor) ETHDKGINVALIDORDUPLICATEDPARTICIPANT(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ETHDKGErrorCodes.contract.Transact(opts, "ETHDKG_INVALID_OR_DUPLICATED_PARTICIPANT")
}

// ETHDKGINVALIDORDUPLICATEDPARTICIPANT is a paid mutator transaction binding the contract method 0x2f5e98de.
//
// Solidity: function ETHDKG_INVALID_OR_DUPLICATED_PARTICIPANT() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesSession) ETHDKGINVALIDORDUPLICATEDPARTICIPANT() (*types.Transaction, error) {
	return _ETHDKGErrorCodes.Contract.ETHDKGINVALIDORDUPLICATEDPARTICIPANT(&_ETHDKGErrorCodes.TransactOpts)
}

// ETHDKGINVALIDORDUPLICATEDPARTICIPANT is a paid mutator transaction binding the contract method 0x2f5e98de.
//
// Solidity: function ETHDKG_INVALID_OR_DUPLICATED_PARTICIPANT() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesTransactorSession) ETHDKGINVALIDORDUPLICATEDPARTICIPANT() (*types.Transaction, error) {
	return _ETHDKGErrorCodes.Contract.ETHDKGINVALIDORDUPLICATEDPARTICIPANT(&_ETHDKGErrorCodes.TransactOpts)
}

// ETHDKGINVALIDSHARESORCOMMITMENTS is a paid mutator transaction binding the contract method 0x6d1e4818.
//
// Solidity: function ETHDKG_INVALID_SHARES_OR_COMMITMENTS() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesTransactor) ETHDKGINVALIDSHARESORCOMMITMENTS(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ETHDKGErrorCodes.contract.Transact(opts, "ETHDKG_INVALID_SHARES_OR_COMMITMENTS")
}

// ETHDKGINVALIDSHARESORCOMMITMENTS is a paid mutator transaction binding the contract method 0x6d1e4818.
//
// Solidity: function ETHDKG_INVALID_SHARES_OR_COMMITMENTS() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesSession) ETHDKGINVALIDSHARESORCOMMITMENTS() (*types.Transaction, error) {
	return _ETHDKGErrorCodes.Contract.ETHDKGINVALIDSHARESORCOMMITMENTS(&_ETHDKGErrorCodes.TransactOpts)
}

// ETHDKGINVALIDSHARESORCOMMITMENTS is a paid mutator transaction binding the contract method 0x6d1e4818.
//
// Solidity: function ETHDKG_INVALID_SHARES_OR_COMMITMENTS() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesTransactorSession) ETHDKGINVALIDSHARESORCOMMITMENTS() (*types.Transaction, error) {
	return _ETHDKGErrorCodes.Contract.ETHDKGINVALIDSHARESORCOMMITMENTS(&_ETHDKGErrorCodes.TransactOpts)
}

// ETHDKGKEYSHAREPHASEINVALIDNONCE is a paid mutator transaction binding the contract method 0xc169fbc4.
//
// Solidity: function ETHDKG_KEYSHARE_PHASE_INVALID_NONCE() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesTransactor) ETHDKGKEYSHAREPHASEINVALIDNONCE(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ETHDKGErrorCodes.contract.Transact(opts, "ETHDKG_KEYSHARE_PHASE_INVALID_NONCE")
}

// ETHDKGKEYSHAREPHASEINVALIDNONCE is a paid mutator transaction binding the contract method 0xc169fbc4.
//
// Solidity: function ETHDKG_KEYSHARE_PHASE_INVALID_NONCE() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesSession) ETHDKGKEYSHAREPHASEINVALIDNONCE() (*types.Transaction, error) {
	return _ETHDKGErrorCodes.Contract.ETHDKGKEYSHAREPHASEINVALIDNONCE(&_ETHDKGErrorCodes.TransactOpts)
}

// ETHDKGKEYSHAREPHASEINVALIDNONCE is a paid mutator transaction binding the contract method 0xc169fbc4.
//
// Solidity: function ETHDKG_KEYSHARE_PHASE_INVALID_NONCE() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesTransactorSession) ETHDKGKEYSHAREPHASEINVALIDNONCE() (*types.Transaction, error) {
	return _ETHDKGErrorCodes.Contract.ETHDKGKEYSHAREPHASEINVALIDNONCE(&_ETHDKGErrorCodes.TransactOpts)
}

// ETHDKGMASTERPUBLICKEYPAIRINGCHECKFAILURE is a paid mutator transaction binding the contract method 0x2ef8cd6e.
//
// Solidity: function ETHDKG_MASTER_PUBLIC_KEY_PAIRING_CHECK_FAILURE() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesTransactor) ETHDKGMASTERPUBLICKEYPAIRINGCHECKFAILURE(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ETHDKGErrorCodes.contract.Transact(opts, "ETHDKG_MASTER_PUBLIC_KEY_PAIRING_CHECK_FAILURE")
}

// ETHDKGMASTERPUBLICKEYPAIRINGCHECKFAILURE is a paid mutator transaction binding the contract method 0x2ef8cd6e.
//
// Solidity: function ETHDKG_MASTER_PUBLIC_KEY_PAIRING_CHECK_FAILURE() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesSession) ETHDKGMASTERPUBLICKEYPAIRINGCHECKFAILURE() (*types.Transaction, error) {
	return _ETHDKGErrorCodes.Contract.ETHDKGMASTERPUBLICKEYPAIRINGCHECKFAILURE(&_ETHDKGErrorCodes.TransactOpts)
}

// ETHDKGMASTERPUBLICKEYPAIRINGCHECKFAILURE is a paid mutator transaction binding the contract method 0x2ef8cd6e.
//
// Solidity: function ETHDKG_MASTER_PUBLIC_KEY_PAIRING_CHECK_FAILURE() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesTransactorSession) ETHDKGMASTERPUBLICKEYPAIRINGCHECKFAILURE() (*types.Transaction, error) {
	return _ETHDKGErrorCodes.Contract.ETHDKGMASTERPUBLICKEYPAIRINGCHECKFAILURE(&_ETHDKGErrorCodes.TransactOpts)
}

// ETHDKGMIGRATIONINPUTDATAMISMATCH is a paid mutator transaction binding the contract method 0x64a20638.
//
// Solidity: function ETHDKG_MIGRATION_INPUT_DATA_MISMATCH() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesTransactor) ETHDKGMIGRATIONINPUTDATAMISMATCH(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ETHDKGErrorCodes.contract.Transact(opts, "ETHDKG_MIGRATION_INPUT_DATA_MISMATCH")
}

// ETHDKGMIGRATIONINPUTDATAMISMATCH is a paid mutator transaction binding the contract method 0x64a20638.
//
// Solidity: function ETHDKG_MIGRATION_INPUT_DATA_MISMATCH() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesSession) ETHDKGMIGRATIONINPUTDATAMISMATCH() (*types.Transaction, error) {
	return _ETHDKGErrorCodes.Contract.ETHDKGMIGRATIONINPUTDATAMISMATCH(&_ETHDKGErrorCodes.TransactOpts)
}

// ETHDKGMIGRATIONINPUTDATAMISMATCH is a paid mutator transaction binding the contract method 0x64a20638.
//
// Solidity: function ETHDKG_MIGRATION_INPUT_DATA_MISMATCH() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesTransactorSession) ETHDKGMIGRATIONINPUTDATAMISMATCH() (*types.Transaction, error) {
	return _ETHDKGErrorCodes.Contract.ETHDKGMIGRATIONINPUTDATAMISMATCH(&_ETHDKGErrorCodes.TransactOpts)
}

// ETHDKGMIGRATIONINVALIDNONCE is a paid mutator transaction binding the contract method 0x06a621df.
//
// Solidity: function ETHDKG_MIGRATION_INVALID_NONCE() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesTransactor) ETHDKGMIGRATIONINVALIDNONCE(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ETHDKGErrorCodes.contract.Transact(opts, "ETHDKG_MIGRATION_INVALID_NONCE")
}

// ETHDKGMIGRATIONINVALIDNONCE is a paid mutator transaction binding the contract method 0x06a621df.
//
// Solidity: function ETHDKG_MIGRATION_INVALID_NONCE() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesSession) ETHDKGMIGRATIONINVALIDNONCE() (*types.Transaction, error) {
	return _ETHDKGErrorCodes.Contract.ETHDKGMIGRATIONINVALIDNONCE(&_ETHDKGErrorCodes.TransactOpts)
}

// ETHDKGMIGRATIONINVALIDNONCE is a paid mutator transaction binding the contract method 0x06a621df.
//
// Solidity: function ETHDKG_MIGRATION_INVALID_NONCE() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesTransactorSession) ETHDKGMIGRATIONINVALIDNONCE() (*types.Transaction, error) {
	return _ETHDKGErrorCodes.Contract.ETHDKGMIGRATIONINVALIDNONCE(&_ETHDKGErrorCodes.TransactOpts)
}

// ETHDKGMINVALIDATORSNOTMET is a paid mutator transaction binding the contract method 0x7e9f3983.
//
// Solidity: function ETHDKG_MIN_VALIDATORS_NOT_MET() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesTransactor) ETHDKGMINVALIDATORSNOTMET(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ETHDKGErrorCodes.contract.Transact(opts, "ETHDKG_MIN_VALIDATORS_NOT_MET")
}

// ETHDKGMINVALIDATORSNOTMET is a paid mutator transaction binding the contract method 0x7e9f3983.
//
// Solidity: function ETHDKG_MIN_VALIDATORS_NOT_MET() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesSession) ETHDKGMINVALIDATORSNOTMET() (*types.Transaction, error) {
	return _ETHDKGErrorCodes.Contract.ETHDKGMINVALIDATORSNOTMET(&_ETHDKGErrorCodes.TransactOpts)
}

// ETHDKGMINVALIDATORSNOTMET is a paid mutator transaction binding the contract method 0x7e9f3983.
//
// Solidity: function ETHDKG_MIN_VALIDATORS_NOT_MET() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesTransactorSession) ETHDKGMINVALIDATORSNOTMET() (*types.Transaction, error) {
	return _ETHDKGErrorCodes.Contract.ETHDKGMINVALIDATORSNOTMET(&_ETHDKGErrorCodes.TransactOpts)
}

// ETHDKGNOTINDISPUTEPHASE is a paid mutator transaction binding the contract method 0x4c4cfd75.
//
// Solidity: function ETHDKG_NOT_IN_DISPUTE_PHASE() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesTransactor) ETHDKGNOTINDISPUTEPHASE(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ETHDKGErrorCodes.contract.Transact(opts, "ETHDKG_NOT_IN_DISPUTE_PHASE")
}

// ETHDKGNOTINDISPUTEPHASE is a paid mutator transaction binding the contract method 0x4c4cfd75.
//
// Solidity: function ETHDKG_NOT_IN_DISPUTE_PHASE() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesSession) ETHDKGNOTINDISPUTEPHASE() (*types.Transaction, error) {
	return _ETHDKGErrorCodes.Contract.ETHDKGNOTINDISPUTEPHASE(&_ETHDKGErrorCodes.TransactOpts)
}

// ETHDKGNOTINDISPUTEPHASE is a paid mutator transaction binding the contract method 0x4c4cfd75.
//
// Solidity: function ETHDKG_NOT_IN_DISPUTE_PHASE() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesTransactorSession) ETHDKGNOTINDISPUTEPHASE() (*types.Transaction, error) {
	return _ETHDKGErrorCodes.Contract.ETHDKGNOTINDISPUTEPHASE(&_ETHDKGErrorCodes.TransactOpts)
}

// ETHDKGNOTINGPKJSUBMISSIONPHASE is a paid mutator transaction binding the contract method 0x53187f6a.
//
// Solidity: function ETHDKG_NOT_IN_GPKJ_SUBMISSION_PHASE() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesTransactor) ETHDKGNOTINGPKJSUBMISSIONPHASE(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ETHDKGErrorCodes.contract.Transact(opts, "ETHDKG_NOT_IN_GPKJ_SUBMISSION_PHASE")
}

// ETHDKGNOTINGPKJSUBMISSIONPHASE is a paid mutator transaction binding the contract method 0x53187f6a.
//
// Solidity: function ETHDKG_NOT_IN_GPKJ_SUBMISSION_PHASE() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesSession) ETHDKGNOTINGPKJSUBMISSIONPHASE() (*types.Transaction, error) {
	return _ETHDKGErrorCodes.Contract.ETHDKGNOTINGPKJSUBMISSIONPHASE(&_ETHDKGErrorCodes.TransactOpts)
}

// ETHDKGNOTINGPKJSUBMISSIONPHASE is a paid mutator transaction binding the contract method 0x53187f6a.
//
// Solidity: function ETHDKG_NOT_IN_GPKJ_SUBMISSION_PHASE() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesTransactorSession) ETHDKGNOTINGPKJSUBMISSIONPHASE() (*types.Transaction, error) {
	return _ETHDKGErrorCodes.Contract.ETHDKGNOTINGPKJSUBMISSIONPHASE(&_ETHDKGErrorCodes.TransactOpts)
}

// ETHDKGNOTINKEYSHARESUBMISSIONPHASE is a paid mutator transaction binding the contract method 0x8789ca22.
//
// Solidity: function ETHDKG_NOT_IN_KEYSHARE_SUBMISSION_PHASE() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesTransactor) ETHDKGNOTINKEYSHARESUBMISSIONPHASE(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ETHDKGErrorCodes.contract.Transact(opts, "ETHDKG_NOT_IN_KEYSHARE_SUBMISSION_PHASE")
}

// ETHDKGNOTINKEYSHARESUBMISSIONPHASE is a paid mutator transaction binding the contract method 0x8789ca22.
//
// Solidity: function ETHDKG_NOT_IN_KEYSHARE_SUBMISSION_PHASE() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesSession) ETHDKGNOTINKEYSHARESUBMISSIONPHASE() (*types.Transaction, error) {
	return _ETHDKGErrorCodes.Contract.ETHDKGNOTINKEYSHARESUBMISSIONPHASE(&_ETHDKGErrorCodes.TransactOpts)
}

// ETHDKGNOTINKEYSHARESUBMISSIONPHASE is a paid mutator transaction binding the contract method 0x8789ca22.
//
// Solidity: function ETHDKG_NOT_IN_KEYSHARE_SUBMISSION_PHASE() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesTransactorSession) ETHDKGNOTINKEYSHARESUBMISSIONPHASE() (*types.Transaction, error) {
	return _ETHDKGErrorCodes.Contract.ETHDKGNOTINKEYSHARESUBMISSIONPHASE(&_ETHDKGErrorCodes.TransactOpts)
}

// ETHDKGNOTINMASTERPUBLICKEYSUBMISSIONPHASE is a paid mutator transaction binding the contract method 0x3072e363.
//
// Solidity: function ETHDKG_NOT_IN_MASTER_PUBLIC_KEY_SUBMISSION_PHASE() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesTransactor) ETHDKGNOTINMASTERPUBLICKEYSUBMISSIONPHASE(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ETHDKGErrorCodes.contract.Transact(opts, "ETHDKG_NOT_IN_MASTER_PUBLIC_KEY_SUBMISSION_PHASE")
}

// ETHDKGNOTINMASTERPUBLICKEYSUBMISSIONPHASE is a paid mutator transaction binding the contract method 0x3072e363.
//
// Solidity: function ETHDKG_NOT_IN_MASTER_PUBLIC_KEY_SUBMISSION_PHASE() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesSession) ETHDKGNOTINMASTERPUBLICKEYSUBMISSIONPHASE() (*types.Transaction, error) {
	return _ETHDKGErrorCodes.Contract.ETHDKGNOTINMASTERPUBLICKEYSUBMISSIONPHASE(&_ETHDKGErrorCodes.TransactOpts)
}

// ETHDKGNOTINMASTERPUBLICKEYSUBMISSIONPHASE is a paid mutator transaction binding the contract method 0x3072e363.
//
// Solidity: function ETHDKG_NOT_IN_MASTER_PUBLIC_KEY_SUBMISSION_PHASE() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesTransactorSession) ETHDKGNOTINMASTERPUBLICKEYSUBMISSIONPHASE() (*types.Transaction, error) {
	return _ETHDKGErrorCodes.Contract.ETHDKGNOTINMASTERPUBLICKEYSUBMISSIONPHASE(&_ETHDKGErrorCodes.TransactOpts)
}

// ETHDKGNOTINPOSTGPKJDISPUTEPHASE is a paid mutator transaction binding the contract method 0xc4423781.
//
// Solidity: function ETHDKG_NOT_IN_POST_GPKJ_DISPUTE_PHASE() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesTransactor) ETHDKGNOTINPOSTGPKJDISPUTEPHASE(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ETHDKGErrorCodes.contract.Transact(opts, "ETHDKG_NOT_IN_POST_GPKJ_DISPUTE_PHASE")
}

// ETHDKGNOTINPOSTGPKJDISPUTEPHASE is a paid mutator transaction binding the contract method 0xc4423781.
//
// Solidity: function ETHDKG_NOT_IN_POST_GPKJ_DISPUTE_PHASE() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesSession) ETHDKGNOTINPOSTGPKJDISPUTEPHASE() (*types.Transaction, error) {
	return _ETHDKGErrorCodes.Contract.ETHDKGNOTINPOSTGPKJDISPUTEPHASE(&_ETHDKGErrorCodes.TransactOpts)
}

// ETHDKGNOTINPOSTGPKJDISPUTEPHASE is a paid mutator transaction binding the contract method 0xc4423781.
//
// Solidity: function ETHDKG_NOT_IN_POST_GPKJ_DISPUTE_PHASE() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesTransactorSession) ETHDKGNOTINPOSTGPKJDISPUTEPHASE() (*types.Transaction, error) {
	return _ETHDKGErrorCodes.Contract.ETHDKGNOTINPOSTGPKJDISPUTEPHASE(&_ETHDKGErrorCodes.TransactOpts)
}

// ETHDKGNOTINPOSTGPKJSUBMISSIONPHASE is a paid mutator transaction binding the contract method 0x6d429ef2.
//
// Solidity: function ETHDKG_NOT_IN_POST_GPKJ_SUBMISSION_PHASE() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesTransactor) ETHDKGNOTINPOSTGPKJSUBMISSIONPHASE(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ETHDKGErrorCodes.contract.Transact(opts, "ETHDKG_NOT_IN_POST_GPKJ_SUBMISSION_PHASE")
}

// ETHDKGNOTINPOSTGPKJSUBMISSIONPHASE is a paid mutator transaction binding the contract method 0x6d429ef2.
//
// Solidity: function ETHDKG_NOT_IN_POST_GPKJ_SUBMISSION_PHASE() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesSession) ETHDKGNOTINPOSTGPKJSUBMISSIONPHASE() (*types.Transaction, error) {
	return _ETHDKGErrorCodes.Contract.ETHDKGNOTINPOSTGPKJSUBMISSIONPHASE(&_ETHDKGErrorCodes.TransactOpts)
}

// ETHDKGNOTINPOSTGPKJSUBMISSIONPHASE is a paid mutator transaction binding the contract method 0x6d429ef2.
//
// Solidity: function ETHDKG_NOT_IN_POST_GPKJ_SUBMISSION_PHASE() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesTransactorSession) ETHDKGNOTINPOSTGPKJSUBMISSIONPHASE() (*types.Transaction, error) {
	return _ETHDKGErrorCodes.Contract.ETHDKGNOTINPOSTGPKJSUBMISSIONPHASE(&_ETHDKGErrorCodes.TransactOpts)
}

// ETHDKGNOTINPOSTKEYSHARESUBMISSIONPHASE is a paid mutator transaction binding the contract method 0x60987646.
//
// Solidity: function ETHDKG_NOT_IN_POST_KEYSHARE_SUBMISSION_PHASE() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesTransactor) ETHDKGNOTINPOSTKEYSHARESUBMISSIONPHASE(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ETHDKGErrorCodes.contract.Transact(opts, "ETHDKG_NOT_IN_POST_KEYSHARE_SUBMISSION_PHASE")
}

// ETHDKGNOTINPOSTKEYSHARESUBMISSIONPHASE is a paid mutator transaction binding the contract method 0x60987646.
//
// Solidity: function ETHDKG_NOT_IN_POST_KEYSHARE_SUBMISSION_PHASE() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesSession) ETHDKGNOTINPOSTKEYSHARESUBMISSIONPHASE() (*types.Transaction, error) {
	return _ETHDKGErrorCodes.Contract.ETHDKGNOTINPOSTKEYSHARESUBMISSIONPHASE(&_ETHDKGErrorCodes.TransactOpts)
}

// ETHDKGNOTINPOSTKEYSHARESUBMISSIONPHASE is a paid mutator transaction binding the contract method 0x60987646.
//
// Solidity: function ETHDKG_NOT_IN_POST_KEYSHARE_SUBMISSION_PHASE() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesTransactorSession) ETHDKGNOTINPOSTKEYSHARESUBMISSIONPHASE() (*types.Transaction, error) {
	return _ETHDKGErrorCodes.Contract.ETHDKGNOTINPOSTKEYSHARESUBMISSIONPHASE(&_ETHDKGErrorCodes.TransactOpts)
}

// ETHDKGNOTINPOSTREGISTRATIONPHASE is a paid mutator transaction binding the contract method 0x7385db5d.
//
// Solidity: function ETHDKG_NOT_IN_POST_REGISTRATION_PHASE() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesTransactor) ETHDKGNOTINPOSTREGISTRATIONPHASE(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ETHDKGErrorCodes.contract.Transact(opts, "ETHDKG_NOT_IN_POST_REGISTRATION_PHASE")
}

// ETHDKGNOTINPOSTREGISTRATIONPHASE is a paid mutator transaction binding the contract method 0x7385db5d.
//
// Solidity: function ETHDKG_NOT_IN_POST_REGISTRATION_PHASE() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesSession) ETHDKGNOTINPOSTREGISTRATIONPHASE() (*types.Transaction, error) {
	return _ETHDKGErrorCodes.Contract.ETHDKGNOTINPOSTREGISTRATIONPHASE(&_ETHDKGErrorCodes.TransactOpts)
}

// ETHDKGNOTINPOSTREGISTRATIONPHASE is a paid mutator transaction binding the contract method 0x7385db5d.
//
// Solidity: function ETHDKG_NOT_IN_POST_REGISTRATION_PHASE() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesTransactorSession) ETHDKGNOTINPOSTREGISTRATIONPHASE() (*types.Transaction, error) {
	return _ETHDKGErrorCodes.Contract.ETHDKGNOTINPOSTREGISTRATIONPHASE(&_ETHDKGErrorCodes.TransactOpts)
}

// ETHDKGNOTINPOSTSHAREDDISTRIBUTIONPHASE is a paid mutator transaction binding the contract method 0x8e25d1e1.
//
// Solidity: function ETHDKG_NOT_IN_POST_SHARED_DISTRIBUTION_PHASE() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesTransactor) ETHDKGNOTINPOSTSHAREDDISTRIBUTIONPHASE(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ETHDKGErrorCodes.contract.Transact(opts, "ETHDKG_NOT_IN_POST_SHARED_DISTRIBUTION_PHASE")
}

// ETHDKGNOTINPOSTSHAREDDISTRIBUTIONPHASE is a paid mutator transaction binding the contract method 0x8e25d1e1.
//
// Solidity: function ETHDKG_NOT_IN_POST_SHARED_DISTRIBUTION_PHASE() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesSession) ETHDKGNOTINPOSTSHAREDDISTRIBUTIONPHASE() (*types.Transaction, error) {
	return _ETHDKGErrorCodes.Contract.ETHDKGNOTINPOSTSHAREDDISTRIBUTIONPHASE(&_ETHDKGErrorCodes.TransactOpts)
}

// ETHDKGNOTINPOSTSHAREDDISTRIBUTIONPHASE is a paid mutator transaction binding the contract method 0x8e25d1e1.
//
// Solidity: function ETHDKG_NOT_IN_POST_SHARED_DISTRIBUTION_PHASE() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesTransactorSession) ETHDKGNOTINPOSTSHAREDDISTRIBUTIONPHASE() (*types.Transaction, error) {
	return _ETHDKGErrorCodes.Contract.ETHDKGNOTINPOSTSHAREDDISTRIBUTIONPHASE(&_ETHDKGErrorCodes.TransactOpts)
}

// ETHDKGNOTINREGISTRATIONPHASE is a paid mutator transaction binding the contract method 0xf4edda56.
//
// Solidity: function ETHDKG_NOT_IN_REGISTRATION_PHASE() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesTransactor) ETHDKGNOTINREGISTRATIONPHASE(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ETHDKGErrorCodes.contract.Transact(opts, "ETHDKG_NOT_IN_REGISTRATION_PHASE")
}

// ETHDKGNOTINREGISTRATIONPHASE is a paid mutator transaction binding the contract method 0xf4edda56.
//
// Solidity: function ETHDKG_NOT_IN_REGISTRATION_PHASE() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesSession) ETHDKGNOTINREGISTRATIONPHASE() (*types.Transaction, error) {
	return _ETHDKGErrorCodes.Contract.ETHDKGNOTINREGISTRATIONPHASE(&_ETHDKGErrorCodes.TransactOpts)
}

// ETHDKGNOTINREGISTRATIONPHASE is a paid mutator transaction binding the contract method 0xf4edda56.
//
// Solidity: function ETHDKG_NOT_IN_REGISTRATION_PHASE() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesTransactorSession) ETHDKGNOTINREGISTRATIONPHASE() (*types.Transaction, error) {
	return _ETHDKGErrorCodes.Contract.ETHDKGNOTINREGISTRATIONPHASE(&_ETHDKGErrorCodes.TransactOpts)
}

// ETHDKGNOTINSHAREDDISTRIBUTIONPHASE is a paid mutator transaction binding the contract method 0xba50ad20.
//
// Solidity: function ETHDKG_NOT_IN_SHARED_DISTRIBUTION_PHASE() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesTransactor) ETHDKGNOTINSHAREDDISTRIBUTIONPHASE(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ETHDKGErrorCodes.contract.Transact(opts, "ETHDKG_NOT_IN_SHARED_DISTRIBUTION_PHASE")
}

// ETHDKGNOTINSHAREDDISTRIBUTIONPHASE is a paid mutator transaction binding the contract method 0xba50ad20.
//
// Solidity: function ETHDKG_NOT_IN_SHARED_DISTRIBUTION_PHASE() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesSession) ETHDKGNOTINSHAREDDISTRIBUTIONPHASE() (*types.Transaction, error) {
	return _ETHDKGErrorCodes.Contract.ETHDKGNOTINSHAREDDISTRIBUTIONPHASE(&_ETHDKGErrorCodes.TransactOpts)
}

// ETHDKGNOTINSHAREDDISTRIBUTIONPHASE is a paid mutator transaction binding the contract method 0xba50ad20.
//
// Solidity: function ETHDKG_NOT_IN_SHARED_DISTRIBUTION_PHASE() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesTransactorSession) ETHDKGNOTINSHAREDDISTRIBUTIONPHASE() (*types.Transaction, error) {
	return _ETHDKGErrorCodes.Contract.ETHDKGNOTINSHAREDDISTRIBUTIONPHASE(&_ETHDKGErrorCodes.TransactOpts)
}

// ETHDKGONLYVALIDATORSALLOWED is a paid mutator transaction binding the contract method 0x83c069e4.
//
// Solidity: function ETHDKG_ONLY_VALIDATORS_ALLOWED() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesTransactor) ETHDKGONLYVALIDATORSALLOWED(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ETHDKGErrorCodes.contract.Transact(opts, "ETHDKG_ONLY_VALIDATORS_ALLOWED")
}

// ETHDKGONLYVALIDATORSALLOWED is a paid mutator transaction binding the contract method 0x83c069e4.
//
// Solidity: function ETHDKG_ONLY_VALIDATORS_ALLOWED() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesSession) ETHDKGONLYVALIDATORSALLOWED() (*types.Transaction, error) {
	return _ETHDKGErrorCodes.Contract.ETHDKGONLYVALIDATORSALLOWED(&_ETHDKGErrorCodes.TransactOpts)
}

// ETHDKGONLYVALIDATORSALLOWED is a paid mutator transaction binding the contract method 0x83c069e4.
//
// Solidity: function ETHDKG_ONLY_VALIDATORS_ALLOWED() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesTransactorSession) ETHDKGONLYVALIDATORSALLOWED() (*types.Transaction, error) {
	return _ETHDKGErrorCodes.Contract.ETHDKGONLYVALIDATORSALLOWED(&_ETHDKGErrorCodes.TransactOpts)
}

// ETHDKGPARTICIPANTDISTRIBUTEDSHARESINROUND is a paid mutator transaction binding the contract method 0xbdf1d0a8.
//
// Solidity: function ETHDKG_PARTICIPANT_DISTRIBUTED_SHARES_IN_ROUND() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesTransactor) ETHDKGPARTICIPANTDISTRIBUTEDSHARESINROUND(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ETHDKGErrorCodes.contract.Transact(opts, "ETHDKG_PARTICIPANT_DISTRIBUTED_SHARES_IN_ROUND")
}

// ETHDKGPARTICIPANTDISTRIBUTEDSHARESINROUND is a paid mutator transaction binding the contract method 0xbdf1d0a8.
//
// Solidity: function ETHDKG_PARTICIPANT_DISTRIBUTED_SHARES_IN_ROUND() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesSession) ETHDKGPARTICIPANTDISTRIBUTEDSHARESINROUND() (*types.Transaction, error) {
	return _ETHDKGErrorCodes.Contract.ETHDKGPARTICIPANTDISTRIBUTEDSHARESINROUND(&_ETHDKGErrorCodes.TransactOpts)
}

// ETHDKGPARTICIPANTDISTRIBUTEDSHARESINROUND is a paid mutator transaction binding the contract method 0xbdf1d0a8.
//
// Solidity: function ETHDKG_PARTICIPANT_DISTRIBUTED_SHARES_IN_ROUND() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesTransactorSession) ETHDKGPARTICIPANTDISTRIBUTEDSHARESINROUND() (*types.Transaction, error) {
	return _ETHDKGErrorCodes.Contract.ETHDKGPARTICIPANTDISTRIBUTEDSHARESINROUND(&_ETHDKGErrorCodes.TransactOpts)
}

// ETHDKGPARTICIPANTPARTICIPATINGINROUND is a paid mutator transaction binding the contract method 0x5bbeb11b.
//
// Solidity: function ETHDKG_PARTICIPANT_PARTICIPATING_IN_ROUND() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesTransactor) ETHDKGPARTICIPANTPARTICIPATINGINROUND(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ETHDKGErrorCodes.contract.Transact(opts, "ETHDKG_PARTICIPANT_PARTICIPATING_IN_ROUND")
}

// ETHDKGPARTICIPANTPARTICIPATINGINROUND is a paid mutator transaction binding the contract method 0x5bbeb11b.
//
// Solidity: function ETHDKG_PARTICIPANT_PARTICIPATING_IN_ROUND() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesSession) ETHDKGPARTICIPANTPARTICIPATINGINROUND() (*types.Transaction, error) {
	return _ETHDKGErrorCodes.Contract.ETHDKGPARTICIPANTPARTICIPATINGINROUND(&_ETHDKGErrorCodes.TransactOpts)
}

// ETHDKGPARTICIPANTPARTICIPATINGINROUND is a paid mutator transaction binding the contract method 0x5bbeb11b.
//
// Solidity: function ETHDKG_PARTICIPANT_PARTICIPATING_IN_ROUND() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesTransactorSession) ETHDKGPARTICIPANTPARTICIPATINGINROUND() (*types.Transaction, error) {
	return _ETHDKGErrorCodes.Contract.ETHDKGPARTICIPANTPARTICIPATINGINROUND(&_ETHDKGErrorCodes.TransactOpts)
}

// ETHDKGPARTICIPANTSUBMITTEDGPKJINROUND is a paid mutator transaction binding the contract method 0xc43f133a.
//
// Solidity: function ETHDKG_PARTICIPANT_SUBMITTED_GPKJ_IN_ROUND() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesTransactor) ETHDKGPARTICIPANTSUBMITTEDGPKJINROUND(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ETHDKGErrorCodes.contract.Transact(opts, "ETHDKG_PARTICIPANT_SUBMITTED_GPKJ_IN_ROUND")
}

// ETHDKGPARTICIPANTSUBMITTEDGPKJINROUND is a paid mutator transaction binding the contract method 0xc43f133a.
//
// Solidity: function ETHDKG_PARTICIPANT_SUBMITTED_GPKJ_IN_ROUND() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesSession) ETHDKGPARTICIPANTSUBMITTEDGPKJINROUND() (*types.Transaction, error) {
	return _ETHDKGErrorCodes.Contract.ETHDKGPARTICIPANTSUBMITTEDGPKJINROUND(&_ETHDKGErrorCodes.TransactOpts)
}

// ETHDKGPARTICIPANTSUBMITTEDGPKJINROUND is a paid mutator transaction binding the contract method 0xc43f133a.
//
// Solidity: function ETHDKG_PARTICIPANT_SUBMITTED_GPKJ_IN_ROUND() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesTransactorSession) ETHDKGPARTICIPANTSUBMITTEDGPKJINROUND() (*types.Transaction, error) {
	return _ETHDKGErrorCodes.Contract.ETHDKGPARTICIPANTSUBMITTEDGPKJINROUND(&_ETHDKGErrorCodes.TransactOpts)
}

// ETHDKGPARTICIPANTSUBMITTEDKEYSHARESINROUND is a paid mutator transaction binding the contract method 0x1723b084.
//
// Solidity: function ETHDKG_PARTICIPANT_SUBMITTED_KEYSHARES_IN_ROUND() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesTransactor) ETHDKGPARTICIPANTSUBMITTEDKEYSHARESINROUND(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ETHDKGErrorCodes.contract.Transact(opts, "ETHDKG_PARTICIPANT_SUBMITTED_KEYSHARES_IN_ROUND")
}

// ETHDKGPARTICIPANTSUBMITTEDKEYSHARESINROUND is a paid mutator transaction binding the contract method 0x1723b084.
//
// Solidity: function ETHDKG_PARTICIPANT_SUBMITTED_KEYSHARES_IN_ROUND() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesSession) ETHDKGPARTICIPANTSUBMITTEDKEYSHARESINROUND() (*types.Transaction, error) {
	return _ETHDKGErrorCodes.Contract.ETHDKGPARTICIPANTSUBMITTEDKEYSHARESINROUND(&_ETHDKGErrorCodes.TransactOpts)
}

// ETHDKGPARTICIPANTSUBMITTEDKEYSHARESINROUND is a paid mutator transaction binding the contract method 0x1723b084.
//
// Solidity: function ETHDKG_PARTICIPANT_SUBMITTED_KEYSHARES_IN_ROUND() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesTransactorSession) ETHDKGPARTICIPANTSUBMITTEDKEYSHARESINROUND() (*types.Transaction, error) {
	return _ETHDKGErrorCodes.Contract.ETHDKGPARTICIPANTSUBMITTEDKEYSHARESINROUND(&_ETHDKGErrorCodes.TransactOpts)
}

// ETHDKGPUBLICKEYNOTONCURVE is a paid mutator transaction binding the contract method 0x29896fd4.
//
// Solidity: function ETHDKG_PUBLIC_KEY_NOT_ON_CURVE() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesTransactor) ETHDKGPUBLICKEYNOTONCURVE(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ETHDKGErrorCodes.contract.Transact(opts, "ETHDKG_PUBLIC_KEY_NOT_ON_CURVE")
}

// ETHDKGPUBLICKEYNOTONCURVE is a paid mutator transaction binding the contract method 0x29896fd4.
//
// Solidity: function ETHDKG_PUBLIC_KEY_NOT_ON_CURVE() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesSession) ETHDKGPUBLICKEYNOTONCURVE() (*types.Transaction, error) {
	return _ETHDKGErrorCodes.Contract.ETHDKGPUBLICKEYNOTONCURVE(&_ETHDKGErrorCodes.TransactOpts)
}

// ETHDKGPUBLICKEYNOTONCURVE is a paid mutator transaction binding the contract method 0x29896fd4.
//
// Solidity: function ETHDKG_PUBLIC_KEY_NOT_ON_CURVE() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesTransactorSession) ETHDKGPUBLICKEYNOTONCURVE() (*types.Transaction, error) {
	return _ETHDKGErrorCodes.Contract.ETHDKGPUBLICKEYNOTONCURVE(&_ETHDKGErrorCodes.TransactOpts)
}

// ETHDKGPUBLICKEYZERO is a paid mutator transaction binding the contract method 0x54ad808e.
//
// Solidity: function ETHDKG_PUBLIC_KEY_ZERO() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesTransactor) ETHDKGPUBLICKEYZERO(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ETHDKGErrorCodes.contract.Transact(opts, "ETHDKG_PUBLIC_KEY_ZERO")
}

// ETHDKGPUBLICKEYZERO is a paid mutator transaction binding the contract method 0x54ad808e.
//
// Solidity: function ETHDKG_PUBLIC_KEY_ZERO() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesSession) ETHDKGPUBLICKEYZERO() (*types.Transaction, error) {
	return _ETHDKGErrorCodes.Contract.ETHDKGPUBLICKEYZERO(&_ETHDKGErrorCodes.TransactOpts)
}

// ETHDKGPUBLICKEYZERO is a paid mutator transaction binding the contract method 0x54ad808e.
//
// Solidity: function ETHDKG_PUBLIC_KEY_ZERO() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesTransactorSession) ETHDKGPUBLICKEYZERO() (*types.Transaction, error) {
	return _ETHDKGErrorCodes.Contract.ETHDKGPUBLICKEYZERO(&_ETHDKGErrorCodes.TransactOpts)
}

// ETHDKGREQUISITESINCOMPLETE is a paid mutator transaction binding the contract method 0x23277f87.
//
// Solidity: function ETHDKG_REQUISITES_INCOMPLETE() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesTransactor) ETHDKGREQUISITESINCOMPLETE(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ETHDKGErrorCodes.contract.Transact(opts, "ETHDKG_REQUISITES_INCOMPLETE")
}

// ETHDKGREQUISITESINCOMPLETE is a paid mutator transaction binding the contract method 0x23277f87.
//
// Solidity: function ETHDKG_REQUISITES_INCOMPLETE() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesSession) ETHDKGREQUISITESINCOMPLETE() (*types.Transaction, error) {
	return _ETHDKGErrorCodes.Contract.ETHDKGREQUISITESINCOMPLETE(&_ETHDKGErrorCodes.TransactOpts)
}

// ETHDKGREQUISITESINCOMPLETE is a paid mutator transaction binding the contract method 0x23277f87.
//
// Solidity: function ETHDKG_REQUISITES_INCOMPLETE() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesTransactorSession) ETHDKGREQUISITESINCOMPLETE() (*types.Transaction, error) {
	return _ETHDKGErrorCodes.Contract.ETHDKGREQUISITESINCOMPLETE(&_ETHDKGErrorCodes.TransactOpts)
}

// ETHDKGSHARESANDCOMMITMENTSMISMATCH is a paid mutator transaction binding the contract method 0xe0c1afd8.
//
// Solidity: function ETHDKG_SHARES_AND_COMMITMENTS_MISMATCH() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesTransactor) ETHDKGSHARESANDCOMMITMENTSMISMATCH(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ETHDKGErrorCodes.contract.Transact(opts, "ETHDKG_SHARES_AND_COMMITMENTS_MISMATCH")
}

// ETHDKGSHARESANDCOMMITMENTSMISMATCH is a paid mutator transaction binding the contract method 0xe0c1afd8.
//
// Solidity: function ETHDKG_SHARES_AND_COMMITMENTS_MISMATCH() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesSession) ETHDKGSHARESANDCOMMITMENTSMISMATCH() (*types.Transaction, error) {
	return _ETHDKGErrorCodes.Contract.ETHDKGSHARESANDCOMMITMENTSMISMATCH(&_ETHDKGErrorCodes.TransactOpts)
}

// ETHDKGSHARESANDCOMMITMENTSMISMATCH is a paid mutator transaction binding the contract method 0xe0c1afd8.
//
// Solidity: function ETHDKG_SHARES_AND_COMMITMENTS_MISMATCH() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesTransactorSession) ETHDKGSHARESANDCOMMITMENTSMISMATCH() (*types.Transaction, error) {
	return _ETHDKGErrorCodes.Contract.ETHDKGSHARESANDCOMMITMENTSMISMATCH(&_ETHDKGErrorCodes.TransactOpts)
}

// ETHDKGVARIABLECANNOTBESETWHILERUNNING is a paid mutator transaction binding the contract method 0x79ec8296.
//
// Solidity: function ETHDKG_VARIABLE_CANNOT_BE_SET_WHILE_RUNNING() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesTransactor) ETHDKGVARIABLECANNOTBESETWHILERUNNING(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ETHDKGErrorCodes.contract.Transact(opts, "ETHDKG_VARIABLE_CANNOT_BE_SET_WHILE_RUNNING")
}

// ETHDKGVARIABLECANNOTBESETWHILERUNNING is a paid mutator transaction binding the contract method 0x79ec8296.
//
// Solidity: function ETHDKG_VARIABLE_CANNOT_BE_SET_WHILE_RUNNING() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesSession) ETHDKGVARIABLECANNOTBESETWHILERUNNING() (*types.Transaction, error) {
	return _ETHDKGErrorCodes.Contract.ETHDKGVARIABLECANNOTBESETWHILERUNNING(&_ETHDKGErrorCodes.TransactOpts)
}

// ETHDKGVARIABLECANNOTBESETWHILERUNNING is a paid mutator transaction binding the contract method 0x79ec8296.
//
// Solidity: function ETHDKG_VARIABLE_CANNOT_BE_SET_WHILE_RUNNING() returns(bytes32)
func (_ETHDKGErrorCodes *ETHDKGErrorCodesTransactorSession) ETHDKGVARIABLECANNOTBESETWHILERUNNING() (*types.Transaction, error) {
	return _ETHDKGErrorCodes.Contract.ETHDKGVARIABLECANNOTBESETWHILERUNNING(&_ETHDKGErrorCodes.TransactOpts)
}
