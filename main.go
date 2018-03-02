package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)
func main() {
	// Open up our database connection.
	//db, err := sql.Open("mysql", "user:password@tcp(hostIP:port)/database")
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/auxilium_db")

	// if there is an error opening the connection, handle it
	if err != nil {
		log.Print(err.Error())
	}
	
	// Execute the query
	results, err := db.Query("SELECT * FROM users")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	for results.Next() {
		var user User
		// for each row, scan the result into our tag composite object
		err = results.Scan(&user.ID, &user.lastName, &user.firstName, &user.identificationDocument, &user.email, &user.passwd)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// and then print out the tag's Name attribute
		user.showUserInfo()
	}

	// defer the close till after the main function has finished
	// executing
	defer db.Close()
}