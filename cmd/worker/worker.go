package worker

import (
	"fmt"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/fox-one/hodl/config"
	"github.com/fox-one/hodl/worker"
	"github.com/fox-one/hodl/worker/events"
	"github.com/fox-one/mixin-sdk-go"
	"github.com/spf13/cobra"
	"golang.org/x/sync/errgroup"
)

func NewCmd(cfg *config.Config) *cobra.Command {
	client, err := mixin.NewFromKeystore(&cfg.Dapp)
	if err != nil {
		panic(fmt.Errorf("create mixin client failed: %w", err))
	}

	conn, err := ethclient.Dial(cfg.MVM.RPC)
	if err != nil {
		panic(err)
	}

	cmd := &cobra.Command{
		Use: "worker",
		RunE: func(cmd *cobra.Command, args []string) error {
			workers := []worker.Worker{
				events.New(client, conn, cfg.System.ProcessID, cfg.MVM.ContractAddress),
			}

			var g errgroup.Group

			for idx := range workers {
				idx := idx
				g.Go(func() error {
					return workers[idx].Run(cmd.Context())
				})
			}

			return g.Wait()
		},
	}

	return cmd
}
