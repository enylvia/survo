package user

type User struct {
	Id         int64  `json:"id"`
	FullName   string `json:"fullName"`
	Email      string `json:"email"`
	Username   string `json:"username"`
	Occupation string `json:"occupation"`
	Role string `json:"role"`
	Password   string `json:"password"`
	Image      string `json:"image"`
	Phone string `json:"phone"`
	Birthday string `json:"birthday"`
}
