package config

import (
	"github.com/fox-one/hodl/core"
	"github.com/fox-one/mixin-sdk-go"
	jsoniter "github.com/json-iterator/go"
	"github.com/shopspring/decimal"
	"github.com/spf13/viper"
)

type Config struct {
	Group  core.Group     `json:"group,omitempty"`
	System System         `json:"system,omitempty"`
	Dapp   mixin.Keystore `json:"dapp,omitempty"`
	MVM    MVM            `json:"mvm,omitempty"`
}

type System struct {
	ProcessID  string          `json:"process_id,omitempty"`
	GasAssetID string          `json:"gas_asset_id,omitempty"`
	GasAmount  decimal.Decimal `json:"gas_amount,omitempty"`
}

type MVM struct {
	RPC             string `json:"rpc,omitempty"`
	ContractAddress string `json:"contract_address,omitempty"`
}

func Load() (*Config, error) {
	b, err := jsoniter.Marshal(viper.AllSettings())
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := jsoniter.Unmarshal(b, &cfg); err != nil {
		return nil, err
	}

	// default config
	if cfg.System.GasAssetID == "" {
		cfg.System.GasAssetID = "965e5c6e-434c-3fa9-b780-c50f43cd955c" // cnb
	}

	if !cfg.System.GasAmount.IsPositive() {
		cfg.System.GasAmount = decimal.NewFromInt(1)
	}

	return &cfg, nil
}
