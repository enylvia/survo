package user

type UserFormatter struct {
	ID         int               `json:"id"`
	FullName   string            `json:"fullName"`
	Occupation string            `json:"occupation"`
	Email      string            `json:"email"`
	Token      string            `json:"token"`
	Attribute  AttributFormatter `json:"attribute"`
}
type UserFormatterLogin struct {
	ID         int    `json:"id"`
	FullName   string `json:"fullName"`
	Occupation string `json:"occupation"`
	Email      string `json:"email"`
	Token      string `json:"token"`
}
type UserFormatterRegister struct {
	ID         int    `json:"id"`
	FullName   string `json:"fullName"`
	Occupation string `json:"occupation"`
	Email      string `json:"email"`
}
type UserDetail struct {
	ID         int               `json:"id"`
	FullName   string            `json:"fullName"`
	Email      string            `json:"email"`
	Username   string            `json:"username"`
	Image      string            `json:"image_path"`
	Phone      string            `json:"phone"`
	Birthday   string            `json:"birthday"`
	Occupation string            `json:"occupation"`
	IsAdmin    string            `json:"is_admin"`
	Attribute  AttributFormatter `json:"attribute"`
}
type AttributFormatter struct {
	UserId            uint `json:"user_id"`
	PostedSurvey      int  `json:"posted_survey"`
	TotalRespondent   int  `json:"total_respondent"`
	ParticipateSurvey int  `json:"participate_survey"`
	IsPremium         bool `json:"is_premium"`
	Balance           int  `json:"balance"`
}

func FormatUser(user User, token string) UserFormatterLogin {
	formatter := UserFormatterLogin{
		ID:         int(user.Id),
		FullName:   user.FullName,
		Occupation: user.Occupation,
		Email:      user.Email,
		Token:      token,
	}
	return formatter
}
func FormatUserRegister(user User) UserFormatterRegister {
	formatter := UserFormatterRegister{
		ID:         int(user.Id),
		FullName:   user.FullName,
		Occupation: user.Occupation,
		Email:      user.Email,
	}
	return formatter
}

func FormatDetailUser(user User) UserDetail {
	formatter := UserDetail{}
	formatter.ID = int(user.Id)
	formatter.FullName = user.FullName
	formatter.Email = user.Email
	formatter.Username = user.Username
	formatter.Image = user.Image
	formatter.Phone = user.Phone
	formatter.Birthday = user.Birthday
	formatter.Occupation = user.Occupation
	formatter.IsAdmin = user.IsAdmin

	attrib := user.Attribut
	attributFormatter := AttributFormatter{}
	attributFormatter.UserId = uint(user.Id)
	attributFormatter.PostedSurvey = attrib.PostedSurvey
	attributFormatter.TotalRespondent = attrib.TotalRespondent
	attributFormatter.ParticipateSurvey = attrib.ParticipateSurvey
	attributFormatter.IsPremium = attrib.IsPremium
	attributFormatter.Balance = attrib.Balance

	formatter.Attribute = attributFormatter

	return formatter
}
