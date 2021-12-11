package list

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/fox-one/hodl/core"
	"github.com/spf13/cobra"
)

func NewCmd(vats core.VaultStore, userID string) *cobra.Command {
	cmd := &cobra.Command{
		Use: "list",
		RunE: func(cmd *cobra.Command, args []string) error {
			vaults, err := vats.List(userID)
			if err != nil {
				log.Println("list vaults error:", err)
				return err
			}

			var b strings.Builder
			b.WriteString(fmt.Sprintf("%d vaults\n\n", len(vaults)))

			for _, v := range vaults {
				b.WriteString(fmt.Sprintf("#%d %s %s ", v.ID, v.Amount, v.Symbol))

				if remain := time.Until(v.End); remain > 0 {
					b.WriteString(remain.String())
				} else {
					b.WriteString("expired")
				}

				b.WriteByte('\n')
			}

			cmd.Println(b.String())
			return nil
		},
	}

	return cmd
}
