// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package controller

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

// SharedStructsAgreement is an auto generated low-level Go binding around an user-defined struct.
type SharedStructsAgreement struct {
	State                    uint8
	ResourceProviderAgreedAt *big.Int
	JobCreatorAgreedAt       *big.Int
	DealCreatedAt            *big.Int
	DealAgreedAt             *big.Int
	ResultsSubmittedAt       *big.Int
	ResultsAcceptedAt        *big.Int
	ResultsCheckedAt         *big.Int
	MediationAcceptedAt      *big.Int
	MediationRejectedAt      *big.Int
	TimeoutAgreeAt           *big.Int
	TimeoutSubmitResultsAt   *big.Int
	TimeoutJudgeResultsAt    *big.Int
	TimeoutMediateResultsAt  *big.Int
}

// SharedStructsDealMembers is an auto generated low-level Go binding around an user-defined struct.
type SharedStructsDealMembers struct {
	Solver           common.Address
	JobCreator       common.Address
	ResourceProvider common.Address
	Mediators        []common.Address
}

// SharedStructsDealPricing is an auto generated low-level Go binding around an user-defined struct.
type SharedStructsDealPricing struct {
	InstructionPrice          *big.Int
	PaymentCollateral         *big.Int
	ResultsCollateralMultiple *big.Int
	MediationFee              *big.Int
}

// SharedStructsDealTimeout is an auto generated low-level Go binding around an user-defined struct.
type SharedStructsDealTimeout struct {
	Timeout    *big.Int
	Collateral *big.Int
}

// SharedStructsDealTimeouts is an auto generated low-level Go binding around an user-defined struct.
type SharedStructsDealTimeouts struct {
	Agree          SharedStructsDealTimeout
	SubmitResults  SharedStructsDealTimeout
	JudgeResults   SharedStructsDealTimeout
	MediateResults SharedStructsDealTimeout
}

// ControllerMetaData contains all meta data concerning the Controller contract.
var ControllerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"}],\"name\":\"acceptResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"resultsId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"dataId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"instructionCount\",\"type\":\"uint256\"}],\"name\":\"addResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"solver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"jobCreator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"resourceProvider\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"mediators\",\"type\":\"address[]\"}],\"internalType\":\"structSharedStructs.DealMembers\",\"name\":\"members\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealTimeout\",\"name\":\"agree\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealTimeout\",\"name\":\"submitResults\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealTimeout\",\"name\":\"judgeResults\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealTimeout\",\"name\":\"mediateResults\",\"type\":\"tuple\"}],\"internalType\":\"structSharedStructs.DealTimeouts\",\"name\":\"timeouts\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"instructionPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paymentCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsCollateralMultiple\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationFee\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealPricing\",\"name\":\"pricing\",\"type\":\"tuple\"}],\"name\":\"agree\",\"outputs\":[{\"components\":[{\"internalType\":\"enumSharedStructs.AgreementState\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"resourceProviderAgreedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"jobCreatorAgreedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dealCreatedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dealAgreedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsSubmittedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsAcceptedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsCheckedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationAcceptedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationRejectedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutAgreeAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutSubmitResultsAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutJudgeResultsAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutMediateResultsAt\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.Agreement\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"}],\"name\":\"checkResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getJobCreatorAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMediationAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPaymentsAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStorageAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUsersAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_storageAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_usersAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_paymentsAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_mediationAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_jobCreatorAddress\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"}],\"name\":\"mediationAcceptResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"}],\"name\":\"mediationRejectResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_jobCreatorAddress\",\"type\":\"address\"}],\"name\":\"setJobCreatorAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_mediationAddress\",\"type\":\"address\"}],\"name\":\"setMediationAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_paymentsAddress\",\"type\":\"address\"}],\"name\":\"setPaymentsAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_storageAddress\",\"type\":\"address\"}],\"name\":\"setStorageAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_usersAddress\",\"type\":\"address\"}],\"name\":\"setUsersAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"}],\"name\":\"timeoutAgree\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"}],\"name\":\"timeoutJudgeResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"}],\"name\":\"timeoutMediateResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"}],\"name\":\"timeoutSubmitResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ControllerABI is the input ABI used to generate the binding from.
// Deprecated: Use ControllerMetaData.ABI instead.
var ControllerABI = ControllerMetaData.ABI

