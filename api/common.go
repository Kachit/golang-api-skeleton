package api

type ResponseBodyInterface interface {
	IsSuccess() bool
	IsFailed() bool
}

type ResponseBodyMetaInterface interface {
	GetMetaType() string
}

type ResponseBodyMetaPagination struct {
	Total int64 `json:"total"`
}

func (p *ResponseBodyMetaPagination) GetMetaType() string {
	return "pagination"
}

type ResponseBody struct {
	Result bool                      `json:"result"`
	Data   interface{}               `json:"data"`
	Meta   ResponseBodyMetaInterface `json:"meta"`
	Error  string                    `json:"error"`
}

func NewResponseBody(data interface{}) *ResponseBody {
	return &ResponseBody{Result: true, Data: data}
}

func NewResponseBodyCollection(data interface{}, total int64) *ResponseBody {
	meta := &ResponseBodyMetaPagination{Total: total}
	rb := NewResponseBody(data)
	rb.Meta = meta
	return rb
}

func NewResponseBodyError(error error) ResponseBody {
	return ResponseBody{Result: false, Error: error.Error()}
}
