package role

import "net/http"

type Handler interface {
	// GET /api/admin/roles
	ListRoles(w http.ResponseWriter, r *http.Request)
	// GET /api/admin/users/{userID}/roles
	GetUserRoles(w http.ResponseWriter, r *http.Request)
	// POST /api/admin/users/{userID}/roles {"role": "superadmin"}
	GrantRole(w http.ResponseWriter, r *http.Request)
	// DELETE /api/admin/users/{userID}/roles/{role}
	RevokeRole(w http.ResponseWriter, r *http.Request)
}
