// Code generated by go-swagger; DO NOT EDIT.

package main

import (
	"flag"
	"log"
	"os"

	dbpkg "github.com/aeckard87/WornOut/db"
	model "github.com/aeckard87/WornOut/models"
	loads "github.com/go-openapi/loads"
	flags "github.com/jessevdk/go-flags"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/aeckard87/WornOut/restapi"
	"github.com/aeckard87/WornOut/restapi/operations"
)

// This file was generated by the swagger tool.
// Make sure not to overwrite this file after you generated it because all your edits would be lost!
type Product struct {
	gorm.Model
	Code  string
	Price uint
}

var (
	port = os.Getenv("PORT")
	host = flag.String("host", "0.0.0.0", "HTTP listener host")
)

func main() {

	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewWornOutAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()

	parser := flags.NewParser(server, flags.Default)
	parser.ShortDescription = "Worn Out"
	parser.LongDescription = "Item API to help define a thing belonging to users, for example clothing."

	//init db
	db := dbpkg.Connect()

	defer db.Close()

	db.AutoMigrate(&model.Category{}, &model.SubCategory{}, &model.Detail{}, &model.Descriptor{}, &model.User{}, &model.Item{})

	//init server
	server.ConfigureFlags()
	for _, optsGroup := range api.CommandLineOptionsGroups {
		_, err := parser.AddGroup(optsGroup.ShortDescription, optsGroup.LongDescription, optsGroup.Options)
		if err != nil {
			log.Fatalln(err)
		}
	}

	if _, err := parser.Parse(); err != nil {
		code := 1
		if fe, ok := err.(*flags.Error); ok {
			if fe.Type == flags.ErrHelp {
				code = 0
			}
		}
		os.Exit(code)
	}

	server.ConfigureAPI()

	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}

}
