package db

import (
	"io/ioutil"
	"log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"ordermanagement/models"
)

func Init() *gorm.DB {
	dbURL := "host=localhost user=pg password=pass dbname=orders port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Run raw SQL migrations from migrations.sql
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	migrationFile := "pkg/db/migrations.sql"
	migrations, err := ioutil.ReadFile(migrationFile)
	if err != nil {
		log.Fatalf("failed to read migration file: %v", err)
	}
	_, err = sqlDB.Exec(string(migrations))
	if err != nil {
		log.Fatalf("failed to execute migrations: %v", err)
	}

	// Also run GORM automigrate for model changes
	db.AutoMigrate(&models.Product{}, &models.Order{}, &models.OrderProduct{})
	return db
}