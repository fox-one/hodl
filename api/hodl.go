// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package api

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

// BLSMetaData contains all meta data concerning the BLS contract.
var BLSMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_x\",\"type\":\"bytes32\"}],\"name\":\"mapToPoint\",\"outputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"p\",\"type\":\"uint256[2]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_x\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"expected_roots\",\"type\":\"uint256[]\"}],\"name\":\"mapToPointWithHelp\",\"outputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"p\",\"type\":\"uint256[2]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"signature\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[4]\",\"name\":\"pubkey\",\"type\":\"uint256[4]\"},{\"internalType\":\"uint256[2]\",\"name\":\"message\",\"type\":\"uint256[2]\"}],\"name\":\"verifySingle\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"c8ce9571": "mapToPoint(bytes32)",
		"f25c3485": "mapToPointWithHelp(bytes32,uint256[])",
		"ebbdac91": "verifySingle(uint256[2],uint256[4],uint256[2])",
	},
	Bin: "0x61091761003a600b82828239805160001a60731461002d57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe730000000000000000000000000000000000000000301460806040526004361061004b5760003560e01c8063c8ce957114610050578063ebbdac9114610079578063f25c34851461009c575b600080fd5b61006361005e3660046105e4565b6100af565b60405161007091906105fd565b60405180910390f35b61008c6100873660046106e5565b610159565b6040519015158152602001610070565b6100636100aa366004610785565b61034a565b6100b76105a8565b60006100d16000805160206108c283398151915284610837565b90506000805b6000805160206108c283398151915283840991506000805160206108c283398151915283830991506000805160206108c283398151915260038308915061011d826104d2565b909250905080156101375782845260208401829052610151565b6000805160206108c28339815191526001840892506100d7565b505050919050565b6000806040518061018001604052808660006002811061017b5761017b610859565b602002015181526020018660016002811061019857610198610859565b602002015181526020017f198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c281526020017f1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed81526020017f275dc4a288d1afb3cbb1ac09187524c7db36395df7be3b99e673b13a075a65ec81526020017f1d9befcd05a5323e6da4d435f3b617cdb3af83285c2df711ef39c01571827f9d81526020018460006002811061024d5761024d610859565b602002015181526020018460016002811061026a5761026a610859565b602002015181526020018560016004811061028757610287610859565b60200201518152602001856000600481106102a4576102a4610859565b60200201518152602001856003600481106102c1576102c1610859565b60200201518152602001856002600481106102de576102de610859565b6020020151905290506102ef6105c6565b60006020826101808560086107d05a03fa905080801561030e57610310565bfe5b508061033d5760405162461bcd60e51b815260206004820152600060248201526044015b60405180910390fd5b5051151595945050505050565b6103526105a8565b600061036c6000805160206108c283398151915285610837565b9050600080805b6000805160206108c283398151915284850991506000805160206108c283398151915284830991506000805160206108c28339815191526003830891506000805160206108c2833981519152868460ff16815181106103d4576103d4610859565b6020026020010151878560ff16815181106103f1576103f1610859565b602002602001015109905081811415610443578385528551869060ff851690811061041e5761041e610859565b60200260200101518560016002811061043957610439610859565b60200201526104c8565b8161045c826000805160206108c2833981519152610885565b1415610489576000805160206108c283398151915260018508935061048260018461089c565b9250610373565b60405162461bcd60e51b81526020600482015260146024820152732bb937b7339032bc3832b1ba32b2103937b7ba1760611b6044820152606401610334565b5050505092915050565b600080600060405160208152602080820152602060408201528460608201527f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f5260808201526000805160206108c283398151915260a082015260208160c08360056107d05a03fa9051935090506000805160206108c283398151915283800984149150806105a25760405162461bcd60e51b815260206004820152601c60248201527f424c533a2073717274206d6f646578702063616c6c206661696c6564000000006044820152606401610334565b50915091565b60405180604001604052806002906020820280368337509192915050565b60405180602001604052806001906020820280368337509192915050565b6000602082840312156105f657600080fd5b5035919050565b60408101818360005b6002811015610625578151835260209283019290910190600101610606565b50505092915050565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff8111828210171561066d5761066d61062e565b604052919050565b600082601f83011261068657600080fd5b6040516040810181811067ffffffffffffffff821117156106a9576106a961062e565b80604052508060408401858111156106c057600080fd5b845b818110156106da5780358352602092830192016106c2565b509195945050505050565b600080600061010084860312156106fb57600080fd5b6107058585610675565b925084605f85011261071657600080fd5b6040516080810181811067ffffffffffffffff821117156107395761073961062e565b6040528060c086018781111561074e57600080fd5b604087015b8181101561076b578035835260209283019201610753565b508294506107798882610675565b93505050509250925092565b6000806040838503121561079857600080fd5b8235915060208084013567ffffffffffffffff808211156107b857600080fd5b818601915086601f8301126107cc57600080fd5b8135818111156107de576107de61062e565b8060051b91506107ef848301610644565b818152918301840191848101908984111561080957600080fd5b938501935b838510156108275784358252938501939085019061080e565b8096505050505050509250929050565b60008261085457634e487b7160e01b600052601260045260246000fd5b500690565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b6000828210156108975761089761086f565b500390565b600060ff821660ff84168060ff038211156108b9576108b961086f565b01939250505056fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47a2646970667358221220a39abf74535e16919c623982fbc5147cc9f165571106ca5b8cc99739a4f55a6d64736f6c634300080a0033",
}

// BLSABI is the input ABI used to generate the binding from.
// Deprecated: Use BLSMetaData.ABI instead.
var BLSABI = BLSMetaData.ABI

// Deprecated: Use BLSMetaData.Sigs instead.
// BLSFuncSigs maps the 4-byte function signature to its string representation.
var BLSFuncSigs = BLSMetaData.Sigs

// BLSBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BLSMetaData.Bin instead.
var BLSBin = BLSMetaData.Bin

// DeployBLS deploys a new Ethereum contract, binding an instance of BLS to it.
func DeployBLS(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BLS, error) {
	parsed, err := BLSMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BLSBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BLS{BLSCaller: BLSCaller{contract: contract}, BLSTransactor: BLSTransactor{contract: contract}, BLSFilterer: BLSFilterer{contract: contract}}, nil
}

