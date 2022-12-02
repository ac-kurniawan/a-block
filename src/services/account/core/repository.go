package core

import "context"

type IAccountRepository interface {
	CreateAccount(ctx context.Context, account Account) (*Account, error)
	GetAccountById(ctx context.Context, id string) (*Account, error)
}
