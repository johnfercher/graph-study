package entity

type Vertex struct {
	Id            string   `json:"id"`
	Name          string   `json:"name,omitempty"`
	ParentId      string   `json:"parent_id,omitempty"`
	Vertices      []Vertex `json:"vertices,omitempty"`
	VerticesCount uint64   `json:"vertices_count,omitempty"`
}
