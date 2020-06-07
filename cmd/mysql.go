package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/johnfercher/graph-study/graph/entity"
	"log"
)

func main() {
	// Open up our database connection.
	db, err := sql.Open("mysql", "AdminUser:AdminPassword@tcp(127.0.0.1:3307)/Vertex")

	// if there is an error opening the connection, handle it
	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()

	// Execute the query
	results, err := db.Query(`SELECT id, name FROM vertex WHERE id = "c37e04a4-01e6-4b6f-bb34-94bc60dd1495"`)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	for results.Next() {
		var vertex entity.Vertex
		// for each row, scan the result into our tag composite object
		err = results.Scan(&vertex.Id, &vertex.Name)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		// and then print out the tag's Name attribute
		log.Printf(vertex.Name)
	}

}
