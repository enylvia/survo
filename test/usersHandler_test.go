package test

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"net/http/httptest"
	"survorest/auth"
	"survorest/handler"
	"survorest/user"
	"testing"
)

func getRouter() *gin.Engine{
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	return r
}

func GenerateJWT(id int , email string) (token string,err error){
	jwtWrapper := auth.JwtWrapper{
		SecretKey:       "survosecret",
		Issuer:          "AuthService",
		ExpirationHours: 2,
	}

	generatedToken, err := jwtWrapper.GenerateToken(id, email)
	if err != nil {
		return "",err
	}
	return generatedToken,nil
}
func GetConnection() (*gorm.DB, error) {
	dsn := "root:@tcp(127.0.0.1:3306)/survo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func MigrateTable(db *gorm.DB) {
	if db.Migrator().HasTable("users") || db.Migrator().HasTable("attributs") {

		db.Migrator().DropTable("users")
		db.Migrator().DropTable("attributs")
		log.Printf("Table users dropped")
		return
	}
	db.Migrator().CreateTable(&user.User{})
	db.Migrator().CreateTable(&user.Attribut{})
}
func TruncateTable(db *gorm.DB) {
	db.Migrator().DropTable("users")
	db.Migrator().DropTable("attributs")
}
func TestMigrate(t *testing.T) {
	db ,_ := GetConnection()
	MigrateTable(db)
}
func TestRegisterUser_ValidationForm(t *testing.T) {
	router := getRouter()
	db, _ := GetConnection()
	authRepository := user.NewRepository(db)
	authService := user.NewService(authRepository)
	userHandler := handler.NewUserHandler(authService)

	MigrateTable(db)
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

	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnprocessableEntity, w.Code, "Status code should be 422")
	assert.Equal(t, "application/json", req.Header.Get("Content-Type"), "Content-Type should be application/json")
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
		Email:                "johns@mail.com",
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

	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, w.Code, "Status code should be 201")
	assert.Equal(t, "application/json", req.Header.Get("Content-Type"), "Content-Type should be application/json")
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

	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, w.Code, "Status code should be 201")
	assert.Equal(t, "application/json", req.Header.Get("Content-Type"), "Content-Type should be application/json")
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

	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnprocessableEntity, w.Code, "Status code should be 422")
	assert.Equal(t, "application/json", req.Header.Get("Content-Type"), "Content-Type should be application/json")
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
		Email:    "johns@mail.com",
		Password: "123456789",
	}
	var buf bytes.Buffer

	err := json.NewEncoder(&buf).Encode(input)
	if err != nil {
		log.Fatal(err)
	}
	req, err := http.NewRequest("POST", "http://localhost:8080/api/v1/login", &buf)

	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, w.Code, "Status code should be 200")
	assert.Equal(t, "application/json", req.Header.Get("Content-Type"), "Content-Type should be application/json")
}

func TestUpdateUser_InvalidToken(t *testing.T) {
	router := getRouter()

	db, _ := GetConnection()
	authRepository := user.NewRepository(db)
	authService := user.NewService(authRepository)
	userHandler := handler.NewUserHandler(authService)
	var inputID string
	inputID = "2"

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

	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnauthorized, w.Code, "Status code should be 401")
	assert.Equal(t, "application/json", req.Header.Get("Content-Type"), "Content-Type should be application/json")
}
func TestUpdateUser_Success(t *testing.T) {
	router := getRouter()

	db, _ := GetConnection()
	authRepository := user.NewRepository(db)
	authService := user.NewService(authRepository)
	userHandler := handler.NewUserHandler(authService)
	var inputID string
	inputID = "2"

	jwtWrapper := auth.JwtWrapper{
		SecretKey:       "survosecret",
		Issuer:          "AuthService",
		ExpirationHours: 2,
	}

	generatedToken, _:= jwtWrapper.GenerateToken(2, "johns@mail.com")
	validToken := "Bearer "+generatedToken

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
	req.Header.Set("Authorization", validToken)

	router.ServeHTTP(w, req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, w.Code, "Status code should be 200")
	assert.Equal(t, "application/json", req.Header.Get("Content-Type"), "Content-Type should be application/json")
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
	generateToken,_ := jwtWrapper.GenerateToken(1,"john@mail.com")
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

	generatedToken, _:= jwtWrapper.GenerateToken(2, "johns@mail.com")
	validToken := "Bearer "+generatedToken

	router.Use(auth.AuthMiddleware(authService))

	router.GET("/api/v1/profile/:id", userHandler.GetProfile)
	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "http://localhost:8080/api/v1/profile/"+inputID, nil)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", validToken)

	router.ServeHTTP(w, req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnauthorized, w.Code, "Status code should be 401")
	assert.Equal(t, "application/json", req.Header.Get("Content-Type"), "Content-Type should be application/json")

}
func TestGetUserById_Success(t *testing.T) {
	router := getRouter()

	db, _ := GetConnection()
	authRepository := user.NewRepository(db)
	authService := user.NewService(authRepository)
	userHandler := handler.NewUserHandler(authService)
	var inputID string
	inputID = "2"

	jwtWrapper := auth.JwtWrapper{
		SecretKey:       "survosecret",
		Issuer:          "AuthService",
		ExpirationHours: 2,
	}

	generatedToken, _:= jwtWrapper.GenerateToken(2, "johns@mail.com")
	validToken := "Bearer "+generatedToken

	router.Use(auth.AuthMiddleware(authService))

	router.GET("/api/v1/profile/:id", userHandler.GetProfile)
	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "http://localhost:8080/api/v1/profile/"+inputID, nil)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", validToken)

	router.ServeHTTP(w, req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, w.Code, "Status code should be 200")
	assert.Equal(t, "application/json", req.Header.Get("Content-Type"), "Content-Type should be application/json")


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

	generatedToken, _:= jwtWrapper.GenerateToken(1, "john@mail.com")
	validToken := "Bearer "+generatedToken
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
func TestTruncateTable(t *testing.T) {
	db, _ := GetConnection()
	TruncateTable(db)
}