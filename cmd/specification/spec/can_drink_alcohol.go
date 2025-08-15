package spec

import "go_system_programming/cmd/specification/domain"

type CanDrinkAlcoholSpecification struct {
	User         *domain.User
	Subscription *domain.Subscription
}

func (s CanDrinkAlcoholSpecification) IsSatisfied() bool {
	if s.User == nil || s.Subscription == nil {
		return false
	}
	return s.User.IsAdult() && s.User.IsMale() && s.Subscription.IsActive()
}
