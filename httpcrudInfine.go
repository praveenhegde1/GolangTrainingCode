package main

import (
	"encoding/json"
	"log"
	"net/http"

	"os"
	"runtime/pprof"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

var users []User

func main() {

	
	f,err := os.Create("cpu.prof")
	if err != nil {	
	
		log.Fatal("could not create CPU profile: ", err)
	}
	defer f.Close()

	if err := pprof.StartCPUProfile(f); err!= nil {	
		log.Fatal("could not start CPU profile: ", err)
    }
	defer pprof.StopCPUProfile()
	
	http.HandleFunc("/users", getUsers)
	http.HandleFunc("/users/add", addUser)
	http.HandleFunc("/users/update", updateUser)
	http.HandleFunc("/users/delete", deleteUser)

	log.Println(http.ListenAndServe(":8080", nil))
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(users)
	
}

func addUser(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	for i :=0; i<100000000; i++ {
	users = append(users, user)
	json.NewEncoder(w).Encode(user)
	}
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	var updatedUser User
	err := json.NewDecoder(r.Body).Decode(&updatedUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for i, user := range users {
		if user.ID == updatedUser.ID {
			users[i] = updatedUser
			break
		}
	}

	json.NewEncoder(w).Encode(updatedUser)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	var userToDelete User
	err := json.NewDecoder(r.Body).Decode(&userToDelete)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for i, user := range users {
		if user.ID == userToDelete.ID {
			users = append(users[:i], users[i+1:]...)
			break
		}
	}

	json.NewEncoder(w).Encode(userToDelete)
}
