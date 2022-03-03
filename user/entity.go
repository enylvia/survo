package user

type User struct {
	Id         int64  `json:"id" gorm:"primaryKey; not null"`
	FullName   string `json:"fullName" gorm:"type:varchar(100); not null"`
	Email      string `json:"email" gorm:"type:varchar(100); not null"`
	Username   string `json:"username"  gorm:"type:varchar(100); not null"`
	Occupation string `json:"occupation" gorm:"type:varchar(100); not null"`
	Password   string `json:"password" gorm:"type:varchar(100); not null"`
	Token      string `json:"token" gorm:"type:varchar(244); nullable"`
	Image      string `json:"image" gorm:"type:varchar(100); nullable"`
	Phone      string `json:"phone" gorm:"type:varchar(100); nullable"`
	Birthday   string `json:"birthday" gorm:"type:varchar(100); nullable"`
}
