package main

import (
	"database"
	"log"
	"mytest"
	"os"
)

func main() {
	defer database.DB.Close()

	// add database
	_, err := database.Init()
	if err != nil {
		log.Println("connection to DB failed, aborting...")
		log.Fatal(err)
	}

	mytest.Testmeout()

	log.Println("connected to DB")

	// print env
	env := os.Getenv("APP_ENV")
	if env == "production" {
		log.Println("Running api server in production mode")
	} else {
		log.Println("Running api server in dev mode")
	}

	// http.ListenAndServe(":8080", router)
}