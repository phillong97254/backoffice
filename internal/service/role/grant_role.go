package role

import (
	"context"

	authzGwy "github.com/phillong97254/backoffice/internal/gateway/authz"
	"github.com/phillong97254/backoffice/internal/service/grpcerr"
)

func (i *impl) GrantRole(ctx context.Context, input UserRoleInput) (UserRolesOutput, error) {
	ctx, cancel := context.WithTimeout(ctx, i.timeout)
	defer cancel()
	res, err := i.gwy.Authz().GrantRole(ctx, authzGwy.UserRoleInput{UserID: input.UserID, Role: input.Role})
	if err != nil {
		return UserRolesOutput{}, grpcerr.FromGRPC(err)
	}
	return toUserRolesOutput(res), nil
}
