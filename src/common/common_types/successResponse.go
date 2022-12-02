package commontypes

type SuccessResponse[T any] struct {
	Status string `json:"status"`
	Code   string `json:"code"`
	Data   T      `json:"data"`
}
