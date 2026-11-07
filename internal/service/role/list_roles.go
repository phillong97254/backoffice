package role

import (
	"context"

	"github.com/phillong97254/backoffice/internal/service/grpcerr"
)

func (i *impl) ListRoles(ctx context.Context) (ListRolesOutput, error) {
	ctx, cancel := context.WithTimeout(ctx, i.timeout)
	defer cancel()
	res, err := i.gwy.Authz().ListRoles(ctx)
	if err != nil {
		return ListRolesOutput{}, grpcerr.FromGRPC(err)
	}
	out := ListRolesOutput{Items: make([]RoleOutput, 0, len(res.Items))}
	for _, r := range res.Items {
		out.Items = append(out.Items, RoleOutput{Name: r.Name, Description: r.Description, Scopes: r.Scopes})
	}
	return out, nil
}
