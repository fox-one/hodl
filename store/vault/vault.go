package vault

import (
	"sort"
	"sync"

	"github.com/fox-one/hodl/core"
)

type vaultStore struct {
	vats map[uint64]*core.Vault
	mux  sync.Mutex
}

func New() core.VaultStore {
	return &vaultStore{
		vats: make(map[uint64]*core.Vault),
	}
}

func (s *vaultStore) Save(vault *core.Vault) error {
	s.mux.Lock()
	defer s.mux.Unlock()

	s.vats[vault.ID] = vault
	return nil
}

func (s *vaultStore) Delete(id uint64) (*core.Vault, error) {
	s.mux.Lock()
	defer s.mux.Unlock()

	vault := s.vats[id]
	delete(s.vats, id)
	return vault, nil
}

func (s *vaultStore) List(userID string) ([]*core.Vault, error) {
	s.mux.Lock()
	defer s.mux.Unlock()

	var vats []*core.Vault
	for _, v := range s.vats {
		if v.UserID == userID {
			vats = append(vats, v)
		}
	}

	sort.Slice(vats, func(i, j int) bool {
		return vats[i].ID < vats[j].ID
	})
	return vats, nil
}
