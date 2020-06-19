package controller

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/johnfercher/graph-study/graph/entity"
	"github.com/johnfercher/graph-study/graph/service"
	"net/http"
)

type GraphController struct {
	service service.GraphService
}

func NewGraphController(service service.GraphService) *GraphController {
	return &GraphController{
		service: service,
	}
}

func (self *GraphController) CreateEdge(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	ctx := context.TODO()

	params := mux.Vars(request)

	response := self.service.CreateEdge(ctx, params["parent_id"], params["id"])
	json.NewEncoder(writer).Encode(response)
}

func (self *GraphController) DeleteEdge(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	ctx := context.TODO()

	params := mux.Vars(request)

	response := self.service.DeleteEdge(ctx, params["parent_id"], params["id"])
	json.NewEncoder(writer).Encode(response)
}

func (self *GraphController) UpdateVertex(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	ctx := context.TODO()

	vertex := &entity.Vertex{}
	err := json.NewDecoder(request.Body).Decode(vertex)
	if err != nil {
		json.NewEncoder(writer).Encode(err)
		return
	}

	params := mux.Vars(request)
	vertex.Id = params["id"]

	response := self.service.UpdateVertex(ctx, vertex)
	json.NewEncoder(writer).Encode(response)
}

func (self *GraphController) CreateVertex(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	ctx := context.TODO()

	vertex := &entity.Vertex{}
	err := json.NewDecoder(request.Body).Decode(vertex)
	if err != nil {
		json.NewEncoder(writer).Encode(err)
		return
	}

	id, _ := uuid.NewRandom()
	vertex.Id = id.String()

	response := self.service.CreateVertex(ctx, vertex)
	json.NewEncoder(writer).Encode(response)
}

func (self *GraphController) GetVertexById(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	ctx := context.TODO()

	params := mux.Vars(request)

	response := self.service.GetVertexById(ctx, params["id"])
	json.NewEncoder(writer).Encode(response)
}

func (self *GraphController) GetAllVertices(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	ctx := context.TODO()

	response := self.service.GetAllVertices(ctx)
	json.NewEncoder(writer).Encode(response)
}

func (self *GraphController) DeleteVertex(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	ctx := context.TODO()

	params := mux.Vars(request)

	response := self.service.DeleteVertex(ctx, params["id"])
	json.NewEncoder(writer).Encode(response)
}
