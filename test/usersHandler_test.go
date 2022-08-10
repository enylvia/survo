package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"survorest/auth"
	"survorest/handler"
	"survorest/migrations"
	"survorest/survey"
	"survorest/transactions"
	"survorest/user"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func getRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	return r
}

func GenerateJWT(id int, email string) (token string, err error) {
	jwtWrapper := auth.JwtWrapper{
		SecretKey:       "survosecret",
		Issuer:          "AuthService",
		ExpirationHours: 2,
	}

	generatedToken, err := jwtWrapper.GenerateToken(id, email)
	if err != nil {
		return "", err
	}
	return generatedToken, nil
}
func GetConnection() (*gorm.DB, error) {
	// dsn := "postgres://ndgownkmqbplmm:e9ae287ceccf8d993e76540c09f9297328db128f5be24ce932a9a9bf8bb65e4f@ec2-23-23-151-191.compute-1.amazonaws.com:5432/d3mhbf33iu0k5b"
	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	dsn := "root:@tcp(127.0.0.1:3306)/survo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}
	return db, nil
}

func MigrateTable(db *gorm.DB) {
	if db.Migrator().HasTable("users") || db.Migrator().HasTable("attributs") || db.Migrator().HasTable("surveys") || db.Migrator().HasTable("questions") || db.Migrator().HasTable("answers") || db.Migrator().HasTable("transactions") {

		db.Migrator().DropTable("users")
		db.Migrator().DropTable("attributs")
		db.Migrator().DropTable("surveys")
		db.Migrator().DropTable("questions")
		db.Migrator().DropTable("answers")
		db.Migrator().DropTable("transactions")
		log.Printf("Table users dropped")
		log.Printf("Table Survey dropped")
		log.Printf("Table transactions dropped")
		return
	}

	db.Migrator().CreateTable(&migrations.User{})
	db.Migrator().CreateTable(&migrations.Survey{})
	db.Migrator().CreateTable(&migrations.Question{})
	db.Migrator().CreateTable(&migrations.Answer{})
	db.Migrator().CreateTable(&migrations.Attribut{})
	db.Migrator().CreateTable(&migrations.Transaction{})
}
func TestMigrate(t *testing.T) {
	db, _ := GetConnection()
	MigrateTable(db)
}
func TruncateTable(db *gorm.DB) {
	db.Migrator().DropTable("users")
	db.Migrator().DropTable("attributs")
	db.Migrator().DropTable("surveys")
	db.Migrator().DropTable("questions")
	db.Migrator().DropTable("answers")
}
func TestRegisterUser_ValidationForm(t *testing.T) {
	router := getRouter()
	db, _ := GetConnection()
	authRepository := user.NewRepository(db)
	authService := user.NewService(authRepository)
	userHandler := handler.NewUserHandler(authService)

	router.POST("/api/v1/register", userHandler.RegisterUser)
	w := httptest.NewRecorder()
	input := user.RegisterInput{
		FullName:             "john doe",
		Email:                "",
		Username:             "example",
		Occupation:           "Software Engineer",
		Password:             "12345678",
		PasswordConfirmation: "12345678",
	}
	var buf bytes.Buffer

	err := json.NewEncoder(&buf).Encode(input)
	if err != nil {
		log.Fatal(err)
	}
	req, err := http.NewRequest("POST", "http://localhost:8080/api/v1/register", &buf)

	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)
	var responseBody map[string]interface{}
	body, _ := ioutil.ReadAll(w.Body)

	json.Unmarshal(body, &responseBody)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnprocessableEntity, w.Code, "Status code should be 422")
	assert.Equal(t, "Key: 'RegisterInput.Email' Error:Field validation for 'Email' failed on the 'required' tag", responseBody["data"].(map[string]interface{})["errors"].([]interface{})[0].(string))
}

func TestRegisterUser_ValidationFormSuccess(t *testing.T) {
	router := getRouter()
	db, _ := GetConnection()
	authRepository := user.NewRepository(db)
	authService := user.NewService(authRepository)
	userHandler := handler.NewUserHandler(authService)

	router.POST("/api/v1/register", userHandler.RegisterUser)
	w := httptest.NewRecorder()
	input := user.RegisterInput{
		FullName:             "johns doe",
		Email:                "johser1@mail.com",
		Username:             "example",
		Occupation:           "Software Engineer",
		Password:             "123456789",
		PasswordConfirmation: "123456789",
	}
	var buf bytes.Buffer

	err := json.NewEncoder(&buf).Encode(input)
	if err != nil {
		log.Fatal(err)
	}
	req, err := http.NewRequest("POST", "http://localhost:8080/api/v1/register", &buf)

	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	var responseBody map[string]interface{}
	body, _ := ioutil.ReadAll(w.Body)

	json.Unmarshal(body, &responseBody)

	fmt.Println(responseBody)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, w.Code, "Status code should be 201")
	assert.Equal(t, "success", responseBody["meta"].(map[string]interface{})["status"], "Status code should be success")
	assert.Equal(t, input.Email, responseBody["data"].(map[string]interface{})["email"].(string), "Email should be the same and unique")
}

