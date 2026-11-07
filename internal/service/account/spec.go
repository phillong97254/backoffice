package account

import "context"

type Service interface {
	// GetMe: claims trong token + hồ sơ lấy từ user-service.
	GetMe(ctx context.Context) (GetMeOutput, error)
}
