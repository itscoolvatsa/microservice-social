package errors

import "errors"

var ErrNotFound = errors.New("not found")
var ErrInternalServer = errors.New("internal server error")
var ErrAlreadyExists = errors.New("user already exists")
var ErrInvalidCredentials = errors.New("invalid credentials")
var ErrNoAuthHeader = errors.New("authorization header does not exists")
var ErrInvalidAuthHeader = errors.New("invalid authorization header format")
var ErrInvalidAuthData = errors.New("invalid authorization header")
