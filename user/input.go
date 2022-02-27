package user

type RegisterInput struct {
	FullName             string `json:"fullName" binding:"required"`
	Email                string `json:"email" binding:"required"`
	Username             string `json:"username" binding:"required"`
	Occupation           string `json:"occupation" binding:"required"`
	Password             string `json:"password" binding:"required,min=8,max=32"`
	PasswordConfirmation string `json:"passwordConfirmation" binding:"required,eqfield=Password"`
}

type LoginInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UpdateInput struct {
	FullName string `json:"fullName" binding:"required"`
	Email	string `json:"email" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password             string `json:"password" binding:"required"`
	PasswordConfirmation string `json:"passwordConfirmation" binding:"required"`
	Image string `json:"image" binding:"required"`
	Phone string `json:"phone" binding:"required"`
	Birthday string `json:"birthday" binding:"required"`
}

type DetailUserInput struct {
	ID int `uri:"id" binding:"required"`
}
