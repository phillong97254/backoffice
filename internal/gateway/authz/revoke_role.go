package authz

import (
	"context"

	authzv1 "github.com/phillong97254/backoffice/pkg/gen/authz/v1"
)

func (i *impl) RevokeRole(ctx context.Context, input UserRoleInput) (UserRolesOutput, error) {
	resp, err := i.client.RevokeRole(ctx, &authzv1.RevokeRoleRequest{UserId: input.UserID, Role: input.Role})
	if err != nil {
		return UserRolesOutput{}, err
	}
	return toUserRolesOutput(resp), nil
}
