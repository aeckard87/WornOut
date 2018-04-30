package main

import (
	// WARNING!
	// Change this to a fully-qualified import path
	// once you place this file into your project.
	// For example,
	//
	//    sw "github.com/myname/myrepo/go"
	//
	"log"
	"net/http"
	"os"

	sw "github.com/aeckard87/WornOut/rest"
	"github.com/rs/cors"
)

func main() {
	log.Printf("Server started")
	origin := os.Getenv("WORN_OUT_ALLOWED_ORIGIN")
	allowedPort := os.Getenv("WORN_OUT_ALLOWED_PORT")

	router := sw.NewRouter()
	handler := cors.Default().Handler(router)
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://" + origin + ":" + allowedPort},
		AllowCredentials: true,
		// Enable Debugging for testing, consider disabling in production
		Debug:              true,
		AllowedMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		OptionsPassthrough: true,
	})

	// Insert the middleware
	handler = c.Handler(handler)

	port := os.Getenv("WORN_OUT_PORT")
	// log.Fatal(http.ListenAndServe(":8081", handler))
	log.Fatal(http.ListenAndServe(":"+port, handler))
}

// package main

// import (
// 	// WARNING!
// 	// Change this to a fully-qualified import path
// 	// once you place this file into your project.
// 	// For example,
// 	//
// 	//    sw "github.com/myname/myrepo/go"
// 	//
// 	"log"
// 	"net/http"

// 	"github.com/julienschmidt/httprouter"
// )

// func main() {
// 	log.Printf("Server started")

// 	router := httprouter.New()

// 	log.Fatal(http.ListenAndServe(":8081", router))
// }
