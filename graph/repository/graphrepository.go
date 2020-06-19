package repository

import (
	"context"
	"github.com/johnfercher/graph-study/graph/entity"
)

type GraphRepository interface {
	GetVertexById(ctx context.Context, id string) (*entity.Vertex, error)
	GetAllVertices(ctx context.Context) ([]*entity.Vertex, error)
	CreateVertex(ctx context.Context, vertex *entity.Vertex) error
	UpdateVertex(ctx context.Context, vertex *entity.Vertex) error
	DeleteVertex(ctx context.Context, id string) error
	CreateEdge(ctx context.Context, parentId string, id string) error
	DeleteEdge(ctx context.Context, parentId string, id string) error
}
