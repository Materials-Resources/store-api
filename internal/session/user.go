package session

type UserSession struct {
	Profile     *Profile
	AccessToken string
}

type Profile struct {
	ID        string
	Email     string
	ContactID string
	BranchID  string
}
