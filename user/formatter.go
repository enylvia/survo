package user

type UserFormatter struct {
	ID         int    `json:"id"`
	FullName   string `json:"fullName"`
	Occupation string `json:"occupation"`
	Email      string `json:"email"`
	ImageURL   string `json:"image_url"`
}
func FormatUser(user User) UserFormatter {
	formatter := UserFormatter{
		ID: int(user.Id),
		FullName: user.FullName,
		Occupation: user.Occupation,
		Email: user.Email,
	}

	return formatter
}
