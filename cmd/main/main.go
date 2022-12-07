package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/michaelzdrav/go-bookstore/pkg/routes"
)

func main() {
	fmt.Println("RUNNING - Starting application ...")
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	fmt.Println("SUCCESS - Listening on port 9010 ...")
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
