package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func HandleRoot(writer http.ResponseWriter, router *http.Request) {
	fmt.Fprintf(writer, "Hello Root")
}

func HandleHome(writer http.ResponseWriter, router *http.Request) {
	fmt.Fprintf(writer, "Hello Home")
}

func UserPostRequest(writer http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	var user User
	err := decoder.Decode(&user)
	if err != nil {
		fmt.Fprintf(writer, "error: %v", err)
		return
	}
	response, err := user.ToJson()
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(response)
}

func PostRequest(writer http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	var metadata Metadata
	err := decoder.Decode(&metadata)
	if err != nil {
		fmt.Fprintf(writer, "error: %v", err)
		return
	}
	fmt.Fprintf(writer, "Payload %v\n", metadata)
}
