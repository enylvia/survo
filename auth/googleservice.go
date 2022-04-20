package auth

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"log"
	"net/http"
	"survorest/helper"
	"survorest/user"
)

type NewGoogleService interface {
	HandleLogin(c *gin.Context)
	HandleCallback(c *gin.Context)
}

type googleservice struct {
	repository user.Repository
}

var (
	OauthConfGl = &oauth2.Config{
		RedirectURL:  "http://localhost:8080/api/v1/callback",
		ClientID:     "782280680980-mg73o4fhqllch96s65mbqkcp47mvckkk.apps.googleusercontent.com",
		ClientSecret: "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",
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
	oauthStateStringGl = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="
	JWTWrappergl       = JwtWrapper{
		SecretKey:       "survosecret",
		Issuer:          "AuthService",
		ExpirationHours: 2,
	}
)

func GoogleService(repository user.Repository) *googleservice {
	return &googleservice{repository: repository}
}

func HandleLogin(c *gin.Context) {
	url := OauthConfGl.AuthCodeURL(oauthStateStringGl)
	http.Redirect(c.Writer, c.Request, url, http.StatusTemporaryRedirect)
}
func (r *googleservice) HandleCallback(c *gin.Context) {
	if c.Request.FormValue("state") != oauthStateStringGl {
		log.Printf("error, state not valid")
		http.Redirect(c.Writer, c.Request, "/api/v1/google/login", http.StatusTemporaryRedirect)
		return
	}
	token, err := OauthConfGl.Exchange(oauth2.NoContext, c.Request.FormValue("code"))
	if err != nil {
		log.Printf("could not get token %s\n", err.Error())
		http.Redirect(c.Writer, c.Request, "/api/v1/google/login", http.StatusTemporaryRedirect)
		return
	}
	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		log.Printf("could not get request %s\n", err.Error())
		http.Redirect(c.Writer, c.Request, "/api/v1/google/login", http.StatusTemporaryRedirect)
		return
	}

	defer resp.Body.Close()
	var Formatter GoogleFormatter
	if err = json.NewDecoder(resp.Body).Decode(&Formatter); err != nil {
		log.Printf("could not parse response %s\n", err.Error())
		http.Redirect(c.Writer, c.Request, "/api/v1/google/login", http.StatusTemporaryRedirect)
		return
	}
	checkIfUserExist, err := r.repository.FindByEmail(Formatter.Email)
	checkIfUserExist.Email = Formatter.Email
	if err != nil {
		checkIfUserExist.Email = Formatter.Email
		r.repository.Create(checkIfUserExist)
		newtoken := JwtWrapper{
			SecretKey:       "survosecret",
			Issuer:          "AuthService",
			ExpirationHours: 2,
		}
		tokenString, err := newtoken.GenerateToken(int(checkIfUserExist.Id),checkIfUserExist.Email)
		if err != nil {
			errorMessage := gin.H{"error": err.Error()}
			response := helper.ApiResponse("Failed to generate token", http.StatusUnprocessableEntity, "error", errorMessage)
			c.JSON(http.StatusUnprocessableEntity, response)
			return
		}
		formatter := FormatGoogle(checkIfUserExist.Email, tokenString)

		response := helper.ApiResponse("Login Successfully", http.StatusOK, "Successfully login using Gmail", formatter)
		c.JSON(http.StatusOK, response)
		return
	}
	newtoken := JwtWrapper{
		SecretKey:       "survosecret",
		Issuer:          "AuthService",
		ExpirationHours: 2,
	}
	tokenString, err := newtoken.GenerateToken(int(checkIfUserExist.Id),checkIfUserExist.Email)
	if err != nil {
		errorMessage := gin.H{"error": err.Error()}
		response := helper.ApiResponse("Failed to generate token", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	formatter := FormatGoogle(checkIfUserExist.Email, tokenString)

	response := helper.ApiResponse("Login Successfully", http.StatusOK, "Successfully login using Gmail", formatter)
	c.JSON(http.StatusOK, response)
	return
}
