package domain

type Gender int

const (
	GenderMale Gender = iota
	GenderFemale
)

type User struct {
	ID     int
	Name   string
	Email  string
	Age    int
	Gender Gender
}

func (u *User) IsAdult() bool {
	return u.Age >= 20
}

func (u *User) IsMale() bool {
	return u.Gender == GenderMale
}
