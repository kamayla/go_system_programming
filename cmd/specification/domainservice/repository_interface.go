package domainservice

import "go_system_programming/cmd/specification/domain"

type IUserRepository interface {
	FindUserByID(id int) (*domain.User, error)
}

type ISubscriptionRepository interface {
	FindSubscriptionByUserID(userID int) (*domain.Subscription, error)
}