// Controller is an auto generated Go binding around an Ethereum contract.
type Controller struct {
	ControllerCaller     // Read-only binding to the contract
	ControllerTransactor // Write-only binding to the contract
	ControllerFilterer   // Log filterer for contract events
}

// ControllerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ControllerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ControllerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ControllerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ControllerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ControllerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ControllerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ControllerSession struct {
	Contract     *Controller       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ControllerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ControllerCallerSession struct {
	Contract *ControllerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ControllerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ControllerTransactorSession struct {
	Contract     *ControllerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ControllerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ControllerRaw struct {
	Contract *Controller // Generic contract binding to access the raw methods on
}

// ControllerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ControllerCallerRaw struct {
	Contract *ControllerCaller // Generic read-only contract binding to access the raw methods on
}

// ControllerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ControllerTransactorRaw struct {
	Contract *ControllerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewController creates a new instance of Controller, bound to a specific deployed contract.
func NewController(address common.Address, backend bind.ContractBackend) (*Controller, error) {
	contract, err := bindController(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Controller{ControllerCaller: ControllerCaller{contract: contract}, ControllerTransactor: ControllerTransactor{contract: contract}, ControllerFilterer: ControllerFilterer{contract: contract}}, nil
}

// NewControllerCaller creates a new read-only instance of Controller, bound to a specific deployed contract.
func NewControllerCaller(address common.Address, caller bind.ContractCaller) (*ControllerCaller, error) {
	contract, err := bindController(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ControllerCaller{contract: contract}, nil
}

// NewControllerTransactor creates a new write-only instance of Controller, bound to a specific deployed contract.
func NewControllerTransactor(address common.Address, transactor bind.ContractTransactor) (*ControllerTransactor, error) {
	contract, err := bindController(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ControllerTransactor{contract: contract}, nil
}

// NewControllerFilterer creates a new log filterer instance of Controller, bound to a specific deployed contract.
func NewControllerFilterer(address common.Address, filterer bind.ContractFilterer) (*ControllerFilterer, error) {
	contract, err := bindController(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ControllerFilterer{contract: contract}, nil
}

// bindController binds a generic wrapper to an already deployed contract.
func bindController(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ControllerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Controller *ControllerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Controller.Contract.ControllerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Controller *ControllerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Controller.Contract.ControllerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Controller *ControllerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Controller.Contract.ControllerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Controller *ControllerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Controller.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Controller *ControllerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Controller.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Controller *ControllerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Controller.Contract.contract.Transact(opts, method, params...)
}

// GetJobCreatorAddress is a free data retrieval call binding the contract method 0x0aca35ce.
//
// Solidity: function getJobCreatorAddress() view returns(address)
func (_Controller *ControllerCaller) GetJobCreatorAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "getJobCreatorAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetJobCreatorAddress is a free data retrieval call binding the contract method 0x0aca35ce.
//
// Solidity: function getJobCreatorAddress() view returns(address)
func (_Controller *ControllerSession) GetJobCreatorAddress() (common.Address, error) {
	return _Controller.Contract.GetJobCreatorAddress(&_Controller.CallOpts)
}

// GetJobCreatorAddress is a free data retrieval call binding the contract method 0x0aca35ce.
//
// Solidity: function getJobCreatorAddress() view returns(address)
func (_Controller *ControllerCallerSession) GetJobCreatorAddress() (common.Address, error) {
	return _Controller.Contract.GetJobCreatorAddress(&_Controller.CallOpts)
}

// GetMediationAddress is a free data retrieval call binding the contract method 0x155329ea.
//
// Solidity: function getMediationAddress() view returns(address)
func (_Controller *ControllerCaller) GetMediationAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "getMediationAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetMediationAddress is a free data retrieval call binding the contract method 0x155329ea.
//
// Solidity: function getMediationAddress() view returns(address)
func (_Controller *ControllerSession) GetMediationAddress() (common.Address, error) {
	return _Controller.Contract.GetMediationAddress(&_Controller.CallOpts)
}

// GetMediationAddress is a free data retrieval call binding the contract method 0x155329ea.
//
// Solidity: function getMediationAddress() view returns(address)
func (_Controller *ControllerCallerSession) GetMediationAddress() (common.Address, error) {
	return _Controller.Contract.GetMediationAddress(&_Controller.CallOpts)
}

// GetPaymentsAddress is a free data retrieval call binding the contract method 0xd48b1084.
//
// Solidity: function getPaymentsAddress() view returns(address)
func (_Controller *ControllerCaller) GetPaymentsAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "getPaymentsAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPaymentsAddress is a free data retrieval call binding the contract method 0xd48b1084.
//
// Solidity: function getPaymentsAddress() view returns(address)
func (_Controller *ControllerSession) GetPaymentsAddress() (common.Address, error) {
	return _Controller.Contract.GetPaymentsAddress(&_Controller.CallOpts)
}

// GetPaymentsAddress is a free data retrieval call binding the contract method 0xd48b1084.
//
// Solidity: function getPaymentsAddress() view returns(address)
func (_Controller *ControllerCallerSession) GetPaymentsAddress() (common.Address, error) {
	return _Controller.Contract.GetPaymentsAddress(&_Controller.CallOpts)
}

// GetStorageAddress is a free data retrieval call binding the contract method 0x393a4d34.
//
// Solidity: function getStorageAddress() view returns(address)
func (_Controller *ControllerCaller) GetStorageAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "getStorageAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetStorageAddress is a free data retrieval call binding the contract method 0x393a4d34.
//
// Solidity: function getStorageAddress() view returns(address)
func (_Controller *ControllerSession) GetStorageAddress() (common.Address, error) {
	return _Controller.Contract.GetStorageAddress(&_Controller.CallOpts)
}

// GetStorageAddress is a free data retrieval call binding the contract method 0x393a4d34.
//
// Solidity: function getStorageAddress() view returns(address)
func (_Controller *ControllerCallerSession) GetStorageAddress() (common.Address, error) {
	return _Controller.Contract.GetStorageAddress(&_Controller.CallOpts)
}

// GetUsersAddress is a free data retrieval call binding the contract method 0x93dbed3e.
//
// Solidity: function getUsersAddress() view returns(address)
func (_Controller *ControllerCaller) GetUsersAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "getUsersAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetUsersAddress is a free data retrieval call binding the contract method 0x93dbed3e.
//
// Solidity: function getUsersAddress() view returns(address)
func (_Controller *ControllerSession) GetUsersAddress() (common.Address, error) {
	return _Controller.Contract.GetUsersAddress(&_Controller.CallOpts)
}

// GetUsersAddress is a free data retrieval call binding the contract method 0x93dbed3e.
//
// Solidity: function getUsersAddress() view returns(address)
func (_Controller *ControllerCallerSession) GetUsersAddress() (common.Address, error) {
	return _Controller.Contract.GetUsersAddress(&_Controller.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Controller *ControllerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Controller *ControllerSession) Owner() (common.Address, error) {
	return _Controller.Contract.Owner(&_Controller.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Controller *ControllerCallerSession) Owner() (common.Address, error) {
	return _Controller.Contract.Owner(&_Controller.CallOpts)
}

// AcceptResult is a paid mutator transaction binding the contract method 0x297f9e55.
//
// Solidity: function acceptResult(string dealId) returns()
func (_Controller *ControllerTransactor) AcceptResult(opts *bind.TransactOpts, dealId string) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "acceptResult", dealId)
}

// AcceptResult is a paid mutator transaction binding the contract method 0x297f9e55.
//
// Solidity: function acceptResult(string dealId) returns()
func (_Controller *ControllerSession) AcceptResult(dealId string) (*types.Transaction, error) {
	return _Controller.Contract.AcceptResult(&_Controller.TransactOpts, dealId)
}

// AcceptResult is a paid mutator transaction binding the contract method 0x297f9e55.
//
// Solidity: function acceptResult(string dealId) returns()
func (_Controller *ControllerTransactorSession) AcceptResult(dealId string) (*types.Transaction, error) {
	return _Controller.Contract.AcceptResult(&_Controller.TransactOpts, dealId)
}

// AddResult is a paid mutator transaction binding the contract method 0x3955548e.
//
// Solidity: function addResult(string dealId, string resultsId, string dataId, uint256 instructionCount) returns()
func (_Controller *ControllerTransactor) AddResult(opts *bind.TransactOpts, dealId string, resultsId string, dataId string, instructionCount *big.Int) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "addResult", dealId, resultsId, dataId, instructionCount)
}

// AddResult is a paid mutator transaction binding the contract method 0x3955548e.
//
// Solidity: function addResult(string dealId, string resultsId, string dataId, uint256 instructionCount) returns()
func (_Controller *ControllerSession) AddResult(dealId string, resultsId string, dataId string, instructionCount *big.Int) (*types.Transaction, error) {
	return _Controller.Contract.AddResult(&_Controller.TransactOpts, dealId, resultsId, dataId, instructionCount)
}

// AddResult is a paid mutator transaction binding the contract method 0x3955548e.
//
// Solidity: function addResult(string dealId, string resultsId, string dataId, uint256 instructionCount) returns()
func (_Controller *ControllerTransactorSession) AddResult(dealId string, resultsId string, dataId string, instructionCount *big.Int) (*types.Transaction, error) {
	return _Controller.Contract.AddResult(&_Controller.TransactOpts, dealId, resultsId, dataId, instructionCount)
}

// Agree is a paid mutator transaction binding the contract method 0xf583b125.
//
// Solidity: function agree(string dealId, (address,address,address,address[]) members, ((uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256,uint256)) timeouts, (uint256,uint256,uint256,uint256) pricing) returns((uint8,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_Controller *ControllerTransactor) Agree(opts *bind.TransactOpts, dealId string, members SharedStructsDealMembers, timeouts SharedStructsDealTimeouts, pricing SharedStructsDealPricing) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "agree", dealId, members, timeouts, pricing)
}

// Agree is a paid mutator transaction binding the contract method 0xf583b125.
//
// Solidity: function agree(string dealId, (address,address,address,address[]) members, ((uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256,uint256)) timeouts, (uint256,uint256,uint256,uint256) pricing) returns((uint8,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_Controller *ControllerSession) Agree(dealId string, members SharedStructsDealMembers, timeouts SharedStructsDealTimeouts, pricing SharedStructsDealPricing) (*types.Transaction, error) {
	return _Controller.Contract.Agree(&_Controller.TransactOpts, dealId, members, timeouts, pricing)
}

// Agree is a paid mutator transaction binding the contract method 0xf583b125.
//
// Solidity: function agree(string dealId, (address,address,address,address[]) members, ((uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256,uint256)) timeouts, (uint256,uint256,uint256,uint256) pricing) returns((uint8,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_Controller *ControllerTransactorSession) Agree(dealId string, members SharedStructsDealMembers, timeouts SharedStructsDealTimeouts, pricing SharedStructsDealPricing) (*types.Transaction, error) {
	return _Controller.Contract.Agree(&_Controller.TransactOpts, dealId, members, timeouts, pricing)
}

