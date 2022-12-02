package core

import "context"

type IAccountMessaging interface {
	AccountCreatedEvent(ctx context.Context, account Account)
}
