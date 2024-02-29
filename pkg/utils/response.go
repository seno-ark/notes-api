package utils

type Response struct {
	Message        string        `json:"message"`
	Data           any           `json:"data"`
	Meta           *ResponseMeta `json:"meta,omitempty"`
	ValidationErrs []string      `json:"validation_errors,omitempty"`
}

type ResponseMeta struct {
	Total int64 `json:"total"`
	Page  int   `json:"page"`
	Count int   `json:"count"`
}

func NewResponse() *Response {
	return &Response{}
}

func (res *Response) Set(message string, data any) *Response {
	res.Message = message
	if data == nil {
		data = struct{}{}
	}
	res.Data = data
	return res
}

func (res *Response) AddMeta(page, count int, total int64) *Response {
	res.Meta = &ResponseMeta{
		Total: total,
		Page:  page,
		Count: count,
	}

	return res
}

func (res *Response) AddErrValidation(errs []string) *Response {
	res.ValidationErrs = errs
	return res
}
