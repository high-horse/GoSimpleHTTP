package http1

type HandlerFunc func(request *Request) *Response

type Server struct {
	handlers map[string]HandlerFunc
}