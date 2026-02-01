package dto

type UpdateThoughtInput struct {
	ID     int64  `json:"id"`
	Text   string `json:"text"`
	Author string `json:"author"`
}
