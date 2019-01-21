package main

import (
	"flag"
	"net/http"
	"strconv"
	"text/template"

	"github.com/gorilla/mux"
)

var (
	dbServer = flag.String("dbserver", "localhost", "database server ip")
	host     = flag.String("host", "0.0.0.0:8000", "host to serve web")
	port     = flag.Int("port", 5432, "port number")
	user     = flag.String("user", "testdbmaster", "postgres username")
	password = flag.String("password", "testdbmasterpass", "postgres password")
	dbName   = flag.String("dbname", "testdb1", "database name")
)

func main() {
	flag.Parse()
	db, err := OpenDB(*dbServer)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	InitStore(&dbStore{db: db})
	r := mux.NewRouter()
	r.HandleFunc("/", indexHandler).Methods("GET")
	r.HandleFunc("/add", addHandler).Methods("POST")
	r.HandleFunc("/delete", deleteHandler).Methods("POST")
	http.ListenAndServe(*host, r)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	var list []Entry
	list, err := store.ListEntries()
	if err != nil {
		panic(err)
	}
	t, _ := template.ParseFiles("template.tmpl")
	t.ExecuteTemplate(w, "main", &list)
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.Form["name"]
	store.AddEntry(name[0])
	http.Redirect(w, r, r.Header.Get("Referer"), 302)
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id, err := strconv.Atoi(r.Form["id"][0])
	if err != nil {
		panic(err)
	}
	err = store.DeleteEntry(id)
	if err != nil {
		panic(err)
	}
	http.Redirect(w, r, r.Header.Get("Referer"), 302)
}
