package dto

type CreateThoughtRequest struct {
	Text   string `json:"text"`
	Author string `json:"author"`
}
type CreateThoughtResponse struct {
	ID int64 `json:"id"`
}
