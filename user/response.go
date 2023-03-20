package user

type userResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Fullname string `json:"fullname"`
}

func FormatUser(user User) *userResponse {
	return &userResponse{
		ID:       user.ID,
		Username: user.Username,
		Fullname: user.FullName,
	}
}
