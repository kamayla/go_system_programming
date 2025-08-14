package domain

type Subscription struct {
	ID     int
	User   User
	IsPaid bool
}
