package model

type Owner struct {
	ID   string `gorm:primarykey`
	Name string
}
