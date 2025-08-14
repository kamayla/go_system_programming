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
