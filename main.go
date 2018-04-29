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

	sw "github.com/aeckard87/WornOut/rest"
	"github.com/rs/cors"
)

func main() {
	log.Printf("Server started")

	router := sw.NewRouter()
	handler := cors.Default().Handler(router)
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://10.0.0.13:8100"},
		AllowCredentials: true,
		// Enable Debugging for testing, consider disabling in production
		Debug:              true,
		AllowedMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		OptionsPassthrough: true,
	})

	// Insert the middleware
	handler = c.Handler(handler)

	log.Fatal(http.ListenAndServe(":8081", handler))
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
