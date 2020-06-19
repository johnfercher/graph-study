package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/johnfercher/graph-study/graph/entity"
	"github.com/johnfercher/graph-study/graph/repository"
	"time"
)

type MysqlService struct {
	mysql repository.GraphRepository
}

func NewMySqlService(repository repository.GraphRepository) *MysqlService {
	return &MysqlService{
		mysql: repository,
	}
}

func (self *MysqlService) CreateEdge(ctx context.Context, parentId, id string) []entity.Response {
	begin := time.Now()

	err := self.mysql.CreateEdge(ctx, parentId, id)

	response := entity.Response{
		Content:  err,
		Source:   "mysql",
		Duration: time.Now().Sub(begin).String(),
	}

	return []entity.Response{response}
}

func (self *MysqlService) DeleteEdge(ctx context.Context, parentId, id string) []entity.Response {
	begin := time.Now()

	err := self.mysql.DeleteEdge(ctx, parentId, id)

	response := entity.Response{
		Content:  err,
		Source:   "mysql",
		Duration: time.Now().Sub(begin).String(),
	}

	return []entity.Response{response}
}

func (self *MysqlService) CreateVertex(ctx context.Context, vertex *entity.Vertex) []entity.Response {
	begin := time.Now()

	id, _ := uuid.NewRandom()
	vertex.Id = id.String()

	err := self.mysql.CreateVertex(ctx, vertex)

	response := entity.Response{
		Content:  err,
		Source:   "mysql",
		Duration: time.Now().Sub(begin).String(),
	}

	return []entity.Response{response}
}

func (self *MysqlService) UpdateVertex(ctx context.Context, vertex *entity.Vertex) []entity.Response {
	begin := time.Now()

	err := self.mysql.UpdateVertex(ctx, vertex)

	response := entity.Response{
		Content:  err,
		Source:   "mysql",
		Duration: time.Now().Sub(begin).String(),
	}

	return []entity.Response{response}
}

func (self *MysqlService) GetVertexById(ctx context.Context, id string) []entity.Response {
	begin := time.Now()

	vertex, err := self.mysql.GetVertexById(ctx, id)

	var content interface{}
	if err != nil {
		content = err
	} else {
		content = vertex
	}

	response := entity.Response{
		Content:  content,
		Source:   "mysql",
		Duration: time.Now().Sub(begin).String(),
	}

	return []entity.Response{response}
}

func (self *MysqlService) GetAllVertices(ctx context.Context) []entity.Response {
	begin := time.Now()

	vertex, err := self.mysql.GetAllVertices(ctx)

	var content interface{}
	if err != nil {
		content = err
	} else {
		content = vertex
	}

	response := entity.Response{
		Content:  content,
		Source:   "mysql",
		Duration: time.Now().Sub(begin).String(),
	}

	return []entity.Response{response}
}

func (self *MysqlService) DeleteVertex(ctx context.Context, id string) []entity.Response {
	begin := time.Now()

	err := self.mysql.DeleteVertex(ctx, id)

	response := entity.Response{
		Content:  err,
		Source:   "mysql",
		Duration: time.Now().Sub(begin).String(),
	}

	return []entity.Response{response}
}
