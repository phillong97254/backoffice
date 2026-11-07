package admin

import "github.com/phillong97254/backoffice/internal/handler/rest/v1/admin/role"

// Handler cho các endpoint quản trị (cần token có scope roles:*, ghi cần superadmin).
type Handler interface {
	GetRoleHandler() role.Handler
}