func TestRegisterUser_CreatUserSuccess(t *testing.T) {

	router := getRouter()

	db, _ := GetConnection()
	authRepository := user.NewRepository(db)
	authService := user.NewService(authRepository)
	userHandler := handler.NewUserHandler(authService)

	router.POST("/api/v1/register", userHandler.RegisterUser)
	w := httptest.NewRecorder()
	input := user.RegisterInput{
		FullName:             "john doe",
		Email:                "john@mail.com",
		Username:             "example",
		Occupation:           "Software Engineer",
		Password:             "12345678",
		PasswordConfirmation: "12345678",
	}
	var buf bytes.Buffer

	err := json.NewEncoder(&buf).Encode(input)
	if err != nil {
		log.Fatal(err)
	}
	req, err := http.NewRequest("POST", "http://localhost:8080/api/v1/register", &buf)

	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	var responseBody map[string]interface{}
	body, _ := ioutil.ReadAll(w.Body)

	json.Unmarshal(body, &responseBody)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, w.Code, "Status code should be 201")
	assert.Equal(t, "success", responseBody["meta"].(map[string]interface{})["status"], "Status code should be success")
	assert.Equal(t, input.Email, responseBody["data"].(map[string]interface{})["email"].(string), "Email should be the same and unique")
}
func TestLoginUser_ValidationFailed(t *testing.T) {
	router := getRouter()

	db, _ := GetConnection()
	authRepository := user.NewRepository(db)
	authService := user.NewService(authRepository)
	userHandler := handler.NewUserHandler(authService)

	router.POST("/api/v1/login", userHandler.LoginUser)
	w := httptest.NewRecorder()
	input := user.LoginInput{
		Email:    "",
		Password: "12345678",
	}
	var buf bytes.Buffer

	err := json.NewEncoder(&buf).Encode(input)
	if err != nil {
		log.Fatal(err)
	}
	req, err := http.NewRequest("POST", "http://localhost:8080/api/v1/login", &buf)

	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	var responseBody map[string]interface{}
	body, _ := ioutil.ReadAll(w.Body)

	json.Unmarshal(body, &responseBody)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnprocessableEntity, w.Code, "Status code should be 422")
	assert.Equal(t, "Key: 'LoginInput.Email' Error:Field validation for 'Email' failed on the 'required' tag", responseBody["data"].(map[string]interface{})["errors"].([]interface{})[0].(string), "Email should be not filled")
	assert.Equal(t, "error", responseBody["meta"].(map[string]interface{})["status"], "Status code should be error")
}

func TestLoginUser_ValidationAndLoginSuccess(t *testing.T) {
	router := getRouter()

	db, _ := GetConnection()
	authRepository := user.NewRepository(db)
	authService := user.NewService(authRepository)
	userHandler := handler.NewUserHandler(authService)

	router.POST("/api/v1/login", userHandler.LoginUser)
	w := httptest.NewRecorder()
	input := user.LoginInput{
		Email:    "john@mail.com",
		Password: "12345678",
	}
	var buf bytes.Buffer

	err := json.NewEncoder(&buf).Encode(input)
	if err != nil {
		log.Fatal(err)
	}
	req, err := http.NewRequest("POST", "http://localhost:8080/api/v1/login", &buf)

	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)
	var responseBody map[string]interface{}
	body, _ := ioutil.ReadAll(w.Body)

	json.Unmarshal(body, &responseBody)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, w.Code, "Status code should be 200")
	assert.Equal(t, input.Email, responseBody["data"].(map[string]interface{})["email"].(string), "Email should be same to the input")
	assert.Equal(t, "success", responseBody["meta"].(map[string]interface{})["status"], "Status code should be success")
}

func TestUpdateUser_InvalidToken(t *testing.T) {
	router := getRouter()

	db, _ := GetConnection()
	authRepository := user.NewRepository(db)
	authService := user.NewService(authRepository)
	userHandler := handler.NewUserHandler(authService)
	var inputID string
	inputID = "1"

	invalidToken := "token"

	router.Use(auth.AuthMiddleware(authService))

	router.PUT("/api/v1/update/:id", userHandler.UpdateProfile)
	w := httptest.NewRecorder()
	newData := user.UpdateInput{
		FullName:             "New John",
		Email:                "example@mail.com",
		Username:             "example",
		Password:             "12345678",
		PasswordConfirmation: "12345678",
		Phone:                "081234567882",
		Birthday:             "06-01-2001",
	}
	var buf bytes.Buffer

	err := json.NewEncoder(&buf).Encode(newData)
	if err != nil {
		log.Fatal(err)
	}
	req, err := http.NewRequest("PUT", "http://localhost:8080/api/v1/update/"+inputID, &buf)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", invalidToken)

	router.ServeHTTP(w, req)
	var responseBody map[string]interface{}
	body, _ := ioutil.ReadAll(w.Body)

	json.Unmarshal(body, &responseBody)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnauthorized, w.Code, "Status code should be 401")
	assert.Equal(t, "Invalid Authorization Format", responseBody["meta"].(map[string]interface{})["message"].(string), "Message should be Invalid Authorization Format")
	assert.Equal(t, "error", responseBody["meta"].(map[string]interface{})["status"], "Status code should be error")
}
func TestUpdateUser_Success(t *testing.T) {
	router := getRouter()

	db, _ := GetConnection()
	authRepository := user.NewRepository(db)
	authService := user.NewService(authRepository)
	userHandler := handler.NewUserHandler(authService)
	var inputID string
	inputID = "1"

	jwtWrapper := auth.JwtWrapper{
		SecretKey:       "survosecret",
		Issuer:          "AuthService",
		ExpirationHours: 2,
	}

	generatedToken, _ := jwtWrapper.GenerateToken(1, "john@mail.com")
	validToken := "Bearer " + generatedToken

	router.Use(auth.AuthMiddleware(authService))

	router.PUT("/api/v1/update/:id", userHandler.UpdateProfile)
	w := httptest.NewRecorder()
	newData := user.UpdateInput{
		FullName:             "New John",
		Email:                "example@mail.com",
		Username:             "example",
		Password:             "123456789",
		PasswordConfirmation: "123456789",
		Phone:                "081234567882",
		Birthday:             "06-01-2001",
	}
	var buf bytes.Buffer

	err := json.NewEncoder(&buf).Encode(newData)
	if err != nil {
		log.Fatal(err)
	}
	req, err := http.NewRequest("PUT", "http://localhost:8080/api/v1/update/"+inputID, &buf)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", validToken)

	router.ServeHTTP(w, req)
	var responseBody map[string]interface{}
	body, _ := ioutil.ReadAll(w.Body)

	json.Unmarshal(body, &responseBody)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, w.Code, "Status code should be 200")
	assert.Equal(t, newData.FullName, responseBody["data"].(map[string]interface{})["fullName"].(string), "FullName should be the same as newData")
	assert.Equal(t, "success", responseBody["meta"].(map[string]interface{})["status"], "Status code should be Success")
}

