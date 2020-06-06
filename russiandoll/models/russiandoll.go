package models

type RussianDoll struct {
	Id                 string        `json:"id"`
	Name               string        `json:"name,omitempty"`
	ParentId           string        `json:"parent_id,omitempty"`
	RussianDolls       []RussianDoll `json:"russian_dolls,omitempty"`
	RussianDolslsCount uint64        `json:"russian_dolls_count,omitempty"`
}
