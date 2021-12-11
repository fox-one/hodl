package events

import (
	"context"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/fox-one/hodl/api"
	"github.com/fox-one/hodl/core"
	"github.com/fox-one/mixin-sdk-go"
	"github.com/gofrs/uuid"
	"github.com/shopspring/decimal"
)

type Events struct {
	client *mixin.Client
	hodl   *api.Hodl
	groups map[string]*core.Group

	offset uint64
}

func New(client *mixin.Client, conn *ethclient.Client, address string) *Events {
	hodl, err := api.NewHodl(common.HexToAddress(address), conn)
	if err != nil {
		log.Fatal(err)
	}

	return &Events{
		client: client,
		hodl:   hodl,
		groups: make(map[string]*core.Group),
	}
}

func (w *Events) Run(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(time.Second):
			_ = w.run(ctx)
		}
	}
}

func (w *Events) addressToGroup(ctx context.Context, address common.Address) (*core.Group, error) {
	if group, ok := w.groups[address.String()]; ok {
		return group, nil
	}

	opts := &bind.CallOpts{
		Context: ctx,
	}

	b, err := w.hodl.Members(opts, address)
	if err != nil {
		return nil, err
	}

	g, err := core.DecodeGroup(b)
	if err != nil {
		return nil, err
	}

	w.groups[address.String()] = g
	return g, nil
}

func (w *Events) run(ctx context.Context) error {
	opts := &bind.FilterOpts{
		Start:   w.offset,
		Context: ctx,
	}

	iter, err := w.hodl.FilterMixinEvent(opts, nil)
	if err != nil {
		log.Println("filter mixin event error:", err)
		return err
	}

	defer iter.Close()

	for iter.Next() {
		if err := w.handleMixinEvent(ctx, iter.Event); err != nil {
			return err
		}

		w.offset = iter.Event.Raw.BlockNumber + 1
	}

	return iter.Error()
}

func (w *Events) handleMixinEvent(ctx context.Context, e *api.HodlMixinEvent) error {
	nonce := e.Nonce.Uint64()
	assetID := uuid.FromBytesOrNil(e.Asset.Bytes()).String()
	amount := decimal.NewFromBigInt(e.Amount, -8)

	group, err := w.addressToGroup(ctx, e.Sender)
	if err != nil {
		log.Println("addressToGroup", err)
		return err
	}

	if len(group.Members) > 1 {
		return nil
	}

	userID := group.Members[0]
	asset, err := mixin.ReadNetworkAsset(ctx, assetID)
	if err != nil {
		log.Println("ReadNetworkAsset", assetID, err)
		return err
	}

	body, err := core.DecodeEvent(e.Extra)
	if err != nil {
		log.Println("DecodeExtra", hex.EncodeToString(e.Extra), err)
		return err
	}

	log.Println("handle Mixin Event", e.Raw.BlockNumber, body.Action, body.VaultID, userID)

	var msg string

	switch body.Action {
	case core.ActionLock:
		id := nonce + 1
		msg = fmt.Sprintf("Vault #%d locked with %s %s", id, amount, asset.Symbol)
	case core.ActionUnlock:
		msg = fmt.Sprintf("Vault #%d unlocked", body.VaultID)
	default:
		return nil
	}

	conversationID := mixin.UniqueConversationID(w.client.ClientID, userID)
	messageID := uuid.NewV5(uuid.FromStringOrNil(conversationID), e.Nonce.String()).String()
	log.Println("send message", messageID, userID, msg)
	if err := w.client.SendMessage(ctx, &mixin.MessageRequest{
		ConversationID: conversationID,
		RecipientID:    userID,
		MessageID:      messageID,
		Category:       mixin.MessageCategoryPlainText,
		Data:           base64.StdEncoding.EncodeToString([]byte(msg)),
	}); err != nil {
		log.Println("SendMessage", err)
		return err
	}

	return nil
}
