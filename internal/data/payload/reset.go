package payload

type RequestReset struct {
	Email    string
	Password string
}

type ResetPassword struct {
	Email           string
	NewPassword     string
	NewPasswordCopy string
}
