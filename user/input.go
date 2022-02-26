package user

type RegisterInput struct {
	FullName             string `json:"fullName" validate:"required"`
	Email                string `json:"email" validate:"required"`
	Username             string `json:"username" validate:"required"`
	Occupation           string `json:"occupation" validate:"required"`
	Password             string `json:"password" validate:"required,min=8,max=32"`
	PasswordConfirmation string `json:"passwordConfirmation" validate:"required,eqfield=Password"`
}

type LoginInput struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UpdateInput struct {
	FullName string `json:"fullName"`
	Email	string `json:"email"`
	Username string `json:"username"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"passwordConfirmation"`
	Phone string `json:"phone"`
	Birthday string `json:"birthday"`
}
