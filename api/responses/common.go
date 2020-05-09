package responses

type CommonResponses struct {
	IsSuccessful bool `json:"is_successful"`
	HasNext      bool `json:"has_next"`
}

type CommonResponse struct {
	IsSuccessful bool `json:"is_successful"`
}
