package main

type Workshop struct {
	Title       string       `json:"title"`
	Instructors []Instructor `json:"instructors"`
}

type Instructor struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
