package spec

import "go_system_programming/cmd/specification/domain"

// サブスクリプションが支払い済みかどうかをチェックする仕様
type SubscriptionPaidSpecification struct {
	Subscription *domain.Subscription
}

func (s SubscriptionPaidSpecification) IsSatisfied() bool {
	return s.Subscription.IsPaid
}
