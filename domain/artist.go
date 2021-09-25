package domain

type Artist struct {
	Id     string  `json:"id"`
	Name   string  `json:"name"`
	Images []Image `json:"images"`
}
