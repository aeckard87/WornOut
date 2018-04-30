package db

import (
	"fmt"
	"log"
	"os"

	"github.com/aeckard87/WornOut/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	// setting up mysql conn with gorm
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Connect returns *gorm.DB connection
func Connect() *gorm.DB {
	dbURL := getConnectionString()
	if dbURL == "" {
		return nil
	}

	db, err := gorm.Open("mysql", dbURL)
	if err != nil {
		log.Fatalf("Got error when connect database, the error is '%v'", err)
	}

	db.LogMode(false)

	if gin.IsDebugging() {
		db.LogMode(true)
	}

	if os.Getenv("AUTOMIGRATE") == "1" {
		db.AutoMigrate(
			&models.Category{},
		)
	}

	return db
}

// DBInstance returns *gorm.DB instance
func DBInstance(c *gin.Context) *gorm.DB {
	return c.MustGet("DB").(*gorm.DB)
}

func getConnectionString() string {
	dbUsername := os.Getenv("WORN_OUT_DB_USERNAME")
	dbPassword := os.Getenv("WORN_OUT_DB_PASSWORD")
	// dbHostname := os.Getenv("WORN_OUT_DB_HOSTNAME")
	dbName := os.Getenv("WORN_OUT_DB_NAME")
	// return fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", dbUsername, dbPassword, dbHostname, dbName)
	return fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", dbUsername, dbPassword, dbName)
}
