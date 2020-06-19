package repository

import (
	"context"
	"github.com/johnfercher/graph-study/graph/entity"
	"github.com/neo4j/neo4j-go-driver/neo4j"
)

type neo4jRepository struct {
	driver neo4j.Driver
}

func NewNeo4jRepository(drive neo4j.Driver) *neo4jRepository {
	return &neo4jRepository{
		driver: drive,
	}
}

func (self *neo4jRepository) GetVertexById(ctx context.Context, id string) (*entity.Vertex, error) {
	return &entity.Vertex{}, nil
}

func (self *neo4jRepository) GetAllVertices(ctx context.Context) ([]*entity.Vertex, error) {
	return []*entity.Vertex{}, nil
}

func (self *neo4jRepository) CreateVertex(ctx context.Context, vertex *entity.Vertex) error {
	return nil
}

func (self *neo4jRepository) UpdateVertex(ctx context.Context, vertex *entity.Vertex) error {
	return nil
}

func (self *neo4jRepository) DeleteVertex(ctx context.Context, id string) error {
	return nil
}

func (self *neo4jRepository) CreateEdge(ctx context.Context, parentId string, id string) error {
	return nil
}

func (self *neo4jRepository) DeleteEdge(ctx context.Context, parentId string, id string) error {
	return nil
}

func (self *neo4jRepository) BuildTree(parentId string, vertices map[string]entity.Vertex) entity.Vertex {
	return entity.Vertex{}
}
