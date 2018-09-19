// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// JNSABI is the input ABI used to generate the binding from.
const JNSABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"getNickname\",\"outputs\":[{\"name\":\"_nickname\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_nickname\",\"type\":\"string\"},{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"registerAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_nickname\",\"type\":\"string\"}],\"name\":\"getAddress\",\"outputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_wallet\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// JNSBin is the compiled bytecode used for deploying new contracts.
const JNSBin = `0x608060405234801561001057600080fd5b506040516020806107e38339810160405251600160a060020a038116151561009957604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f596f75206d75737420696e666f726d20612076616c6964206164647265737300604482015290519081900360640190fd5b60028054600160a060020a031916600160a060020a039290921691909117905561071b806100c86000396000f3006080604052600436106100565763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663adaccd74811461005b578063b13f5f19146100f1578063bf40fac11461015c575b600080fd5b34801561006757600080fd5b5061007c600160a060020a03600435166101d1565b6040805160208082528351818301528351919283929083019185019080838360005b838110156100b657818101518382015260200161009e565b50505050905090810190601f1680156100e35780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6040805160206004803580820135601f810184900484028501840190955284845261014894369492936024939284019190819084018382808284375094975050509235600160a060020a0316935061027b92505050565b604080519115158252519081900360200190f35b34801561016857600080fd5b506040805160206004803580820135601f81018490048402850184019095528484526101b59436949293602493928401919081908401838280828437509497506105e49650505050505050565b60408051600160a060020a039092168252519081900360200190f35b600160a060020a03811660009081526001602081815260409283902080548451600294821615610100026000190190911693909304601f8101839004830284018301909452838352606093909183018282801561026f5780601f106102445761010080835404028352916020019161026f565b820191906000526020600020905b81548152906001019060200180831161025257829003601f168201915b50505050509050919050565b600066038d7ea4c6800034116102f257604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600f60248201527f53656e64206d6f7265206d6f6e65790000000000000000000000000000000000604482015290519081900360640190fd5b6000600160a060020a03166000846040518082805190602001908083835b6020831061032f5780518252601f199092019160209182019101610310565b51815160209384036101000a6000190180199092169116179052920194855250604051938490030190922054600160a060020a03169290921491506103d7905057604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f4e616d6520616c72656164792072656769737465726564000000000000000000604482015290519081900360640190fd5b60408051600160a060020a03841660009081526001602081905292902080547fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470939192918291849160026101009183161591909102600019019091160480156104775780601f10610455576101008083540402835291820191610477565b820191906000526020600020905b815481529060010190602001808311610463575b50509150506040518091039020600019161415156104f657604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601a60248201527f4164647265737320616c72656164792072656769737465726564000000000000604482015290519081900360640190fd5b816000846040518082805190602001908083835b602083106105295780518252601f19909201916020918201910161050a565b51815160209384036101000a600019018019909216911617905292019485525060408051948590038201909420805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a039687161790559386166000908152600185529290922086516105a0949193509087019150610654565b50600254604051600160a060020a03909116903480156108fc02916000818181858888f193505050501580156105da573d6000803e3d6000fd5b5060019392505050565b600080826040518082805190602001908083835b602083106106175780518252601f1990920191602091820191016105f8565b51815160209384036101000a6000190180199092169116179052920194855250604051938490030190922054600160a060020a0316949350505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061069557805160ff19168380011785556106c2565b828001600101855582156106c2579182015b828111156106c25782518255916020019190600101906106a7565b506106ce9291506106d2565b5090565b6106ec91905b808211156106ce57600081556001016106d8565b905600a165627a7a723058202ae6cd4339c02167643cc619c479cbbe2c671aeb1cb084a8cd1490a4638353b60029`

// DeployJNS deploys a new Ethereum contract, binding an instance of JNS to it.
func DeployJNS(auth *bind.TransactOpts, backend bind.ContractBackend, _wallet common.Address) (common.Address, *types.Transaction, *JNS, error) {
	parsed, err := abi.JSON(strings.NewReader(JNSABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(JNSBin), backend, _wallet)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &JNS{JNSCaller: JNSCaller{contract: contract}, JNSTransactor: JNSTransactor{contract: contract}, JNSFilterer: JNSFilterer{contract: contract}}, nil
}

// JNS is an auto generated Go binding around an Ethereum contract.
type JNS struct {
	JNSCaller     // Read-only binding to the contract
	JNSTransactor // Write-only binding to the contract
	JNSFilterer   // Log filterer for contract events
}

// JNSCaller is an auto generated read-only Go binding around an Ethereum contract.
type JNSCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JNSTransactor is an auto generated write-only Go binding around an Ethereum contract.
type JNSTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JNSFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type JNSFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JNSSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type JNSSession struct {
	Contract     *JNS              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// JNSCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type JNSCallerSession struct {
	Contract *JNSCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// JNSTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type JNSTransactorSession struct {
	Contract     *JNSTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// JNSRaw is an auto generated low-level Go binding around an Ethereum contract.
type JNSRaw struct {
	Contract *JNS // Generic contract binding to access the raw methods on
}

// JNSCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type JNSCallerRaw struct {
	Contract *JNSCaller // Generic read-only contract binding to access the raw methods on
}

// JNSTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type JNSTransactorRaw struct {
	Contract *JNSTransactor // Generic write-only contract binding to access the raw methods on
}

