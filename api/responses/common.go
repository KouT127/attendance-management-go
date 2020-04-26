package responses

type CommonResponses struct {
	IsSuccessful bool `json:"isSuccessful"`
	HasNext      bool `json:"hasNext"`
}

type CommonResponse struct {
	IsSuccessful bool `json:"isSuccessful"`
}
