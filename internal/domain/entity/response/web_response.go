package response

type WebResponse[T any] struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Data   T      `json:"data"`
}

type WebResponseError struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Error  string `json:"error"`
}
