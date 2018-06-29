package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/viper"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

/*
func Authenticate(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Authenticate token !")
}

func LdapLookup(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "LDAP Lookup!")
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
*/
// The new router function creates the router and
// returns it to us. We can now use this function
// to instantiate and test the router outside of the main function
func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/hello", handler).Methods("GET")

	// Declare the static file directory and point it to the directory we just made
	staticFileDirectory := http.Dir("./assets/")
	// Declare the handler, that routes requests to their respective filename.
	// The fileserver is wrapped in the `stripPrefix` method, because we want to
	// remove the "/assets/" prefix when looking for files.
	// For example, if we type "/assets/index.html" in our browser, the file server
	// will look for only "index.html" inside the directory declared above.
	// If we did not strip the prefix, the file server would look for "./assets/assets/index.html", and yield an error
	staticFileHandler := http.StripPrefix("/assets/", http.FileServer(staticFileDirectory))
	// The "PathPrefix" method acts as a matcher, and matches all routes starting
	// with "/assets/", instead of the absolute route itself
	r.PathPrefix("/assets/").Handler(staticFileHandler).Methods("GET")
	r.HandleFunc("/images", ImageList).Methods("GET")
	r.HandleFunc("/images", ImageDownload).Methods("POST")
	return r
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
	/*	db, err := models.NewDB(dbcon.String())
		if err != nil {
			log.Panic(err)
		}

		InitStore(&dbStore{db: db})
	*/
	fmt.Println("Starting server...")
	db, err := sql.Open("mysql", dbcon.String())

	if err != nil {
		panic(err)
	}
	err = db.Ping()

	if err != nil {
		panic(err)
	}

	InitStore(&dbStore{db: db})

	//Set routes using Gorilla

	/*
		r := mux.NewRouter()
		r.HandleFunc("/authenticate", Authenticate).Methods("GET")
		r.HandleFunc("/ldap", LdapLookup).Methods("POST")
		r.HandleFunc("/images", ImageList).Methods("GET")
		r.HandleFunc("/images/download", ImageDownload).Methods("POST")
		r.HandleFunc("/logs/twistlock", LogsTwistlock).Methods("GET")
		r.HandleFunc("/logs/main", LogsMainApp).Methods("GET")
		r.HandleFunc("/logs/db", LogsDB).Methods("GET")
		if err := http.ListenAndServe(":3005", r); err != nil {
			log.Fatal(err)
		}

	*/

	r := newRouter()
	fmt.Println("Serving on port 3005")
	http.ListenAndServe(":3005", r)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}
