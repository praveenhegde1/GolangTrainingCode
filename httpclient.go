package main

import (
	"bytes"
	"encoding/json"
	
	"fmt"
	"log"
	"net/http"
	"os"
	
	"runtime/pprof"
)
http.
type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}


func main() {

	f,err := os.Create("cpuclient.prof")
	if err != nil {	
	
		log.Fatal("could not create CPU profile: ", err)
	}
	defer f.Close()

	if err := pprof.StartCPUProfile(f); err!= nil {	
		log.Fatal("could not start CPU profile: ", err)
    }
	defer pprof.StopCPUProfile()

	newUser := User{
		Name:     "Praveen Hegde",
		Email:    "praveen@saastech.com",
		Password: "password123",
	}
	createUser(newUser)

	userID := 1
	getUser(userID)

	updatedUser := User{
		ID:       1,
		Name:     "Praveen H",
		Email:    "praveen@fudotech.in",
		Password: "newpassword",
	}
	updateUser(updatedUser)

	deleteUser(userID)

	
}

func createUser(user User) {
	jsonData, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	resp, err := http.Post("http://localhost:8080/users", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating user:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("User created successfully")
}

func getUser(userID int) {
	resp, err := http.Get(fmt.Sprintf("http://localhost:8080/users/%d", userID))
	if err != nil {
		fmt.Println("Error getting user:", err)
		return
	}
	defer resp.Body.Close()

	var user User
	err = json.NewDecoder(resp.Body).Decode(&user)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	fmt.Println("User:", user)
}

func updateUser(user User) {
	jsonData, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("http://localhost:8080/users/update/%d", user.ID), bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error updating user:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("User updated successfully")
}

func deleteUser(userID int) {
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("http://localhost:8080/users/delete/%d", userID), nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error deleting user:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("User deleted successfully")
}