// CheckResult is a paid mutator transaction binding the contract method 0x46834d1e.
//
// Solidity: function checkResult(string dealId) returns()
func (_Controller *ControllerTransactor) CheckResult(opts *bind.TransactOpts, dealId string) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "checkResult", dealId)
}

// CheckResult is a paid mutator transaction binding the contract method 0x46834d1e.
//
// Solidity: function checkResult(string dealId) returns()
func (_Controller *ControllerSession) CheckResult(dealId string) (*types.Transaction, error) {
	return _Controller.Contract.CheckResult(&_Controller.TransactOpts, dealId)
}

// CheckResult is a paid mutator transaction binding the contract method 0x46834d1e.
//
// Solidity: function checkResult(string dealId) returns()
func (_Controller *ControllerTransactorSession) CheckResult(dealId string) (*types.Transaction, error) {
	return _Controller.Contract.CheckResult(&_Controller.TransactOpts, dealId)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address _storageAddress, address _usersAddress, address _paymentsAddress, address _mediationAddress, address _jobCreatorAddress) returns()
func (_Controller *ControllerTransactor) Initialize(opts *bind.TransactOpts, _storageAddress common.Address, _usersAddress common.Address, _paymentsAddress common.Address, _mediationAddress common.Address, _jobCreatorAddress common.Address) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "initialize", _storageAddress, _usersAddress, _paymentsAddress, _mediationAddress, _jobCreatorAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address _storageAddress, address _usersAddress, address _paymentsAddress, address _mediationAddress, address _jobCreatorAddress) returns()
func (_Controller *ControllerSession) Initialize(_storageAddress common.Address, _usersAddress common.Address, _paymentsAddress common.Address, _mediationAddress common.Address, _jobCreatorAddress common.Address) (*types.Transaction, error) {
	return _Controller.Contract.Initialize(&_Controller.TransactOpts, _storageAddress, _usersAddress, _paymentsAddress, _mediationAddress, _jobCreatorAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address _storageAddress, address _usersAddress, address _paymentsAddress, address _mediationAddress, address _jobCreatorAddress) returns()
func (_Controller *ControllerTransactorSession) Initialize(_storageAddress common.Address, _usersAddress common.Address, _paymentsAddress common.Address, _mediationAddress common.Address, _jobCreatorAddress common.Address) (*types.Transaction, error) {
	return _Controller.Contract.Initialize(&_Controller.TransactOpts, _storageAddress, _usersAddress, _paymentsAddress, _mediationAddress, _jobCreatorAddress)
}

// MediationAcceptResult is a paid mutator transaction binding the contract method 0x824518aa.
//
// Solidity: function mediationAcceptResult(string dealId) returns()
func (_Controller *ControllerTransactor) MediationAcceptResult(opts *bind.TransactOpts, dealId string) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "mediationAcceptResult", dealId)
}

