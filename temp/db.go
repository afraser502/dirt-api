package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Tag struct {
	Name string `json:"name"`
}

type Image struct {
	Repository string `json:"repository"`
	Id         string `json:"id"`
	Tag        string `json:"tag"`
	Size       string `json:"size"`
	Created    string `json:"created"`
	Requestor  string `json:"requestor"`
	Status     string `json:"status"`
}

var db string

func DbConnect() {

	fmt.Println("Go MySQL Tutorial")

	// Open up our database connection.
	// I've set up a database on my local machine using phpmyadmin.
	// The database is called testDb
	db, err := sql.Open("mysql", "root:password@tcp(172.28.128.3:3307)/mysql-db")

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	// defer the close till after the main function has finished
	// executing
	defer db.Close()

	// perform a db.Query insert
	/*insert, err := db.Query("INSERT INTO user VALUES ( 'af2' )")

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}
	// be careful deferring Queries if you are using transactions
	defer insert.Close()

	fmt.Println("successfully added user")*/

	// Execute the query
	results, err := db.Query("SELECT username FROM user")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	for results.Next() {
		var tag Tag
		// for each row, scan the result into our tag composite object
		err = results.Scan(&tag.Name)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// and then print out the tag's Name attribute
		fmt.Println(tag.Name)
	}
}

func DbImagesList() {

	db, err := sql.Open("mysql", "root:password@tcp(172.28.128.3:3307)/mysql-db")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	// Execute the query
	results, err := db.Query("SELECT repository, id, tag, size, created, requestor, status FROM images")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	for results.Next() {
		var image Image
		// for each row, scan the result into our tag composite object
		err = results.Scan(&image.Repository, &image.Id, &image.Tag, &image.Size, &image.Created, &image.Requestor, &image.Status)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// and then print out the tag's Name attribute
		fmt.Println(image.Repository, image.Id, image.Tag, image.Size, image.Created, image.Requestor, image.Status)
	}

}
