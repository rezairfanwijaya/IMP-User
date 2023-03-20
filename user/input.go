package user

type InputNewUser struct {
	Username string `json:"username" binding:"required,min=2"`
	Password string `json:"password" binding:"required,min=5"`
	FullName string `json:"fullname" binding:"required"`
}
