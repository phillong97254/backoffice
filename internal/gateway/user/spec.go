package user

import "context"

type Gateway interface {
	GetProfile(ctx context.Context, input GetProfileInput) (GetProfileOutput, error)
}
