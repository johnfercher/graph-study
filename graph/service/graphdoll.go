package service

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/johnfercher/graph-study/graph/entity"
	"github.com/johnfercher/graph-study/graph/repository"
	"net/http"
)

type GraphService struct {
	repository *repository.GraphRepository
}

func NewGraphService(repository *repository.GraphRepository) *GraphService {
	return &GraphService{
		repository: repository,
	}
}

func (self *GraphService) CreateEdge(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	ctx := context.TODO()

	params := mux.Vars(request)

	err := self.repository.CreateEdge(ctx, params["parent_id"], params["id"])

	if err != nil {
		json.NewEncoder(writer).Encode(err)
		return
	}

	json.NewEncoder(writer).Encode(nil)
}

func (self *GraphService) DeleteEdge(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	ctx := context.TODO()

	params := mux.Vars(request)

	err := self.repository.DeleteEdge(ctx, params["parent_id"], params["id"])

	if err != nil {
		json.NewEncoder(writer).Encode(err)
		return
	}

	json.NewEncoder(writer).Encode(nil)
}

func (self *GraphService) UpdateVertex(writer http.ResponseWriter, request *http.Request) {
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

	err = self.repository.UpdateVertex(ctx, vertex)
	if err != nil {
		json.NewEncoder(writer).Encode(err)
		return
	}

	json.NewEncoder(writer).Encode(vertex)
}

func (self *GraphService) CreateVertex(writer http.ResponseWriter, request *http.Request) {
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

	err = self.repository.CreateVertex(ctx, vertex)
	if err != nil {
		json.NewEncoder(writer).Encode(err)
		return
	}

	json.NewEncoder(writer).Encode(vertex)
}

func (self *GraphService) GetVertexById(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	ctx := context.TODO()

	params := mux.Vars(request)

	vertex, err := self.repository.GetVertexById(ctx, params["id"])

	if err != nil {
		json.NewEncoder(writer).Encode(err)
		return
	}

	json.NewEncoder(writer).Encode(vertex)
}

func (self *GraphService) GetAllVertices(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	ctx := context.TODO()

	vertices, err := self.repository.GetAllVertices(ctx)

	if err != nil {
		json.NewEncoder(writer).Encode(err)
		return
	}

	json.NewEncoder(writer).Encode(vertices)
}

func (self *GraphService) DeleteVertex(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	ctx := context.TODO()

	params := mux.Vars(request)

	err := self.repository.DeleteVertex(ctx, params["id"])

	if err != nil {
		json.NewEncoder(writer).Encode(err)
		return
	}

	json.NewEncoder(writer).Encode(nil)
}
