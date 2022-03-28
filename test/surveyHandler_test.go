package test

import (
	"gorm.io/gorm"
	"log"
	"survorest/survey"
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

func TestCreateSurvey(t *testing.T) {

}
