package main

import (
	"github.com/gorilla/mux"
	"github.com/johnfercher/graph-study/russiandoll/controllers"
	"log"
	"net/http"
)

func main() {
	// Router
	router := mux.NewRouter()

	// Controllers
	controller := controllers.RussianDollController{}

	// Binds
	controller.Bind(router)

	log.Fatal(http.ListenAndServe(":8080", router))
}
