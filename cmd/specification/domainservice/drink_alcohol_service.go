package domainservice

import (
	"fmt"
	"go_system_programming/cmd/specification/spec"
)

type DrinkAlcoholService struct {
	specFacade             spec.SpecFacade
	UserRepository         IUserRepository
	SubscriptionRepository ISubscriptionRepository
}

// 飲めるようなら飲む
func (s DrinkAlcoholService) DrinkAlcohol(userID int) {
	user, err := s.UserRepository.FindUserByID(userID)
	if err != nil {
		fmt.Println("User not found")
		return
	}

	subscription, err := s.SubscriptionRepository.FindSubscriptionByUserID(userID)
	if err != nil {
		fmt.Println("Subscription not found")
		return
	}

	if s.specFacade.CanEatBreakfast(user, subscription) {
		fmt.Println("朝食もどうぞ！！！")
	}

	if s.specFacade.Combination(user, subscription) {
		fmt.Println("飲めるし食べれる🍺🥐🥐🥐")
	}

	if s.specFacade.CanDrinkAlcohol(user, subscription) {
		fmt.Println("お酒を飲んでください🍺")
		return
	}

	fmt.Println("お酒は飲めません🈲🈲🈲")
}
