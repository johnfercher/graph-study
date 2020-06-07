package entity

type Vertex struct {
	Id            string   `json:"id"`
	Type          string   `json:"type,omitempty"`
	ParentId      string   `json:"parent_id,omitempty"`
	Vertices      []Vertex `json:"vertices,omitempty"`
	VerticesCount uint64   `json:"vertices_count,omitempty"`
}
