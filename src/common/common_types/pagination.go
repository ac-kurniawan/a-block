package commontypes

type Pagination[T any] struct {
	NextPage *string `json:"nextPage" bson:"nextPage"`
	Size     uint64  `json:"size" bson:"size"`
	Total    uint64  `json:"total" bson:"total"`
	Contents T       `json:"contents" bson:"contents"`
}

type Page[T any] interface {
	Result() Pagination[T]
}

func (p Pagination[T]) Result() Pagination[T] {
	return p
}

func NewPagination[T any](module Pagination[T]) Page[T] {
	return &module
}
