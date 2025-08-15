package domain

type SubscriptionStatus string

const (
	SubscriptionStatusActive   SubscriptionStatus = "active"
	SubscriptionStatusInactive SubscriptionStatus = "inactive"
)

type Subscription struct {
	ID             int
	User           User
	Status         SubscriptionStatus
	SubscribeMonth int // サブスクリプションの月数
}

func (s *Subscription) IsLongTimeSubscribe() bool {
	return s.SubscribeMonth >= 12
}

func (s *Subscription) IsActive() bool {
	return s.Status == SubscriptionStatusActive
}
