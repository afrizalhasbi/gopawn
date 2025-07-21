package payload

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Delete Login
type RequestReset Login
