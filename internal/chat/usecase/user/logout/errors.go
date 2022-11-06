package logout

import "errors"

var (
	ErrSessionNotFound  = errors.New("session  is not found")
	ErrSessionNotActive = errors.New("session is not active")
)
