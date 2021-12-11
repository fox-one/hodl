package blaze

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/fox-one/hodl/cmd/lock"
	"github.com/fox-one/hodl/cmd/unlock"
	"github.com/fox-one/hodl/config"
	"github.com/fox-one/hodl/writer/mm"
	"github.com/fox-one/mixin-sdk-go"
	"github.com/spf13/cobra"
)

func NewCmd(cfg *config.Config) *cobra.Command {
	client, err := mixin.NewFromKeystore(&cfg.Dapp)
	if err != nil {
		panic(fmt.Errorf("create mixin client failed: %w", err))
	}

	var handler mixin.BlazeListenFunc = func(ctx context.Context, msg *mixin.MessageView, userID string) error {
		if msg.Category != mixin.MessageCategoryPlainText {
			return nil
		}

		b, _ := base64.StdEncoding.DecodeString(msg.Data)
		args := strings.Fields(string(b))

		root := &cobra.Command{Use: "blaze"}
		root.AddCommand(lock.NewCmd(cfg))
		root.AddCommand(unlock.NewCmd(cfg))
		root.SetArgs(args)
		w := mm.Reply(client, msg.UserID)
		root.SetOut(w)

		if err := root.ExecuteContext(ctx); err != nil {
			log.Printf("execute command %q failed: %v", string(b), err)
		}

		return w.Flush(ctx)
	}

	cmd := &cobra.Command{
		Use:           "blaze",
		SilenceErrors: true,
		SilenceUsage:  true,
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()

			for {
				select {
				case <-ctx.Done():
					return ctx.Err()
				case <-time.After(time.Second):
					if err := client.LoopBlaze(ctx, handler); err != nil {
						log.Println("LoopBlaze", err)
					}
				}
			}
		},
	}

	return cmd
}