package user

type User struct {
	Id         int64  `json:"id"`
	FullName   string `json:"fullname"`
	Email      string `json:"email"`
	Username   string `json:"username"`
	Occupation string `json:"occupation"`
	Password   string `json:"password"`
}
