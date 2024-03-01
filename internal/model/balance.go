package model

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Balance struct {
	ID     uuid.UUID
	Amount decimal.Decimal
}
