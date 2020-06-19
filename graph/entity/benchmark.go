package entity

type Response struct {
	Content  interface{} `json:"content,omitempty"`
	Source   string      `json:"source,omitempty"`
	Duration string      `json:"duration,omitempty"`
}
