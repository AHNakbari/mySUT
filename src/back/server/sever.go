package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"google.golang.org/grpc"
	"log"
	pb "mysut/server/pb"
	"net/http"
	"regexp"
	"strconv"
	"time"
)

var c2 pb.DatabaseServiceClient

type form_submission struct {
	Username    string `json:"username" form:"username"`
	Password    string `json:"password" form:"password"`
	NewUsername string `json:"newUsername" form:"newUsername"`
	NewPassword string `json:"newPassword" form:"newPassword"`
	StudentID   string `json:"studentID" form:"studentID"`
}

func startServer() {
	conn, err := grpc.Dial("localhost:5062", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)
	c2 = pb.NewDatabaseServiceClient(conn)
	router := gin.Default()
	// Define the routes for the API Gateway
	router.POST("/create-user", handleCreateUser)
	router.POST("/login", handleLogin)
	router.POST("/get-user", handleGetUser)
	router.Run(":8080")
}

type Response struct {
	Message   string   `json:"message"`
	Username  string   `json:"username"`
	StudentID int32    `json:"number"`
	Role      int32    `json:"role"`
	Groups    []string `json:"groups"`
	Courses   []string `json:"courses"`
	Field     string   `json:"field"`
	Year      string   `json:"year"`
}

func main() {
	startServer()
}

func handleLogin(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.Header("Access-Control-Allow-Origin", "*")
	var r form_submission
	err := c.ShouldBindWith(&r, binding.FormMultipart)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	response := Response{
		Message: "Form data received by Go server",
	}
	if !usernameCheck(r.Username) {
		response = Response{Message: "username format is not correct"}
		c.JSON(http.StatusBadRequest, response)
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	replyMsg, err := c2.GetUser(ctx, &pb.UserRequest{
		UserId: r.Username,
	})
	if replyMsg == nil || r.Password != replyMsg.User.Password {
		response = Response{Message: "user name or password are wrong"}
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response = Response{Message: "login successful", Username: replyMsg.User.UserId, Role: replyMsg.User.Role, Groups: replyMsg.User.Groups,
		Courses: replyMsg.User.Courses, Field: replyMsg.User.Reshte, Year: replyMsg.User.Vorudi}
	c.JSON(http.StatusAccepted, response)
	return

}

func handleCreateUser(c *gin.Context) {
	// Set the appropriate headers
	c.Header("Content-Type", "application/json")
	c.Header("Access-Control-Allow-Origin", "*")
	var r form_submission
	err := c.ShouldBindWith(&r, binding.FormMultipart)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	response := Response{
		Message: "Form data received by Go server",
	}
	if !passwordCheck(r.NewPassword) {
		response = Response{Message: "password format is not correct"}
		c.JSON(http.StatusBadRequest, response)
		return
	}
	if !usernameCheck(r.NewUsername) {
		response = Response{Message: "username format is not correct"}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	replyMsg, err := c2.GetUser(ctx, &pb.UserRequest{
		UserId: r.NewUsername,
	})
	if err != nil {
		fmt.Println(err)
	}
	if replyMsg != nil {
		response = Response{Message: "this username is taken"}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	id, err := strconv.Atoi(r.StudentID)
	if err != nil {
		response = Response{Message: "studentId format is not correct"}
		c.JSON(http.StatusBadRequest, response)
		return
	}
	_, err = c2.SendUser(ctx, &pb.UserResponse{
		User: &pb.User{UserId: r.NewUsername, Password: r.NewPassword, Number: int32(id), Role: 3},
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	response = Response{Message: "user created"}
	c.JSON(http.StatusCreated, response)
	if err != nil {
		return
	}
}
func handleGetUser(c *gin.Context) {
	// Set the appropriate headers
	c.Header("Content-Type", "application/json")
	c.Header("Access-Control-Allow-Origin", "*")
	var r form_submission
	err := c.ShouldBindWith(&r, binding.FormMultipart)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	response := Response{
		Message: "Form data received by Go server",
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	replyMsg, err := c2.GetUser(ctx, &pb.UserRequest{
		UserId: r.Username,
	})
	if err != nil {
		fmt.Println(err)
	}
	response = Response{Message: "login successful", Username: replyMsg.User.UserId, Role: replyMsg.User.Role, Groups: replyMsg.User.Groups,
		Courses: replyMsg.User.Courses, Field: replyMsg.User.Reshte, Year: replyMsg.User.Vorudi, StudentID: replyMsg.User.Number}
	c.JSON(http.StatusAccepted, response)
	if err != nil {
		return
	}
}
func passwordCheck(pass string) bool {
	return len(pass) >= 1
}
func usernameCheck(username string) bool {
	return regexp.MustCompile(`^[a-zA-Z0-9_.-]*$`).MatchString(username)
}
