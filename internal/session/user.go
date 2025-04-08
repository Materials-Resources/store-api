package session

type UserSession struct {
	Profile     *Profile
	AccessToken string
}

type Profile struct {
	UserID    string
	Email     string
	ContactID string
	BranchID  string
}

var AnonymousUserSession = &UserSession{}

// IsAnonymous checks if user is anonymous
func (u *UserSession) IsAnonymous() bool {
	return u == AnonymousUserSession
}
