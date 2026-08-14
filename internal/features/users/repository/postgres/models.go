package users_postgres_repository

import "pollify/internal/core/domain"

type UserModel struct {
	ID          int
	FullName    string
	Email       string
	PhoneNumber *string
}

func userDomainsFromModels(users []UserModel) []domain.User {
	userDomains := make([]domain.User, len(users))

	for i, user := range users {
		userDomains[i] = domain.NewUser(
			user.ID,
			user.FullName,
			user.Email,
			user.PhoneNumber,
		)
	}

	return userDomains
}
