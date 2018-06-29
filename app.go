package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/viper"

	"github.com/afraser502/dirt-api/models"
	"github.com/gorilla/mux"
)

type Env struct {
	db *sql.DB
}

func Authenticate(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Authenticate token !")
}

func LdapLookup(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "LDAP Lookup!")
}

func (env *Env) ImageList(w http.ResponseWriter, r *http.Request) {

	//DbImagesList()
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	imgs, err := models.AllImages(env.db)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		log.Fatal(err)
		return
	}

	for _, img := range imgs {
		fmt.Println(img.Repository, img.Id, img.Tag, img.Size, img.Created, img.Requestor, img.Status)
	}

	fmt.Fprintln(w, "List images here!")
}

func (env *Env) ImageDownload(w http.ResponseWriter, r *http.Request) {
	imgdl, err := models.DownloadImage(env.db)

	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		log.Fatal(err)
		return
	}

	fmt.Fprintln(w, "Download images from here !")
}

func LogsTwistlock(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Twistlock logs!")
}

func LogsMainApp(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Main app logs!")
}

func LogsDB(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "DB logs!")
}

func main() {
	//Set variable connection path using Viper
	viper.SetConfigFile("./configs/env.json") //Set config file

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	//Build DB connection string
	var dbcon bytes.Buffer
	dbcon.WriteString(viper.GetString("db.user") + ":")
	dbcon.WriteString(viper.GetString("db.pass") + "@tcp(")
	dbcon.WriteString(viper.GetString("db.host") + ":")
	dbcon.WriteString(viper.GetString("db.port") + ")/")
	dbcon.WriteString(viper.GetString("db.name"))

	fmt.Println(dbcon.String())

	//Open DB Connection
	db, err := models.NewDB(dbcon.String())
	if err != nil {
		log.Panic(err)
	}
	env := &Env{db: db}

	//Set routes using Gorilla
	r := mux.NewRouter()
	r.HandleFunc("/authenticate", Authenticate).Methods("GET")
	r.HandleFunc("/ldap", LdapLookup).Methods("POST")
	r.HandleFunc("/images", env.ImageList).Methods("GET")
	r.HandleFunc("/images/download", env.ImageDownload).Methods("POST")
	r.HandleFunc("/logs/twistlock", LogsTwistlock).Methods("GET")
	r.HandleFunc("/logs/main", LogsMainApp).Methods("GET")
	r.HandleFunc("/logs/db", LogsDB).Methods("GET")
	if err := http.ListenAndServe(":3005", r); err != nil {
		log.Fatal(err)
	}

}
