package core

import (
	"time"

	"github.com/shopspring/decimal"
)

type (
	Vault struct {
		ID     uint64
		UserID string
		Symbol string
		Amount decimal.Decimal
		End    time.Time
	}

	VaultStore interface {
		Save(vault *Vault) error
		Delete(id uint64) (*Vault, error)
		List(userID string) ([]*Vault, error)
	}
)
