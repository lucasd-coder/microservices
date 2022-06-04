package model

type User struct {
	ID string `json:"id"`

	AuthUserId string `json:"authUserId"`

	Enrollments []*Enrollment `json:"enrollments"`
}

func (User) IsEntity() {}
