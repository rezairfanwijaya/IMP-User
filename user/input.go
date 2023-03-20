package user

type InputNewUser struct {
	Username string `json:"username" binding:"required,min=2"`
	Password string `json:"password" binding:"required,min=5"`
	FullName string `json:"fullname" binding:"required"`
}

type InputLoginUser struct {
	Username string `json:"username" binding:"required,min=2"`
	Password string `json:"password" binding:"required,min=5"`
}

type ParamsGetAllUsers struct {
	Page  int
	Order string
	Limit int
}
