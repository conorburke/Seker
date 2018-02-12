package main

import(
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// util
func getFrontendUrl() string {
	if os.Getenv("APP_ENV") == "production" {
		return "http://localhost:3000" // change this to production domain
	} else {
		return "http://localhost:3000"
	}
}

func setCors(w http.ResponseWriter) {
	frontendUrl := getFrontendUrl()
	w.Header().Set("Access-Control-Allow-Origin", frontendUrl)
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func indexHandler(res http.ResponseWriter, req *http.Request) {
	setCors(res)

	testUser := User{firstName: "conor", lastName: "burke", email: "cjburke89@gmail.com"}
	
	testTool := Tool{Owner: testUser, Name: "Mjolnir"}

	js, err := json.Marshal(testTool)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	os.Stdout.Write(js)
	res.Header().Set("Content-Type", "application/json")
	res.Write(js)
}

func toolsIndexHandler(res http.ResponseWriter, req *http.Request) {
	setCors(res)
	var tools []Tool
	DB.Find(&tools)
	js, err := json.Marshal(tools)
	if err != nil {
		http.Error(res, err.Error(), 500)
		return
	}

	res.Write(js)
}

func myPrint(){
	fmt.Println("my package printer")
}

