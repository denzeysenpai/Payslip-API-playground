package main

import (
	"database/sql"
	"fmt"
	"gomysql/cmd"
	"gomysql/routes"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	var cmd cmd.Value // this is for making MySQL queries
	// var insert sample.Value
	user := "root"                //username for the db (registered)
	pass := "coupedegrace"        //password of the user
	host := "tcp(127.0.0.1:3306)" //host where your db is in
	dbName := "dbTest"            //db name
	dsn := user + ":" + pass + "@" + host + "/" + dbName
	db, err := sql.Open("mysql", dsn) //mysql start
	if err != nil {
		log.Fatal("Error: ", err)
	}

	_ = db
	_ = cmd
	r := mux.NewRouter()   // routing
	routes.AttachRoutes(r) // for serving static html files

	fmt.Printf("Domain Expansion: Serving at 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r)) // serve
}
