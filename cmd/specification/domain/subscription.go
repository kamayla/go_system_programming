package domain

type Subscription struct {
	ID             int
	User           User
	IsPaid         bool
	SubscribeMonth int // サブスクリプションの月数
}

func (s *Subscription) IsLongTimeSubscribe() bool {
	return s.SubscribeMonth >= 12
}
