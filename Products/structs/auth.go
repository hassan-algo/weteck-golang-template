package structs

type Auth struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Authenticate struct {
	Token string `json:"token"`
}
