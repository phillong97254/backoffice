package authz

import (
	"context"

	authzv1 "github.com/phillong97254/backoffice/pkg/gen/authz/v1"
)

func (i *impl) GetUserRoles(ctx context.Context, input UserRoleInput) (UserRolesOutput, error) {
	resp, err := i.client.GetUserRoles(ctx, &authzv1.GetUserRolesRequest{UserId: input.UserID})
	if err != nil {
		return UserRolesOutput{}, err
	}
	return toUserRolesOutput(resp), nil
}