func TestGenerateToken(t *testing.T) {

	jwtWrapper := auth.JwtWrapper{
		SecretKey:       "survosecret",
		Issuer:          "AuthService",
		ExpirationHours: 2,
	}

	generatedToken, err := jwtWrapper.GenerateToken(1, "john@mail.com")
	assert.NoError(t, err)

	log.Printf("Generated Token: %s", generatedToken)
}

func TestValidateJwtToken(t *testing.T) {
	jwtWrapper := auth.JwtWrapper{
		SecretKey: "survosecret",
		Issuer:    "AuthService",
	}
	generateToken, _ := jwtWrapper.GenerateToken(1, "john@mail.com")
	signedToken := generateToken

	claims, err := jwtWrapper.ValidateToken(signedToken)
	assert.NoError(t, err)
	assert.Equal(t, 1, claims.UserID)
	assert.Equal(t, "john@mail.com", claims.Email)
	assert.Equal(t, "AuthService", claims.Issuer)
}

func TestGetUserById_FailedNotAuthorization(t *testing.T) {
	router := getRouter()

	db, _ := GetConnection()
	authRepository := user.NewRepository(db)
	authService := user.NewService(authRepository)
	userHandler := handler.NewUserHandler(authService)
	var inputID string
	inputID = "1"

	jwtWrapper := auth.JwtWrapper{
		SecretKey:       "survosecret",
		Issuer:          "AuthService",
		ExpirationHours: 2,
	}

	generatedToken, _ := jwtWrapper.GenerateToken(2, "johns@mail.com")
	validToken := "Bearer " + generatedToken

	router.Use(auth.AuthMiddleware(authService))

	router.GET("/api/v1/profile/:id", userHandler.GetProfile)
	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "http://localhost:8080/api/v1/profile/"+inputID, nil)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", validToken)

	router.ServeHTTP(w, req)
	var responseBody map[string]interface{}
	body, _ := ioutil.ReadAll(w.Body)

	json.Unmarshal(body, &responseBody)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnauthorized, w.Code, "Status code should be 401")
	assert.Equal(t, "error", responseBody["meta"].(map[string]interface{})["status"], "Status code should be error")

}
func TestGetUserById_Success(t *testing.T) {
	router := getRouter()

	db, _ := GetConnection()
	authRepository := user.NewRepository(db)
	authService := user.NewService(authRepository)
	userHandler := handler.NewUserHandler(authService)
	var inputID string
	inputID = "1"

	jwtWrapper := auth.JwtWrapper{
		SecretKey:       "survosecret",
		Issuer:          "AuthService",
		ExpirationHours: 2,
	}

	generatedToken, _ := jwtWrapper.GenerateToken(1, "john@mail.com")
	validToken := "Bearer " + generatedToken

	router.Use(auth.AuthMiddleware(authService))

	router.GET("/api/v1/profile/:id", userHandler.GetProfile)
	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "http://localhost:8080/api/v1/profile/"+inputID, nil)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", validToken)

	router.ServeHTTP(w, req)
	var responseBody map[string]interface{}
	body, _ := ioutil.ReadAll(w.Body)

	json.Unmarshal(body, &responseBody)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, w.Code, "Status code should be 200")
	assert.Equal(t, "New John", responseBody["data"].(map[string]interface{})["fullName"].(string), "FullName should be the same as newData")
	assert.Equal(t, "success", responseBody["meta"].(map[string]interface{})["status"], "Status code should be Success")

}

func TestOauthLogin(t *testing.T) {
	log.Print("TestOauthLogin")
	oauthStateStringGl := "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

	client := &http.Client{}
	newUrl := auth.OauthConfGl.AuthCodeURL(oauthStateStringGl)

	req, _ := http.NewRequest("GET", newUrl, nil)
	_, err := client.Do(req)

	assert.Equal(t, req.FormValue("client_id"), auth.OauthConfGl.ClientID)

	assert.Equal(t, req.FormValue("state"), oauthStateStringGl)

	assert.NoError(t, err)
}

func TestMiddlewareWithNoHeader(t *testing.T) {
	router := gin.Default()

	db, _ := GetConnection()
	authRepository := user.NewRepository(db)
	authService := user.NewService(authRepository)
	authHandler := handler.NewUserHandler(authService)

	router.Use(auth.AuthMiddleware(authService))

	router.GET("/api/v1/profile/:id", authHandler.GetProfile)
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/api/v1/profile/1", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 401, w.Code)

}

