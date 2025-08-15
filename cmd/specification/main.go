package main

import (
	"go_system_programming/cmd/specification/domain"
	"go_system_programming/cmd/specification/domainservice"
)

func main() {
	user := domain.User{
		ID:     1,
		Name:   "John Doe",
		Email:  "john.doe@example.com",
		Age:    50,
		Gender: domain.GenderMale,
	}

	subscription := domain.Subscription{
		ID:             1,
		User:           user,
		Status:         domain.SubscriptionStatusActive,
		SubscribeMonth: 6,
	}

	service := domainservice.DrinkAlcoholService{}
	service.DrinkAlcohol(user, subscription)
}
