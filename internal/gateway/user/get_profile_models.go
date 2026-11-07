package user

type GetProfileInput struct {
	ID string
}

// GetProfileOutput theo định dạng JSON của user-service.
type GetProfileOutput struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Department string `json:"department"`
}
