package authz

import "context"

// Gateway gọi authz-service qua gRPC, forward token của admin đang đăng nhập.
// Lỗi trả về là gRPC status gốc (không wrap) để tầng service đọc được status code.
type Gateway interface {
	ListRoles(ctx context.Context) (ListRolesOutput, error)
	GetUserRoles(ctx context.Context, input UserRoleInput) (UserRolesOutput, error)
	GrantRole(ctx context.Context, input UserRoleInput) (UserRolesOutput, error)
	RevokeRole(ctx context.Context, input UserRoleInput) (UserRolesOutput, error)
	Close() error
}
