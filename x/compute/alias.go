// nolint
// autogenerated code using github.com/rigelrozanski/multitool
// aliases generated for the following subdirectories:
// ALIASGEN: github.com/enigmampc/SecretNetwork/x/compute/internal/types
// ALIASGEN: github.com/enigmampc/SecretNetwork/x/compute/internal/keeper
package compute

import (
	"github.com/enigmampc/SecretNetwork/x/compute/internal/keeper"
	"github.com/enigmampc/SecretNetwork/x/compute/internal/types"
)

const (
	ModuleName                    = types.ModuleName
	StoreKey                      = types.StoreKey
	TStoreKey                     = types.TStoreKey
	QuerierRoute                  = types.QuerierRoute
	RouterKey                     = types.RouterKey
	MaxWasmSize                   = types.MaxWasmSize
	GasMultiplier                 = keeper.GasMultiplier
	MaxGas                        = keeper.MaxGas
	QueryGetContract              = keeper.QueryGetContract
	QueryGetContractState         = keeper.QueryGetContractState
	QueryGetCode                  = keeper.QueryGetCode
	QueryListCode                 = keeper.QueryListCode
	QueryListContractByCode       = keeper.QueryListContractByCode
	QueryMethodContractStateSmart = keeper.QueryMethodContractStateSmart
	QueryMethodContractStateAll   = keeper.QueryMethodContractStateAll
	QueryMethodContractStateRaw   = keeper.QueryMethodContractStateRaw
)

var (
	// functions aliases
	RegisterCodec             = types.RegisterCodec
	ValidateGenesis           = types.ValidateGenesis
	GetCodeKey                = types.GetCodeKey
	GetContractAddressKey     = types.GetContractAddressKey
	GetContractStorePrefixKey = types.GetContractStorePrefixKey
	NewCodeInfo               = types.NewCodeInfo
	NewParams                 = types.NewParams
	NewWasmCoins              = types.NewWasmCoins
	NewContractInfo           = types.NewContractInfo
	CosmosResult              = types.CosmosResult
	DefaultWasmConfig         = types.DefaultWasmConfig
	InitGenesis               = keeper.InitGenesis
	ExportGenesis             = keeper.ExportGenesis
	NewKeeper                 = keeper.NewKeeper
	NewQuerier                = keeper.NewQuerier
	MakeTestCodec             = keeper.MakeTestCodec
	CreateTestInput           = keeper.CreateTestInput

	// variable aliases
	ModuleCdc            = types.ModuleCdc
	DefaultCodespace     = types.DefaultCodespace
	ErrCreateFailed      = types.ErrCreateFailed
	ErrAccountExists     = types.ErrAccountExists
	ErrInstantiateFailed = types.ErrInstantiateFailed
	ErrExecuteFailed     = types.ErrExecuteFailed
	ErrGasLimit          = types.ErrGasLimit
	ErrInvalidGenesis    = types.ErrInvalidGenesis
	ErrNotFound          = types.ErrNotFound
	ErrQueryFailed       = types.ErrQueryFailed
	KeyLastCodeID        = types.KeyLastCodeID
	KeyLastInstanceID    = types.KeyLastInstanceID
	CodeKeyPrefix        = types.CodeKeyPrefix
	ContractKeyPrefix    = types.ContractKeyPrefix
	ContractStorePrefix  = types.ContractStorePrefix
)

type (
	GenesisState            = types.GenesisState
	Code                    = types.Code
	Contract                = types.Contract
	MsgStoreCode            = types.MsgStoreCode
	MsgInstantiateContract  = types.MsgInstantiateContract
	MsgExecuteContract      = types.MsgExecuteContract
	Model                   = types.Model
	CodeInfo                = types.CodeInfo
	ContractInfo            = types.ContractInfo
	WasmConfig              = types.WasmConfig
	Keeper                  = keeper.Keeper
	ContractInfoWithAddress = keeper.ContractInfoWithAddress
	GetCodeResponse         = keeper.GetCodeResponse
	ListCodeResponse        = keeper.ListCodeResponse
)
