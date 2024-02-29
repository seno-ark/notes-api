// Package utils provides helper functions
package utils

// Response represents json api response
type Response struct {
	Message        string        `json:"message"`
	Data           any           `json:"data"`
	Meta           *ResponseMeta `json:"meta,omitempty"`
	ValidationErrs []string      `json:"validation_errors,omitempty"`
}

// ResponseMeta represents json api for meta response
type ResponseMeta struct {
	Total int64 `json:"total"`
	Page  int   `json:"page"`
	Count int   `json:"count"`
}

// NewResponse returns new Response struct
func NewResponse() *Response {
	return &Response{}
}

// Set add message and data value to existing Response struct and returns it
func (res *Response) Set(message string, data any) *Response {
	res.Message = message
	if data == nil {
		data = struct{}{}
	}
	res.Data = data
	return res
}

// AddMeta add page, count, and total value to existing ResponseMeta struct and returns it
func (res *Response) AddMeta(page, count int, total int64) *Response {
	res.Meta = &ResponseMeta{
		Total: total,
		Page:  page,
		Count: count,
	}

	return res
}

// AddErrValidation add list of errors generated from validator to existing ValidationErrs list and returns it
func (res *Response) AddErrValidation(errs []string) *Response {
	res.ValidationErrs = errs
	return res
}
