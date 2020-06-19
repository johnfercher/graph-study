package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/johnfercher/graph-study/graph/entity"
)

type mysqlRepository struct {
	db *sql.DB
}

func NewMysqlRepository(db *sql.DB) *mysqlRepository {
	return &mysqlRepository{
		db: db,
	}
}

func (self *mysqlRepository) GetVertexById(ctx context.Context, id string) (*entity.Vertex, error) {
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
SELECT c.id, c.type, parent_id
FROM get_vertices gc
         JOIN vertex c on c.id = gc.vertex_id`, id))
	if err != nil {
		return nil, err
	}

	vertices := make(map[string]entity.Vertex)
	var idRead sql.NullString
	var vertexType sql.NullString
	var parentId sql.NullString

	for results.Next() {
		err = results.Scan(&idRead, &vertexType, &parentId)
		if err != nil {
			return nil, err
		}

		vertex := entity.Vertex{
			Id:       idRead.String,
			Type:     vertexType.String,
			ParentId: parentId.String,
		}

		vertices[vertex.Id] = vertex
	}

	vertex := self.BuildTree(id, vertices)

	return &vertex, nil
}

func (self *mysqlRepository) GetAllVertices(ctx context.Context) ([]*entity.Vertex, error) {
	results, err := self.db.QueryContext(ctx, `SELECT id, type, parent_id FROM vertex LEFT JOIN edge as p ON p.vertex_id = id`)
	if err != nil {
		return nil, err
	}

	vertices := []*entity.Vertex{}

	for results.Next() {
		var idRead sql.NullString
		var vertexType sql.NullString
		var parentId sql.NullString

		err = results.Scan(&idRead, &vertexType, &parentId)
		if err != nil {
			return nil, err
		}

		vertex := &entity.Vertex{
			Id:       idRead.String,
			Type:     vertexType.String,
			ParentId: parentId.String,
		}

		vertices = append(vertices, vertex)
	}

	return vertices, nil
}

func (self *mysqlRepository) CreateVertex(ctx context.Context, vertex *entity.Vertex) error {
	_, err := self.db.ExecContext(ctx, fmt.Sprintf(`INSERT INTO vertex (id, type) values ('%s', '%s')`, vertex.Id, vertex.Type))
	if err != nil {
		return err
	}

	return nil
}

func (self *mysqlRepository) UpdateVertex(ctx context.Context, vertex *entity.Vertex) error {
	_, err := self.db.ExecContext(ctx, fmt.Sprintf(`UPDATE vertex SET type = '%s' WHERE id = '%s'`, vertex.Type, vertex.Id))
	if err != nil {
		return err
	}

	return nil
}

func (self *mysqlRepository) DeleteVertex(ctx context.Context, id string) error {
	_, err := self.db.ExecContext(ctx, fmt.Sprintf(`DELETE FROM vertex WHERE id = '%s'`, id))
	if err != nil {
		return err
	}

	return nil
}

func (self *mysqlRepository) CreateEdge(ctx context.Context, parentId string, id string) error {
	_, err := self.db.ExecContext(ctx, fmt.Sprintf(`INSERT INTO edge (vertex_id, parent_id) VALUES ('%s', '%s')`, id, parentId))
	if err != nil {
		return err
	}

	return nil
}

func (self *mysqlRepository) DeleteEdge(ctx context.Context, parentId string, id string) error {
	_, err := self.db.ExecContext(ctx, fmt.Sprintf(`DELETE FROM edge WHERE parent_id = '%s' AND vertex_id='%s'`, parentId, id))
	if err != nil {
		return err
	}

	return nil
}

func (self *mysqlRepository) BuildTree(parentId string, vertices map[string]entity.Vertex) entity.Vertex {
	parent := vertices[parentId]
	delete(vertices, parentId)

	vertex, _ := self.buildTree(parent, vertices)
	return vertex
}

func (self *mysqlRepository) buildTree(parent entity.Vertex, vertices map[string]entity.Vertex) (entity.Vertex, map[string]entity.Vertex) {
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
