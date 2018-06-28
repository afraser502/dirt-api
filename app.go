package main

import (
	"fmt"
	"log"
	"net/http"

	. "github.com/afraser502/dirt-api/db"
	"github.com/gorilla/mux"
)

func Authenticate(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Authenticate token !")
}

func LdapLookup(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "LDAP Lookup!")
}

func ImageList(w http.ResponseWriter, r *http.Request) {

	DbImagesList()
	fmt.Fprintln(w, "List images here!")

}

func ImageDownload(w http.ResponseWriter, r *http.Request) {
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

}
