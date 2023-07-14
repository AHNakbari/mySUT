package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func main() {
	http.HandleFunc("/submit-form", handleFormSubmission)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleFormSubmission(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form data", http.StatusBadRequest)
		return
	}
	log.Println(r)
	// Access the form field values
	username := r.Form.Get("username")
	password := r.Form.Get("password")
	newUsername := r.Form.Get("newUsername")
	newPassword := r.Form.Get("newPassword")
	repeatPassword := r.Form.Get("repeatPassword")
	studentID := r.Form.Get("studentID")
	log.Println(r.Form)

	log.Println(username)
	log.Println(password)
	log.Println(newUsername)
	log.Println(newPassword)
	log.Println(repeatPassword)
	log.Println(studentID)

	response := Response{
		Message: "Form data received by Go server",
	}
	// Convert response to JSON
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Failed to create response", http.StatusInternalServerError)
		return
	}

	// Set the appropriate headers
	w.Header().Set("Content-Type", "application/json")

	// Write the response content
	_, err = w.Write(jsonResponse)
	if err != nil {
		return
	}
}
