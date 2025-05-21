package banana

import "errors"

var (
	UserConflict = errors.New("User already exists")
	SignUpFail   = errors.New("Registration failed")
)
