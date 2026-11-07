package authz

import (
	"context"

	authzv1 "github.com/phillong97254/backoffice/pkg/gen/authz/v1"
)

func (i *impl) ListRoles(ctx context.Context) (ListRolesOutput, error) {
	resp, err := i.client.ListRoles(ctx, &authzv1.ListRolesRequest{})
	if err != nil {
		return ListRolesOutput{}, err
	}
	out := ListRolesOutput{Items: make([]RoleOutput, 0, len(resp.GetItems()))}
	for _, r := range resp.GetItems() {
		out.Items = append(out.Items, RoleOutput{Name: r.GetName(), Description: r.GetDescription(), Scopes: r.GetScopes()})
	}
	return out, nil
}
