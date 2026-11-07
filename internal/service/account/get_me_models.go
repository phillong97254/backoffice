package account

type GetMeOutput struct {
	Subject string
	Roles   []string
	Scope   string
	Profile ProfileOutput
}

type ProfileOutput struct {
	ID         string
	Name       string
	Email      string
	Department string
}
