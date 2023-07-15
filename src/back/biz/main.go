package main

import (
	"database/sql"
	// "errors"
	// "fmt"
	"log"
	"net"

	// "time"

	// "sync"

	v1 "sina/pb"

	// "github.com/jackc/pgx/pgtype"
	// "github.com/jackc/pgx/v5/pgtype"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

const (
	port         = "5432"
	dbConnection = "host=localhost port=5432 user=postgres password=password dbname=users sslmode=disable"
)

type server struct {
	v1.UnimplementedDatabaseServiceServer
	db *sql.DB
}

func (s *server) GetUser(ctx context.Context, req *v1.UserRequest) (*v1.UserResponse, error) {
	// Retrieve the user_id from the request
	userID := req.GetUserId()

	// Prepare the SQL statement for selecting a user by user_id
	stmt, err := s.db.Prepare("SELECT user_id, name, number, password, reshte, vorudi, courses, groups, role FROM users WHERE user_id = $1")
	if err != nil {
		log.Fatalf("Failed to prepare SQL statement: %v", err)
	}
	defer stmt.Close()

	// Execute the SQL statement to select the user by user_id
	row := stmt.QueryRow(userID)

	// Variables to store the retrieved user details
	var user_id string
	var name string
	var number int32
	var password string
	var reshte string
	var vorudi string
	var courses []string
	var groups []string
	var role int32
	// Scan the retrieved row into the variables
	err = row.Scan(&user_id, &name, &number, &password, &reshte, &vorudi, pq.Array(&courses), pq.Array(&groups), &role)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, "User not found")
		}
		log.Fatalf("Failed to retrieve user details: %v", err)
	}

	// Create and return the UserResponse with the retrieved user details
	userr := &v1.User{
		UserId:   user_id,
		Name:     name,
		Number:   number,
		Password: password,
		Reshte:   reshte,
		Vorudi:   vorudi,
		Courses:  courses,
		Groups:   groups,
		Role:     role,
	}
	resp := &v1.UserResponse{
		User: userr,
	}

	return resp, nil
}

func (s *server) SendUser(ctx context.Context, req *v1.UserResponse) (*v1.UserRequest, error) {
	User := req.GetUser()
	userId := User.GetUserId()
	name := User.GetName()
	number := User.GetNumber()
	password := User.GetPassword()
	reshte := User.GetReshte()
	vorudi := User.GetVorudi()
	courses := User.GetCourses()
	groups := User.GetGroups()
	role := User.GetRole()

	// Prepare the SQL statement for inserting a new row into the users table
	stmt, err := s.db.Prepare("INSERT INTO users(user_id, name, number, password, reshte, vorudi, courses, groups, role) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9)")
	if err != nil {
		log.Fatalf("Failed to prepare SQL statement: %v", err)
	}
	defer stmt.Close()

	// Execute the SQL statement with the user details
	_, err = stmt.Exec(userId, name, number, password, reshte, vorudi, pq.Array(courses), pq.Array(groups), role)
	if err != nil {
		log.Fatalf("Failed to execute SQL statement: %v", err)
	}

	// Create and return the UserResponse
	resp := &v1.UserRequest{
		UserId: userId,
	}

	return resp, nil
}

func (s *server) GetGroup(ctx context.Context, req *v1.GroupRequest) (*v1.GroupResponse, error) {
	// Retrieve the user_id from the request
	namee := req.GetName()

	// Prepare the SQL statement for selecting a user by user_id
	stmt, err := s.db.Prepare("SELECT name, subgroups, courses, members, owner, news FROM groups WHERE name = $1")
	if err != nil {
		log.Fatalf("Failed to prepare SQL statement: %v", err)
	}
	defer stmt.Close()

	// Execute the SQL statement to select the user by user_id
	row := stmt.QueryRow(namee)

	// Variables to store the retrieved user details
	var name string
	var subgroups []string
	var courses []string
	var members []string
	var owner string
	var news []string
	// Scan the retrieved row into the variables
	err = row.Scan(&name, pq.Array(&subgroups), pq.Array(&courses), pq.Array(&members), &owner, pq.Array(&news))
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, "User not found")
		}
		log.Fatalf("Failed to retrieve user details: %v", err)
	}

	// Create and return the UserResponse with the retrieved user details
	groupp := &v1.Group{
		Name:      name,
		Subgroups: subgroups,
		Courses:   courses,
		Members:   members,
		Owner:     owner,
		News:      news,
	}
	resp := &v1.GroupResponse{
		Group: groupp,
	}

	return resp, nil
}

