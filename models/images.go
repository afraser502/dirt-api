package models

import "database/sql"

/*
type Book struct {
	Isbn   string
	Title  string
	Author string
	Price  float32
}
*/
type Image struct {
	Repository string `json:"repository"`
	Id         string `json:"id"`
	Tag        string `json:"tag"`
	Size       string `json:"size"`
	Created    string `json:"created"`
	Requestor  string `json:"requestor"`
	Status     string `json:"status"`
}

func AllImages(db *sql.DB) ([]*Image, error) {
	rows, err := db.Query("SELECT repository, id, tag, size, created, requestor, status FROM images")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	imgs := make([]*Image, 0)
	for rows.Next() {
		img := new(Image)
		err := rows.Scan(&img.Repository, &img.Id, &img.Tag, &img.Size, &img.Created, &img.Requestor, &img.Status)

		if err != nil {
			return nil, err
		}
		imgs = append(imgs, img)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return imgs, nil
}

func DownloadImage(db *sql.DB) (image *Image) {

	_, err := db.Query("INSERT INTO user VALUES (?)")

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}

	return err

}
