package core

import (
	"github.com/MixinNetwork/mixin/common"
	"github.com/gofrs/uuid"
)

type Group struct {
	Members   []string `json:"members,omitempty"`
	Threshold uint8    `json:"threshold,omitempty"`
}

func (g *Group) Encode() []byte {
	enc := common.NewEncoder()

	if len(g.Members) > 64 {
		panic(len(g.Members))
	}

	enc.WriteInt(len(g.Members))
	for _, m := range g.Members {
		enc.Write(uuid.FromStringOrNil(m).Bytes())
	}

	enc.WriteInt(int(g.Threshold))
	return enc.Bytes()
}

func DecodeGroup(b []byte) (*Group, error) {
	dec := common.NewDecoder(b)

	n, err := dec.ReadInt()
	if err != nil {
		return nil, err
	}

	members := make([]string, n)
	for i := 0; i < n; i++ {
		var b [16]byte
		if err := dec.Read(b[:]); err != nil {
			return nil, err
		}

		members[i] = uuid.FromBytesOrNil(b[:]).String()
	}

	threshold, err := dec.ReadInt()
	if err != nil {
		return nil, err
	}

	return &Group{
		Members:   members,
		Threshold: uint8(threshold),
	}, nil
}