// MediationAcceptResult is a paid mutator transaction binding the contract method 0x824518aa.
//
// Solidity: function mediationAcceptResult(string dealId) returns()
func (_Controller *ControllerSession) MediationAcceptResult(dealId string) (*types.Transaction, error) {
	return _Controller.Contract.MediationAcceptResult(&_Controller.TransactOpts, dealId)
}

// MediationAcceptResult is a paid mutator transaction binding the contract method 0x824518aa.
//
// Solidity: function mediationAcceptResult(string dealId) returns()
func (_Controller *ControllerTransactorSession) MediationAcceptResult(dealId string) (*types.Transaction, error) {
	return _Controller.Contract.MediationAcceptResult(&_Controller.TransactOpts, dealId)
}

// MediationRejectResult is a paid mutator transaction binding the contract method 0x80ffdfe0.
//
// Solidity: function mediationRejectResult(string dealId) returns()
func (_Controller *ControllerTransactor) MediationRejectResult(opts *bind.TransactOpts, dealId string) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "mediationRejectResult", dealId)
}

// MediationRejectResult is a paid mutator transaction binding the contract method 0x80ffdfe0.
//
// Solidity: function mediationRejectResult(string dealId) returns()
func (_Controller *ControllerSession) MediationRejectResult(dealId string) (*types.Transaction, error) {
	return _Controller.Contract.MediationRejectResult(&_Controller.TransactOpts, dealId)
}

