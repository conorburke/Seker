package main

import (
	"fmt"
	"net/http"
	"log"
	"mytest"
	"os"
)

func main() {
	defer DB.Close()

	// add database
	_, err := Init()
	if err != nil {
		log.Println("connection to DB failed, aborting...")
		log.Fatal(err)
	}

	myPrint()

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
	port := "8080"
	fmt.Println("server up and running on port", port)

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/tools", toolsIndexHandler)
	
	http.ListenAndServe(":"+port, nil)
}