func (s *server) SendGroup(ctx context.Context, req *v1.GroupResponse) (*v1.GroupRequest, error) {
	Group := req.GetGroup()
	name := Group.GetName()
	subgroups := Group.GetSubgroups()
	courses := Group.GetCourses()
	members := Group.GetMembers()
	owner := Group.GetOwner()
	news := Group.GetNews()

	// Prepare the SQL statement for inserting a new row into the users table
	stmt, err := s.db.Prepare("INSERT INTO groups(name, subgroups, courses, members, owner, news) VALUES($1, $2, $3, $4, $5, $6)")
	if err != nil {
		log.Fatalf("Failed to prepare SQL statement: %v", err)
	}
	defer stmt.Close()

	// Execute the SQL statement with the user details
	_, err = stmt.Exec(name, pq.Array(subgroups), pq.Array(courses), pq.Array(members), owner, pq.Array(news))
	if err != nil {
		log.Fatalf("Failed to execute SQL statement: %v", err)
	}

	// Create and return the UserResponse
	resp := &v1.GroupRequest{
		Name: name,
	}

	return resp, nil
}

func (s *server) GetSubgroup(ctx context.Context, req *v1.SubRequest) (*v1.SubResponse, error) {
	// Retrieve the user_id from the request
	namee := req.GetName()

	// Prepare the SQL statement for selecting a user by user_id
	stmt, err := s.db.Prepare("SELECT name, members, courses, owner FROM subgroups WHERE name = $1")
	if err != nil {
		log.Fatalf("Failed to prepare SQL statement: %v", err)
	}
	defer stmt.Close()

	// Execute the SQL statement to select the user by user_id
	row := stmt.QueryRow(namee)

	// Variables to store the retrieved user details
	var name string
	var members []string
	var courses []string
	var owner string
	// Scan the retrieved row into the variables
	err = row.Scan(&name, pq.Array(&members), pq.Array(&courses), &owner)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, "User not found")
		}
		log.Fatalf("Failed to retrieve user details: %v", err)
	}

	// Create and return the UserResponse with the retrieved user details
	subgroupp := &v1.Sub{
		Name:    name,
		Members: members,
		Courses: courses,
		Owner:   owner,
	}
	resp := &v1.SubResponse{
		Sub: subgroupp,
	}

	return resp, nil
}

func (s *server) SendSubgroup(ctx context.Context, req *v1.SubResponse) (*v1.SubRequest, error) {
	Subgroup := req.GetSub()
	name := Subgroup.GetName()
	courses := Subgroup.GetCourses()
	members := Subgroup.GetMembers()
	owner := Subgroup.GetOwner()

	// Prepare the SQL statement for inserting a new row into the users table
	stmt, err := s.db.Prepare("INSERT INTO subgroups(name, members, courses, owner) VALUES($1, $2, $3, $4)")
	if err != nil {
		log.Fatalf("Failed to prepare SQL statement: %v", err)
	}
	defer stmt.Close()

	// Execute the SQL statement with the user details
	_, err = stmt.Exec(name, pq.Array(members), pq.Array(courses), owner)
	if err != nil {
		log.Fatalf("Failed to execute SQL statement: %v", err)
	}

	// Create and return the UserResponse
	resp := &v1.SubRequest{
		Name: name,
	}

	return resp, nil
}

