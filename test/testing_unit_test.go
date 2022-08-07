package test

import (
	"survorest/survey"
	"survorest/transactions"
	"survorest/user"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func InitializeUser() user.Service {
	db, _ := GetConnection()
	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)

	return userService
}
func InitializeSurvey() survey.Service {
	db, _ := GetConnection()
	userRepository := user.NewRepository(db)
	surveyRepository := survey.NewRepository(db)
	surveyService := survey.NewService(surveyRepository, userRepository)

	return surveyService
}
func InitializeTrx() transactions.Service {
	db, _ := GetConnection()
	trxRepository := transactions.NewRepository(db)
	trxService := transactions.NewService(trxRepository)

	return trxService
}

/*
Unit Testing for Package User
*/
func TestServiceRegisterUserFailed(t *testing.T) {
	service := InitializeUser()

	var inputUser = user.RegisterInput{
		FullName:             "Mocking Test",
		Email:                "",
		Password:             "mocking",
		PasswordConfirmation: "mocking",
		Username:             "mocking Test",
		Occupation:           "mocking Test",
	}

	_, err := service.RegisterUserForm(inputUser)
	assert.EqualError(t, err, "email is required", "Error Message should be email is required")

}
func TestServiceRegisterUserSuccess(t *testing.T) {
	//TODO implement me
	service := InitializeUser()
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte("mocking"), bcrypt.MinCost)
	var inputUser = user.RegisterInput{
		FullName:             "Mocking Test",
		Email:                "mockingtest@test.mail",
		Password:             string(hashPassword),
		PasswordConfirmation: string(hashPassword),
		Username:             "mocking Test",
		Occupation:           "mocking Test",
	}

	createUser, err := service.RegisterUserForm(inputUser)

	assert.NoError(t, err, "Error Message should be nil")
	assert.Equal(t, inputUser.FullName, createUser.FullName, "Error Message should be nil")
	assert.Equal(t, inputUser.Email, createUser.Email, "Error Message should be nil")
}
func TestServiceLoginUserFailed(t *testing.T) {
	service := InitializeUser()

	var inputLogin = user.LoginInput{
		Email:    "mockingTest@mail.com",
		Password: "secrettest",
	}

	_, err := service.LoginUserForm(inputLogin)
	assert.ErrorContains(t, err, "not found", "Error Message should be not found")
}
func TestServiceLoginUserSuccess(t *testing.T) {
	service := InitializeUser()

	var inputLogin = user.LoginInput{
		Email:    "john@mail.com",
		Password: "12345678",
	}

	login, err := service.LoginUserForm(inputLogin)

	assert.NoError(t, err, "Error Message should be nil")
	assert.Equal(t, inputLogin.Email, login.Email, "Error Message should be nil")

}
func TestServiceUpdateUserFailed(t *testing.T) {
	service := InitializeUser()
	var inputUser = user.UpdateInput{
		FullName:             "Mocking Test Update",
		Email:                "mockingemail@update.test",
		Password:             "mocking",
		PasswordConfirmation: "mockings",
		Username:             "mockingupdate",
		Phone:                "0812345678",
		Birthday:             "2001-01-01",
	}
	var user = user.DetailUserInput{
		ID: 4,
	}
	_, err := service.UpdateUserForm(user, inputUser)
	assert.ErrorContains(t, err, "password", "Error Message should be password not match")
}
func TestServiceUpdateUserSuccess(t *testing.T) {
	service := InitializeUser()
	var inputUser = user.UpdateInput{
		FullName:             "Mocking Test Update",
		Email:                "mockingemail@update.test",
		Password:             "mocking",
		PasswordConfirmation: "mocking",
		Username:             "mockingupdate",
		Phone:                "0812345678",
		Birthday:             "2001-01-01",
	}
	var user = user.DetailUserInput{
		ID: 8,
	}
	updateProfil, err := service.UpdateUserForm(user, inputUser)
	assert.NoError(t, err, "Error Message should be nil")
	assert.Equal(t, inputUser.FullName, updateProfil.FullName, "Error Message should be nil")
	assert.Equal(t, inputUser.Email, updateProfil.Email, "Error Message should be nil")
}
func TestServiceGetUserByEmailFailed(t *testing.T) {
	service := InitializeUser()
	email := ""

	_, err := service.GetUserByEmail(email)
	assert.ErrorContains(t, err, "email", "Error Message should be email not found")

}
func TestServiceGetUserByEmailSuccess(t *testing.T) {
	service := InitializeUser()
	email := "john@mail.com"

	user, err := service.GetUserByEmail(email)
	assert.NoError(t, err, "Error Message should be nil")
	assert.Equal(t, email, user.Email, "Error Message should be nil")
}
func TestServiceGetUserByIDFailed(t *testing.T) {
	service := InitializeUser()
	user := 0

	_, err := service.GetUserByID(user)
	assert.ErrorContains(t, err, "not found", "Error Message should be not found")
}
func TestServiceGetUserByIDSuccess(t *testing.T) {
	service := InitializeUser()
	user := 1

	getUser, err := service.GetUserByID(user)
	assert.NoError(t, err, "Error Message should be nil")
	assert.Equal(t, user, int(getUser.Id), "Error Message should be nil")
}
func TestGetAllUser(t *testing.T) {
	service := InitializeUser()
	data, err := service.GetAllUser()
	assert.NoError(t, err, "Error Message should be nil")
	assert.Greater(t, len(data), 0, "Error Message should be nil")
}
func TestServiceDeleteUserFailed(t *testing.T) {
	service := InitializeUser()

	userID := 0
	err := service.DeleteUser(userID)
	assert.ErrorContains(t, err, "not found", "Error Message should be not found")
}
func TestServiceDeleteUserSuccess(t *testing.T) {
	service := InitializeUser()

	userID := 2

	err := service.DeleteUser(userID)
	assert.NoError(t, err, "Error Message should be nil")
}
func TestCreateSurvey(t *testing.T) {
	service := InitializeSurvey()
	var inputSurvey = survey.CreateSurveyInput{
		UserId:            1,
		SurveyTitle:       "Mocking Test Survey",
		SurveyDescription: "Mocking Test Survey Description",
		SurveyCategory:    "Mocking Test Survey Category",
		Target:            10,
		Question: []survey.QuestionInput{
			{
				UserId:         1,
				SurveyQuestion: "apakah anda pernah mengikuti kursus?",
				QuestionType:   "Checkbox",
				OptionName:     "Option 1,Option 2,Option 3",
			},
			{
				UserId:         1,
				SurveyQuestion: "dimana anda mengikuti kursus tersebut?",
				QuestionType:   "Checkbox",
				OptionName:     "Option 1,Option 2,Option 3",
			},
			{
				UserId:         1,
				SurveyQuestion: "apakah anda ingin mengikuti kursus?",
				QuestionType:   "Checkbox",
				OptionName:     "Option 1,Option 2,Option 3",
			},
		},
	}
	createSurvey, err := service.CreateSurveyForm(inputSurvey)
	assert.NoError(t, err, "Error Message should be nil")
	assert.Equal(t, inputSurvey.SurveyTitle, createSurvey.Title, "Error Message should be nil")
}
func TestGetSurveyDetailFailed(t *testing.T) {
	service := InitializeSurvey()
	surveyID := 0
	_, err := service.GetSurveyDetail(surveyID)
	assert.ErrorContains(t, err, "Not Found", "Error Message should be not found")
}
func TestGetSurveyDetailSuccess(t *testing.T) {
	service := InitializeSurvey()
	surveyID := 1
	survey, err := service.GetSurveyDetail(surveyID)
	assert.NoError(t, err, "Error Message should be nil")
	assert.Equal(t, surveyID, int(survey.Id), "Error Message should be nil")
}
func TestGetSurveyListSuccess(t *testing.T) {
	service := InitializeSurvey()
	data, err := service.GetSurveyList(0)
	assert.NoError(t, err, "Error Message should be nil")
	assert.Greater(t, len(data), 0, "Error Message should be nil")
}
func TestServiceAnswerQuestionSuccess(t *testing.T) {
	service := InitializeSurvey()
	var inputAnswer = []survey.AnswerInput{
		{
			Id:         1,
			SurveyId:   1,
			UserId:     1,
			QuestionId: 1,
			Respond:    "Option 1",
		},
		{
			Id:         1,
			SurveyId:   1,
			UserId:     1,
			QuestionId: 2,
			Respond:    "Option 2",
		},
		{
			Id:         1,
			SurveyId:   1,
			UserId:     1,
			QuestionId: 3,
			Respond:    "Option 3",
		},
	}
	answer, err := service.AnswerQuestion(inputAnswer)
	assert.NoError(t, err, "Error Message should be nil")
	assert.Equal(t, inputAnswer[2].Respond, answer.Respond, "Error Message should be nil")
}
func TestGetRespondSurveyFailed(t *testing.T) {
	service := InitializeSurvey()
	surveyID := 0
	_, err := service.GetRespondSurvey(surveyID)
	assert.ErrorContains(t, err, "not found", "Error Message should be not found")
}
func TestGetRespondSurveySuccess(t *testing.T) {
	service := InitializeSurvey()
	surveyID := 1
	respond, err := service.GetRespondSurvey(surveyID)
	assert.NoError(t, err, "Error Message should be nil")
	assert.Equal(t, surveyID, int(respond[0].SurveyId), "Error Message should be nil")
}
func TestGetAllTransactionSuccess(t *testing.T) {
	service := InitializeTrx()
	data, err := service.GetAllTransaction()

	assert.GreaterOrEqual(t, len(data), 0, "Error Message should be nil")
	assert.NoError(t, err, "Error Message should be nil")
}
func TestGetTransactionByIDUserFailed(t *testing.T) {
	service := InitializeTrx()
	userID := transactions.GetTransactionUserInput{
		ID: 0,
	}
	_, err := service.GetDataTransactionByIDUser(userID)
	assert.ErrorContains(t, err, "invalid", "Error Message should be invalid")
}
func TestGetTransactionByIDUserSuccess(t *testing.T) {
	service := InitializeTrx()
	userID := transactions.GetTransactionUserInput{
		ID: 1,
	}
	data, err := service.GetDataTransactionByIDUser(userID)
	assert.NoError(t, err, "Error Message should be nil")
	assert.Greater(t, len(data), 0, "Error Message should be nil")
}
func TestCreateTransactionWithdrawFailed(t *testing.T) {
	service := InitializeTrx()
	var inputTransaction = transactions.CreateTransactionInput{
		UserID: 0,
		Amount: 10000,
	}
	_, err := service.CreateTransactionWithdraw(inputTransaction)
	assert.ErrorContains(t, err, "invalid", "Error Message should be invalid")
}
func TestCreateTransactionWithdrawSuccess(t *testing.T) {
	service := InitializeTrx()
	var inputTransaction = transactions.CreateTransactionInput{
		UserID: 1,
		Amount: 10000,
	}
	data, err := service.CreateTransactionWithdraw(inputTransaction)
	assert.NoError(t, err, "Error Message should be nil")
	assert.Equal(t, inputTransaction.Amount, data.Amount, "Error Message should be nil")
}
func TestCreatePremiumTransactionFailed(t *testing.T) {
	service := InitializeTrx()
	var inputTransaction = transactions.CreateTransactionPremium{
		ID: 0,
	}
	_, err := service.CreateTransactionPremium(inputTransaction)
	assert.ErrorContains(t, err, "invalid", "Error Message should be invalid")
}
func TestCreatePremiumTransactionSuccess(t *testing.T) {
	service := InitializeTrx()
	var inputTransaction = transactions.CreateTransactionPremium{
		ID: 1,
	}
	data, err := service.CreateTransactionPremium(inputTransaction)
	assert.NoError(t, err, "Error Message should be nil")
	assert.Equal(t, inputTransaction.ID, data.UserId, "Error Message should be nil")
}
func TestConfirmationTransactionSuccess(t *testing.T) {
	service := InitializeTrx()
	userID := 1
	data, err := service.ConfirmationTransaction(userID)
	assert.NoError(t, err, "Error Message should be nil")
	assert.Equal(t, userID, data.ID, "Error Message should be nil")
}
func TestConfirmationTransactionFailed(t *testing.T) {
	service := InitializeTrx()
	userID := 0
	_, err := service.ConfirmationTransaction(userID)
	assert.ErrorContains(t, err, "Invalid", "Error Message should be invalid")
}
