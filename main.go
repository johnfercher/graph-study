package main

import (
	"github.com/gorilla/mux"
	"github.com/johnfercher/graph-study/russiandoll/infra"
	"github.com/johnfercher/graph-study/russiandoll/repositories"
	"github.com/johnfercher/graph-study/russiandoll/services"
	"log"
	"net/http"
)

func main() {
	// Router
	router := mux.NewRouter()

	// Infra
	db, err := infra.GetMysqlConnection()
	if err != nil {
		panic(err)
	}

	// Repositories
	repository := repositories.NewRussianDollRepository(db)

	// Services
	service := services.NewRussianDollService(repository)

	// CRUD
	router.HandleFunc("/api/russian-doll", service.GetAll).Methods("GET")
	router.HandleFunc("/api/russian-doll/{id}", service.GetById).Methods("GET")
	router.HandleFunc("/api/russian-doll", service.Create).Methods("POST")
	router.HandleFunc("/api/russian-doll/{id}", service.Update).Methods("PUT")
	router.HandleFunc("/api/russian-doll/{id}", service.Delete).Methods("DELETE")

	// Relations
	router.HandleFunc("/api/russian-doll/{parent_id}/russian-doll/{id}", service.AddToRussianDoll).Methods("POST")
	router.HandleFunc("/api/russian-doll/{parent_id}/russian-doll/{id}", service.RemoveFromRussianDoll).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}
