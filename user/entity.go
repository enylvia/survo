package user

type User struct {
	Id         int64  `json:"id" gorm:"primaryKey; not null"`
	FullName   string `json:"fullName" gorm:"type:varchar(100); not null"`
	Email      string `json:"email" gorm:"type:varchar(100); not null"`
	Username   string `json:"username"  gorm:"type:varchar(100); not null"`
	Occupation string `json:"occupation" gorm:"type:varchar(100); not null"`
	Role       string `json:"role" gorm:"type:varchar(100); not null"`
	Password   string `json:"password" gorm:"type:varchar(100); not null"`
	Image      string `json:"image" gorm:"type:varchar(100); not null"`
	Phone      string `json:"phone" gorm:"type:varchar(100); not null"`
	Birthday   string `json:"birthday" gorm:"type:varchar(100); not null"`
}
