package serializers

type CommonResponses struct {
	IsSuccessful bool `json:"isSuccessful"`
	HasNext      bool `json:"hasNext"`
}

type CommonResponse struct {
	IsSuccessful bool `json:"isSuccessful"`
}

type CommonError struct {
	Errors map[string]interface{} `json:"errors"`
}

func NewError(key string, err error) CommonError {
	res := CommonError{}
	res.Errors = make(map[string]interface{})
	res.Errors[key] = err.Error()
	return res
}
