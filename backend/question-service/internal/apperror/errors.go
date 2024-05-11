package apperror

import "errors"

var (
	ErrQuestionNotFound = errors.New("question not found")
	ErrInternalServer   = errors.New("internal server error")
)
