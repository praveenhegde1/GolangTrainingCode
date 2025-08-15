package main

import (
	"encoding/json"
	"log"
	"net/http"
	_ "net/http/pprof"
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"

)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

var users []User
var connStr string
func main() {
	dbHost := "localhost"
	dbPort := 3306
	dbUser := "root"
	dbPassword := "welcome"
	dbName := "golang"

	connStr = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)


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
	fmt.Println(user.ID)
	fmt.Println(user.Name)
	fmt.Println(user.Email)
	fmt.Println(user.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println(connStr)
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		fmt.Println("Failed to connect to the database:", err)
		return
	}
	defer db.Close()
	
	query := `INSERT INTO goex (id, name, email,password) VALUES (?, ?, ?,?)`
	_, err = db.Exec(query, user.ID, user.Name, user.Email, user.Password)
	if err != nil {
		fmt.Println("Failed to execute the query:", err)
		return
	}
	defer rows.Close()
	users = append(users, user)
	json.NewEncoder(w).Encode(user)
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
