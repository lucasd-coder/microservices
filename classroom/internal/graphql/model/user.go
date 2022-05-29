package model

type User struct {
	ID string `json:"id"`

	AuthUserId string

	Enrollments []Enrollment `json:"enrollments"`
}

func (User) IsEntity() {}
