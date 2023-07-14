package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

type form_submission struct {
	Username    string `json:"username" form:"username"`
	Password    string `json:"password" form:"password"`
	NewUsername string `json:"newUsername" form:"newUsername"`
	NewPassword string `json:"newPassword" form:"newPassword"`
	StudentID   string `json:"studentID" form:"studentID"`
}

func startServer() {
	router := gin.Default()

	// Define the routes for the API Gateway
	router.POST("/submit-form", handleFormSubmission)
	router.Run(":8080")
}

type Response struct {
	Message string `json:"message"`
}

func main() {
	startServer()
}

func handleFormSubmission(c *gin.Context) {
	fmt.Println("req")
	var r form_submission
	err := c.BindWith(&r, binding.FormMultipart)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	response := Response{
		Message: "Form data received by Go server",
	}

	// Convert response to JSON
	/*jsonResponse, err := json.Marshal(response)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}*/

	// Set the appropriate headers
	c.Header("Content-Type", "application/json")
	c.Header("Access-Control-Allow-Origin", "*")

	// Write the response content
	c.JSON(http.StatusCreated, response)
	if err != nil {
		return
	}
}
