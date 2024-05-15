package httpmodels

type Error struct {
	Code    string
	Message string
}

func IncorrectReqBodyErr(e error) Error {
	return Error{
		Code:    "INCORRECT_REQUEST_BODY",
		Message: e.Error(),
	}

}

func ValidationFailedErr(e error) Error {
	return Error{
		Code:    "VALIDATION_FAILED",
		Message: e.Error(),
	}

}

func IncorrectParameter(e error) Error {
	return Error{
		Code:    "INCORRECT_PARAMETER",
		Message: e.Error(),
	}

}
func InternalServerErrorErr(e error) Error {
	return Error{
		Code:    "INTERNAL_ERROR",
		Message: e.Error(),
	}
}
