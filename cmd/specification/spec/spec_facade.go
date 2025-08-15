package spec

import (
	"go_system_programming/cmd/specification/domain"
)

type SpecFacade struct{}

func (b *SpecFacade) CanDrinkAlcohol(user *domain.User, subscription *domain.Subscription) bool {
	spec := CanDrinkAlcoholSpecification{
		User:         user,
		Subscription: subscription,
	}
	return spec.IsSatisfied()
}

func (b *SpecFacade) CanEatBreakfast(user *domain.User, subscription *domain.Subscription) bool {
	spec := CanEatBreakfastSpecification{
		User:         user,
		Subscription: subscription,
	}
	return spec.IsSatisfied()
}

func (b *SpecFacade) Combination(user *domain.User, subscription *domain.Subscription) bool {
	spec := CombinationSpecification{
		User:         user,
		Subscription: subscription,
	}
	return spec.IsSatisfied()
}
