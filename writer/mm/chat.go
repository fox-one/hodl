package mm

import (
	"context"
	"encoding/base64"
	"encoding/json"

	"github.com/fox-one/hodl/writer"
	"github.com/fox-one/hodl/writer/mm/color"
	"github.com/fox-one/mixin-sdk-go"
)

type Replier struct {
	client      *mixin.Client
	RecipientID string
	messages    []*mixin.MessageRequest
}

func Reply(client *mixin.Client, recipientID string) *Replier {
	return &Replier{
		client:      client,
		RecipientID: recipientID,
	}
}

func (r *Replier) Write(p []byte) (n int, err error) {
	category := mixin.MessageCategoryPlainText

	if raw, label, ok := writer.ExtractLabel(string(p)); ok {
		buttons := mixin.AppButtonGroupMessage{
			{
				Label:  label,
				Action: raw,
				Color:  color.Random(),
			},
		}

		p, _ = json.Marshal(buttons)
		category = mixin.MessageCategoryAppButtonGroup
	}

	msg := &mixin.MessageRequest{
		ConversationID: mixin.UniqueConversationID(r.client.ClientID, r.RecipientID),
		MessageID:      mixin.RandomTraceID(),
		RecipientID:    r.RecipientID,
		Category:       category,
		Data:           base64.StdEncoding.EncodeToString(p),
	}

	r.messages = append(r.messages, msg)
	return len(p), nil
}

func (r *Replier) Flush(ctx context.Context) error {
	if len(r.messages) == 0 {
		return nil
	}

	return r.client.SendMessages(ctx, r.messages)
}
