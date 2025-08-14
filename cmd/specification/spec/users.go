package spec

import "go_system_programming/cmd/specification/domain"

// ユーザーが成人かどうかをチェックする仕様
type UserAdultSpecification struct {
	User *domain.User
}

func (s UserAdultSpecification) IsSatisfied() bool {
	return s.User.Age >= 20
}

// ユーザーが男性かどうかをチェックする仕様
type UserMaleSpecification struct {
	User *domain.User
}

func (s UserMaleSpecification) IsSatisfied() bool {
	return s.User.Gender == domain.GenderMale
}
