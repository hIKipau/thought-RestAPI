package dto

type CreateThoughtInput struct {
	Text   string `json:"text"`
	Author string `json:"author"`
}
type CreateThoughtOutput struct {
	ID int64 `json:"id"`
}
