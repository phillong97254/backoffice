package role

import roleSvc "github.com/phillong97254/backoffice/internal/service/role"

type UserRolesResponse struct {
	UserID   string            `json:"user_id"`
	Roles    []string          `json:"roles"`
	Scope    string            `json:"scope"`
	Bindings []BindingResponse `json:"bindings"`
}

type BindingResponse struct {
	Role          string `json:"role"`
	GrantedBy     string `json:"granted_by"`
	GrantedAtUnix int64  `json:"granted_at_unix"`
}

func toUserRolesResponse(out roleSvc.UserRolesOutput) UserRolesResponse {
	resp := UserRolesResponse{UserID: out.UserID, Roles: out.Roles, Scope: out.Scope, Bindings: make([]BindingResponse, 0, len(out.Bindings))}
	if resp.Roles == nil {
		resp.Roles = []string{}
	}
	for _, b := range out.Bindings {
		resp.Bindings = append(resp.Bindings, BindingResponse{Role: b.Role, GrantedBy: b.GrantedBy, GrantedAtUnix: b.GrantedAtUnix})
	}
	return resp
}
