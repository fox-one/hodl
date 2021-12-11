package core

import (
	"github.com/MixinNetwork/mixin/common"
)

const (
	ActionLock   = 1
	ActionUnlock = 2
)

type Event struct {
	Action  int
	VaultID uint64
	Exp     uint64
}

func (e *Event) Encode() []byte {
	enc := common.NewEncoder()
	enc.WriteInt(e.Action)
	enc.WriteUint64(e.VaultID)
	enc.WriteUint64(e.Exp)
	return enc.Bytes()
}
