package test

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"strconv"
	"survorest/helper"
	"survorest/user"
	"testing"
)

func TestRegisterUser_ValidationForm(t *testing.T) {

	log.Print("TestRegisterUser_ValidationFormFailed")

	db, err := GetConnection()
	helper.ErrorNotNil(err)

	MigrateTable(db)
	defer TruncateTable(db)

	client := &http.Client{}

	//var bodyData []user.UserFormatter

	input := user.RegisterInput{
		FullName:             "john doe",
		Email:                "",
		Username:             "example",
		Occupation:           "Software Engineer",
		Password:             "12345678",
		PasswordConfirmation: "12345678",
	}
	var buf bytes.Buffer

	err = json.NewEncoder(&buf).Encode(input)
	if err != nil {
		log.Fatal(err)
	}
	req, err := http.NewRequest("POST", "http://localhost:8080/api/v1/register", &buf)

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)

	defer resp.Body.Close()

	assert.Equal(t, http.StatusUnprocessableEntity, resp.StatusCode)
	assert.Equal(t, "application/json", req.Header.Get("Content-Type"))

}

func TestRegisterUser_ValidationFormSuccess(t *testing.T) {

	log.Print("TestRegisterUser_ValidationFormSuccess")
	db, err := GetConnection()
	helper.ErrorNotNil(err)

	MigrateTable(db)
	defer TruncateTable(db)

	client := &http.Client{}

	input := user.RegisterInput{
		FullName:             "john doe",
		Email:                "example@mail.com",
		Username:             "example",
		Occupation:           "Software Engineer",
		Password:             "12345678",
		PasswordConfirmation: "12345678",
	}
	var buf bytes.Buffer

	err = json.NewEncoder(&buf).Encode(input)
	if err != nil {
		log.Fatal(err)
	}
	req, err := http.NewRequest("POST", "http://localhost:8080/api/v1/register", &buf)

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)

	defer resp.Body.Close()

	assert.Equal(t, http.StatusCreated, resp.StatusCode)
	assert.Equal(t, "application/json", req.Header.Get("Content-Type"))

}

func TestRegisterUser_CreatUserSuccess(t *testing.T) {

	log.Print("TestRegisterUser_CreateUserSuccess")
	db, err := GetConnection()
	helper.ErrorNotNil(err)

	MigrateTable(db)

	client := &http.Client{}

	input := user.RegisterInput{
		FullName:             "john doe",
		Email:                "example@mail.com",
		Username:             "example",
		Occupation:           "Software Engineer",
		Password:             "12345678",
		PasswordConfirmation: "12345678",
	}
	var buf bytes.Buffer

	err = json.NewEncoder(&buf).Encode(input)
	if err != nil {
		log.Fatal(err)
	}
	req, err := http.NewRequest("POST", "http://localhost:8080/api/v1/register", &buf)

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)

	defer resp.Body.Close()

	assert.Equal(t, http.StatusCreated, resp.StatusCode)
	assert.Equal(t, "application/json", req.Header.Get("Content-Type"))

}

func TestLoginUser_ValidationFailed(t *testing.T) {
	log.Print("TestRegisterUser_CreateUserSuccess")

	client := &http.Client{}

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

	resp, err := client.Do(req)

	defer resp.Body.Close()

	assert.Equal(t, http.StatusUnprocessableEntity, resp.StatusCode)
	assert.Equal(t, "application/json", req.Header.Get("Content-Type"))
}

func TestLoginUser_ValidationAndLoginSuccess(t *testing.T) {
	log.Print("TestRegisterUser_ValidationAndLoginSuccess")
	client := &http.Client{}

	input := user.LoginInput{
		Email:    "example@mail.com",
		Password: "12345678",
	}
	var buf bytes.Buffer

	err := json.NewEncoder(&buf).Encode(input)
	if err != nil {
		log.Fatal(err)
	}
	req, err := http.NewRequest("POST", "http://localhost:8080/api/v1/login", &buf)

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)

	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "application/json", req.Header.Get("Content-Type"))
}

func TestUpdateUser_FormValidation(t *testing.T) {
	log.Print("TestUpdateUser_FormValidation")
	client := &http.Client{}

	inputID := 1
	inputIDs := strconv.Itoa(inputID)

	inputData := user.UpdateInput{
		FullName:             "",
		Email:                "example@mail.com",
		Username:             "example",
		Password:             "12345678",
		PasswordConfirmation: "12345678",
		Image: 				  "image.jpg",
		Phone:                "081234567882",
		Birthday:             "06-01-2001",
	}

	var buf bytes.Buffer

	err := json.NewEncoder(&buf).Encode(inputData)

	if err != nil {
		log.Fatal(err)
	}

	req,_ := http.NewRequest("PUT","http://localhost:8080/api/v1/update/"+inputIDs,&buf)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)

	defer resp.Body.Close()

	assert.Equal(t, http.StatusUnprocessableEntity, resp.StatusCode)
	assert.Equal(t, "application/json", req.Header.Get("Content-Type"))

}

func TestUpdateUser_Success(t *testing.T) {
	log.Print("TestUpdateUser_Success")
	//db, err := GetConnection()
	//helper.ErrorNotNil(err)
	//
	//defer TruncateTable(db)

	client := &http.Client{}

	inputID := 1
	inputIDs := strconv.Itoa(inputID)

	inputData := user.UpdateInput{
		FullName:             "Aditya",
		Email:                "example@mail.com",
		Username:             "example",
		Password:             "12345678",
		PasswordConfirmation: "12345678",
		Image: 				  "image.jpg",
		Phone:                "081234567882",
		Birthday:             "06-01-2001",
	}

	var buf bytes.Buffer

	err := json.NewEncoder(&buf).Encode(inputData)

	if err != nil {
		log.Fatal(err)
	}

	req,_ := http.NewRequest("PUT","http://localhost:8080/api/v1/update/"+inputIDs,&buf)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)

	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "application/json", req.Header.Get("Content-Type"))

}
