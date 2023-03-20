package user

type PaginationUser struct {
	Limit        int         `json:"limit"`
	Page         int         `json:"page"`
	TotalData    int         `json:"total_data"`
	TotalPage    int         `json:"total_page"`
	Users        interface{} `json:"users"`
	FirstPage    string      `json:"first_page"`
	PreviousPage string      `json:"previous_page"`
	NextPage     string      `json:"next_page"`
	LastPage     string      `json:"last_page"`
}
