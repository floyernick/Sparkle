package errors

import "errors"

var InternalError = errors.New("INTERNAL_ERROR")
var BadRequest = errors.New("BAD_REQUEST")
var InvalidParams = errors.New("INVALID_PARAMS")

var ActionNotAllowed = errors.New("ACTION_NOT_ALLOWED")

var UsernameUsed = errors.New("USERNAME_USED")
var InvalidCredentials = errors.New("INVALID_CREDENTIALS")

var PostNotFound = errors.New("POST_NOT_FOUND")
