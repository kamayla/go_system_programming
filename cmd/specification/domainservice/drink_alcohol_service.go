package domainservice

import (
	"fmt"
	"go_system_programming/cmd/specification/domain"
	"go_system_programming/cmd/specification/spec"
)

type DrinkAlcoholService struct {
	CanDrinkAlcoholSpecification spec.Specification
}

// 飲めるようなら飲む
func (s DrinkAlcoholService) DrinkAlcohol(user domain.User, subscription domain.Subscription) {
	if s.CanDrinkAlcoholSpecification.IsSatisfied() {
		fmt.Println("User can drink alcohol🍺🍺🍺")
		return
	}

	fmt.Println("User can't drink alcohol☠️")
}
