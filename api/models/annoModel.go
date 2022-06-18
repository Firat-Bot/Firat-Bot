package models

type Event struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Url         string `json:"url"`
}
type Lecturer struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Workspace string `json:"workspace"`
}
