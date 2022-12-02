package core

import "context"

type IAccountService interface {
	CreateAccount(ctx context.Context, account Account) (*Account, error)
	GetAccountById(ctx context.Context, id string) (*Account, error)
}

type AccountService struct {
	AccountRepository IAccountRepository
	AccountMessaging  IAccountMessaging
}

func (a AccountService) CreateAccount(ctx context.Context, account Account) (*Account, error) {
	return nil, nil
}

func (a AccountService) GetAccountById(ctx context.Context, id string) (*Account, error) {
	return nil, nil
}

func NewAccountService(module AccountService) IAccountService {
	return &module
}
