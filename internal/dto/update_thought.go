package dto

type UpdateThoughtRequest struct {
	Text   string `json:"text"`
	Author string `json:"author"`
}
