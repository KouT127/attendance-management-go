package responses

type CommonError struct {
	Errors map[string]interface{} `json:"errors"`
}

func NewError(key string, err error) CommonError {
	res := CommonError{}
	res.Errors = make(map[string]interface{})
	res.Errors[key] = err.Error()
	return res
}

const (
	InvalidValueError = "指定した値が正しくありません"
	BadAccessError    = "不正な値です"
)