// MediationRejectResult is a paid mutator transaction binding the contract method 0x80ffdfe0.
//
// Solidity: function mediationRejectResult(string dealId) returns()
func (_Controller *ControllerTransactorSession) MediationRejectResult(dealId string) (*types.Transaction, error) {
	return _Controller.Contract.MediationRejectResult(&_Controller.TransactOpts, dealId)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Controller *ControllerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Controller *ControllerSession) RenounceOwnership() (*types.Transaction, error) {
	return _Controller.Contract.RenounceOwnership(&_Controller.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Controller *ControllerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Controller.Contract.RenounceOwnership(&_Controller.TransactOpts)
}

// SetJobCreatorAddress is a paid mutator transaction binding the contract method 0xb4031e54.
//
// Solidity: function setJobCreatorAddress(address _jobCreatorAddress) returns()
func (_Controller *ControllerTransactor) SetJobCreatorAddress(opts *bind.TransactOpts, _jobCreatorAddress common.Address) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "setJobCreatorAddress", _jobCreatorAddress)
}

// SetJobCreatorAddress is a paid mutator transaction binding the contract method 0xb4031e54.
//
// Solidity: function setJobCreatorAddress(address _jobCreatorAddress) returns()
func (_Controller *ControllerSession) SetJobCreatorAddress(_jobCreatorAddress common.Address) (*types.Transaction, error) {
	return _Controller.Contract.SetJobCreatorAddress(&_Controller.TransactOpts, _jobCreatorAddress)
}

