package model

// R 常规的响应类型
type R[T any] struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    T      `json:"data,omitempty"`
}

// PR 分页的响应类型
type PR[T any] struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Page    int    `json:"page"`
	Count   int    `json:"count"`
	Total   int    `json:"total"`
	Data    []T    `json:"data,omitempty"`
}

type RBuilder struct {
}

func Success[T any]() *R[T] {
	r := &R[T]{}
	return r
}

func SuccessGeneric[T any](data T) *R[T] {
	r := &R[T]{}
	r.Data = data
	return r
}

func Failed[T any](code int, message string) *R[T] {
	r := &R[T]{}
	r.Code = code
	r.Message = message
	return r
}

func SuccessPage[T any](data []T, page int, count int, total int) *PR[T] {
	r := &PR[T]{}
	r.Page = page
	r.Count = count
	r.Total = total
	r.Data = data
	return r
}

func FailedPage[T any](code int, message string) *PR[T] {
	r := &PR[T]{}
	r.Code = code
	r.Message = message
	return r
}
