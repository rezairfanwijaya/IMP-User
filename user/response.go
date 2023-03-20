package user

type userResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Fullname string `json:"fullname"`
}

type userReponseLogin struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Fullname string `json:"fullname"`
	Token    string `json:"token"`
}

func FormatUser(user User) *userResponse {
	return &userResponse{
		ID:       user.ID,
		Username: user.Username,
		Fullname: user.FullName,
	}
}

func FormatUserLogin(user User, token string) *userReponseLogin {
	return &userReponseLogin{
		ID:       user.ID,
		Username: user.Username,
		Fullname: user.FullName,
		Token:    token,
	}
}