func TestMiddlewareInvalidFormatToken(t *testing.T) {
	router := gin.Default()

	db, _ := GetConnection()
	authRepository := user.NewRepository(db)
	authService := user.NewService(authRepository)
	authHandler := handler.NewUserHandler(authService)

	router.Use(auth.AuthMiddleware(authService))

	router.GET("/api/v1/profile/:id", authHandler.GetProfile)
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/api/v1/profile/1", nil)
	req.Header.Set("Authorization", "TestFormatToken")

	router.ServeHTTP(w, req)

	assert.Equal(t, 401, w.Code)
}

func TestMiddlewareInvalidToken(t *testing.T) {
	invalidToken := "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJlbWFpbCI6ImV4YW1wbGVAbWFpbC5jb20iLCJleHAiOjE2NDY0MDQ0NTksImlzcyI6IkF1dGhTZXJ2aWNlIn0.AGQ-dn4T2hXF-FLF0ZLA21qd8gmWEyarZdqYEZqiFdM"
	router := gin.Default()

	db, _ := GetConnection()
	authRepository := user.NewRepository(db)
	authService := user.NewService(authRepository)
	authHandler := handler.NewUserHandler(authService)

	router.Use(auth.AuthMiddleware(authService))

	router.GET("/api/v1/profile/:id", authHandler.GetProfile)
	w := httptest.NewRecorder()

	req, err := http.NewRequest("GET", "/api/v1/profile/1", nil)
	req.Header.Set("Authorization", invalidToken)
	router.ServeHTTP(w, req)

	assert.NoError(t, err)
	assert.Equal(t, 401, w.Code)
}

func TestTokenisValid(t *testing.T) {
	jwtWrapper := auth.JwtWrapper{
		SecretKey:       "survosecret",
		Issuer:          "AuthService",
		ExpirationHours: 2,
	}

	generatedToken, _ := jwtWrapper.GenerateToken(1, "john@mail.com")
	validToken := "Bearer " + generatedToken
	router := gin.Default()

	db, _ := GetConnection()
	authRepository := user.NewRepository(db)
	authService := user.NewService(authRepository)
	authHandler := handler.NewUserHandler(authService)

	router.Use(auth.AuthMiddleware(authService))

	router.GET("/api/v1/profile/:id", authHandler.GetProfile)
	w := httptest.NewRecorder()

	req, err := http.NewRequest("GET", "/api/v1/profile/1", nil)
	req.Header.Set("Authorization", validToken)
	router.ServeHTTP(w, req)

	assert.NoError(t, err)
	assert.Equal(t, 200, w.Code)
}
func TestCreateSurveySuccess(t *testing.T) {
	router := getRouter()
	db, _ := GetConnection()
	surveyRepository := survey.NewRepository(db)
	userRepository := user.NewRepository(db)
	surveyService := survey.NewService(surveyRepository, userRepository)
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
		UserId:            1,
		SurveyCategory:    "Education",
		SurveyTitle:       "Education for People",
		SurveyDescription: "Education its very important",
		Target:            25,
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

	var buf bytes.Buffer

	err := json.NewEncoder(&buf).Encode(inputSurvey)

	if err != nil {
		t.Errorf("Error encoding json")
	}

	req, err := http.NewRequest("POST", "http://localhost:8080/api/v1/createsurvey", &buf)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", validToken)
	router.ServeHTTP(w, req)

	assert.NoError(t, err)
	assert.Equal(t, 200, w.Code)

	var responseBody map[string]interface{}
	body, _ := ioutil.ReadAll(w.Body)

	json.Unmarshal(body, &responseBody)

	assert.Equal(t, "success", responseBody["meta"].(map[string]interface{})["status"], "Status code should be success")
	assert.Equal(t, "Create Survey Successfully", responseBody["meta"].(map[string]interface{})["message"], "Message code should be Survey created successfully")

}

func TestCreateSurveyFailedAuthorized(t *testing.T) {
	router := getRouter()
	db, _ := GetConnection()
	surveyRepository := survey.NewRepository(db)
	userRepository := user.NewRepository(db)
	surveyService := survey.NewService(surveyRepository, userRepository)
	surveyHandler := handler.NewSurveyHandler(surveyService)
	authRepository := user.NewRepository(db)
	authService := user.NewService(authRepository)
	validToken := "Bearer FailedToken"
	router.Use(auth.AuthMiddleware(authService))
	router.POST("/api/v1/createsurvey", surveyHandler.CreateSurvey)
	w := httptest.NewRecorder()

	inputSurvey := survey.CreateSurveyInput{
		UserId:            1,
		SurveyCategory:    "Education",
		SurveyTitle:       "Education for People",
		SurveyDescription: "Education its very important",
		Target:            25,
		Question: []survey.QuestionInput{
			{
				UserId:         1,
				SurveyQuestion: "apakah anda pernah mengikuti kursus?",
				QuestionType:   "Checkbox",
				OptionName:     "Option 1",
			},
			{
				UserId:         1,
				SurveyQuestion: "dimana anda mengikuti kursus tersebut?",
				QuestionType:   "Checkbox",
				OptionName:     "Option 1",
			},
			{
				UserId:         1,
				SurveyQuestion: "apakah anda ingin mengikuti kursus?",
				QuestionType:   "Checkbox",
				OptionName:     "Option 2",
			},
		},
	}

	var buf bytes.Buffer

	err := json.NewEncoder(&buf).Encode(inputSurvey)

	if err != nil {
		t.Errorf("Error encoding json")
	}

	req, err := http.NewRequest("POST", "http://localhost:8080/api/v1/createsurvey", &buf)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", validToken)
	router.ServeHTTP(w, req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnauthorized, w.Code)

	var responseBody map[string]interface{}
	body, _ := ioutil.ReadAll(w.Body)

	json.Unmarshal(body, &responseBody)

	assert.Equal(t, "error", responseBody["meta"].(map[string]interface{})["status"], "Status code should be error")
	assert.Equal(t, "Unauthorize", responseBody["meta"].(map[string]interface{})["message"], "Message code should be Unauthorize")

}