func (s *server) GetCourse(ctx context.Context, req *v1.CourseRequest) (*v1.CourseResponse, error) {
	// Retrieve the user_id from the request
	namee := req.GetName()

	// Prepare the SQL statement for selecting a user by user_id
	stmt, err := s.db.Prepare("SELECT name, exercises, members, owner FROM courses WHERE name = $1")
	if err != nil {
		log.Fatalf("Failed to prepare SQL statement: %v", err)
	}
	defer stmt.Close()

	// Execute the SQL statement to select the user by user_id
	row := stmt.QueryRow(namee)

	// Variables to store the retrieved user details
	var name string
	var exercises []string
	var members []string
	var owner string
	// Scan the retrieved row into the variables
	err = row.Scan(&name, pq.Array(&exercises), pq.Array(&members), &owner)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, "User not found")
		}
		log.Fatalf("Failed to retrieve user details: %v", err)
	}

	// Create and return the UserResponse with the retrieved user details
	coursee := &v1.Course{
		Name:      name,
		Exercises: exercises,
		Members:   members,
		Owner:     owner,
	}
	resp := &v1.CourseResponse{
		Course: coursee,
	}

	return resp, nil
}

func (s *server) SendCourse(ctx context.Context, req *v1.CourseResponse) (*v1.CourseRequest, error) {
	Course := req.GetCourse()
	name := Course.GetName()
	exercises := Course.GetExercises()
	members := Course.GetMembers()
	owner := Course.GetOwner()

	// Prepare the SQL statement for inserting a new row into the users table
	stmt, err := s.db.Prepare("INSERT INTO courses(name, exercises, members, owner) VALUES($1, $2, $3, $4)")
	if err != nil {
		log.Fatalf("Failed to prepare SQL statement: %v", err)
	}
	defer stmt.Close()

	// Execute the SQL statement with the user details
	_, err = stmt.Exec(name, pq.Array(exercises), pq.Array(members), owner)
	if err != nil {
		log.Fatalf("Failed to execute SQL statement: %v", err)
	}

	// Create and return the UserResponse
	resp := &v1.CourseRequest{
		Name: name,
	}

	return resp, nil
}

func (s *server) DeleteUser(ctx context.Context, req *v1.UserRequest) (*v1.DeleteResponse, error) {
	// Retrieve the user_id from the request
	userID := req.GetUserId()

	// Prepare the SQL statement for deleting a user by user_id
	stmt, err := s.db.Prepare("DELETE FROM users WHERE user_id = $1")
	if err != nil {
		log.Fatalf("Failed to prepare SQL statement: %v", err)
	}
	defer stmt.Close()

	// Execute the SQL statement to delete the user by user_id
	result, err := stmt.Exec(userID)
	if err != nil {
		log.Fatalf("Failed to execute SQL statement: %v", err)
	}

	// Check the affected rows count to determine if the deletion was successful
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Fatalf("Failed to retrieve rows affected count: %v", err)
	}

	success := int32(0)
	if rowsAffected > 0 {
		success = int32(1)
	}
	// Create and return the UserDeleteResponse with the deletion success status
	resp := &v1.DeleteResponse{
		IsTrue: success,
	}

	return resp, nil
}

func (s *server) DeleteGroup(ctx context.Context, req *v1.GroupRequest) (*v1.DeleteResponse, error) {
	// Retrieve the user_id from the request
	name := req.GetName()

	// Prepare the SQL statement for deleting a user by user_id
	stmt, err := s.db.Prepare("DELETE FROM groups WHERE name = $1")
	if err != nil {
		log.Fatalf("Failed to prepare SQL statement: %v", err)
	}
	defer stmt.Close()

	// Execute the SQL statement to delete the user by user_id
	result, err := stmt.Exec(name)
	if err != nil {
		log.Fatalf("Failed to execute SQL statement: %v", err)
	}

	// Check the affected rows count to determine if the deletion was successful
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Fatalf("Failed to retrieve rows affected count: %v", err)
	}

	success := int32(0)
	if rowsAffected > 0 {
		success = int32(1)
	}
	// Create and return the UserDeleteResponse with the deletion success status
	resp := &v1.DeleteResponse{
		IsTrue: success,
	}

	return resp, nil
}

