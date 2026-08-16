package domain

import (
	"fmt"
	core_errors "pollify/internal/core/errors"
	"regexp"
)

type User struct {
	ID          int
	FullName    string
	Email       string
	PhoneNumber *string
	Password    string
}

func NewUser(
	id int,
	fullName string,
	email string,
	phoneNumber *string,
	password string,
) User {
	return User{
		ID:          id,
		FullName:    fullName,
		Email:       email,
		PhoneNumber: phoneNumber,
		Password:    password,
	}
}

func NewUserUnInitialized(
	fullName string,
	email string,
	phoneNumber *string,
	password string,
) User {
	return NewUser(
		UnInitializedID,
		fullName,
		email,
		phoneNumber,
		password,
	)
}

func (u *User) Validate() error {
	fullNameLength := len([]rune(u.FullName))
	if fullNameLength < 3 || fullNameLength > 100 {
		return fmt.Errorf(
			"invalid user `FullName` length: %d: %w",
			fullNameLength,
			core_errors.ErrInvalidArgument,
		)
	}
	if _, err := regexp.MatchString("@", u.Email); err != nil {
		return fmt.Errorf("invalid user `Email`: %w", err)
	}

	if u.PhoneNumber != nil {
		phoneNumberLength := len([]rune(*u.PhoneNumber))
		if phoneNumberLength < 10 || phoneNumberLength > 15 {
			return fmt.Errorf(
				"invalid `PhoneNumber` lenght: %d: %w",
				phoneNumberLength,
				core_errors.ErrInvalidArgument,
			)
		}

		re := regexp.MustCompile(`^\+[0-9]+$`)

		if !re.MatchString(*u.PhoneNumber) {
			return fmt.Errorf(
				"invalid `PhoneNumber` format: %w",
				core_errors.ErrInvalidArgument,
			)
		}
	}
	return nil
}

type UserPatch struct {
	FullName    Nullable[string]
	Email       Nullable[string]
	PhoneNumber Nullable[string]
}

func NewUserPatch(
	fullName Nullable[string],
	email Nullable[string],
	phoneNumber Nullable[string],
) UserPatch {
	return UserPatch{
		FullName:    fullName,
		Email:       email,
		PhoneNumber: phoneNumber,
	}
}

func (p *UserPatch) Validate() error {
	if p.FullName.Set && p.FullName.Value == nil {
		return fmt.Errorf(
			"`FullName` can't be patched to NULL: %w",
			core_errors.ErrInvalidArgument,
		)
	}
	return nil
}

func (u *User) ApplyPatch(patch UserPatch) error {
	if err := patch.Validate(); err != nil {
		return fmt.Errorf("validate user patch: %w", err)
	}
	tmp := *u

	if patch.FullName.Set {
		tmp.FullName = *patch.FullName.Value
	}
	if patch.PhoneNumber.Set {
		tmp.PhoneNumber = patch.PhoneNumber.Value
	}

	if err := tmp.Validate(); err != nil {
		return fmt.Errorf("validate patched user: %w", err)
	}

	*u = tmp

	return nil
}

type UserAuthorize struct {
	Email    string
	Password string
}

func NewAuthorizeUser(email, password string) UserAuthorize {
	return UserAuthorize{
		Email:    email,
		Password: password,
	}
}
