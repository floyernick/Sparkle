package errors

import "errors"

var InternalError = errors.New("INTERNAL_ERROR")
var BadRequest = errors.New("BAD_REQUEST")
var InvalidParams = errors.New("INVALID_PARAMS")

var UsernameUsed = errors.New("USERNAME_USED")
var InvalidCredentials = errors.New("INVALID_CREDENTIALS")
