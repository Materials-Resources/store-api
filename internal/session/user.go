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

var AnonymousUserSession = &UserSession{}

// IsAnonymous checks if user is anonymous
func (u *UserSession) IsAnonymous() bool {
	return u == AnonymousUserSession
}
