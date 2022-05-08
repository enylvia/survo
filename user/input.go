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
	FullName             string `json:"fullName" form:"fullName"`
	Email                string `json:"email" form:"email"`
	Username             string `json:"username" form:"username"`
	Password             string `json:"password" form:"password"`
	PasswordConfirmation string `json:"passwordConfirmation" form:"passwordConfirmation"`
	Phone                string `json:"phone" form:"phone"`
	Birthday             string `json:"birthday" form:"birthday"`
}

type DetailUserInput struct {
	ID int `uri:"id" binding:"required"`
}

type UserImageInput struct {
	UserId int `form:"user_id" binding:"required"`
}

type UpdateAttributInput struct {

}