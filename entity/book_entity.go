package entity

type BookEntity struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Owner  string `json:"owner"`
}
