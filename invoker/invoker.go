package invoker

import (
	"context"
	"encoding/base64"

	"github.com/fox-one/hodl/core"
	"github.com/fox-one/mixin-sdk-go"
	"github.com/shopspring/decimal"
	"github.com/MixinNetwork/trusted-group/mvm/encoding"
	"github.com/asaskevich/govalidator"
)

type Config struct {
	Process string `validate:"uuid,required"`
	Members []string `validate:"required"`
	Threshold uint8 `validate:"required"`
}

func New(client *mixin.Client,cfg Config) core.Invoker {
	if _,err := govalidator.ValidateStruct(cfg); err != nil {
		panic(err)
	}

	return &invoker{
		client: client,
		cfg: cfg,
	}
}

type invoker struct {
	client *mixin.Client
	cfg    Config
}

func (s *invoker) Payment(ctx context.Context, assetID string, amount decimal.Decimal, extra []byte) (string, error) {
	op := &encoding.Operation{
		Purpose: encoding.OperationPurposeGroupEvent,
		Process: s.cfg.Process,
		Extra:   extra,
	}

	input := mixin.TransferInput{
		AssetID: assetID,
		Amount:  amount,
		TraceID: mixin.RandomTraceID(),
	}

	input.OpponentMultisig.Receivers = s.cfg.Members
	input.OpponentMultisig.Threshold = s.cfg.Threshold
	input.Memo = base64.RawURLEncoding.EncodeToString(op.Encode())
	payment, err := s.client.VerifyPayment(ctx, input)
	if err != nil {
		return "", err
	}

	return payment.CodeID,nil
}
