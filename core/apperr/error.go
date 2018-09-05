package apperr

type AppError struct {
	Msg string
	Err error
}

func (e AppError) Error() string {
	return e.Err.Error()
}
