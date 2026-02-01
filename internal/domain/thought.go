package domain

type Thought struct {
	ID     int64  `json:"id"`
	Text   string `json:"text"`
	Author string `json:"author"`
}
