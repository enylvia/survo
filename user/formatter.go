package user

type UserFormatter struct {
	ID         int    `json:"id"`
	FullName   string `json:"fullName"`
	Occupation string `json:"occupation"`
	Email      string `json:"email"`
	Token      string `json:"token"`
}
type UserDetail struct {
	ID       int    `json:"id"`
	FullName string `json:"fullName"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Image    string `json:"image_path"`
	Phone    string `json:"phone"`
	Birthday string `json:"birthday"`
}

func FormatUser(user User, token string) UserFormatter {
	formatter := UserFormatter{
		ID:         int(user.Id),
		FullName:   user.FullName,
		Occupation: user.Occupation,
		Email:      user.Email,
		Token:      token,
	}
	return formatter
}

func FormatDetailUser(user User) UserDetail {
	formatter := UserDetail{
		ID:       int(user.Id),
		FullName: user.FullName,
		Email:    user.Email,
		Username: user.Username,
		Image:    user.Image,
		Phone:    user.Phone,
		Birthday: user.Birthday,
	}
	return formatter
}
