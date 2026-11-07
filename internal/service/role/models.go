package role

import authzGwy "github.com/phillong97254/backoffice/internal/gateway/authz"

type UserRoleInput struct {
	UserID string
	Role   string // bỏ trống với GetUserRoles
}

type UserRolesOutput struct {
	UserID   string
	Roles    []string
	Scope    string
	Bindings []BindingOutput
}

type BindingOutput struct {
	Role          string
	GrantedBy     string
	GrantedAtUnix int64
}

type ListRolesOutput struct {
	Items []RoleOutput
}

type RoleOutput struct {
	Name        string
	Description string
	Scopes      []string
}

func toUserRolesOutput(r authzGwy.UserRolesOutput) UserRolesOutput {
	out := UserRolesOutput{UserID: r.UserID, Roles: r.Roles, Scope: r.Scope, Bindings: make([]BindingOutput, 0, len(r.Bindings))}
	for _, b := range r.Bindings {
		out.Bindings = append(out.Bindings, BindingOutput{Role: b.Role, GrantedBy: b.GrantedBy, GrantedAtUnix: b.GrantedAtUnix})
	}
	return out
}
