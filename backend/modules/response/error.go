package response

type CommonError struct {
	IsSuccessful bool   `json:"is_successful"`
	Message      string `json:"message"`
}

type ValidationError struct {
	IsSuccessful bool                   `json:"is_successful"`
	Message      string                 `json:"message"`
	Errors       map[string]interface{} `json:"errors"`
}

func NewError(msg string) CommonError {
	return CommonError{
		IsSuccessful: false,
		Message:      msg,
	}
}

func NewValidationError(key string, err error) ValidationError {
	res := ValidationError{}
	res.Errors = make(map[string]interface{})
	res.Errors[key] = err.Error()
	return res
}

const (
	InvalidValueError = "指定した値が正しくありません"
	BadAccessError    = "不正な値です"
)
