package vault

import (
	"github.com/fox-one/hodl/core"
)

type vaultStore struct {
	vats map[uint64]*core.Vault
}

func New() core.VaultStore {
	return &vaultStore{
		vats: make(map[uint64]*core.Vault),
	}
}

func (s *vaultStore) Save(vault *core.Vault) error {
	s.vats[vault.ID] = vault
	return nil
}

func (s *vaultStore) Delete(id uint64) (*core.Vault, error) {
	vault := s.vats[id]
	delete(s.vats, id)
	return vault, nil
}

func (s *vaultStore) List(userID string) ([]*core.Vault, error) {
	var vats []*core.Vault
	for _, v := range s.vats {
		if v.UserID == userID {
			vats = append(vats, v)
		}
	}

	return vats, nil
}
