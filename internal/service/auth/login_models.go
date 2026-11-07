package auth

type LoginInput struct {
	Username string
	Password string
}

type LoginOutput struct {
	AccessToken string
	TokenType   string
	ExpiresIn   int
	Scope       string
}
