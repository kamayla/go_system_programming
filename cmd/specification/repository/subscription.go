package repository

import (
	"go_system_programming/cmd/specification/domain"
	"go_system_programming/cmd/specification/domainservice"
)

type subscriptionRepository struct{}

func (*subscriptionRepository) FindSubscriptionByUserID(userID int) (*domain.Subscription, error) {
	return &domain.Subscription{
		ID:             1,
		User:           domain.User{ID: userID},
		Status:         domain.SubscriptionStatusActive,
		SubscribeMonth: 30,
	}, nil
}

func NewSubscriptionRepository() domainservice.ISubscriptionRepository {
	return &subscriptionRepository{}
}
