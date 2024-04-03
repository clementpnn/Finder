package model

type SearchInput struct {
	Input string
	Page  int32
}

type SearchOutput struct {
	Url         string
	Title       string
	Description string
	Image       string
}

type MetaData struct {
	Charset  string `json:"charset,omitempty"`
	Name     string `json:"name,omitempty"`
	Content  string `json:"content,omitempty"`
	Inertia  string `json:"inertia,omitempty"`
	Property string `json:"property,omitempty"`
}
