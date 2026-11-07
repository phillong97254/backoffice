package authz

import (
	"context"

	authzv1 "github.com/phillong97254/backoffice/pkg/gen/authz/v1"
)

func (i *impl) GrantRole(ctx context.Context, input UserRoleInput) (UserRolesOutput, error) {
	resp, err := i.client.GrantRole(ctx, &authzv1.GrantRoleRequest{UserId: input.UserID, Role: input.Role})
	if err != nil {
		return UserRolesOutput{}, err
	}
	return toUserRolesOutput(resp), nil
}
