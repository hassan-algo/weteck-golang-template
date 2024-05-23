package structs

type Auth struct {
	Email      string `json:"email"`
	Password   string `json:"password"`
	UserGuid   string `json:"user_guid"`
	LoginToken string `json:"login_token"`
}

type Authenticate struct {
	Token string `json:"token"`
}
type Response struct {
	Valid   bool        `json:"valid"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
type ResponseUserWithToken struct {
	// UserGuid   string `json:"userguid"`
	Name       string `json:"name"`
	ProfilePic string `json:"profilepic"`
	Email      string `json:"email"`
	Token      string `json:"token"`
}

type Credentials struct {
	UserGuid    string `json:"userguid"`
	FullName    string `json:"fullname"`
	ProfilePic  string `json:"profilepic"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	UserType    string `json:"usertype"`
	Login_Token string `json:"login_token"`
}
