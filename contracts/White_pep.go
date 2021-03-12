// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// WhitePepABI is the input ABI used to generate the binding from.
const WhitePepABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"GovernanceTransferred\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[],\"name\":\"Open\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"_governance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"pep\",\"type\":\"address\"}],\"name\":\"get_Whiteis\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"is_Openall\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"pep\",\"type\":\"address\"}],\"name\":\"remove_Address\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"governance\",\"type\":\"address\"}],\"name\":\"setGovernance\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"pep\",\"type\":\"address\"}],\"name\":\"set_Address\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// WhitePepBin is the compiled bytecode used for deploying new contracts.
var WhitePepBin = "0x60806040526000600260006101000a81548160ff021916908315150217905550326000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506108528061006e6000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c8063ab033ea91161005b578063ab033ea9146100f8578063b0162e411461013c578063c056722614610180578063d70d3125146101c45761007d565b80631c2f3e3d1461008257806359ebeb90146100cc57806362bc012b146100d6575b600080fd5b61008a610220565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6100d4610245565b005b6100de610333565b604051808215151515815260200191505060405180910390f35b61013a6004803603602081101561010e57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610346565b005b61017e6004803603602081101561015257600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610569565b005b6101c26004803603602081101561019657600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610686565b005b610206600480360360208110156101da57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506107a2565b604051808215151515815260200191505060405180910390f35b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610307576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600e8152602001807f6e6f7420676f7665726e616e636500000000000000000000000000000000000081525060200191505060405180910390fd5b600260009054906101000a900460ff1615600260006101000a81548160ff021916908315150217905550565b600260009054906101000a900460ff1681565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610408576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600e8152602001807f6e6f7420676f7665726e616e636500000000000000000000000000000000000081525060200191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614156104ab576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601f8152602001807f6e657720676f7665726e616e636520746865207a65726f20616464726573730081525060200191505060405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f5f56bee8cffbe9a78652a74a60705edede02af10b0bbb888ca44b79a0d42ce8060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461062b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600e8152602001807f6e6f7420676f7665726e616e636500000000000000000000000000000000000081525060200191505060405180910390fd5b6000600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610748576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600e8152602001807f6e6f7420676f7665726e616e636500000000000000000000000000000000000081525060200191505060405180910390fd5b60018060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555050565b6000801515600260009054906101000a900460ff161515141561081357600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff169050610818565b600190505b91905056fea265627a7a72315820912dcf1b4f5004a7e6bd6c743b6d53a891babd43e4d1177d738adc5f5865822864736f6c63430005110032"

