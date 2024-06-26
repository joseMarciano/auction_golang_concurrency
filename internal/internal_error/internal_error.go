package internal_error

type InternalError struct {
	Message string
	Err     string
}

func (e *InternalError) Error() string {
	return e.Message
}

func NewNotFoundError(message string) *InternalError {
	return &InternalError{
		Message: message,
		Err:     "not found",
	}
}

func NewInternalError(message string) *InternalError {
	return &InternalError{
		Message: message,
		Err:     "internal error",
	}
}
