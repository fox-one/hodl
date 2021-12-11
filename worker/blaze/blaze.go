package blaze

import (
	"context"
	"encoding/base64"
	"log"
	"strings"
	"time"

	"github.com/fox-one/hodl/cmd/list"
	"github.com/fox-one/hodl/cmd/lock"
	"github.com/fox-one/hodl/cmd/unlock"
	"github.com/fox-one/hodl/config"
	"github.com/fox-one/hodl/core"
	"github.com/fox-one/hodl/writer/mm"
	"github.com/fox-one/mixin-sdk-go"
	"github.com/spf13/cobra"
)

const usageMessage = `欢迎体验大牢宝 MVM 版本，支持以下命令：

将 0.1 btc 冻结到大牢宝1小时/30分钟/60秒 
lock 0.1 btc 1h/30m/60s

解冻 ID 为 1 的金库
unlock 1 		

查看我的金库
list
`

type Blaze struct {
	client *mixin.Client
	vats   core.VaultStore
	cfg    *config.Config
}

func New(client *mixin.Client, vats core.VaultStore, cfg *config.Config) *Blaze {
	return &Blaze{
		client: client,
		vats:   vats,
		cfg:    cfg,
	}
}

func (w *Blaze) Run(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(time.Second):
			if err := w.client.LoopBlaze(ctx, w); err != nil {
				log.Println("LoopBlaze", err)
			}
		}
	}
}

func (w *Blaze) OnAckReceipt(ctx context.Context, msg *mixin.MessageView, userID string) error {
	return nil
}

func (w *Blaze) OnMessage(ctx context.Context, msg *mixin.MessageView, userID string) error {
	if msg.Category != mixin.MessageCategoryPlainText {
		return nil
	}

	b, _ := base64.StdEncoding.DecodeString(msg.Data)
	args := strings.Fields(string(b))

	root := &cobra.Command{
		Use: "blaze",
	}

	root.SetUsageTemplate(usageMessage)
	root.AddCommand(lock.NewCmd(w.cfg))
	root.AddCommand(unlock.NewCmd(w.cfg))
	root.AddCommand(list.NewCmd(w.vats, msg.UserID))
	root.SetArgs(args)

	ww := mm.Reply(w.client, msg.UserID)
	root.SetOut(ww)

	if err := root.ExecuteContext(ctx); err != nil {
		log.Printf("execute command %q failed: %v", string(b), err)
		_ = root.Usage()
	}

	return ww.Flush(ctx)
}
