package routes

import (
	"net/http"

	v1 "github.com/phillong97254/backoffice/internal/handler/rest/v1"
	auth "github.com/phillong97254/go-lib/authkit"
)

// /api/admin/*: quản lý role (gọi authz-service). Chặn sớm ở đây, authz-service kiểm tra lại.
//   - đọc: scope roles:read
//   - ghi: scope roles:write + role superadmin
func v1AdminRoutes(api *http.ServeMux, h v1.Handler) {
	rolesRead := auth.Policy{Scopes: []string{"roles:read"}}
	rolesWrite := auth.Policy{Scopes: []string{"roles:write"}, Roles: []string{"superadmin"}}

	roleHandler := h.GetAdminHandler().GetRoleHandler()
	api.HandleFunc("GET /api/admin/roles", auth.Require(rolesRead, roleHandler.ListRoles))
	api.HandleFunc("GET /api/admin/users/{userID}/roles", auth.Require(rolesRead, roleHandler.GetUserRoles))
	api.HandleFunc("POST /api/admin/users/{userID}/roles", auth.Require(rolesWrite, roleHandler.GrantRole))
	api.HandleFunc("DELETE /api/admin/users/{userID}/roles/{role}", auth.Require(rolesWrite, roleHandler.RevokeRole))
}
