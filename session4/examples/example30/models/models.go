package models

type Instructor struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Workshop struct {
	Title       string       `json:"title"`
	Instructors []Instructor `json:"instructors"`
}
