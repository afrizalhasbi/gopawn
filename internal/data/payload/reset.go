package payload

type ResetPassword struct {
	Email           string
	NewPassword     string
	NewPasswordCopy string
}
