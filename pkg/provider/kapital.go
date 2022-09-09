package provider

type KapitalResponseData[T any] struct {
	ResponseData T               `json:"responseData"`
	Response     KapitalResponse `json:"response"`
}
type KapitalResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
