package main

import (
	"go_system_programming/cmd/specification/domainservice"
	"go_system_programming/cmd/specification/repository"
)

func main() {
	service := domainservice.DrinkAlcoholService{
		UserRepository:         repository.NewUserRepository(),
		SubscriptionRepository: repository.NewSubscriptionRepository(),
	}
	service.DrinkAlcohol(1)
}
