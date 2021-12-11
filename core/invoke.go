package core

import (
	"context"

	"github.com/shopspring/decimal"
)

type Invoker interface {
	// Payment generate a mixin payment code
	Payment(ctx context.Context, assetID string,amount decimal.Decimal,extra []byte) (string, error)
}


