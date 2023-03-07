package jwt_prac2

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Info     string `json:"info"`
}

//constructor
func NewUser(u string, p string, info string) *User {
	return &User{
		Username: u,
		Password: p,
		Info:     info,
	}
}
