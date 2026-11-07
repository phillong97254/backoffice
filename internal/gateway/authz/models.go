package authz

import authzv1 "github.com/phillong97254/backoffice/pkg/gen/authz/v1"

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

func toUserRolesOutput(r *authzv1.UserRoles) UserRolesOutput {
	out := UserRolesOutput{UserID: r.GetUserId(), Roles: r.GetRoles(), Scope: r.GetScope()}
	for _, b := range r.GetBindings() {
		out.Bindings = append(out.Bindings, BindingOutput{Role: b.GetRole(), GrantedBy: b.GetGrantedBy(), GrantedAtUnix: b.GetGrantedAtUnix()})
	}
	return out
}
