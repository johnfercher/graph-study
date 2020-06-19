package controller

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/johnfercher/graph-study/graph/entity"
	"github.com/johnfercher/graph-study/graph/repository"
	"net/http"
)

type GraphController struct {
	mysql repository.GraphRepository
}

func NewMySqlService(repository repository.GraphRepository) *GraphController {
	return &GraphController{
		mysql: repository,
	}
}

func (self *GraphController) CreateEdge(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	ctx := context.TODO()

	params := mux.Vars(request)

	err := self.mysql.CreateEdge(ctx, params["parent_id"], params["id"])

	if err != nil {
		json.NewEncoder(writer).Encode(err)
		return
	}

	json.NewEncoder(writer).Encode(nil)
}

func (self *GraphController) DeleteEdge(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	ctx := context.TODO()

	params := mux.Vars(request)

	err := self.mysql.DeleteEdge(ctx, params["parent_id"], params["id"])

	if err != nil {
		json.NewEncoder(writer).Encode(err)
		return
	}

	json.NewEncoder(writer).Encode(nil)
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

	err = self.mysql.UpdateVertex(ctx, vertex)
	if err != nil {
		json.NewEncoder(writer).Encode(err)
		return
	}

	json.NewEncoder(writer).Encode(vertex)
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

	err = self.mysql.CreateVertex(ctx, vertex)
	if err != nil {
		json.NewEncoder(writer).Encode(err)
		return
	}

	json.NewEncoder(writer).Encode(vertex)
}

func (self *GraphController) GetVertexById(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	ctx := context.TODO()

	params := mux.Vars(request)

	vertex, err := self.mysql.GetVertexById(ctx, params["id"])

	if err != nil {
		json.NewEncoder(writer).Encode(err)
		return
	}

	json.NewEncoder(writer).Encode(vertex)
}

func (self *GraphController) GetAllVertices(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	ctx := context.TODO()

	vertices, err := self.mysql.GetAllVertices(ctx)

	if err != nil {
		json.NewEncoder(writer).Encode(err)
		return
	}

	json.NewEncoder(writer).Encode(vertices)
}

func (self *GraphController) DeleteVertex(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	ctx := context.TODO()

	params := mux.Vars(request)

	err := self.mysql.DeleteVertex(ctx, params["id"])

	if err != nil {
		json.NewEncoder(writer).Encode(err)
		return
	}

	json.NewEncoder(writer).Encode(nil)
}
