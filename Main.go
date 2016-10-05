package main

import (
	"log"
	"net/http"

	routers "github.com/SivaShhankar/CMS_1/Routers"
)

func main() {

	// Load the configuration stuffs.
	//config.LoadAppConfig()

	// Initiate the database information.
	//config.CreateDBSession()

	// Add the neccessary indexes.
	//config.AddIndexes()

	// Created the routes of this application/
	mux := http.NewServeMux()
	mux = routers.SetCandidateRoutes(mux)

	log.Println("Listening...")

	// Listen the server.
	http.ListenAndServe(":8080", mux)
}

// // GetPort -- get the Port from the Dynamic environment
// func GetPort() string {
// 	var port = os.Getenv("PORT")
// 	if port == "" {
// 		port = "4747"
// 		fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
// 	}

// 	fmt.Println("Running Port is" + port)
// 	return ":" + port
// }
