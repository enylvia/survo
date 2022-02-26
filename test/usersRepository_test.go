package test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"survorest/helper"
	"survorest/user"
	"testing"
)

func GetConnection() (*gorm.DB, error) {
	dsn := "root:@tcp(127.0.0.1:3306)/survo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func MigrateTable(db *gorm.DB) {
	if db.Migrator().HasTable("users") {
		db.Migrator().DropTable("users")
		log.Printf("Table users dropped")
		return
	}
	db.Migrator().CreateTable(&user.User{})
}

func TruncateTable(db *gorm.DB) {
	db.Migrator().DropTable("users")
}

func TestMigrateTableUser(t *testing.T) {
	db, err := GetConnection()
	helper.ErrorNotNil(err)

	MigrateTable(db)

	defer TruncateTable(db)
	assert.Equal(t, true, db.Migrator().HasTable("users"))
	assert.NoError(t, err)
}

func TestCreateRepository_withDummyData(t *testing.T) {
	db, err := GetConnection()
	helper.ErrorNotNil(err)

	MigrateTable(db)
	defer TruncateTable(db)

	createData := user.NewRepository(db)
	input := user.User{
		Id:         1,
		FullName:   "Adit",
		Email:      "adit@mail.com",
		Username:   "addityap",
		Occupation: "Mahasiswa",
		Password:   "12345678",
		Image:      "image.jpg",
		Phone:      "0820222",
		Birthday:   "06-01-2001",
	}

	result, err := createData.Create(input)

	assert.Equal(t, true, result.Id > 0)
	assert.NoError(t, err)

}

func TestFindByEmail_withDummyEmail(t *testing.T) {
	db, err := GetConnection()
	helper.ErrorNotNil(err)

	MigrateTable(db)
	defer TruncateTable(db)

	findByEmail := user.NewRepository(db)
	input := user.User{
		Id:         1,
		FullName:   "Adit",
		Email:      "adit@mail.com",
		Username:   "addityap",
		Occupation: "Mahasiswa",
		Password:   "12345678",
		Image:      "image.jpg",
		Phone:      "0820222",
		Birthday:   "06-01-2001",
	}
	_, err = findByEmail.Create(input)

	dummyEmail := "adit@mail.com"
	result, err := findByEmail.FindByEmail(dummyEmail)

	assert.Equal(t, dummyEmail, result.Email)
	assert.NoError(t, err)
}

func TestFindByID_withDummyIdUser(t *testing.T) {
	db, err := GetConnection()
	helper.ErrorNotNil(err)

	MigrateTable(db)
	defer TruncateTable(db)

	findByID := user.NewRepository(db)
	input := user.User{
		Id:         1,
		FullName:   "Adit",
		Email:      "adit@mail.com",
		Username:   "addityap",
		Occupation: "Mahasiswa",
		Password:   "12345678",
		Image:      "image.jpg",
		Phone:      "0820222",
		Birthday:   "06-01-2001",
	}
	_, err = findByID.Create(input)

	dummyID := input.Id

	result, err := findByID.FindByID(int(dummyID))

	assert.Equal(t, dummyID, result.Id)
	assert.NoError(t, err)
}

func TestUpdateData_withDummyData(t *testing.T) {
	db, err := GetConnection()
	helper.ErrorNotNil(err)

	MigrateTable(db)
	defer TruncateTable(db)

	updateData := user.NewRepository(db)
	input := user.User{
		Id:         1,
		FullName:   "Adit",
		Email:      "adit@mail.com",
		Username:   "addityap",
		Occupation: "Mahasiswa",
		Password:   "12345678",
		Image:      "image.jpg",
		Phone:      "0820222",
		Birthday:   "06-01-2001",
	}
	result, err := updateData.Create(input)

	fmt.Println(result)
		data , err := updateData.FindByID(int(input.Id))
		data.FullName=   "Adist"
		data.Email=      "adist@mail.com"
		data.Username=   "addistyap"
		data.Occupation= "Mahasiswas"
		data.Password=   "12345678s"
		data.Image=      "image.jpg"
		data.Phone=      "0820222"
		data.Birthday=  "06-01-2001"

		db.Save(data)

	assert.NotEqual(t, input, data)
	assert.NoError(t, err)
}
