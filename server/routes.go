package main

import(
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type tool struct {
	Owner string
	Name string
}

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

func D(res http.ResponseWriter, req *http.Request) {
	setCors(res)

	conor := tool{
		Owner: "contr0n",
		Name: "hammer",
	}

	js, err := json.Marshal(conor)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	os.Stdout.Write(js)
	res.Header().Set("Content-Type", "application/json")
	res.Write(js)
}

func myPrint(){
	fmt.Println("my package printer")
}

