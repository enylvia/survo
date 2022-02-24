package test

import (
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
		Role:       "Surveyor",
		Password:   "12345678",
		Image:      "image.jpg",
		Phone:      "0820222",
		Birthday:   "06-01-2001",
	}

	result , err := createData.Create(input)

	assert.Equal(t, true, result.Id>0)
	assert.NoError(t, err)


}
