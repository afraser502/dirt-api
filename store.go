package main

import (
	"database/sql"
	"fmt"
)

//Store interfacte for functions
type Store interface {
	CreateDownload(images *Image) error
	GetImageDownloads() ([]*Image, error)
}

type dbStore struct {
	db *sql.DB
}

func (store *dbStore) CreateDownload(image *Image) error {
	// 'Bird' is a simple struct which has "species" and "description" attributes
	// THe first underscore means that we don't care about what's returned from
	// this insert query. We just want to know if it was inserted correctly,
	// and the error will be populated if it wasn't
	//_, err := store.db.Exec("INSERT INTO images (repository, id, tag, size, created, requestor, status) VALUES (repository, id, tag, size, created, requestor, status)", image.Repository, image.Id, image.Tag, image.Size, image.Created, image.Requestor, image.Status)
	stmt, err := store.db.Prepare("INSERT images SET repository=?, id=?, tag=?, size=?, created=?, requestor=?, status=?")
	res, err := stmt.Exec(image.Repository, image.Id, image.Tag, image.Size, image.Created, image.Requestor, image.Status)

	if err != nil {
		panic(err)
	}

	fmt.Println(res)

	return err
}

func (store *dbStore) GetImageDownloads() ([]*Image, error) {
	// Query the database for all birds, and return the result to the
	// `rows` object
	rows, err := store.db.Query("SELECT repository, id, tag, size, created, requestor, status FROM images")
	// We return incase of an error, and defer the closing of the row structure
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Create the data structure that is returned from the function.
	// By default, this will be an empty array of birds
	imgs := []*Image{}
	for rows.Next() {
		// For each row returned by the table, create a pointer to a bird,
		img := &Image{}
		// Populate the `Species` and `Description` attributes of the bird,
		// and return incase of an error
		if err := rows.Scan(&img.Repository, &img.Id, &img.Tag, &img.Size, &img.Created, &img.Requestor, &img.Status); err != nil {
			return nil, err
		}
		// Finally, append the result to the returned array, and repeat for
		// the next row
		imgs = append(imgs, img)
	}
	return imgs, nil
}

// The store variable is a package level variable that will be available for
// use throughout our application code
var store Store

/*
We will need to call the InitStore method to initialize the store. This will
typically be done at the beginning of our application (in this case, when the server starts up)
This can also be used to set up the store as a mock, which we will be observing
later on
*/
func InitStore(s Store) {
	store = s
}
