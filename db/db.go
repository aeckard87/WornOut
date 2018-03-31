package db

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/aeckard87/WornOut/models"
	"github.com/serenize/snaker"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

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

func DBInstance(c *gin.Context) *gorm.DB {
	return c.MustGet("DB").(*gorm.DB)
}

func (self *Parameter) SetPreloads(db *gorm.DB) *gorm.DB {
	if self.Preloads == "" {
		return db
	}

	for _, preload := range strings.Split(self.Preloads, ",") {
		var a []string

		for _, s := range strings.Split(preload, ".") {
			a = append(a, snaker.SnakeToCamel(s))
		}

		db = db.Preload(strings.Join(a, "."))
	}

	return db
}

func getConnectionString() string {
	dbUsername := os.Getenv("WORN_OUT_DB_USERNAME")
	dbPassword := os.Getenv("WORN_OUT_DB_PASSWORD")
	// dbHostname := os.Getenv("WORN_OUT_DB_HOSTNAME")
	dbName := os.Getenv("WORN_OUT_DB_NAME")
	// return fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", dbUsername, dbPassword, dbHostname, dbName)
	return fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", dbUsername, dbPassword, dbName)
}
