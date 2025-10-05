package status_codes

type LoginStatusCode int

func (l LoginStatusCode) String() string {
	return LoginStatusCodeToString(l)
}

func (l LoginStatusCode) Int() int {
	return int(l)
}

const (
	LoginSuccess LoginStatusCode = iota
	LoginFailure
	LoginInvalidCredentials
	LoginUserNotFound
)

func LoginStatusCodeToString(code LoginStatusCode) string {
	switch code {
	case LoginSuccess:
		return "SUCCESS"
	case LoginFailure:
		return "FAILURE"
	case LoginInvalidCredentials:
		return "INVALID_CREDENTIALS"
	case LoginUserNotFound:
		return "USER_NOT_FOUND"
	default:
		return "UNKNOWN_ERROR"
	}
}
