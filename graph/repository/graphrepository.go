package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/johnfercher/graph-study/graph/entity"
)

type GraphRepository struct {
	db *sql.DB
}

func NewGraphRepository(db *sql.DB) *GraphRepository {
	return &GraphRepository{
		db: db,
	}
}

func (self *GraphRepository) GetVertexById(ctx context.Context, id string) (*entity.Vertex, error) {
	results, err := self.db.QueryContext(ctx, fmt.Sprintf(`WITH RECURSIVE get_vertices AS (
    SELECT
        c.id as vertex_id
         , cp.parent_id
    FROM vertex as c
             LEFT JOIN edge cp on cp.vertex_id = c.id
    WHERE
            c.id = '%s'
    UNION ALL
    SELECT
        cp2.vertex_id
         , cp2.parent_id
    FROM edge cp2
             INNER JOIN get_vertices o on o.vertex_id = cp2.parent_id
)
SELECT c.id, c.name, parent_id
FROM get_vertices gc
         JOIN vertex c on c.id = gc.vertex_id`, id))
	if err != nil {
		return nil, err
	}

	vertices := make(map[string]entity.Vertex)
	var idRead sql.NullString
	var name sql.NullString
	var parentId sql.NullString

	for results.Next() {
		err = results.Scan(&idRead, &name, &parentId)
		if err != nil {
			return nil, err
		}

		vertex := entity.Vertex{
			Id:       idRead.String,
			Name:     name.String,
			ParentId: parentId.String,
		}

		vertices[vertex.Id] = vertex
	}

	vertex := self.BuildTree(id, vertices)

	return &vertex, nil
}

func (self *GraphRepository) GetAllVertices(ctx context.Context) ([]*entity.Vertex, error) {
	results, err := self.db.QueryContext(ctx, `SELECT id, name, parent_id FROM vertex LEFT JOIN edge as p ON p.vertex_id = id`)
	if err != nil {
		return nil, err
	}

	vertices := []*entity.Vertex{}

	for results.Next() {
		var idRead sql.NullString
		var name sql.NullString
		var parentId sql.NullString

		err = results.Scan(&idRead, &name, &parentId)
		if err != nil {
			return nil, err
		}

		vertex := &entity.Vertex{
			Id:       idRead.String,
			Name:     name.String,
			ParentId: parentId.String,
		}

		vertices = append(vertices, vertex)
	}

	return vertices, nil
}

func (self *GraphRepository) CreateVertex(ctx context.Context, vertex *entity.Vertex) error {
	_, err := self.db.ExecContext(ctx, fmt.Sprintf(`INSERT INTO vertex (id, name) values ('%s', '%s')`, vertex.Id, vertex.Name))
	if err != nil {
		return err
	}

	return nil
}

func (self *GraphRepository) UpdateVertex(ctx context.Context, vertex *entity.Vertex) error {
	_, err := self.db.ExecContext(ctx, fmt.Sprintf(`UPDATE vertex SET name = '%s' WHERE id = '%s'`, vertex.Name, vertex.Id))
	if err != nil {
		return err
	}

	return nil
}

func (self *GraphRepository) DeleteVertex(ctx context.Context, id string) error {
	_, err := self.db.ExecContext(ctx, fmt.Sprintf(`DELETE FROM vertex WHERE id = '%s'`, id))
	if err != nil {
		return err
	}

	return nil
}

func (self *GraphRepository) CreateEdge(ctx context.Context, parentId string, id string) error {
	_, err := self.db.ExecContext(ctx, fmt.Sprintf(`INSERT INTO edge (vertex_id, parent_id) VALUES ('%s', '%s')`, id, parentId))
	if err != nil {
		return err
	}

	return nil
}

func (self *GraphRepository) DeleteEdge(ctx context.Context, parentId string, id string) error {
	_, err := self.db.ExecContext(ctx, fmt.Sprintf(`DELETE FROM edge WHERE parent_id = '%s' AND vertex_id='%s'`, parentId, id))
	if err != nil {
		return err
	}

	return nil
}

func (self *GraphRepository) BuildTree(parentId string, vertices map[string]entity.Vertex) entity.Vertex {
	parent := vertices[parentId]
	delete(vertices, parentId)

	vertex, _ := self.buildTree(parent, vertices)
	return vertex
}

func (self *GraphRepository) buildTree(parent entity.Vertex, vertices map[string]entity.Vertex) (entity.Vertex, map[string]entity.Vertex) {
	var verticesFromVertex []entity.Vertex

	for _, vertex := range vertices {
		if parent.Id == vertex.ParentId {
			verticesFromVertex = append(verticesFromVertex, vertex)
			delete(vertices, vertex.Id)
		}
	}

	for _, parentVertex := range verticesFromVertex {
		builtParentVertex, remain := self.buildTree(parentVertex, vertices)

		vertices = remain
		parent.Vertices = append(parent.Vertices, builtParentVertex)
	}

	return parent, vertices
}
