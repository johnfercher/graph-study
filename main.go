package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/johnfercher/graph-study/graph/controller"
	"github.com/johnfercher/graph-study/graph/infra"
	"github.com/johnfercher/graph-study/graph/repository"
	"log"
	"net/http"
)

func main() {
	// Router
	router := mux.NewRouter()

	// Infra
	mysql, err := infra.GetMysqlConnection()
	if err != nil {
		panic(err)
	}

	neo4j, err := infra.GetNeo4jlConnection()
	if err != nil {
		panic(err)
	}

	// Repositories
	mySqlRepository := repository.NewMysqlRepository(mysql)

	// Services
	service := controller.NewMySqlService(mySqlRepository)

	// CRUD
	router.HandleFunc("/api/vertices", service.GetAllVertices).Methods("GET")
	router.HandleFunc("/api/vertices/{id}", service.GetVertexById).Methods("GET")
	router.HandleFunc("/api/vertices", service.CreateVertex).Methods("POST")
	router.HandleFunc("/api/vertices/{id}", service.UpdateVertex).Methods("PUT")
	router.HandleFunc("/api/vertices/{id}", service.DeleteVertex).Methods("DELETE")

	// Relations
	router.HandleFunc("/api/vertices/{parent_id}/vertices/{id}", service.DeleteEdge).Methods("DELETE")
	router.HandleFunc("/api/vertices/{parent_id}/vertices/{id}", service.CreateEdge).Methods("POST")

	fmt.Println("API ON")

	log.Fatal(http.ListenAndServe(":8080", router))
}