func (s *server) DeleteSubgroup(ctx context.Context, req *v1.SubRequest) (*v1.DeleteResponse, error) {
	// Retrieve the user_id from the request
	name := req.GetName()

	// Prepare the SQL statement for deleting a user by user_id
	stmt, err := s.db.Prepare("DELETE FROM subgroups WHERE name = $1")
	if err != nil {
		log.Fatalf("Failed to prepare SQL statement: %v", err)
	}
	defer stmt.Close()

	// Execute the SQL statement to delete the user by user_id
	result, err := stmt.Exec(name)
	if err != nil {
		log.Fatalf("Failed to execute SQL statement: %v", err)
	}

	// Check the affected rows count to determine if the deletion was successful
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Fatalf("Failed to retrieve rows affected count: %v", err)
	}

	success := int32(0)
	if rowsAffected > 0 {
		success = int32(1)
	}
	// Create and return the UserDeleteResponse with the deletion success status
	resp := &v1.DeleteResponse{
		IsTrue: success,
	}

	return resp, nil
}

func (s *server) DeleteCourse(ctx context.Context, req *v1.CourseRequest) (*v1.DeleteResponse, error) {
	// Retrieve the user_id from the request
	name := req.GetName()

	// Prepare the SQL statement for deleting a user by user_id
	stmt, err := s.db.Prepare("DELETE FROM courses WHERE name = $1")
	if err != nil {
		log.Fatalf("Failed to prepare SQL statement: %v", err)
	}
	defer stmt.Close()

	// Execute the SQL statement to delete the user by user_id
	result, err := stmt.Exec(name)
	if err != nil {
		log.Fatalf("Failed to execute SQL statement: %v", err)
	}

	// Check the affected rows count to determine if the deletion was successful
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Fatalf("Failed to retrieve rows affected count: %v", err)
	}

	success := int32(0)
	if rowsAffected > 0 {
		success = int32(1)
	}
	// Create and return the UserDeleteResponse with the deletion success status
	resp := &v1.DeleteResponse{
		IsTrue: success,
	}

	return resp, nil
}

func (s *server) GetAllUsers(ctx context.Context, req *v1.GetAllUsersRequest) (*v1.GetAllUsersResponse, error) {
	// Prepare the SQL statement for selecting all user_ids from the users table
	stmt, err := s.db.Prepare("SELECT user_id FROM users")
	if err != nil {
		log.Fatalf("Failed to prepare SQL statement: %v", err)
	}
	defer stmt.Close()

	// Execute the SQL statement to retrieve all user_ids
	rows, err := stmt.Query()
	if err != nil {
		log.Fatalf("Failed to execute SQL statement: %v", err)
	}
	defer rows.Close()

	// Slice to store the retrieved user_ids
	var userIDs []string

	// Iterate over the result rows and append the user_ids to the slice
	for rows.Next() {
		var userID string
		err := rows.Scan(&userID)
		if err != nil {
			log.Fatalf("Failed to scan user_id: %v", err)
		}
		userIDs = append(userIDs, userID)
	}
	if err = rows.Err(); err != nil {
		log.Fatalf("Error iterating over rows: %v", err)
	}

	// Create and return the GetAllUserIDsResponse with the retrieved user_ids
	resp := &v1.GetAllUsersResponse{
		UserIds: userIDs,
	}

	return resp, nil
}

func main() {
	db, err := sql.Open("postgres", dbConnection)
	if err != nil {
		log.Fatalf("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa %v", err)
	}

	defer db.Close()
	lis, err := net.Listen("tcp", ":5062")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	v1.RegisterDatabaseServiceServer(s, &server{db: db})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