// SetJobCreatorAddress is a paid mutator transaction binding the contract method 0xb4031e54.
//
// Solidity: function setJobCreatorAddress(address _jobCreatorAddress) returns()
func (_Controller *ControllerTransactorSession) SetJobCreatorAddress(_jobCreatorAddress common.Address) (*types.Transaction, error) {
	return _Controller.Contract.SetJobCreatorAddress(&_Controller.TransactOpts, _jobCreatorAddress)
}

// SetMediationAddress is a paid mutator transaction binding the contract method 0x43391cca.
//
// Solidity: function setMediationAddress(address _mediationAddress) returns()
func (_Controller *ControllerTransactor) SetMediationAddress(opts *bind.TransactOpts, _mediationAddress common.Address) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "setMediationAddress", _mediationAddress)
}

// SetMediationAddress is a paid mutator transaction binding the contract method 0x43391cca.
//
// Solidity: function setMediationAddress(address _mediationAddress) returns()
func (_Controller *ControllerSession) SetMediationAddress(_mediationAddress common.Address) (*types.Transaction, error) {
	return _Controller.Contract.SetMediationAddress(&_Controller.TransactOpts, _mediationAddress)
}

// SetMediationAddress is a paid mutator transaction binding the contract method 0x43391cca.
//
// Solidity: function setMediationAddress(address _mediationAddress) returns()
func (_Controller *ControllerTransactorSession) SetMediationAddress(_mediationAddress common.Address) (*types.Transaction, error) {
	return _Controller.Contract.SetMediationAddress(&_Controller.TransactOpts, _mediationAddress)
}

// SetPaymentsAddress is a paid mutator transaction binding the contract method 0x640e570f.
//
// Solidity: function setPaymentsAddress(address _paymentsAddress) returns()
func (_Controller *ControllerTransactor) SetPaymentsAddress(opts *bind.TransactOpts, _paymentsAddress common.Address) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "setPaymentsAddress", _paymentsAddress)
}

// SetPaymentsAddress is a paid mutator transaction binding the contract method 0x640e570f.
//
// Solidity: function setPaymentsAddress(address _paymentsAddress) returns()
func (_Controller *ControllerSession) SetPaymentsAddress(_paymentsAddress common.Address) (*types.Transaction, error) {
	return _Controller.Contract.SetPaymentsAddress(&_Controller.TransactOpts, _paymentsAddress)
}

// SetPaymentsAddress is a paid mutator transaction binding the contract method 0x640e570f.
//
// Solidity: function setPaymentsAddress(address _paymentsAddress) returns()
func (_Controller *ControllerTransactorSession) SetPaymentsAddress(_paymentsAddress common.Address) (*types.Transaction, error) {
	return _Controller.Contract.SetPaymentsAddress(&_Controller.TransactOpts, _paymentsAddress)
}

// SetStorageAddress is a paid mutator transaction binding the contract method 0x59b910d6.
//
// Solidity: function setStorageAddress(address _storageAddress) returns()
func (_Controller *ControllerTransactor) SetStorageAddress(opts *bind.TransactOpts, _storageAddress common.Address) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "setStorageAddress", _storageAddress)
}

