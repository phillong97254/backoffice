package role

import roleSvc "github.com/phillong97254/backoffice/internal/service/role"

type ListRolesResponse struct {
	Items []RoleResponse `json:"items"`
}

type RoleResponse struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Scopes      []string `json:"scopes"`
}

func toListRolesResponse(out roleSvc.ListRolesOutput) ListRolesResponse {
	resp := ListRolesResponse{Items: make([]RoleResponse, 0, len(out.Items))}
	for _, r := range out.Items {
		resp.Items = append(resp.Items, RoleResponse{Name: r.Name, Description: r.Description, Scopes: r.Scopes})
	}
	return resp
}
