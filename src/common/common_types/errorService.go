package commontypes

type ErrorService struct {
	HttpStatusCode int
	Code           error
	Message        string
}
