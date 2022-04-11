package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/dialects/mysql"
	"github.com/leksyking/go-bookstore/pkg/routes"
	"github.com/leksyking/go-bookstore/pkg/routes.go"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:5000", r))
}
