package main

import (
	"go_system_programming/cmd/specification/domain"
	"go_system_programming/cmd/specification/domainservice"
	"go_system_programming/cmd/specification/spec"
)

func main() {
	user := domain.User{
		ID:     1,
		Name:   "John Doe",
		Email:  "john.doe@example.com",
		Age:    20,
		Gender: domain.GenderMale,
	}

	subscription := domain.Subscription{
		ID:             1,
		User:           user,
		IsPaid:         true,
		SubscribeMonth: 12,
	}

	service := domainservice.DrinkAlcoholService{
		CanDrinkAlcoholSpecification: spec.CanDrinkAlcoholSpecification{
			User:         &user,
			Subscription: &subscription,
		},
		CanEatBreakfastSpecification: spec.CanEatBreakfastSpecification{
			User:         &user,
			Subscription: &subscription,
		},
	}
	service.DrinkAlcohol(user, subscription)
}
