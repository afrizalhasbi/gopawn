package payload

type Register struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Login Register
type ForgotPassword Register
type Delete Register
