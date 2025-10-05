package status_codes

type RegisterStatusCode int

func (r RegisterStatusCode) String() string {
	return RegisterStatusCodeToString(r)
}

func (r RegisterStatusCode) Int() int {
	return int(r)
}

const (
	RegisterSuccess RegisterStatusCode = iota
	RegisterFailure
	RegisterUserAlreadyExist
	RegisterInvalidEmail
	RegisterInvalidName
	RegisterInvalidPassword
	RegisterInvalidCredentials
)

func RegisterStatusCodeToString(code RegisterStatusCode) string {
	switch code {
	case RegisterSuccess:
		return "SUCCESS"
	case RegisterFailure:
		return "FAILURE"
	case RegisterUserAlreadyExist:
		return "USER_ALREADY_EXIST"
	case RegisterInvalidEmail:
		return "INVALID_EMAIL"
	case RegisterInvalidName:
		return "INVALID_NAME"
	case RegisterInvalidPassword:
		return "INVALID_PASSWORD"
	case RegisterInvalidCredentials:
		return "INVALID_CREDENTIALS"
	default:
		return "UNKNOWN"
	}
}