// SetStorageAddress is a paid mutator transaction binding the contract method 0x59b910d6.
//
// Solidity: function setStorageAddress(address _storageAddress) returns()
func (_Controller *ControllerSession) SetStorageAddress(_storageAddress common.Address) (*types.Transaction, error) {
	return _Controller.Contract.SetStorageAddress(&_Controller.TransactOpts, _storageAddress)
}

// SetStorageAddress is a paid mutator transaction binding the contract method 0x59b910d6.
//
// Solidity: function setStorageAddress(address _storageAddress) returns()
func (_Controller *ControllerTransactorSession) SetStorageAddress(_storageAddress common.Address) (*types.Transaction, error) {
	return _Controller.Contract.SetStorageAddress(&_Controller.TransactOpts, _storageAddress)
}

// SetUsersAddress is a paid mutator transaction binding the contract method 0xbbfff47d.
//
// Solidity: function setUsersAddress(address _usersAddress) returns()
func (_Controller *ControllerTransactor) SetUsersAddress(opts *bind.TransactOpts, _usersAddress common.Address) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "setUsersAddress", _usersAddress)
}

// SetUsersAddress is a paid mutator transaction binding the contract method 0xbbfff47d.
//
// Solidity: function setUsersAddress(address _usersAddress) returns()
func (_Controller *ControllerSession) SetUsersAddress(_usersAddress common.Address) (*types.Transaction, error) {
	return _Controller.Contract.SetUsersAddress(&_Controller.TransactOpts, _usersAddress)
}

// SetUsersAddress is a paid mutator transaction binding the contract method 0xbbfff47d.
//
// Solidity: function setUsersAddress(address _usersAddress) returns()
func (_Controller *ControllerTransactorSession) SetUsersAddress(_usersAddress common.Address) (*types.Transaction, error) {
	return _Controller.Contract.SetUsersAddress(&_Controller.TransactOpts, _usersAddress)
}

// TimeoutAgree is a paid mutator transaction binding the contract method 0x795f9abf.
//
// Solidity: function timeoutAgree(string dealId) returns()
func (_Controller *ControllerTransactor) TimeoutAgree(opts *bind.TransactOpts, dealId string) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "timeoutAgree", dealId)
}

// TimeoutAgree is a paid mutator transaction binding the contract method 0x795f9abf.
//
// Solidity: function timeoutAgree(string dealId) returns()
func (_Controller *ControllerSession) TimeoutAgree(dealId string) (*types.Transaction, error) {
	return _Controller.Contract.TimeoutAgree(&_Controller.TransactOpts, dealId)
}

// TimeoutAgree is a paid mutator transaction binding the contract method 0x795f9abf.
//
// Solidity: function timeoutAgree(string dealId) returns()
func (_Controller *ControllerTransactorSession) TimeoutAgree(dealId string) (*types.Transaction, error) {
	return _Controller.Contract.TimeoutAgree(&_Controller.TransactOpts, dealId)
}

// TimeoutJudgeResult is a paid mutator transaction binding the contract method 0xe850be37.
//
// Solidity: function timeoutJudgeResult(string dealId) returns()
func (_Controller *ControllerTransactor) TimeoutJudgeResult(opts *bind.TransactOpts, dealId string) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "timeoutJudgeResult", dealId)
}

// TimeoutJudgeResult is a paid mutator transaction binding the contract method 0xe850be37.
//
// Solidity: function timeoutJudgeResult(string dealId) returns()
func (_Controller *ControllerSession) TimeoutJudgeResult(dealId string) (*types.Transaction, error) {
	return _Controller.Contract.TimeoutJudgeResult(&_Controller.TransactOpts, dealId)
}

// TimeoutJudgeResult is a paid mutator transaction binding the contract method 0xe850be37.
//
// Solidity: function timeoutJudgeResult(string dealId) returns()
func (_Controller *ControllerTransactorSession) TimeoutJudgeResult(dealId string) (*types.Transaction, error) {
	return _Controller.Contract.TimeoutJudgeResult(&_Controller.TransactOpts, dealId)
}

