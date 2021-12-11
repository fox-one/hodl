package unlock

import (
	"fmt"

	"github.com/fox-one/hodl/config"
	"github.com/fox-one/hodl/core"
	"github.com/fox-one/hodl/invoker"
	"github.com/fox-one/hodl/writer"
	"github.com/fox-one/mixin-sdk-go"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

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
		Use:     "unlock",
		Example: "unlock {id}",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()

			id := cast.ToUint64(args[0])
			if id == 0 {
				return fmt.Errorf("invalid id: %s", args[0])
			}

			event := core.Event{
				Action:  core.ActionUnlock,
				VaultID: id,
			}

			code, err := ivk.Payment(ctx, cfg.System.GasAssetID, cfg.System.GasAmount, event.Encode())
			if err != nil {
				return fmt.Errorf("payment failed: %w", err)
			}

			action := mixin.URL.Codes(code)
			label := fmt.Sprintf("unlock #%d", id)
			cmd.Println(writer.WithLabel(action, label))
			return nil
		},
	}

	return cmd
}
