package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mohammedfuta2000/go-bookstore/pkg/routes"
)

func main()  {
	r:= mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	fmt.Println("starting server on port 9010")
	log.Fatal(http.ListenAndServe(":9010",r))
}