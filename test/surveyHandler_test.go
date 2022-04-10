package test

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"survorest/auth"
	"survorest/handler"
	"survorest/survey"
	"survorest/user"
	"testing"
)

func MigrateTableSurvey(db *gorm.DB) {
	if db.Migrator().HasTable("surveys") || db.Migrator().HasTable("questions") || db.Migrator().HasTable("answers") {

		db.Migrator().DropTable("surveys")
		db.Migrator().DropTable("questions")
		db.Migrator().DropTable("answers")
		log.Printf("Table users dropped")
		return
	}
	db.Migrator().CreateTable(&survey.Survey{})
	db.Migrator().CreateTable(&survey.Question{})
	db.Migrator().CreateTable(&survey.Answer{})
}

func TestMigrateSurvey(t *testing.T) {
	db, _ := GetConnection()
	MigrateTableSurvey(db)
}

func TestCreateSurveySuccess(t *testing.T) {
	router := getRouter()
	db, _ := GetConnection()
	surveyRepository := survey.NewRepository(db)
	surveyService := survey.NewService(surveyRepository)
	surveyHandler := handler.NewSurveyHandler(surveyService)
	authRepository := user.NewRepository(db)
	authService := user.NewService(authRepository)
	jwtWrapper := auth.JwtWrapper{
		SecretKey:       "survosecret",
		Issuer:          "AuthService",
		ExpirationHours: 2,
	}

	generatedToken, _ := jwtWrapper.GenerateToken(1, "usre@mail.com")
	validToken := "Bearer " + generatedToken

	router.Use(auth.AuthMiddleware(authService))
	router.POST("/api/v1/createsurvey", surveyHandler.CreateSurvey)
	w := httptest.NewRecorder()

	inputSurvey := survey.CreateSurveyInput{
		UserId: 1,
		SurveyCategory:    "Education",
		SurveyTitle:       "Education for People",
		SurveyDescription: "Education its very important",
		Target: 25,
		Question: []survey.QuestionInput{
			{
				UserId: 1,
				SurveyQuestion: "apakah anda pernah mengikuti kursus?",
				QuestionType: "Checkbox",
				OptionName:   "Option 1",
			},
			{
				UserId: 1,
				SurveyQuestion: "dimana anda mengikuti kursus tersebut?",
				QuestionType: "Checkbox",
				OptionName:   "Option 1",
			},
			{
				UserId: 1,
				SurveyQuestion: "apakah anda ingin mengikuti kursus?",
				QuestionType: "Checkbox",
				OptionName:   "Option 2",
			},
		},
	}

	var buf bytes.Buffer

	err := json.NewEncoder(&buf).Encode(inputSurvey)

	if err != nil {
		t.Errorf("Error encoding json")
	}

	req,err := http.NewRequest("POST", "http://localhost:8080/api/v1/createsurvey", &buf)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", validToken)
	router.ServeHTTP(w, req)

	assert.NoError(t, err)
	assert.Equal(t, 200, w.Code)

	var responseBody map[string]interface{}
	body,_ := ioutil.ReadAll(w.Body)

	json.Unmarshal(body, &responseBody)

	assert.Equal(t, "success", responseBody["meta"].(map[string]interface{})["status"], "Status code should be success")
	assert.Equal(t, "Create Survey Successfully", responseBody["meta"].(map[string]interface{})["message"], "Message code should be Survey created successfully")

}
