package spec

import "go_system_programming/cmd/specification/domain"

// ユーザーが朝食を食べられるかどうかをチェックする仕様
type CanEatBreakfastSpecification struct {
	User         *domain.User
	Subscription *domain.Subscription
}

func (s CanEatBreakfastSpecification) IsSatisfied() bool {
	if s.User == nil || s.Subscription == nil {
		return false
	}
	return !s.User.IsMale() && s.Subscription.IsLongTimeSubscribe()
}
