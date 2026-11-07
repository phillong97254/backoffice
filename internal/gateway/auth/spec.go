package auth

import "context"

type Gateway interface {
	// PasswordToken đổi username/password lấy access token (grant_type=password).
	PasswordToken(ctx context.Context, input PasswordTokenInput) (PasswordTokenOutput, error)
}
