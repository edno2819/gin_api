package models

type Person struct {
	FirstName string `json:"firstname"`
}

type Video struct {
	Title       string `json:"title" xml:"title" binding:"min=2,max=10" validate:"is-cool"`
	Description string `json:"description" validate:"is-cool"`
	URL         string `json:"ulr"`
}
