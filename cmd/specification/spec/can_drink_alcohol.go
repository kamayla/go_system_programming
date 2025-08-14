package spec

import "go_system_programming/cmd/specification/domain"

type CanDrinkAlcoholSpecification struct {
	User         *domain.User
	Subscription *domain.Subscription
}

func (s CanDrinkAlcoholSpecification) IsSatisfied() bool {
	return s.User.IsAdult() && s.User.IsMale() && s.Subscription.IsPaid
}
