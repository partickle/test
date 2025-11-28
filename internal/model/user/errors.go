package user

import "errors"

var (
	ErrTeamNotExists = errors.New("that team does not exist")
	ErrUserNotFound  = errors.New("user nod not found")
)
