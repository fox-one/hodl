package core

import (
	"testing"

	"github.com/MixinNetwork/mixin/common"
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
