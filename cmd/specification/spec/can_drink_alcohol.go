package spec

import "go_system_programming/cmd/specification/domain"

type CanDrinkAlcoholSpecification struct {
	User         *domain.User
	Subscription *domain.Subscription
}

func (s CanDrinkAlcoholSpecification) IsSatisfied() bool {
	return Or(
		UserAdultSpecification{User: s.User},
		UserMaleSpecification{User: s.User},
		SubscriptionPaidSpecification{Subscription: s.Subscription},
	).IsSatisfied()
}
