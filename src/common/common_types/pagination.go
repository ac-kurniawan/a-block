package commontypes

type Pagination[T any] struct {
	NextPage *string `json:"nextPage" bson:"nextPage"`
	Size     uint64  `json:"size" bson:"size"`
	Total    uint64  `json:"total" bson:"total"`
	Contents T       `json:"contents" bson:"contents"`
}
