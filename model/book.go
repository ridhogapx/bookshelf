package model

type Book struct {
	ID     string `gorm:primarykey`
	Title  string
	Author string
	Owner  string
}
