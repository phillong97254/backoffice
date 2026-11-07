package account

import accountSvc "github.com/phillong97254/backoffice/internal/service/account"

type GetMeResponse struct {
	Sub     string          `json:"sub"`
	Roles   []string        `json:"roles"`
	Scope   string          `json:"scope"`
	Profile ProfileResponse `json:"profile"`
}

type ProfileResponse struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Department string `json:"department"`
}

func toGetMeResponse(out accountSvc.GetMeOutput) GetMeResponse {
	return GetMeResponse{
		Sub:   out.Subject,
		Roles: out.Roles,
		Scope: out.Scope,
		Profile: ProfileResponse{
			ID: out.Profile.ID, Name: out.Profile.Name, Email: out.Profile.Email, Department: out.Profile.Department,
		},
	}
}