// NewJNS creates a new instance of JNS, bound to a specific deployed contract.
func NewJNS(address common.Address, backend bind.ContractBackend) (*JNS, error) {
	contract, err := bindJNS(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &JNS{JNSCaller: JNSCaller{contract: contract}, JNSTransactor: JNSTransactor{contract: contract}, JNSFilterer: JNSFilterer{contract: contract}}, nil
}

// NewJNSCaller creates a new read-only instance of JNS, bound to a specific deployed contract.
func NewJNSCaller(address common.Address, caller bind.ContractCaller) (*JNSCaller, error) {
	contract, err := bindJNS(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &JNSCaller{contract: contract}, nil
}

// NewJNSTransactor creates a new write-only instance of JNS, bound to a specific deployed contract.
func NewJNSTransactor(address common.Address, transactor bind.ContractTransactor) (*JNSTransactor, error) {
	contract, err := bindJNS(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &JNSTransactor{contract: contract}, nil
}

// NewJNSFilterer creates a new log filterer instance of JNS, bound to a specific deployed contract.
func NewJNSFilterer(address common.Address, filterer bind.ContractFilterer) (*JNSFilterer, error) {
	contract, err := bindJNS(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &JNSFilterer{contract: contract}, nil
}

// bindJNS binds a generic wrapper to an already deployed contract.
func bindJNS(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(JNSABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_JNS *JNSRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _JNS.Contract.JNSCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_JNS *JNSRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _JNS.Contract.JNSTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_JNS *JNSRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _JNS.Contract.JNSTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_JNS *JNSCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _JNS.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_JNS *JNSTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _JNS.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_JNS *JNSTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _JNS.Contract.contract.Transact(opts, method, params...)
}

// GetAddress is a free data retrieval call binding the contract method 0xbf40fac1.
//
// Solidity: function getAddress(_nickname string) constant returns(_address address)
func (_JNS *JNSCaller) GetAddress(opts *bind.CallOpts, _nickname string) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _JNS.contract.Call(opts, out, "getAddress", _nickname)
	return *ret0, err
}

// GetAddress is a free data retrieval call binding the contract method 0xbf40fac1.
//
// Solidity: function getAddress(_nickname string) constant returns(_address address)
func (_JNS *JNSSession) GetAddress(_nickname string) (common.Address, error) {
	return _JNS.Contract.GetAddress(&_JNS.CallOpts, _nickname)
}

// GetAddress is a free data retrieval call binding the contract method 0xbf40fac1.
//
// Solidity: function getAddress(_nickname string) constant returns(_address address)
func (_JNS *JNSCallerSession) GetAddress(_nickname string) (common.Address, error) {
	return _JNS.Contract.GetAddress(&_JNS.CallOpts, _nickname)
}

// GetNickname is a free data retrieval call binding the contract method 0xadaccd74.
//
// Solidity: function getNickname(_address address) constant returns(_nickname string)
func (_JNS *JNSCaller) GetNickname(opts *bind.CallOpts, _address common.Address) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _JNS.contract.Call(opts, out, "getNickname", _address)
	return *ret0, err
}

// GetNickname is a free data retrieval call binding the contract method 0xadaccd74.
//
// Solidity: function getNickname(_address address) constant returns(_nickname string)
func (_JNS *JNSSession) GetNickname(_address common.Address) (string, error) {
	return _JNS.Contract.GetNickname(&_JNS.CallOpts, _address)
}

// GetNickname is a free data retrieval call binding the contract method 0xadaccd74.
//
// Solidity: function getNickname(_address address) constant returns(_nickname string)
func (_JNS *JNSCallerSession) GetNickname(_address common.Address) (string, error) {
	return _JNS.Contract.GetNickname(&_JNS.CallOpts, _address)
}

// RegisterAddress is a paid mutator transaction binding the contract method 0xb13f5f19.
//
// Solidity: function registerAddress(_nickname string, _address address) returns(bool)
func (_JNS *JNSTransactor) RegisterAddress(opts *bind.TransactOpts, _nickname string, _address common.Address) (*types.Transaction, error) {
	return _JNS.contract.Transact(opts, "registerAddress", _nickname, _address)
}

// RegisterAddress is a paid mutator transaction binding the contract method 0xb13f5f19.
//
// Solidity: function registerAddress(_nickname string, _address address) returns(bool)
func (_JNS *JNSSession) RegisterAddress(_nickname string, _address common.Address) (*types.Transaction, error) {
	return _JNS.Contract.RegisterAddress(&_JNS.TransactOpts, _nickname, _address)
}

// RegisterAddress is a paid mutator transaction binding the contract method 0xb13f5f19.
//
// Solidity: function registerAddress(_nickname string, _address address) returns(bool)
func (_JNS *JNSTransactorSession) RegisterAddress(_nickname string, _address common.Address) (*types.Transaction, error) {
	return _JNS.Contract.RegisterAddress(&_JNS.TransactOpts, _nickname, _address)
}
