package core

import (
	"math/big"
	"testing"

	"github.com/MixinNetwork/mixin/common"
	"github.com/shopspring/decimal"
)

func TestEvent_Encode(t *testing.T) {
	i := common.NewIntegerFromString("1000000.12345678")
	enc := common.NewEncoder()
	enc.WriteInteger(i)
	t.Log(len(enc.Bytes()))

	b := enc.Bytes()

	d := common.NewDecoder(b)
	t.Log(d.ReadInt())
}

func TestDecode(t *testing.T) {
	i := big.NewInt(100)
	d := decimal.NewFromBigInt(i, 2)
	t.Log(d.String())
}
