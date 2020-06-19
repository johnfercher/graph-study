package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/johnfercher/graph-study/graph/entity"
	"github.com/johnfercher/graph-study/graph/repository"
	"time"
)

type GraphService interface {
	GetVertexById(ctx context.Context, id string) []entity.Response
	GetAllVertices(ctx context.Context) []entity.Response
	CreateVertex(ctx context.Context, vertex *entity.Vertex) []entity.Response
	UpdateVertex(ctx context.Context, vertex *entity.Vertex) []entity.Response
	DeleteVertex(ctx context.Context, id string) []entity.Response
	CreateEdge(ctx context.Context, parentId, id string) []entity.Response
	DeleteEdge(ctx context.Context, parentId, id string) []entity.Response
}

type graphService struct {
	source  string
	service repository.GraphRepository
	inner   GraphService
}

func NewGraphService(source string, repository repository.GraphRepository, inner GraphService) *graphService {
	return &graphService{
		source:  source,
		service: repository,
		inner:   inner,
	}
}

func (self *graphService) CreateEdge(ctx context.Context, parentId, id string) []entity.Response {
	begin := time.Now()

	err := self.service.CreateEdge(ctx, parentId, id)

	response := entity.Response{
		Content:  err,
		Source:   self.source,
		Duration: time.Now().Sub(begin).String(),
	}

	if self.inner != nil {
		innerResponse := self.inner.CreateEdge(ctx, parentId, id)
		innerResponse = append(innerResponse, response)
		return innerResponse
	}

	return []entity.Response{response}
}

func (self *graphService) DeleteEdge(ctx context.Context, parentId, id string) []entity.Response {
	begin := time.Now()

	err := self.service.DeleteEdge(ctx, parentId, id)

	response := entity.Response{
		Content:  err,
		Source:   self.source,
		Duration: time.Now().Sub(begin).String(),
	}

	if self.inner != nil {
		innerResponse := self.inner.DeleteEdge(ctx, parentId, id)
		innerResponse = append(innerResponse, response)
		return innerResponse
	}

	return []entity.Response{response}
}

func (self *graphService) CreateVertex(ctx context.Context, vertex *entity.Vertex) []entity.Response {
	begin := time.Now()

	id, _ := uuid.NewRandom()
	vertex.Id = id.String()

	err := self.service.CreateVertex(ctx, vertex)

	response := entity.Response{
		Content:  err,
		Source:   self.source,
		Duration: time.Now().Sub(begin).String(),
	}

	if self.inner != nil {
		innerResponse := self.inner.CreateVertex(ctx, vertex)
		innerResponse = append(innerResponse, response)
		return innerResponse
	}

	return []entity.Response{response}
}

func (self *graphService) UpdateVertex(ctx context.Context, vertex *entity.Vertex) []entity.Response {
	begin := time.Now()

	err := self.service.UpdateVertex(ctx, vertex)

	response := entity.Response{
		Content:  err,
		Source:   self.source,
		Duration: time.Now().Sub(begin).String(),
	}

	if self.inner != nil {
		innerResponse := self.inner.UpdateVertex(ctx, vertex)
		innerResponse = append(innerResponse, response)
		return innerResponse
	}
	return []entity.Response{response}
}

func (self *graphService) GetVertexById(ctx context.Context, id string) []entity.Response {
	begin := time.Now()

	vertex, err := self.service.GetVertexById(ctx, id)

	var content interface{}
	if err != nil {
		content = err
	} else {
		content = vertex
	}

	response := entity.Response{
		Content:  content,
		Source:   self.source,
		Duration: time.Now().Sub(begin).String(),
	}

	if self.inner != nil {
		innerResponse := self.inner.GetVertexById(ctx, id)
		innerResponse = append(innerResponse, response)
		return innerResponse
	}

	return []entity.Response{response}
}

func (self *graphService) GetAllVertices(ctx context.Context) []entity.Response {
	begin := time.Now()

	vertex, err := self.service.GetAllVertices(ctx)

	var content interface{}
	if err != nil {
		content = err
	} else {
		content = vertex
	}

	response := entity.Response{
		Content:  content,
		Source:   self.source,
		Duration: time.Now().Sub(begin).String(),
	}

	if self.inner != nil {
		innerResponse := self.inner.GetAllVertices(ctx)
		innerResponse = append(innerResponse, response)
		return innerResponse
	}

	return []entity.Response{response}
}

func (self *graphService) DeleteVertex(ctx context.Context, id string) []entity.Response {
	begin := time.Now()

	err := self.service.DeleteVertex(ctx, id)

	response := entity.Response{
		Content:  err,
		Source:   self.source,
		Duration: time.Now().Sub(begin).String(),
	}

	if self.inner != nil {
		innerResponse := self.inner.DeleteVertex(ctx, id)
		innerResponse = append(innerResponse, response)
		return innerResponse
	}

	return []entity.Response{response}
}
