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

	fmt.Println("Migrate Success")
}
func TruncateTable(db *gorm.DB) {
	db.Exec("TRUNCATE users")
}

func TestMigrateTableUser(t *testing.T) {
	db, err := GetConnection()
	helper.ErrorNotNil(err)

	MigrateTable(db)
	assert.Equal(t, true, db.Migrator().HasTable("users"))
	assert.NoError(t, err)
}

func TestCreate(t *testing.T) {

}
