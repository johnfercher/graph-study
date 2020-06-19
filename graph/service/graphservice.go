package service

import (
	"context"
	"github.com/johnfercher/graph-study/graph/entity"
)

type GraphService interface {
	CreateEdge(ctx context.Context, parentId, id string) []entity.Response
	DeleteEdge(ctx context.Context, parentId, id string) []entity.Response
	CreateVertex(ctx context.Context, vertex *entity.Vertex) []entity.Response
	UpdateVertex(ctx context.Context, vertex *entity.Vertex) []entity.Response
	GetVertexById(ctx context.Context, id string) []entity.Response
}
