package http1

import (
	"encoding/json"
	"fmt"
	"strings"
)


type Response struct {
	Status string
	Headers map[string]string
	Body string
}

// NewTextResponse creates a text response
func NewTextResponse(statusCode int, body string) *Response {
	status := fmt.Sprintf("HTTP/1.1 %d %s", statusCode, statusText(statusCode))
	return &Response{
		Status: status,
		Headers: map[string]string{"Content-Type": "text/plain", "Content-length": fmt.Sprintf("%d",len(body))},
		Body: body,
	}
}

func NewJSONResponse(statusCode int, data interface{}) (*Response, error) {
	body, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	status := fmt.Sprintf("HTTP/1.1 %d %s", statusCode, statusText(statusCode))
	return &Response{
		Status: status,
		Headers: map[string]string{"Content-Type": "application/json", "Content-length": fmt.Sprintf("%d",len(body))},
		Body: string(body),
	}, nil
}

// ToString converts the Response to an HTTP formatted string
func (r *Response) ToString() string {
	var headers []string
	for key, value := range r.Headers {
		headers = append(headers, fmt.Sprintf("%s: %s", key, value))
	}
	
	return fmt.Sprintf("%s\r\n%s\r\n\r\n%s", r.Status, strings.Join(headers, "\r\n"), r.Body)
}

func statusText(code int) string {
	switch code {
	case 200:
		return "OK"
	case 404:
		return "Not Found"
	case 500:
		return "Internal Server Error"
	default:
		return "Unknown Status"
	}
}
