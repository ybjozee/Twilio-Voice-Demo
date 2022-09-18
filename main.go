package main

import (
	"encoding/json"
	"fmt"
    "github.com/joho/godotenv"
	"log"
	"net/http"
	"voice_demo/helper"
	"voice_demo/model"
)

func main() {
	loadEnv()
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/call", callHandler)

	fmt.Printf("Starting server at port 8000\n")

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}

func loadEnv() {
  err := godotenv.Load(".env.local")
  if err != nil {
    log.Fatal("Error loading .env file")
  }
}

func callHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	response := make(map[string]string)

	var phoneNumber model.PhoneNumber
	json.NewDecoder(request.Body).Decode(&phoneNumber)

	err := phoneNumber.Validate()

	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		response["message"] = err.Error()
	} else {
		res, err := helper.Call(phoneNumber.Value)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			response["message"] = err.Error()
		} else {
			writer.WriteHeader(http.StatusOK)
			response["message"] = res
		}
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}

	writer.Write(jsonResponse)
}
