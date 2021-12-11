package lock

import (
	"errors"
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/fox-one/hodl/config"
	"github.com/fox-one/hodl/core"
	"github.com/fox-one/hodl/invoker"
	"github.com/fox-one/hodl/writer"
	"github.com/fox-one/mixin-sdk-go"
	"github.com/shopspring/decimal"
	"github.com/spf13/cobra"
)

// NewCmd returns a new instance of the lock command.
func NewCmd(cfg *config.Config) *cobra.Command {
	client, err := mixin.NewFromKeystore(&cfg.Dapp)
	if err != nil {
		panic(fmt.Errorf("create mixin client failed: %w", err))
	}

	ivk := invoker.New(client, invoker.Config{
		Process:   cfg.System.ProcessID,
		Members:   cfg.Group.Members,
		Threshold: cfg.Group.Threshold,
	})

	cmd := &cobra.Command{
		Use:     "lock",
		Example: "lock 10 btc 24h",
		Args:    cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()

			amount, _ := decimal.NewFromString(args[0])
			amount = amount.Truncate(8)

			if !amount.IsPositive() {
				return errors.New("amount must be positive")
			}

			symbol := strings.ToLower(args[1])
			assets, err := mixin.ReadNetworkAssetsBySymbol(ctx, symbol)
			if err != nil {
				return fmt.Errorf("read assets by symbol %q failed", symbol)
			}

			if len(assets) == 0 {
				return fmt.Errorf("assets %q not found", symbol)
			}

			sort.Slice(assets, func(i, j int) bool {
				return assets[i].PriceUSD.GreaterThan(assets[j].PriceUSD)
			})

			asset := assets[0]

			exp, err := time.ParseDuration(args[2])
			if err != nil {
				return fmt.Errorf("parse expire duration %q failed", args[2])
			}

			event := core.Event{
				Action:  core.ActionLock,
				VaultID: 0,
				Exp:     uint64(exp),
			}

			code, err := ivk.Payment(ctx, asset.AssetID, amount, event.Encode())
			if err != nil {
				return fmt.Errorf("payment failed: %w", err)
			}

			action := mixin.URL.Codes(code)
			label := fmt.Sprintf("lock %s %s", amount, symbol)
			cmd.Println(writer.WithLabel(action, label))
			return nil
		},
	}

	return cmd
}
