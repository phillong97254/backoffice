package auth

import "context"

type Service interface {
	Login(ctx context.Context, input LoginInput) (LoginOutput, error)
}
