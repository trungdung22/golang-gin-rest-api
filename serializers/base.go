package serializers

type CommonError struct {
	Errors map[string]interface{} `json:"errors"`
}

func ResponseError(key string, err error) CommonError {
	res := CommonError{}
	res.Errors = make(map[string]interface{})
	res.Errors[key] = err.Error()
	return res
}