func TestCreateSurveyFailedForInvalidInput(t *testing.T) {
	router := getRouter()
	db, _ := GetConnection()
	surveyRepository := survey.NewRepository(db)
	userRepository := user.NewRepository(db)
	surveyService := survey.NewService(surveyRepository, userRepository)
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

	inputSurvey := survey.QuestionInput{
		SurveyId:       1,
		UserId:         1,
		SurveyQuestion: "Apakah anda kurang berpengalaman dalam bidangnya?",
		QuestionType:   "Radio Button",
		OptionName:     "Yes",
	}
	var buf bytes.Buffer

	err := json.NewEncoder(&buf).Encode(inputSurvey)

	if err != nil {
		t.Errorf("Error encoding json")
	}

	req, err := http.NewRequest("POST", "http://localhost:8080/api/v1/createsurvey", &buf)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", validToken)
	router.ServeHTTP(w, req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)

	var responseBody map[string]interface{}
	body, _ := ioutil.ReadAll(w.Body)

	json.Unmarshal(body, &responseBody)

	assert.Equal(t, "error", responseBody["meta"].(map[string]interface{})["status"], "Status code should be error")
	assert.Equal(t, "Invalid Input", responseBody["meta"].(map[string]interface{})["message"], "Message code should be Invalid Input")

}

func TestGetSurveyList(t *testing.T) {
	router := getRouter()
	db, _ := GetConnection()
	surveyRepository := survey.NewRepository(db)
	userRepository := user.NewRepository(db)
	surveyService := survey.NewService(surveyRepository, userRepository)
	surveyHandler := handler.NewSurveyHandler(surveyService)
	router.GET("/api/v1/surveylist", surveyHandler.SurveyList)
	w := httptest.NewRecorder()

	req, err := http.NewRequest("GET", "http://localhost:8080/api/v1/surveylist", nil)
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, w.Code)

	var responseBody map[string]interface{}
	body, _ := ioutil.ReadAll(w.Body)

	json.Unmarshal(body, &responseBody)

	assert.Equal(t, "success", responseBody["meta"].(map[string]interface{})["status"], "Status code should be success")
	assert.Equal(t, "Successfully get list survey", responseBody["meta"].(map[string]interface{})["message"], "Message code should be Successfully get list survey")
}
func TestGetSurveyListByIDUser(t *testing.T) {
	router := getRouter()
	db, _ := GetConnection()
	surveyRepository := survey.NewRepository(db)
	userRepository := user.NewRepository(db)
	surveyService := survey.NewService(surveyRepository, userRepository)
	surveyHandler := handler.NewSurveyHandler(surveyService)
	router.GET("/api/v1/surveylist", surveyHandler.SurveyList)
	w := httptest.NewRecorder()
	userID := strconv.Itoa(1)

	req, err := http.NewRequest("GET", "http://localhost:8080/api/v1/surveylist?user_id="+userID, nil)
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, w.Code)

	var responseBody map[string]interface{}
	body, _ := ioutil.ReadAll(w.Body)

	json.Unmarshal(body, &responseBody)

	assert.Equal(t, "success", responseBody["meta"].(map[string]interface{})["status"], "Status code should be success")
	assert.Equal(t, "Successfully get list survey", responseBody["meta"].(map[string]interface{})["message"], "Message code should be Successfully get list survey")

}
func TestGetSurveyDetail(t *testing.T) {
	router := getRouter()
	db, _ := GetConnection()
	surveyRepository := survey.NewRepository(db)
	userRepository := user.NewRepository(db)
	surveyService := survey.NewService(surveyRepository, userRepository)
	surveyHandler := handler.NewSurveyHandler(surveyService)
	router.GET("/api/v1/surveydetail/:id", surveyHandler.GetSurveyDetail)
	w := httptest.NewRecorder()
	userID := strconv.Itoa(1)

	req, err := http.NewRequest("GET", "http://localhost:8080/api/v1/surveydetail/"+userID, nil)
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, w.Code)

	var responseBody map[string]interface{}
	body, _ := ioutil.ReadAll(w.Body)

	json.Unmarshal(body, &responseBody)

	assert.Equal(t, "success", responseBody["meta"].(map[string]interface{})["status"], "Status code should be success")
	assert.Equal(t, "Successfully get survey detail", responseBody["meta"].(map[string]interface{})["message"], "Message code should be Successfully get list survey")

}
func TestGetSurveyDetailInvalidInput(t *testing.T) {
	router := getRouter()
	db, _ := GetConnection()
	surveyRepository := survey.NewRepository(db)
	userRepository := user.NewRepository(db)
	surveyService := survey.NewService(surveyRepository, userRepository)
	surveyHandler := handler.NewSurveyHandler(surveyService)
	router.GET("/api/v1/surveydetail/:id", surveyHandler.GetSurveyDetail)
	w := httptest.NewRecorder()
	userID := "survey_1"

	req, err := http.NewRequest("GET", "http://localhost:8080/api/v1/surveydetail/"+userID, nil)
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)

	var responseBody map[string]interface{}
	body, _ := ioutil.ReadAll(w.Body)

	json.Unmarshal(body, &responseBody)

	assert.Equal(t, "error", responseBody["meta"].(map[string]interface{})["status"], "Status code should be error")
	assert.Equal(t, "Invalid Input", responseBody["meta"].(map[string]interface{})["message"], "Message code should be Invalid Input")

}
func TestGetSurveyDetailNotFound(t *testing.T) {
	router := getRouter()
	db, _ := GetConnection()
	surveyRepository := survey.NewRepository(db)
	userRepository := user.NewRepository(db)
	surveyService := survey.NewService(surveyRepository, userRepository)
	surveyHandler := handler.NewSurveyHandler(surveyService)
	router.GET("/api/v1/surveydetail/:id", surveyHandler.GetSurveyDetail)
	w := httptest.NewRecorder()
	surveyID := "4"

	req, err := http.NewRequest("GET", "http://localhost:8080/api/v1/surveydetail/"+surveyID, nil)
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, w.Code)

	var responseBody map[string]interface{}
	body, _ := ioutil.ReadAll(w.Body)

	json.Unmarshal(body, &responseBody)

	assert.Equal(t, "error", responseBody["meta"].(map[string]interface{})["status"], "Status code should be error")
	assert.Equal(t, "Survey not found", responseBody["meta"].(map[string]interface{})["message"], "Message code should be Invalid Input")

}

