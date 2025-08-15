package domainservice

import (
	"fmt"
	"go_system_programming/cmd/specification/domain"
	"go_system_programming/cmd/specification/spec"
)

type DrinkAlcoholService struct {
	specFactory spec.SpecFacade
}

// 飲めるようなら飲む
func (s DrinkAlcoholService) DrinkAlcohol(user domain.User, subscription domain.Subscription) {

	if s.specFactory.CanEatBreakfast(&user, &subscription) {
		fmt.Println("User can eat breakfast🥐🥐🥐")
	}

	if s.specFactory.Combination(&user, &subscription) {
		fmt.Println("飲めるし食べれる🍺🥐🥐🥐")
	}

	if s.specFactory.CanDrinkAlcohol(&user, &subscription) {
		fmt.Println("User can drink alcohol🍺🍺🍺")
		return
	}

	fmt.Println("User can't drink alcohol☠️")
}