// DeployWhitePep deploys a new Ethereum contract, binding an instance of WhitePep to it.
func DeployWhitePep(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *WhitePep, error) {
	parsed, err := abi.JSON(strings.NewReader(WhitePepABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(WhitePepBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &WhitePep{WhitePepCaller: WhitePepCaller{contract: contract}, WhitePepTransactor: WhitePepTransactor{contract: contract}, WhitePepFilterer: WhitePepFilterer{contract: contract}}, nil
}

// WhitePep is an auto generated Go binding around an Ethereum contract.
type WhitePep struct {
	WhitePepCaller     // Read-only binding to the contract
	WhitePepTransactor // Write-only binding to the contract
	WhitePepFilterer   // Log filterer for contract events
}

// WhitePepCaller is an auto generated read-only Go binding around an Ethereum contract.
type WhitePepCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WhitePepTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WhitePepTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WhitePepFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WhitePepFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WhitePepSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WhitePepSession struct {
	Contract     *WhitePep         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WhitePepCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WhitePepCallerSession struct {
	Contract *WhitePepCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// WhitePepTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WhitePepTransactorSession struct {
	Contract     *WhitePepTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// WhitePepRaw is an auto generated low-level Go binding around an Ethereum contract.
type WhitePepRaw struct {
	Contract *WhitePep // Generic contract binding to access the raw methods on
}

// WhitePepCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WhitePepCallerRaw struct {
	Contract *WhitePepCaller // Generic read-only contract binding to access the raw methods on
}

// WhitePepTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WhitePepTransactorRaw struct {
	Contract *WhitePepTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWhitePep creates a new instance of WhitePep, bound to a specific deployed contract.
func NewWhitePep(address common.Address, backend bind.ContractBackend) (*WhitePep, error) {
	contract, err := bindWhitePep(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WhitePep{WhitePepCaller: WhitePepCaller{contract: contract}, WhitePepTransactor: WhitePepTransactor{contract: contract}, WhitePepFilterer: WhitePepFilterer{contract: contract}}, nil
}

// NewWhitePepCaller creates a new read-only instance of WhitePep, bound to a specific deployed contract.
func NewWhitePepCaller(address common.Address, caller bind.ContractCaller) (*WhitePepCaller, error) {
	contract, err := bindWhitePep(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WhitePepCaller{contract: contract}, nil
}

// NewWhitePepTransactor creates a new write-only instance of WhitePep, bound to a specific deployed contract.
func NewWhitePepTransactor(address common.Address, transactor bind.ContractTransactor) (*WhitePepTransactor, error) {
	contract, err := bindWhitePep(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WhitePepTransactor{contract: contract}, nil
}

// NewWhitePepFilterer creates a new log filterer instance of WhitePep, bound to a specific deployed contract.
func NewWhitePepFilterer(address common.Address, filterer bind.ContractFilterer) (*WhitePepFilterer, error) {
	contract, err := bindWhitePep(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WhitePepFilterer{contract: contract}, nil
}

// bindWhitePep binds a generic wrapper to an already deployed contract.
func bindWhitePep(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(WhitePepABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WhitePep *WhitePepRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WhitePep.Contract.WhitePepCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WhitePep *WhitePepRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WhitePep.Contract.WhitePepTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WhitePep *WhitePepRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WhitePep.Contract.WhitePepTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WhitePep *WhitePepCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WhitePep.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WhitePep *WhitePepTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WhitePep.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WhitePep *WhitePepTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WhitePep.Contract.contract.Transact(opts, method, params...)
}

// Governance is a free data retrieval call binding the contract method 0x1c2f3e3d.
//
// Solidity: function _governance() view returns(address)
func (_WhitePep *WhitePepCaller) Governance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WhitePep.contract.Call(opts, &out, "_governance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Governance is a free data retrieval call binding the contract method 0x1c2f3e3d.
//
// Solidity: function _governance() view returns(address)
func (_WhitePep *WhitePepSession) Governance() (common.Address, error) {
	return _WhitePep.Contract.Governance(&_WhitePep.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x1c2f3e3d.
//
// Solidity: function _governance() view returns(address)
func (_WhitePep *WhitePepCallerSession) Governance() (common.Address, error) {
	return _WhitePep.Contract.Governance(&_WhitePep.CallOpts)
}

// GetWhiteis is a free data retrieval call binding the contract method 0xd70d3125.
//
// Solidity: function get_Whiteis(address pep) view returns(bool)
func (_WhitePep *WhitePepCaller) GetWhiteis(opts *bind.CallOpts, pep common.Address) (bool, error) {
	var out []interface{}
	err := _WhitePep.contract.Call(opts, &out, "get_Whiteis", pep)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetWhiteis is a free data retrieval call binding the contract method 0xd70d3125.
//
// Solidity: function get_Whiteis(address pep) view returns(bool)
func (_WhitePep *WhitePepSession) GetWhiteis(pep common.Address) (bool, error) {
	return _WhitePep.Contract.GetWhiteis(&_WhitePep.CallOpts, pep)
}

// GetWhiteis is a free data retrieval call binding the contract method 0xd70d3125.
//
// Solidity: function get_Whiteis(address pep) view returns(bool)
func (_WhitePep *WhitePepCallerSession) GetWhiteis(pep common.Address) (bool, error) {
	return _WhitePep.Contract.GetWhiteis(&_WhitePep.CallOpts, pep)
}

// IsOpenall is a free data retrieval call binding the contract method 0x62bc012b.
//
// Solidity: function is_Openall() view returns(bool)
func (_WhitePep *WhitePepCaller) IsOpenall(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _WhitePep.contract.Call(opts, &out, "is_Openall")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOpenall is a free data retrieval call binding the contract method 0x62bc012b.
//
// Solidity: function is_Openall() view returns(bool)
func (_WhitePep *WhitePepSession) IsOpenall() (bool, error) {
	return _WhitePep.Contract.IsOpenall(&_WhitePep.CallOpts)
}

// IsOpenall is a free data retrieval call binding the contract method 0x62bc012b.
//
// Solidity: function is_Openall() view returns(bool)
func (_WhitePep *WhitePepCallerSession) IsOpenall() (bool, error) {
	return _WhitePep.Contract.IsOpenall(&_WhitePep.CallOpts)
}

// Open is a paid mutator transaction binding the contract method 0x59ebeb90.
//
// Solidity: function Open() returns()
func (_WhitePep *WhitePepTransactor) Open(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WhitePep.contract.Transact(opts, "Open")
}

// Open is a paid mutator transaction binding the contract method 0x59ebeb90.
//
// Solidity: function Open() returns()
func (_WhitePep *WhitePepSession) Open() (*types.Transaction, error) {
	return _WhitePep.Contract.Open(&_WhitePep.TransactOpts)
}

// Open is a paid mutator transaction binding the contract method 0x59ebeb90.
//
// Solidity: function Open() returns()
func (_WhitePep *WhitePepTransactorSession) Open() (*types.Transaction, error) {
	return _WhitePep.Contract.Open(&_WhitePep.TransactOpts)
}

// RemoveAddress is a paid mutator transaction binding the contract method 0xb0162e41.
//
// Solidity: function remove_Address(address pep) returns()
func (_WhitePep *WhitePepTransactor) RemoveAddress(opts *bind.TransactOpts, pep common.Address) (*types.Transaction, error) {
	return _WhitePep.contract.Transact(opts, "remove_Address", pep)
}

// RemoveAddress is a paid mutator transaction binding the contract method 0xb0162e41.
//
// Solidity: function remove_Address(address pep) returns()
func (_WhitePep *WhitePepSession) RemoveAddress(pep common.Address) (*types.Transaction, error) {
	return _WhitePep.Contract.RemoveAddress(&_WhitePep.TransactOpts, pep)
}

// RemoveAddress is a paid mutator transaction binding the contract method 0xb0162e41.
//
// Solidity: function remove_Address(address pep) returns()
func (_WhitePep *WhitePepTransactorSession) RemoveAddress(pep common.Address) (*types.Transaction, error) {
	return _WhitePep.Contract.RemoveAddress(&_WhitePep.TransactOpts, pep)
}

// SetGovernance is a paid mutator transaction binding the contract method 0xab033ea9.
//
// Solidity: function setGovernance(address governance) returns()
func (_WhitePep *WhitePepTransactor) SetGovernance(opts *bind.TransactOpts, governance common.Address) (*types.Transaction, error) {
	return _WhitePep.contract.Transact(opts, "setGovernance", governance)
}

// SetGovernance is a paid mutator transaction binding the contract method 0xab033ea9.
//
// Solidity: function setGovernance(address governance) returns()
func (_WhitePep *WhitePepSession) SetGovernance(governance common.Address) (*types.Transaction, error) {
	return _WhitePep.Contract.SetGovernance(&_WhitePep.TransactOpts, governance)
}

// SetGovernance is a paid mutator transaction binding the contract method 0xab033ea9.
//
// Solidity: function setGovernance(address governance) returns()
func (_WhitePep *WhitePepTransactorSession) SetGovernance(governance common.Address) (*types.Transaction, error) {
	return _WhitePep.Contract.SetGovernance(&_WhitePep.TransactOpts, governance)
}

// SetAddress is a paid mutator transaction binding the contract method 0xc0567226.
//
// Solidity: function set_Address(address pep) returns()
func (_WhitePep *WhitePepTransactor) SetAddress(opts *bind.TransactOpts, pep common.Address) (*types.Transaction, error) {
	return _WhitePep.contract.Transact(opts, "set_Address", pep)
}

// SetAddress is a paid mutator transaction binding the contract method 0xc0567226.
//
// Solidity: function set_Address(address pep) returns()
func (_WhitePep *WhitePepSession) SetAddress(pep common.Address) (*types.Transaction, error) {
	return _WhitePep.Contract.SetAddress(&_WhitePep.TransactOpts, pep)
}

// SetAddress is a paid mutator transaction binding the contract method 0xc0567226.
//
// Solidity: function set_Address(address pep) returns()
func (_WhitePep *WhitePepTransactorSession) SetAddress(pep common.Address) (*types.Transaction, error) {
	return _WhitePep.Contract.SetAddress(&_WhitePep.TransactOpts, pep)
}

// WhitePepGovernanceTransferredIterator is returned from FilterGovernanceTransferred and is used to iterate over the raw logs and unpacked data for GovernanceTransferred events raised by the WhitePep contract.
type WhitePepGovernanceTransferredIterator struct {
	Event *WhitePepGovernanceTransferred // Event containing the contract specifics and raw log

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
func (it *WhitePepGovernanceTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WhitePepGovernanceTransferred)
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
		it.Event = new(WhitePepGovernanceTransferred)
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
func (it *WhitePepGovernanceTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WhitePepGovernanceTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WhitePepGovernanceTransferred represents a GovernanceTransferred event raised by the WhitePep contract.
type WhitePepGovernanceTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterGovernanceTransferred is a free log retrieval operation binding the contract event 0x5f56bee8cffbe9a78652a74a60705edede02af10b0bbb888ca44b79a0d42ce80.
//
// Solidity: event GovernanceTransferred(address indexed previousOwner, address indexed newOwner)
func (_WhitePep *WhitePepFilterer) FilterGovernanceTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*WhitePepGovernanceTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _WhitePep.contract.FilterLogs(opts, "GovernanceTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &WhitePepGovernanceTransferredIterator{contract: _WhitePep.contract, event: "GovernanceTransferred", logs: logs, sub: sub}, nil
}

// WatchGovernanceTransferred is a free log subscription operation binding the contract event 0x5f56bee8cffbe9a78652a74a60705edede02af10b0bbb888ca44b79a0d42ce80.
//
// Solidity: event GovernanceTransferred(address indexed previousOwner, address indexed newOwner)
func (_WhitePep *WhitePepFilterer) WatchGovernanceTransferred(opts *bind.WatchOpts, sink chan<- *WhitePepGovernanceTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _WhitePep.contract.WatchLogs(opts, "GovernanceTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WhitePepGovernanceTransferred)
				if err := _WhitePep.contract.UnpackLog(event, "GovernanceTransferred", log); err != nil {
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

// ParseGovernanceTransferred is a log parse operation binding the contract event 0x5f56bee8cffbe9a78652a74a60705edede02af10b0bbb888ca44b79a0d42ce80.
//
// Solidity: event GovernanceTransferred(address indexed previousOwner, address indexed newOwner)
func (_WhitePep *WhitePepFilterer) ParseGovernanceTransferred(log types.Log) (*WhitePepGovernanceTransferred, error) {
	event := new(WhitePepGovernanceTransferred)
	if err := _WhitePep.contract.UnpackLog(event, "GovernanceTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