// TimeoutMediateResult is a paid mutator transaction binding the contract method 0xe7b957d1.
//
// Solidity: function timeoutMediateResult(string dealId) returns()
func (_Controller *ControllerTransactor) TimeoutMediateResult(opts *bind.TransactOpts, dealId string) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "timeoutMediateResult", dealId)
}

// TimeoutMediateResult is a paid mutator transaction binding the contract method 0xe7b957d1.
//
// Solidity: function timeoutMediateResult(string dealId) returns()
func (_Controller *ControllerSession) TimeoutMediateResult(dealId string) (*types.Transaction, error) {
	return _Controller.Contract.TimeoutMediateResult(&_Controller.TransactOpts, dealId)
}

// TimeoutMediateResult is a paid mutator transaction binding the contract method 0xe7b957d1.
//
// Solidity: function timeoutMediateResult(string dealId) returns()
func (_Controller *ControllerTransactorSession) TimeoutMediateResult(dealId string) (*types.Transaction, error) {
	return _Controller.Contract.TimeoutMediateResult(&_Controller.TransactOpts, dealId)
}

// TimeoutSubmitResult is a paid mutator transaction binding the contract method 0x511a9f68.
//
// Solidity: function timeoutSubmitResult(string dealId) returns()
func (_Controller *ControllerTransactor) TimeoutSubmitResult(opts *bind.TransactOpts, dealId string) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "timeoutSubmitResult", dealId)
}

// TimeoutSubmitResult is a paid mutator transaction binding the contract method 0x511a9f68.
//
// Solidity: function timeoutSubmitResult(string dealId) returns()
func (_Controller *ControllerSession) TimeoutSubmitResult(dealId string) (*types.Transaction, error) {
	return _Controller.Contract.TimeoutSubmitResult(&_Controller.TransactOpts, dealId)
}

// TimeoutSubmitResult is a paid mutator transaction binding the contract method 0x511a9f68.
//
// Solidity: function timeoutSubmitResult(string dealId) returns()
func (_Controller *ControllerTransactorSession) TimeoutSubmitResult(dealId string) (*types.Transaction, error) {
	return _Controller.Contract.TimeoutSubmitResult(&_Controller.TransactOpts, dealId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Controller *ControllerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Controller *ControllerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Controller.Contract.TransferOwnership(&_Controller.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Controller *ControllerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Controller.Contract.TransferOwnership(&_Controller.TransactOpts, newOwner)
}

// ControllerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Controller contract.
type ControllerInitializedIterator struct {
	Event *ControllerInitialized // Event containing the contract specifics and raw log

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
func (it *ControllerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ControllerInitialized)
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
		it.Event = new(ControllerInitialized)
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
func (it *ControllerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ControllerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ControllerInitialized represents a Initialized event raised by the Controller contract.
type ControllerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Controller *ControllerFilterer) FilterInitialized(opts *bind.FilterOpts) (*ControllerInitializedIterator, error) {

	logs, sub, err := _Controller.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ControllerInitializedIterator{contract: _Controller.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Controller *ControllerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ControllerInitialized) (event.Subscription, error) {

	logs, sub, err := _Controller.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ControllerInitialized)
				if err := _Controller.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Controller *ControllerFilterer) ParseInitialized(log types.Log) (*ControllerInitialized, error) {
	event := new(ControllerInitialized)
	if err := _Controller.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ControllerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Controller contract.
type ControllerOwnershipTransferredIterator struct {
	Event *ControllerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ControllerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ControllerOwnershipTransferred)
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
		it.Event = new(ControllerOwnershipTransferred)
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
func (it *ControllerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ControllerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ControllerOwnershipTransferred represents a OwnershipTransferred event raised by the Controller contract.
type ControllerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Controller *ControllerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ControllerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Controller.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ControllerOwnershipTransferredIterator{contract: _Controller.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Controller *ControllerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ControllerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Controller.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ControllerOwnershipTransferred)
				if err := _Controller.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Controller *ControllerFilterer) ParseOwnershipTransferred(log types.Log) (*ControllerOwnershipTransferred, error) {
	event := new(ControllerOwnershipTransferred)
	if err := _Controller.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
