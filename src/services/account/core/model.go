package core

import "time"

var (
	AccountStatusActive  = "ACTIVE"
	AccountStatusPending = "PENDING"
	AccountStatusBlock   = "BLOCK"

	AccountTypeMain   = "MAIN"
	AccountTypeSaving = "SAVING"
	AccountTypeLock   = "LOCK"
)

type Account struct {
	Id        *string
	Name      string
	Type      string
	Status    string
	CreatedAt time.Time
	UpdatedAt *time.Time
}