// BLS is an auto generated Go binding around an Ethereum contract.
type BLS struct {
	BLSCaller     // Read-only binding to the contract
	BLSTransactor // Write-only binding to the contract
	BLSFilterer   // Log filterer for contract events
}

// BLSCaller is an auto generated read-only Go binding around an Ethereum contract.
type BLSCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BLSTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BLSTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BLSFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BLSFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BLSSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BLSSession struct {
	Contract     *BLS              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BLSCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BLSCallerSession struct {
	Contract *BLSCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BLSTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BLSTransactorSession struct {
	Contract     *BLSTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BLSRaw is an auto generated low-level Go binding around an Ethereum contract.
type BLSRaw struct {
	Contract *BLS // Generic contract binding to access the raw methods on
}

// BLSCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BLSCallerRaw struct {
	Contract *BLSCaller // Generic read-only contract binding to access the raw methods on
}

// BLSTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BLSTransactorRaw struct {
	Contract *BLSTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBLS creates a new instance of BLS, bound to a specific deployed contract.
func NewBLS(address common.Address, backend bind.ContractBackend) (*BLS, error) {
	contract, err := bindBLS(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BLS{BLSCaller: BLSCaller{contract: contract}, BLSTransactor: BLSTransactor{contract: contract}, BLSFilterer: BLSFilterer{contract: contract}}, nil
}

// NewBLSCaller creates a new read-only instance of BLS, bound to a specific deployed contract.
func NewBLSCaller(address common.Address, caller bind.ContractCaller) (*BLSCaller, error) {
	contract, err := bindBLS(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BLSCaller{contract: contract}, nil
}

// NewBLSTransactor creates a new write-only instance of BLS, bound to a specific deployed contract.
func NewBLSTransactor(address common.Address, transactor bind.ContractTransactor) (*BLSTransactor, error) {
	contract, err := bindBLS(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BLSTransactor{contract: contract}, nil
}

// NewBLSFilterer creates a new log filterer instance of BLS, bound to a specific deployed contract.
func NewBLSFilterer(address common.Address, filterer bind.ContractFilterer) (*BLSFilterer, error) {
	contract, err := bindBLS(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BLSFilterer{contract: contract}, nil
}

// bindBLS binds a generic wrapper to an already deployed contract.
func bindBLS(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BLSABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BLS *BLSRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BLS.Contract.BLSCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BLS *BLSRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BLS.Contract.BLSTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BLS *BLSRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BLS.Contract.BLSTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BLS *BLSCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BLS.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BLS *BLSTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BLS.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BLS *BLSTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BLS.Contract.contract.Transact(opts, method, params...)
}

// MapToPoint is a free data retrieval call binding the contract method 0xc8ce9571.
//
// Solidity: function mapToPoint(bytes32 _x) view returns(uint256[2] p)
func (_BLS *BLSCaller) MapToPoint(opts *bind.CallOpts, _x [32]byte) ([2]*big.Int, error) {
	var out []interface{}
	err := _BLS.contract.Call(opts, &out, "mapToPoint", _x)

	if err != nil {
		return *new([2]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([2]*big.Int)).(*[2]*big.Int)

	return out0, err

}

// MapToPoint is a free data retrieval call binding the contract method 0xc8ce9571.
//
// Solidity: function mapToPoint(bytes32 _x) view returns(uint256[2] p)
func (_BLS *BLSSession) MapToPoint(_x [32]byte) ([2]*big.Int, error) {
	return _BLS.Contract.MapToPoint(&_BLS.CallOpts, _x)
}

// MapToPoint is a free data retrieval call binding the contract method 0xc8ce9571.
//
// Solidity: function mapToPoint(bytes32 _x) view returns(uint256[2] p)
func (_BLS *BLSCallerSession) MapToPoint(_x [32]byte) ([2]*big.Int, error) {
	return _BLS.Contract.MapToPoint(&_BLS.CallOpts, _x)
}

// MapToPointWithHelp is a free data retrieval call binding the contract method 0xf25c3485.
//
// Solidity: function mapToPointWithHelp(bytes32 _x, uint256[] expected_roots) pure returns(uint256[2] p)
func (_BLS *BLSCaller) MapToPointWithHelp(opts *bind.CallOpts, _x [32]byte, expected_roots []*big.Int) ([2]*big.Int, error) {
	var out []interface{}
	err := _BLS.contract.Call(opts, &out, "mapToPointWithHelp", _x, expected_roots)

	if err != nil {
		return *new([2]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([2]*big.Int)).(*[2]*big.Int)

	return out0, err

}

// MapToPointWithHelp is a free data retrieval call binding the contract method 0xf25c3485.
//
// Solidity: function mapToPointWithHelp(bytes32 _x, uint256[] expected_roots) pure returns(uint256[2] p)
func (_BLS *BLSSession) MapToPointWithHelp(_x [32]byte, expected_roots []*big.Int) ([2]*big.Int, error) {
	return _BLS.Contract.MapToPointWithHelp(&_BLS.CallOpts, _x, expected_roots)
}

// MapToPointWithHelp is a free data retrieval call binding the contract method 0xf25c3485.
//
// Solidity: function mapToPointWithHelp(bytes32 _x, uint256[] expected_roots) pure returns(uint256[2] p)
func (_BLS *BLSCallerSession) MapToPointWithHelp(_x [32]byte, expected_roots []*big.Int) ([2]*big.Int, error) {
	return _BLS.Contract.MapToPointWithHelp(&_BLS.CallOpts, _x, expected_roots)
}

// VerifySingle is a free data retrieval call binding the contract method 0xebbdac91.
//
// Solidity: function verifySingle(uint256[2] signature, uint256[4] pubkey, uint256[2] message) view returns(bool)
func (_BLS *BLSCaller) VerifySingle(opts *bind.CallOpts, signature [2]*big.Int, pubkey [4]*big.Int, message [2]*big.Int) (bool, error) {
	var out []interface{}
	err := _BLS.contract.Call(opts, &out, "verifySingle", signature, pubkey, message)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifySingle is a free data retrieval call binding the contract method 0xebbdac91.
//
// Solidity: function verifySingle(uint256[2] signature, uint256[4] pubkey, uint256[2] message) view returns(bool)
func (_BLS *BLSSession) VerifySingle(signature [2]*big.Int, pubkey [4]*big.Int, message [2]*big.Int) (bool, error) {
	return _BLS.Contract.VerifySingle(&_BLS.CallOpts, signature, pubkey, message)
}

// VerifySingle is a free data retrieval call binding the contract method 0xebbdac91.
//
// Solidity: function verifySingle(uint256[2] signature, uint256[4] pubkey, uint256[2] message) view returns(bool)
func (_BLS *BLSCallerSession) VerifySingle(signature [2]*big.Int, pubkey [4]*big.Int, message [2]*big.Int) (bool, error) {
	return _BLS.Contract.VerifySingle(&_BLS.CallOpts, signature, pubkey, message)
}

// BytesLibMetaData contains all meta data concerning the BytesLib contract.
var BytesLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220edd6f6ea194a1a6796c67b8fd0f2c9512c8fba162b6f80e4913f54264bec197e64736f6c634300080a0033",
}

// BytesLibABI is the input ABI used to generate the binding from.
// Deprecated: Use BytesLibMetaData.ABI instead.
var BytesLibABI = BytesLibMetaData.ABI

// BytesLibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BytesLibMetaData.Bin instead.
var BytesLibBin = BytesLibMetaData.Bin

// DeployBytesLib deploys a new Ethereum contract, binding an instance of BytesLib to it.
func DeployBytesLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BytesLib, error) {
	parsed, err := BytesLibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BytesLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BytesLib{BytesLibCaller: BytesLibCaller{contract: contract}, BytesLibTransactor: BytesLibTransactor{contract: contract}, BytesLibFilterer: BytesLibFilterer{contract: contract}}, nil
}

// BytesLib is an auto generated Go binding around an Ethereum contract.
type BytesLib struct {
	BytesLibCaller     // Read-only binding to the contract
	BytesLibTransactor // Write-only binding to the contract
	BytesLibFilterer   // Log filterer for contract events
}

// BytesLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type BytesLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BytesLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BytesLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BytesLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BytesLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BytesLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BytesLibSession struct {
	Contract     *BytesLib         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BytesLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BytesLibCallerSession struct {
	Contract *BytesLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// BytesLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BytesLibTransactorSession struct {
	Contract     *BytesLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// BytesLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type BytesLibRaw struct {
	Contract *BytesLib // Generic contract binding to access the raw methods on
}

// BytesLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BytesLibCallerRaw struct {
	Contract *BytesLibCaller // Generic read-only contract binding to access the raw methods on
}

// BytesLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BytesLibTransactorRaw struct {
	Contract *BytesLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBytesLib creates a new instance of BytesLib, bound to a specific deployed contract.
func NewBytesLib(address common.Address, backend bind.ContractBackend) (*BytesLib, error) {
	contract, err := bindBytesLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BytesLib{BytesLibCaller: BytesLibCaller{contract: contract}, BytesLibTransactor: BytesLibTransactor{contract: contract}, BytesLibFilterer: BytesLibFilterer{contract: contract}}, nil
}

// NewBytesLibCaller creates a new read-only instance of BytesLib, bound to a specific deployed contract.
func NewBytesLibCaller(address common.Address, caller bind.ContractCaller) (*BytesLibCaller, error) {
	contract, err := bindBytesLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BytesLibCaller{contract: contract}, nil
}

// NewBytesLibTransactor creates a new write-only instance of BytesLib, bound to a specific deployed contract.
func NewBytesLibTransactor(address common.Address, transactor bind.ContractTransactor) (*BytesLibTransactor, error) {
	contract, err := bindBytesLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BytesLibTransactor{contract: contract}, nil
}

// NewBytesLibFilterer creates a new log filterer instance of BytesLib, bound to a specific deployed contract.
func NewBytesLibFilterer(address common.Address, filterer bind.ContractFilterer) (*BytesLibFilterer, error) {
	contract, err := bindBytesLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BytesLibFilterer{contract: contract}, nil
}

// bindBytesLib binds a generic wrapper to an already deployed contract.
func bindBytesLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BytesLibABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BytesLib *BytesLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BytesLib.Contract.BytesLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BytesLib *BytesLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BytesLib.Contract.BytesLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BytesLib *BytesLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BytesLib.Contract.BytesLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BytesLib *BytesLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BytesLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BytesLib *BytesLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BytesLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BytesLib *BytesLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BytesLib.Contract.contract.Transact(opts, method, params...)
}

// HodlMetaData contains all meta data concerning the Hodl contract.
var HodlMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"pid_\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"asset\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"extra\",\"type\":\"bytes\"}],\"name\":\"MixinEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"MixinTransaction\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"GROUPS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NONCE\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PID\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"name\":\"custodian\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"members\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"raw\",\"type\":\"bytes\"}],\"name\":\"mixin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"vats\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"asset\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"guy\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"end\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"50cc55ae": "GROUPS(uint256)",
		"e091dd1a": "NONCE()",
		"5eaec0e4": "PID()",
		"2d459b6b": "custodian(uint128)",
		"08ae4b0c": "members(address)",
		"5cae8005": "mixin(bytes)",
		"40a4e253": "vats(uint256)",
	},
	Bin: "0x6101006040527f2f741961cea2e88cfa2680eeaac040d41f41f3fedb01e38c06f4c6058fd7e42560809081527e7d68aef83f9690b04f463e13eadd9b18f4869041f1b67e7f1a30c9d1d2c42c60a0527f2a32fa1736807486256ad8dc6a8740dfb91917cf8d15848133819275be92b67360c0527f257ad901f02f8a442ccf4f1b1d0d7d3a8e8fe791102706e575d36de1c2a4a40f60e052620000a690600090600462000124565b50600480546001600160c01b0319166fd402f2c61ba04d82a75c5674fae1acca1790556000600855348015620000db57600080fd5b5060405162001e2a38038062001e2a833981016040819052620000fe916200017e565b600480546001600160801b0319166001600160801b0392909216919091179055620001b0565b826004810192821562000155579160200282015b828111156200015557825182559160200191906001019062000138565b506200016392915062000167565b5090565b5b8082111562000163576000815560010162000168565b6000602082840312156200019157600080fd5b81516001600160801b0381168114620001a957600080fd5b9392505050565b611c6a80620001c06000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c806350cc55ae1161005b57806350cc55ae146101805780635cae8005146101935780635eaec0e4146101a8578063e091dd1a146101d357600080fd5b806308ae4b0c146100825780632d459b6b146100ab57806340a4e253146100d9575b600080fd5b610095610090366004611874565b610205565b6040516100a291906118f1565b60405180910390f35b6100cb6100b9366004611904565b60056020526000908152604090205481565b6040519081526020016100a2565b6101376100e736600461192d565b60076020526000908152604090208054600182015460028301546003909301546001600160801b039092169290916001600160a01b038216916001600160401b03600160a01b9091048116911685565b604080516001600160801b03909616865260208601949094526001600160a01b03909216928401929092526001600160401b03918216606084015216608082015260a0016100a2565b6100cb61018e36600461192d565b61029f565b6101a66101a1366004611946565b6102b6565b005b6004546101bb906001600160801b031681565b6040516001600160801b0390911681526020016100a2565b6004546101ed90600160801b90046001600160401b031681565b6040516001600160401b0390911681526020016100a2565b6006602052600090815260409020805461021e906119b7565b80601f016020809104026020016040519081016040528092919081815260200182805461024a906119b7565b80156102975780601f1061026c57610100808354040283529160200191610297565b820191906000526020600020905b81548152906001019060200180831161027a57829003601f168201915b505050505081565b600081600481106102af57600080fd5b0154905081565b608d8110156103035760405162461bcd60e51b8152602060048201526014602482015273195d995b9d0819185d18481d1bdbc81cdb585b1b60621b60448201526064015b60405180910390fd5b600080600061034b8286868080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092939250506109e99050565b6004549091506001600160801b0380831691161461039d5760405162461bcd60e51b815260206004820152600f60248201526e696e76616c69642070726f6365737360881b60448201526064016102fa565b6103a8826010611a08565b915060006103ef8387878080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509293925050610a479050565b6004549091506001600160401b03808316600160801b90920416146104465760405162461bcd60e51b815260206004820152600d60248201526c696e76616c6964206e6f6e636560981b60448201526064016102fa565b60045461046490600160801b90046001600160401b03166001611a20565b600480546001600160401b0392909216600160801b0267ffffffffffffffff60801b1990921691909117905561049b836008611a08565b925060006104e28488888080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092939250506109e99050565b90506104ef846010611a08565b93506105348488888080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509293925050610aa49050565b61ffff169450610545846002611a08565b935060208511156105905760405162461bcd60e51b8152602060048201526015602482015274696e7465676572206f7574206f6620626f756e647360581b60448201526064016102fa565b600061063860006106326105df888a8d8d8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929493925050610b019050565b6105ea8a6020611a4b565b6001600160401b0381111561060157610601611a62565b6040519080825280601f01601f19166020018201604052801561062b576020820181803683370190505b5090610c0e565b90610c8b565b90506106448686611a08565b94506106898589898080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509293925050610aa49050565b61ffff16955061069a856002611a08565b945060006106e386888b8b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929493925050610b019050565b90506106ef8787611a08565b95506000610736878b8b8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509293925050610a479050565b9050610743876008611a08565b9650610788878b8b8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509293925050610aa49050565b61ffff169750610799886010611a78565b6107a4906002611a08565b6107af906002611a08565b975060006107f8888a8d8d8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929493925050610b019050565b90506108048989611a08565b9750610811886002611a08565b97506108548b8b8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508c9250610ce9915050565b6108945760405162461bcd60e51b8152602060048201526011602482015270696e76616c6964207369676e617475726560781b60448201526064016102fa565b61089f886040611a08565b97508988146108f05760405162461bcd60e51b815260206004820152601860248201527f6d616c666f726d6564206576656e7420656e636f64696e67000000000000000060448201526064016102fa565b6001600160801b038516600090815260056020526040902054610914908590611a08565b6001600160801b038616600090815260056020526040812091909155819060069061094483805160209091012090565b6001600160a01b03166001600160a01b0316815260200190815260200160002090805190602001906109779291906117bd565b50805160208201206001600160a01b03167fed126a3daa455d580cce3268a123812d4eece1deca8bf9b0e6c72eb535e28cd087878786886040516109bf959493929190611a97565b60405180910390a2805160208201206109dc908787878688610dff565b5050505050505050505050565b60006109f6826010611a08565b83511015610a3e5760405162461bcd60e51b8152602060048201526015602482015274746f55696e743132385f6f75744f66426f756e647360581b60448201526064016102fa565b50016010015190565b6000610a54826008611a08565b83511015610a9b5760405162461bcd60e51b8152602060048201526014602482015273746f55696e7436345f6f75744f66426f756e647360601b60448201526064016102fa565b50016008015190565b6000610ab1826002611a08565b83511015610af85760405162461bcd60e51b8152602060048201526014602482015273746f55696e7431365f6f75744f66426f756e647360601b60448201526064016102fa565b50016002015190565b606081610b0f81601f611a08565b1015610b4e5760405162461bcd60e51b815260206004820152600e60248201526d736c6963655f6f766572666c6f7760901b60448201526064016102fa565b610b588284611a08565b84511015610b9c5760405162461bcd60e51b8152602060048201526011602482015270736c6963655f6f75744f66426f756e647360781b60448201526064016102fa565b606082158015610bbb5760405191506000825260208201604052610c05565b6040519150601f8416801560200281840101858101878315602002848b0101015b81831015610bf4578051835260209283019201610bdc565b5050858452601f01601f1916604052505b50949350505050565b6060806040519050835180825260208201818101602087015b81831015610c3f578051835260209283019201610c27565b50855184518101855292509050808201602086015b81831015610c6c578051835260209283019201610c54565b508651929092011591909101601f01601f191660405250905092915050565b6000610c98826020611a08565b83511015610ce05760405162461bcd60e51b8152602060048201526015602482015274746f55696e743235365f6f75744f66426f756e647360581b60448201526064016102fa565b50016020015190565b6000806040518060400160405280610d0a8587610c8b90919063ffffffff16565b8152602001610d26856020610d1f9190611a08565b8790610c8b565b9052604080516002808252818301909252919250600091610d7691610d71919060208201818036833701905050610d6b6000610d6360028a611a4b565b8a9190610b01565b90610c0e565b610e60565b60405163ebbdac9160e01b815290915073__$f8599c2da3d5c7b9f587eddf89abb6c346$__9063ebbdac9190610db59085906000908690600401611b0c565b602060405180830381865af4158015610dd2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610df69190611b5a565b95945050505050565b600080610e0c8382610c8b565b9050610e19826020611a08565b915080610e4b576000610e2c8484610a47565b9050610e4487878b610e3e858a611a20565b8c610e7e565b5050610e56565b610e56818986610f9f565b5050505050505050565b610e68611841565b610e7882805190602001206111cf565b92915050565b60006001600160a01b038416610ec45760405162461bcd60e51b815260206004820152600b60248201526a19dd5e4b5b9bdd0b5cd95d60aa1b60448201526064016102fa565b60001960085410610f025760405162461bcd60e51b81526020600482015260086024820152676f766572666c6f7760c01b60448201526064016102fa565b600860008154610f1190611b7c565b9182905550600081815260076020526040902080546001600160801b0319166001600160801b03989098169790971787556001870195909555506002850180546001600160a01b03949094166001600160e01b031990941693909317600160a01b6001600160401b0393841602179092556003909301805467ffffffffffffffff1916919093161790915590565b6000838152600760205260409020600201546001600160a01b03838116911614610ffc5760405162461bcd60e51b815260206004820152600e60248201526d1b9bdd0b585d5d1a1bdc9a5e995960921b60448201526064016102fa565b6000838152600760205260409020600201546001600160401b03808316600160a01b90920416106110595760405162461bcd60e51b81526020600482015260076024820152661b9bdd0b595b9960ca1b60448201526064016102fa565b600061106484611279565b506000858152600760209081526040808320600381015481546001909201546001600160a01b038a168652600690945291842080549596509394611148946001600160401b03909316936001600160801b039092169287916110c5906119b7565b80601f01602080910402602001604051908101604052809291908181526020018280546110f1906119b7565b801561113e5780601f106111135761010080835404028352916020019161113e565b820191906000526020600020905b81548152906001019060200180831161112157829003601f168201915b5050505050611387565b90507fdb53e751d28ed0d6e3682814bf8d23f7dd7b29c94f74a56fbb7f88e9dca9f39b8160405161117991906118f1565b60405180910390a150505060009182525060076020526040812080546001600160801b031916815560018101919091556002810180546001600160e01b0319169055600301805467ffffffffffffffff19169055565b6111d7611841565b60006111f1600080516020611c1583398151915284611b97565b90506000805b600080516020611c158339815191528384099150600080516020611c158339815191528383099150600080516020611c1583398151915260038308915061123d82611549565b909250905080156112575782845260208401829052611271565b600080516020611c158339815191526001840892506111f7565b505050919050565b6040805160208082528183019092526060916000918291602082018180368337019050509050836000805b60208161ffff16101561134c57828161ffff16602081106112c7576112c7611bb9565b1a60f81b848261ffff16815181106112e1576112e1611bb9565b60200101906001600160f81b031916908160001a905350600060f81b848261ffff168151811061131357611313611bb9565b01602001516001600160f81b031916118015611331575061ffff8216155b1561133a578091505b8061134481611bcf565b9150506112a4565b50600061135a826020611bf1565b905061137b61ffff831661136f846020611bf1565b86919061ffff16610b01565b97909650945050505050565b606060808351106113cc5760405162461bcd60e51b815260206004820152600f60248201526e657874726120746f6f206c6172676560881b60448201526064016102fa565b6001600160801b03851660009081526005602052604090205484111561142d5760405162461bcd60e51b815260206004820152601660248201527534b739bab33334b1b4b2b73a1031bab9ba37b234b0b760511b60448201526064016102fa565b6001600160801b038516600090815260056020526040902054611451908590611a4b565b6001600160801b0380871660009081526005602052604081209290925560045461147b911661161f565b9050611490611489886116af565b8290610c0e565b905061149e6114898761161f565b90506000806114ac87611279565b915091506114c36114bc82611736565b8490610c0e565b92506114cf8383610c0e565b92506114de6114bc8751611736565b92506114ea8387610c0e565b925061152460085b6040519080825280601f01601f19166020018201604052801561151c576020820181803683370190505b508490610c0e565b92506115308386610c0e565b925061153c60026114f2565b9998505050505050505050565b600080600060405160208152602080820152602060408201528460608201527f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f526080820152600080516020611c1583398151915260a082015260208160c08360056107d05a03fa905193509050600080516020611c1583398151915283800984149150806116195760405162461bcd60e51b815260206004820152601c60248201527f424c533a2073717274206d6f646578702063616c6c206661696c65640000000060448201526064016102fa565b50915091565b60408051601080825281830190925260609160009190602082018180368337019050509050608083901b60005b60108110156116a65781816010811061166757611667611bb9565b1a60f81b83828151811061167d5761167d611bb9565b60200101906001600160f81b031916908160001a9053508061169e81611b7c565b91505061164c565b50909392505050565b6040805160088082528183019092526060916000919060208201818036833701905050905060c083901b60005b60088110156116a6578181600881106116f7576116f7611bb9565b1a60f81b83828151811061170d5761170d611bb9565b60200101906001600160f81b031916908160001a9053508061172e81611b7c565b9150506116dc565b6040805160028082528183019092526060916000919060208201818036833701905050905060f083901b60005b60028110156116a65781816002811061177e5761177e611bb9565b1a60f81b83828151811061179457611794611bb9565b60200101906001600160f81b031916908160001a905350806117b581611b7c565b915050611763565b8280546117c9906119b7565b90600052602060002090601f0160209004810192826117eb5760008555611831565b82601f1061180457805160ff1916838001178555611831565b82800160010185558215611831579182015b82811115611831578251825591602001919060010190611816565b5061183d92915061185f565b5090565b60405180604001604052806002906020820280368337509192915050565b5b8082111561183d5760008155600101611860565b60006020828403121561188657600080fd5b81356001600160a01b038116811461189d57600080fd5b9392505050565b6000815180845260005b818110156118ca576020818501810151868301820152016118ae565b818111156118dc576000602083870101525b50601f01601f19169290920160200192915050565b60208152600061189d60208301846118a4565b60006020828403121561191657600080fd5b81356001600160801b038116811461189d57600080fd5b60006020828403121561193f57600080fd5b5035919050565b6000806020838503121561195957600080fd5b82356001600160401b038082111561197057600080fd5b818501915085601f83011261198457600080fd5b81358181111561199357600080fd5b8660208285010111156119a557600080fd5b60209290920196919550909350505050565b600181811c908216806119cb57607f821691505b602082108114156119ec57634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052601160045260246000fd5b60008219821115611a1b57611a1b6119f2565b500190565b60006001600160401b03808316818516808303821115611a4257611a426119f2565b01949350505050565b600082821015611a5d57611a5d6119f2565b500390565b634e487b7160e01b600052604160045260246000fd5b6000816000190483118215151615611a9257611a926119f2565b500290565b60006001600160401b0380881683526001600160801b038716602084015285604084015280851660608401525060a06080830152611ad860a08301846118a4565b979650505050505050565b8060005b6002811015611b06578151845260209384019390910190600101611ae7565b50505050565b6101008101611b1b8286611ae3565b604082018460005b6004811015611b42578154835260209092019160019182019101611b23565b505050611b5260c0830184611ae3565b949350505050565b600060208284031215611b6c57600080fd5b8151801515811461189d57600080fd5b6000600019821415611b9057611b906119f2565b5060010190565b600082611bb457634e487b7160e01b600052601260045260246000fd5b500690565b634e487b7160e01b600052603260045260246000fd5b600061ffff80831681811415611be757611be76119f2565b6001019392505050565b600061ffff83811690831681811015611c0c57611c0c6119f2565b03939250505056fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47a264697066735822122015060dd8bcbd01d7b6a1fcb868ead96118262108a8e5e2b379252a62e7a385fd64736f6c634300080a0033",
}

// HodlABI is the input ABI used to generate the binding from.
// Deprecated: Use HodlMetaData.ABI instead.
var HodlABI = HodlMetaData.ABI

// Deprecated: Use HodlMetaData.Sigs instead.
// HodlFuncSigs maps the 4-byte function signature to its string representation.
var HodlFuncSigs = HodlMetaData.Sigs

// HodlBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use HodlMetaData.Bin instead.
var HodlBin = HodlMetaData.Bin

// DeployHodl deploys a new Ethereum contract, binding an instance of Hodl to it.
func DeployHodl(auth *bind.TransactOpts, backend bind.ContractBackend, pid_ *big.Int) (common.Address, *types.Transaction, *Hodl, error) {
	parsed, err := HodlMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	bLSAddr, _, _, _ := DeployBLS(auth, backend)
	HodlBin = strings.Replace(HodlBin, "__$f8599c2da3d5c7b9f587eddf89abb6c346$__", bLSAddr.String()[2:], -1)

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(HodlBin), backend, pid_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Hodl{HodlCaller: HodlCaller{contract: contract}, HodlTransactor: HodlTransactor{contract: contract}, HodlFilterer: HodlFilterer{contract: contract}}, nil
}

// Hodl is an auto generated Go binding around an Ethereum contract.
type Hodl struct {
	HodlCaller     // Read-only binding to the contract
	HodlTransactor // Write-only binding to the contract
	HodlFilterer   // Log filterer for contract events
}

// HodlCaller is an auto generated read-only Go binding around an Ethereum contract.
type HodlCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HodlTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HodlTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HodlFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HodlFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HodlSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HodlSession struct {
	Contract     *Hodl             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HodlCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HodlCallerSession struct {
	Contract *HodlCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// HodlTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HodlTransactorSession struct {
	Contract     *HodlTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HodlRaw is an auto generated low-level Go binding around an Ethereum contract.
type HodlRaw struct {
	Contract *Hodl // Generic contract binding to access the raw methods on
}

// HodlCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HodlCallerRaw struct {
	Contract *HodlCaller // Generic read-only contract binding to access the raw methods on
}

// HodlTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HodlTransactorRaw struct {
	Contract *HodlTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHodl creates a new instance of Hodl, bound to a specific deployed contract.
func NewHodl(address common.Address, backend bind.ContractBackend) (*Hodl, error) {
	contract, err := bindHodl(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Hodl{HodlCaller: HodlCaller{contract: contract}, HodlTransactor: HodlTransactor{contract: contract}, HodlFilterer: HodlFilterer{contract: contract}}, nil
}

// NewHodlCaller creates a new read-only instance of Hodl, bound to a specific deployed contract.
func NewHodlCaller(address common.Address, caller bind.ContractCaller) (*HodlCaller, error) {
	contract, err := bindHodl(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HodlCaller{contract: contract}, nil
}

// NewHodlTransactor creates a new write-only instance of Hodl, bound to a specific deployed contract.
func NewHodlTransactor(address common.Address, transactor bind.ContractTransactor) (*HodlTransactor, error) {
	contract, err := bindHodl(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HodlTransactor{contract: contract}, nil
}

// NewHodlFilterer creates a new log filterer instance of Hodl, bound to a specific deployed contract.
func NewHodlFilterer(address common.Address, filterer bind.ContractFilterer) (*HodlFilterer, error) {
	contract, err := bindHodl(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HodlFilterer{contract: contract}, nil
}

// bindHodl binds a generic wrapper to an already deployed contract.
func bindHodl(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HodlABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Hodl *HodlRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Hodl.Contract.HodlCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Hodl *HodlRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hodl.Contract.HodlTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Hodl *HodlRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Hodl.Contract.HodlTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Hodl *HodlCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Hodl.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Hodl *HodlTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hodl.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Hodl *HodlTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Hodl.Contract.contract.Transact(opts, method, params...)
}

// GROUPS is a free data retrieval call binding the contract method 0x50cc55ae.
//
// Solidity: function GROUPS(uint256 ) view returns(uint256)
func (_Hodl *HodlCaller) GROUPS(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Hodl.contract.Call(opts, &out, "GROUPS", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GROUPS is a free data retrieval call binding the contract method 0x50cc55ae.
//
// Solidity: function GROUPS(uint256 ) view returns(uint256)
func (_Hodl *HodlSession) GROUPS(arg0 *big.Int) (*big.Int, error) {
	return _Hodl.Contract.GROUPS(&_Hodl.CallOpts, arg0)
}

// GROUPS is a free data retrieval call binding the contract method 0x50cc55ae.
//
// Solidity: function GROUPS(uint256 ) view returns(uint256)
func (_Hodl *HodlCallerSession) GROUPS(arg0 *big.Int) (*big.Int, error) {
	return _Hodl.Contract.GROUPS(&_Hodl.CallOpts, arg0)
}

// NONCE is a free data retrieval call binding the contract method 0xe091dd1a.
//
// Solidity: function NONCE() view returns(uint64)
func (_Hodl *HodlCaller) NONCE(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Hodl.contract.Call(opts, &out, "NONCE")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// NONCE is a free data retrieval call binding the contract method 0xe091dd1a.
//
// Solidity: function NONCE() view returns(uint64)
func (_Hodl *HodlSession) NONCE() (uint64, error) {
	return _Hodl.Contract.NONCE(&_Hodl.CallOpts)
}

// NONCE is a free data retrieval call binding the contract method 0xe091dd1a.
//
// Solidity: function NONCE() view returns(uint64)
func (_Hodl *HodlCallerSession) NONCE() (uint64, error) {
	return _Hodl.Contract.NONCE(&_Hodl.CallOpts)
}

// PID is a free data retrieval call binding the contract method 0x5eaec0e4.
//
// Solidity: function PID() view returns(uint128)
func (_Hodl *HodlCaller) PID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Hodl.contract.Call(opts, &out, "PID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PID is a free data retrieval call binding the contract method 0x5eaec0e4.
//
// Solidity: function PID() view returns(uint128)
func (_Hodl *HodlSession) PID() (*big.Int, error) {
	return _Hodl.Contract.PID(&_Hodl.CallOpts)
}

// PID is a free data retrieval call binding the contract method 0x5eaec0e4.
//
// Solidity: function PID() view returns(uint128)
func (_Hodl *HodlCallerSession) PID() (*big.Int, error) {
	return _Hodl.Contract.PID(&_Hodl.CallOpts)
}

// Custodian is a free data retrieval call binding the contract method 0x2d459b6b.
//
// Solidity: function custodian(uint128 ) view returns(uint256)
func (_Hodl *HodlCaller) Custodian(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Hodl.contract.Call(opts, &out, "custodian", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Custodian is a free data retrieval call binding the contract method 0x2d459b6b.
//
// Solidity: function custodian(uint128 ) view returns(uint256)
func (_Hodl *HodlSession) Custodian(arg0 *big.Int) (*big.Int, error) {
	return _Hodl.Contract.Custodian(&_Hodl.CallOpts, arg0)
}

// Custodian is a free data retrieval call binding the contract method 0x2d459b6b.
//
// Solidity: function custodian(uint128 ) view returns(uint256)
func (_Hodl *HodlCallerSession) Custodian(arg0 *big.Int) (*big.Int, error) {
	return _Hodl.Contract.Custodian(&_Hodl.CallOpts, arg0)
}

// Members is a free data retrieval call binding the contract method 0x08ae4b0c.
//
// Solidity: function members(address ) view returns(bytes)
func (_Hodl *HodlCaller) Members(opts *bind.CallOpts, arg0 common.Address) ([]byte, error) {
	var out []interface{}
	err := _Hodl.contract.Call(opts, &out, "members", arg0)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// Members is a free data retrieval call binding the contract method 0x08ae4b0c.
//
// Solidity: function members(address ) view returns(bytes)
func (_Hodl *HodlSession) Members(arg0 common.Address) ([]byte, error) {
	return _Hodl.Contract.Members(&_Hodl.CallOpts, arg0)
}

// Members is a free data retrieval call binding the contract method 0x08ae4b0c.
//
// Solidity: function members(address ) view returns(bytes)
func (_Hodl *HodlCallerSession) Members(arg0 common.Address) ([]byte, error) {
	return _Hodl.Contract.Members(&_Hodl.CallOpts, arg0)
}

// Vats is a free data retrieval call binding the contract method 0x40a4e253.
//
// Solidity: function vats(uint256 ) view returns(uint128 asset, uint256 amount, address guy, uint64 end, uint64 nonce)
func (_Hodl *HodlCaller) Vats(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Asset  *big.Int
	Amount *big.Int
	Guy    common.Address
	End    uint64
	Nonce  uint64
}, error) {
	var out []interface{}
	err := _Hodl.contract.Call(opts, &out, "vats", arg0)

	outstruct := new(struct {
		Asset  *big.Int
		Amount *big.Int
		Guy    common.Address
		End    uint64
		Nonce  uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Asset = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Amount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Guy = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.End = *abi.ConvertType(out[3], new(uint64)).(*uint64)
	outstruct.Nonce = *abi.ConvertType(out[4], new(uint64)).(*uint64)

	return *outstruct, err

}

// Vats is a free data retrieval call binding the contract method 0x40a4e253.
//
// Solidity: function vats(uint256 ) view returns(uint128 asset, uint256 amount, address guy, uint64 end, uint64 nonce)
func (_Hodl *HodlSession) Vats(arg0 *big.Int) (struct {
	Asset  *big.Int
	Amount *big.Int
	Guy    common.Address
	End    uint64
	Nonce  uint64
}, error) {
	return _Hodl.Contract.Vats(&_Hodl.CallOpts, arg0)
}

// Vats is a free data retrieval call binding the contract method 0x40a4e253.
//
// Solidity: function vats(uint256 ) view returns(uint128 asset, uint256 amount, address guy, uint64 end, uint64 nonce)
func (_Hodl *HodlCallerSession) Vats(arg0 *big.Int) (struct {
	Asset  *big.Int
	Amount *big.Int
	Guy    common.Address
	End    uint64
	Nonce  uint64
}, error) {
	return _Hodl.Contract.Vats(&_Hodl.CallOpts, arg0)
}

// Mixin is a paid mutator transaction binding the contract method 0x5cae8005.
//
// Solidity: function mixin(bytes raw) returns()
func (_Hodl *HodlTransactor) Mixin(opts *bind.TransactOpts, raw []byte) (*types.Transaction, error) {
	return _Hodl.contract.Transact(opts, "mixin", raw)
}

// Mixin is a paid mutator transaction binding the contract method 0x5cae8005.
//
// Solidity: function mixin(bytes raw) returns()
func (_Hodl *HodlSession) Mixin(raw []byte) (*types.Transaction, error) {
	return _Hodl.Contract.Mixin(&_Hodl.TransactOpts, raw)
}

// Mixin is a paid mutator transaction binding the contract method 0x5cae8005.
//
// Solidity: function mixin(bytes raw) returns()
func (_Hodl *HodlTransactorSession) Mixin(raw []byte) (*types.Transaction, error) {
	return _Hodl.Contract.Mixin(&_Hodl.TransactOpts, raw)
}

// HodlMixinEventIterator is returned from FilterMixinEvent and is used to iterate over the raw logs and unpacked data for MixinEvent events raised by the Hodl contract.
type HodlMixinEventIterator struct {
	Event *HodlMixinEvent // Event containing the contract specifics and raw log

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
func (it *HodlMixinEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HodlMixinEvent)
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
		it.Event = new(HodlMixinEvent)
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
func (it *HodlMixinEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HodlMixinEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HodlMixinEvent represents a MixinEvent event raised by the Hodl contract.
type HodlMixinEvent struct {
	Sender    common.Address
	Nonce     *big.Int
	Asset     *big.Int
	Amount    *big.Int
	Timestamp uint64
	Extra     []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMixinEvent is a free log retrieval operation binding the contract event 0xed126a3daa455d580cce3268a123812d4eece1deca8bf9b0e6c72eb535e28cd0.
//
// Solidity: event MixinEvent(address indexed sender, uint256 nonce, uint128 asset, uint256 amount, uint64 timestamp, bytes extra)
func (_Hodl *HodlFilterer) FilterMixinEvent(opts *bind.FilterOpts, sender []common.Address) (*HodlMixinEventIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Hodl.contract.FilterLogs(opts, "MixinEvent", senderRule)
	if err != nil {
		return nil, err
	}
	return &HodlMixinEventIterator{contract: _Hodl.contract, event: "MixinEvent", logs: logs, sub: sub}, nil
}

// WatchMixinEvent is a free log subscription operation binding the contract event 0xed126a3daa455d580cce3268a123812d4eece1deca8bf9b0e6c72eb535e28cd0.
//
// Solidity: event MixinEvent(address indexed sender, uint256 nonce, uint128 asset, uint256 amount, uint64 timestamp, bytes extra)
func (_Hodl *HodlFilterer) WatchMixinEvent(opts *bind.WatchOpts, sink chan<- *HodlMixinEvent, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Hodl.contract.WatchLogs(opts, "MixinEvent", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HodlMixinEvent)
				if err := _Hodl.contract.UnpackLog(event, "MixinEvent", log); err != nil {
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

// ParseMixinEvent is a log parse operation binding the contract event 0xed126a3daa455d580cce3268a123812d4eece1deca8bf9b0e6c72eb535e28cd0.
//
// Solidity: event MixinEvent(address indexed sender, uint256 nonce, uint128 asset, uint256 amount, uint64 timestamp, bytes extra)
func (_Hodl *HodlFilterer) ParseMixinEvent(log types.Log) (*HodlMixinEvent, error) {
	event := new(HodlMixinEvent)
	if err := _Hodl.contract.UnpackLog(event, "MixinEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HodlMixinTransactionIterator is returned from FilterMixinTransaction and is used to iterate over the raw logs and unpacked data for MixinTransaction events raised by the Hodl contract.
type HodlMixinTransactionIterator struct {
	Event *HodlMixinTransaction // Event containing the contract specifics and raw log

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
func (it *HodlMixinTransactionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HodlMixinTransaction)
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
		it.Event = new(HodlMixinTransaction)
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
func (it *HodlMixinTransactionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HodlMixinTransactionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HodlMixinTransaction represents a MixinTransaction event raised by the Hodl contract.
type HodlMixinTransaction struct {
	Arg0 []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterMixinTransaction is a free log retrieval operation binding the contract event 0xdb53e751d28ed0d6e3682814bf8d23f7dd7b29c94f74a56fbb7f88e9dca9f39b.
//
// Solidity: event MixinTransaction(bytes arg0)
func (_Hodl *HodlFilterer) FilterMixinTransaction(opts *bind.FilterOpts) (*HodlMixinTransactionIterator, error) {

	logs, sub, err := _Hodl.contract.FilterLogs(opts, "MixinTransaction")
	if err != nil {
		return nil, err
	}
	return &HodlMixinTransactionIterator{contract: _Hodl.contract, event: "MixinTransaction", logs: logs, sub: sub}, nil
}

// WatchMixinTransaction is a free log subscription operation binding the contract event 0xdb53e751d28ed0d6e3682814bf8d23f7dd7b29c94f74a56fbb7f88e9dca9f39b.
//
// Solidity: event MixinTransaction(bytes arg0)
func (_Hodl *HodlFilterer) WatchMixinTransaction(opts *bind.WatchOpts, sink chan<- *HodlMixinTransaction) (event.Subscription, error) {

	logs, sub, err := _Hodl.contract.WatchLogs(opts, "MixinTransaction")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HodlMixinTransaction)
				if err := _Hodl.contract.UnpackLog(event, "MixinTransaction", log); err != nil {
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

// ParseMixinTransaction is a log parse operation binding the contract event 0xdb53e751d28ed0d6e3682814bf8d23f7dd7b29c94f74a56fbb7f88e9dca9f39b.
//
// Solidity: event MixinTransaction(bytes arg0)
func (_Hodl *HodlFilterer) ParseMixinTransaction(log types.Log) (*HodlMixinTransaction, error) {
	event := new(HodlMixinTransaction)
	if err := _Hodl.contract.UnpackLog(event, "MixinTransaction", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