func TestAnswerQuestionSuccess(t *testing.T) {
	router := getRouter()
	db, _ := GetConnection()
	surveyRepository := survey.NewRepository(db)
	userRepository := user.NewRepository(db)
	surveyService := survey.NewService(surveyRepository, userRepository)
	surveyHandler := handler.NewSurveyHandler(surveyService)

	jwtWrapper := auth.JwtWrapper{
		SecretKey:       "survosecret",
		Issuer:          "AuthService",
		ExpirationHours: 2,
	}

	generatedToken, _ := jwtWrapper.GenerateToken(1, "usre@mail.com")
	validToken := "Bearer " + generatedToken

	router.POST("/api/v1/answerquestion", surveyHandler.AnswerQuestion)
	w := httptest.NewRecorder()

	answerQuestion := []survey.AnswerInput{
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
	var buf bytes.Buffer

	err := json.NewEncoder(&buf).Encode(answerQuestion)

	if err != nil {
		t.Errorf("Error encoding json")
	}

	req, err := http.NewRequest("POST", "http://localhost:8080/api/v1/answerquestion", &buf)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", validToken)
	router.ServeHTTP(w, req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, w.Code)

	var responseBody map[string]interface{}
	body, _ := ioutil.ReadAll(w.Body)

	json.Unmarshal(body, &responseBody)

	assert.Equal(t, "success", responseBody["meta"].(map[string]interface{})["status"], "Status code should be success")
	assert.Equal(t, "Successfully answer question", responseBody["meta"].(map[string]interface{})["message"], "Message code should be Answer Submitted")

}

func TestAnswerWithoutAuthorization(t *testing.T) {
	router := getRouter()
	db, _ := GetConnection()
	surveyRepository := survey.NewRepository(db)
	userRepository := user.NewRepository(db)
	surveyService := survey.NewService(surveyRepository, userRepository)
	surveyHandler := handler.NewSurveyHandler(surveyService)
	authRepository := user.NewRepository(db)
	authService := user.NewService(authRepository)
	router.Use(auth.AuthMiddleware(authService))
	router.POST("/api/v1/answerquestion", surveyHandler.AnswerQuestion)
	w := httptest.NewRecorder()

	answerQuestion := []survey.AnswerInput{
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
	var buf bytes.Buffer

	err := json.NewEncoder(&buf).Encode(answerQuestion)

	if err != nil {
		t.Errorf("Error encoding json")
	}

	req, err := http.NewRequest("POST", "http://localhost:8080/api/v1/answerquestion", &buf)
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnauthorized, w.Code)

	var responseBody map[string]interface{}
	body, _ := ioutil.ReadAll(w.Body)

	json.Unmarshal(body, &responseBody)

	assert.Equal(t, "error", responseBody["meta"].(map[string]interface{})["status"], "Status code should be error")
	assert.Equal(t, "Header not provided", responseBody["meta"].(map[string]interface{})["message"], "Message code should be Header not provided")
}

func TestGetAllTransactionByIDUser_InvalidInput(t *testing.T) {
	router := getRouter()
	db, _ := GetConnection()
	transactionRepository := transactions.NewRepository(db)
	authRepository := user.NewRepository(db)
	authService := user.NewService(authRepository)
	userRepository := user.NewRepository(db)
	transactionService := transactions.NewService(transactionRepository, userRepository)
	transactionHandler := handler.NewTransactionHandler(transactionService)

	jwtWrapper := auth.JwtWrapper{
		SecretKey:       "survosecret",
		Issuer:          "AuthService",
		ExpirationHours: 2,
	}

	generatedToken, _ := jwtWrapper.GenerateToken(1, "usre@mail.com")
	validToken := "Bearer " + generatedToken

	router.Use(auth.AuthMiddleware(authService))
	router.GET("/api/v1/transaction/:id", transactionHandler.GetAllTransactionByIDUser)
	w := httptest.NewRecorder()
	userID := "user_1"

	req, err := http.NewRequest("GET", "http://localhost:8080/api/v1/transaction/"+userID, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", validToken)
	router.ServeHTTP(w, req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)

	var responseBody map[string]interface{}
	body, _ := ioutil.ReadAll(w.Body)

	json.Unmarshal(body, &responseBody)

	assert.Equal(t, "error", responseBody["meta"].(map[string]interface{})["status"], "Status code should be error")
	assert.Equal(t, "Invalid Input", responseBody["meta"].(map[string]interface{})["message"], "Message code should be Invalid Input")

}
func TestGetAllTransactionByIDUser_Success(t *testing.T) {
	router := getRouter()
	db, _ := GetConnection()
	transactionRepository := transactions.NewRepository(db)
	authRepository := user.NewRepository(db)
	authService := user.NewService(authRepository)
	userRepository := user.NewRepository(db)
	transactionService := transactions.NewService(transactionRepository, userRepository)
	transactionHandler := handler.NewTransactionHandler(transactionService)

	jwtWrapper := auth.JwtWrapper{
		SecretKey:       "survosecret",
		Issuer:          "AuthService",
		ExpirationHours: 2,
	}

	generatedToken, _ := jwtWrapper.GenerateToken(1, "usre@mail.com")
	validToken := "Bearer " + generatedToken

	router.Use(auth.AuthMiddleware(authService))
	router.GET("/api/v1/transaction/:id", transactionHandler.GetAllTransactionByIDUser)
	w := httptest.NewRecorder()
	userID := "1"

	req, err := http.NewRequest("GET", "http://localhost:8080/api/v1/transaction/"+userID, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", validToken)
	router.ServeHTTP(w, req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, w.Code)

	var responseBody map[string]interface{}
	body, _ := ioutil.ReadAll(w.Body)

	json.Unmarshal(body, &responseBody)

	assert.Equal(t, "success", responseBody["meta"].(map[string]interface{})["status"], "Status code should be success")
	assert.Equal(t, "Successfully get all transactions", responseBody["meta"].(map[string]interface{})["message"], "Message code should be Successfully get all transactions")

}
func TestCreateTransactionFailedIfNotLogin(t *testing.T) {
	router := getRouter()
	db, _ := GetConnection()
	transactionRepository := transactions.NewRepository(db)
	authRepository := user.NewRepository(db)
	authService := user.NewService(authRepository)
	userRepository := user.NewRepository(db)
	transactionService := transactions.NewService(transactionRepository, userRepository)
	transactionHandler := handler.NewTransactionHandler(transactionService)

	router.Use(auth.AuthMiddleware(authService))
	router.POST("/api/v1/createtransaction", transactionHandler.CreateTransaction)
	w := httptest.NewRecorder()

	transactionInput := transactions.CreateTransactionInput{
		UserID: 1,
		Amount: 10000,
	}
	var buf bytes.Buffer

	err := json.NewEncoder(&buf).Encode(transactionInput)

	if err != nil {
		t.Errorf("Error encoding json")
	}

	req, err := http.NewRequest("POST", "http://localhost:8080/api/v1/createtransaction", &buf)
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnauthorized, w.Code)

	var responseBody map[string]interface{}
	body, _ := ioutil.ReadAll(w.Body)

	json.Unmarshal(body, &responseBody)

	assert.Equal(t, "error", responseBody["meta"].(map[string]interface{})["status"], "Status code should be error")
	assert.Equal(t, "Header not provided", responseBody["meta"].(map[string]interface{})["message"], "Message code should be Header not provided")

}
func TestCreateTransactionFailedIfTokenNotValid(t *testing.T) {
	router := getRouter()
	db, _ := GetConnection()
	transactionRepository := transactions.NewRepository(db)
	authRepository := user.NewRepository(db)
	authService := user.NewService(authRepository)
	userRepository := user.NewRepository(db)
	transactionService := transactions.NewService(transactionRepository, userRepository)
	transactionHandler := handler.NewTransactionHandler(transactionService)

	jwtWrapper := auth.JwtWrapper{
		SecretKey:       "survosecrets",
		Issuer:          "AuthService",
		ExpirationHours: 2,
	}

	generatedToken, _ := jwtWrapper.GenerateToken(1, "usre@mail.com")
	validToken := "Bearer " + generatedToken

	router.Use(auth.AuthMiddleware(authService))
	router.POST("/api/v1/createtransaction", transactionHandler.CreateTransaction)
	w := httptest.NewRecorder()

	transactionInput := transactions.CreateTransactionInput{
		UserID: 1,
		Amount: 10000,
	}
	var buf bytes.Buffer

	err := json.NewEncoder(&buf).Encode(transactionInput)

	if err != nil {
		t.Errorf("Error encoding json")
	}

	req, err := http.NewRequest("POST", "http://localhost:8080/api/v1/createtransaction", &buf)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", validToken)
	router.ServeHTTP(w, req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnauthorized, w.Code)

	var responseBody map[string]interface{}
	body, _ := ioutil.ReadAll(w.Body)

	json.Unmarshal(body, &responseBody)

	assert.Equal(t, "error", responseBody["meta"].(map[string]interface{})["status"], "Status code should be error")
	assert.Equal(t, "Unauthorize", responseBody["meta"].(map[string]interface{})["message"], "Message code should be Unauthorize")

}
func TestCreateTransactionSuccessIfTokenValid(t *testing.T) {
	router := getRouter()
	db, _ := GetConnection()
	transactionRepository := transactions.NewRepository(db)
	authRepository := user.NewRepository(db)
	authService := user.NewService(authRepository)
	userRepository := user.NewRepository(db)
	transactionService := transactions.NewService(transactionRepository, userRepository)
	transactionHandler := handler.NewTransactionHandler(transactionService)

	jwtWrapper := auth.JwtWrapper{
		SecretKey:       "survosecret",
		Issuer:          "AuthService",
		ExpirationHours: 2,
	}

	generatedToken, _ := jwtWrapper.GenerateToken(1, "usre@mail.com")
	validToken := "Bearer " + generatedToken

	router.Use(auth.AuthMiddleware(authService))
	router.POST("/api/v1/createtransaction", transactionHandler.CreateTransaction)
	w := httptest.NewRecorder()

	transactionInput := transactions.CreateTransactionInput{
		UserID: 1,
		Amount: 10000,
	}
	var buf bytes.Buffer

	err := json.NewEncoder(&buf).Encode(transactionInput)

	if err != nil {
		t.Errorf("Error encoding json")
	}

	req, err := http.NewRequest("POST", "http://localhost:8080/api/v1/createtransaction", &buf)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", validToken)
	router.ServeHTTP(w, req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, w.Code)

	var responseBody map[string]interface{}
	body, _ := ioutil.ReadAll(w.Body)

	json.Unmarshal(body, &responseBody)

	assert.Equal(t, "success", responseBody["meta"].(map[string]interface{})["status"], "Status code should be success")
	assert.Equal(t, "Successfully create transaction", responseBody["meta"].(map[string]interface{})["message"], "Message code should be Successfully create transaction")

}
func TestTransactionPremiumFailedIfTokenNotValid(t *testing.T) {
	router := getRouter()
	db, _ := GetConnection()
	transactionRepository := transactions.NewRepository(db)
	authRepository := user.NewRepository(db)
	authService := user.NewService(authRepository)
	userRepository := user.NewRepository(db)
	transactionService := transactions.NewService(transactionRepository, userRepository)
	transactionHandler := handler.NewTransactionHandler(transactionService)

	jwtWrapper := auth.JwtWrapper{
		SecretKey:       "survosecrets",
		Issuer:          "AuthService",
		ExpirationHours: 2,
	}

	generatedToken, _ := jwtWrapper.GenerateToken(1, "usre@mail.com")
	validToken := "Bearer " + generatedToken
	userID := "1"
	router.Use(auth.AuthMiddleware(authService))
	router.POST("/api/v1/transactionpremium/"+userID, transactionHandler.CreateTransactionPremium)
	w := httptest.NewRecorder()

	transactionInput := transactions.CreateTransactionInput{
		UserID: 1,
		Amount: 10000,
	}
	var buf bytes.Buffer

	err := json.NewEncoder(&buf).Encode(transactionInput)

	if err != nil {
		t.Errorf("Error encoding json")
	}

	req, err := http.NewRequest("POST", "http://localhost:8080/api/v1/transactionpremium/"+userID, &buf)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", validToken)
	router.ServeHTTP(w, req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnauthorized, w.Code)

	var responseBody map[string]interface{}
	body, _ := ioutil.ReadAll(w.Body)

	json.Unmarshal(body, &responseBody)

	assert.Equal(t, "error", responseBody["meta"].(map[string]interface{})["status"], "Status code should be error")
	assert.Equal(t, "Unauthorize", responseBody["meta"].(map[string]interface{})["message"], "Message code should be Unauthorize")

}
func TestTransactionPremiumSuccessIfTokenValid(t *testing.T) {
	router := getRouter()
	db, _ := GetConnection()
	transactionRepository := transactions.NewRepository(db)
	authRepository := user.NewRepository(db)
	authService := user.NewService(authRepository)
	userRepository := user.NewRepository(db)
	transactionService := transactions.NewService(transactionRepository, userRepository)
	transactionHandler := handler.NewTransactionHandler(transactionService)

	jwtWrapper := auth.JwtWrapper{
		SecretKey:       "survosecret",
		Issuer:          "AuthService",
		ExpirationHours: 2,
	}
	userID := "1"
	generatedToken, _ := jwtWrapper.GenerateToken(1, "usre@mail.com")
	validToken := "Bearer " + generatedToken

	router.Use(auth.AuthMiddleware(authService))
	router.POST("/api/v1/createtransaction/:id", transactionHandler.CreateTransactionPremium)
	w := httptest.NewRecorder()

	transactionInput := transactions.CreateTransactionInput{
		UserID: 1,
		Amount: 10000,
	}
	var buf bytes.Buffer

	err := json.NewEncoder(&buf).Encode(transactionInput)

	if err != nil {
		t.Errorf("Error encoding json")
	}

	req, err := http.NewRequest("POST", "http://localhost:8080/api/v1/createtransaction/"+userID, &buf)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", validToken)
	router.ServeHTTP(w, req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, w.Code)

	var responseBody map[string]interface{}
	body, _ := ioutil.ReadAll(w.Body)

	json.Unmarshal(body, &responseBody)

	assert.Equal(t, "success", responseBody["meta"].(map[string]interface{})["status"], "Status code should be success")
	assert.Equal(t, "Successfully create transaction", responseBody["meta"].(map[string]interface{})["message"], "Message code should be Successfully create transaction")

}

//func TestTruncateTable(t *testing.T) {
//	db, _ := GetConnection()
//	TruncateTable(db)
//}
