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
	// Exp in nano seconds
	Exp uint64
}

func (e *Event) Encode() []byte {
	enc := common.NewEncoder()
	enc.WriteInt(e.Action)
	enc.WriteUint64(e.VaultID)
	enc.WriteUint64(e.Exp)
	return enc.Bytes()
}

func DecodeEvent(b []byte) (*Event, error) {
	var (
		e   Event
		err error
	)

	dec := common.NewDecoder(b)
	e.Action, err = dec.ReadInt()
	if err != nil {
		return nil, err
	}

	e.VaultID, err = dec.ReadUint64()
	if err != nil {
		return nil, err
	}

	e.Exp, err = dec.ReadUint64()
	if err != nil {
		return nil, err
	}

	return &e, nil
}
