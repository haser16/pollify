package core_postgres_pool

import "errors"

var (
	ErrNoRows             = errors.New("no rows")
	ErrViolatesForeignKey = errors.New("foreign key violation")
	ErrUnknown            = errors.New("unknown")
)
