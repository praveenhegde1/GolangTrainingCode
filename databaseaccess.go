package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	dbHost := "localhost"
	dbPort := 3306
	dbUser := "root"
	dbPassword := "welcome"
	dbName := "golang"

	//connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

	fmt.Println(connStr)
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		fmt.Println("Failed to connect to the database:", err)
		return
	}
	defer db.Close()

	rows, err := db.Query("INSERT INTO goex (id, name, email, password) VALUES (1004, 'praveen', 'pra@saastech.com', 'welcome')")
	if err != nil {
		fmt.Println("Failed to execute the query:", err)
		return
	}
	defer rows.Close()
/*
	for rows.Next() {
		var (
			column1 int
			column2 string
			column3 string
			column4 int
		)
		err := rows.Scan(&column1, &column2, &column3, &column4)
		if err != nil {
			fmt.Println("Failed to scan row:", err)
			return
		}

		fmt.Println(column1, column2)
	}

	if err = rows.Err(); err != nil {
		fmt.Println("Error occurred during iteration:", err)
		return
	}*/
}
