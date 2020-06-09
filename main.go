package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/johnfercher/graph-study/graph/infra"
	"github.com/johnfercher/graph-study/graph/repository"
	"github.com/johnfercher/graph-study/graph/service"
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
	repository := repository.NewGraphRepository(db)

	// Services
	service := service.NewGraphService(repository)

	// CRUD
	router.HandleFunc("/mysql/vertices", service.GetAllVertices).Methods("GET")
	router.HandleFunc("/mysql/vertices/{id}", service.GetVertexById).Methods("GET")
	router.HandleFunc("/mysql/vertices", service.CreateVertex).Methods("POST")
	router.HandleFunc("/mysql/vertices/{id}", service.UpdateVertex).Methods("PUT")
	router.HandleFunc("/mysql/vertices/{id}", service.DeleteVertex).Methods("DELETE")

	// Relations
	router.HandleFunc("/mysql/vertices/{parent_id}/vertices/{id}", service.DeleteEdge).Methods("DELETE")
	router.HandleFunc("/mysql/vertices/{parent_id}/vertices/{id}", service.CreateEdge).Methods("POST")

	fmt.Println("API ON")

	log.Fatal(http.ListenAndServe(":8080", router))
}
