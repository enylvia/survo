package auth

type GoogleFormatter struct {
	Email     string `json:"email"`
	Token	 string `json:"token"`
}

func FormatGoogle(email string, token string) *GoogleFormatter {
	return &GoogleFormatter{
		Email: email,
		Token: token,
	}
}
