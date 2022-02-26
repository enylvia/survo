package test

import (
	"github.com/stretchr/testify/assert"
	"survorest/helper"
	"survorest/user"
	"testing"
)

func TestRegisterUserForm(t *testing.T) {
	db, err := GetConnection()
	helper.ErrorNotNil(err)

	MigrateTable(db)
	defer TruncateTable(db)

	RegisterForm := user.NewRepository(db)
	RegisterService := user.NewService(RegisterForm)

	inputData := user.RegisterInput{
		FullName:             "John Doe",
		Email:                "email@example.com",
		Username:             "example12",
		Occupation:           "Student",
		Password:             "12345678",
		PasswordConfirmation: "12345678",
	}
	structData := user.User{}
	newUserService ,err := RegisterService.RegisterUserForm(inputData)
	helper.ErrorNotNil(err)

	structData.Id = 1
	structData.FullName = inputData.FullName
	structData.Email = inputData.Email
	structData.Username = inputData.Username
	structData.Occupation = inputData.Occupation
	structData.Password = newUserService.Password
	structData.Image = ""
	structData.Phone = ""
	structData.Birthday = ""


	assert.Equal(t, newUserService,structData)
	assert.NoError(t, err)

}
