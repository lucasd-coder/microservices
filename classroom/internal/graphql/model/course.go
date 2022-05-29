package model

type Course struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Slug string `json:"slug"`
}

func (Course) IsEntity() {}

type CreateCourseInput struct {
	Title string `json:"title"`
}