package test

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
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

	//var bodyData []user.UserFormatter

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

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "application/json", req.Header.Get("Content-Type"))

}





