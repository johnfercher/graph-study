package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/johnfercher/graph-study/graph/controller"
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
	neo4jRepository := repository.NewNeo4jRepository(neo4j)

	// Services
	mySqlService := service.NewGraphService("mysql", mySqlRepository, nil)
	neo4jService := service.NewGraphService("neo4j", neo4jRepository, mySqlService)

	// Controllers
	graphController := controller.NewGraphController(neo4jService)

	// CRUD
	registerEndpoint(router, "GET", "/api/vertices", graphController.GetAllVertices)
	registerEndpoint(router, "GET", "/api/vertices/{id}", graphController.GetVertexById)
	registerEndpoint(router, "POST", "/api/vertices", graphController.CreateVertex)
	registerEndpoint(router, "PUT", "/api/vertices/{id}", graphController.UpdateVertex)
	registerEndpoint(router, "DELETE", "/api/vertices/{id}", graphController.DeleteVertex)

	// Relations
	registerEndpoint(router, "DELETE", "/api/vertices/{parent_id}/vertices/{id}", graphController.DeleteEdge)
	registerEndpoint(router, "POST", "/api/vertices/{parent_id}/vertices/{id}", graphController.CreateEdge)

	fmt.Println("API ON: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func registerEndpoint(router *mux.Router, verb, path string, handlerFunc func(writer http.ResponseWriter, request *http.Request)) {
	fmt.Printf("Registred %s: %s\n", verb, path)
	router.HandleFunc(path, handlerFunc).Methods(verb)
}
