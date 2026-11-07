package role

import "context"

// Service quản lý role qua authz-service. Quyền (roles:read / roles:write + superadmin)
// được authz-service kiểm tra lại, backoffice chỉ chặn sớm ở routes.
type Service interface {
	ListRoles(ctx context.Context) (ListRolesOutput, error)
	GetUserRoles(ctx context.Context, input UserRoleInput) (UserRolesOutput, error)
	GrantRole(ctx context.Context, input UserRoleInput) (UserRolesOutput, error)
	RevokeRole(ctx context.Context, input UserRoleInput) (UserRolesOutput, error)
}
