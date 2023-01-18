package jsonresp

import "github.com/gin-gonic/gin"

// Response hold the information about the request
type Response struct {
	StatusCode int
	Status     bool
	Data       interface{}
	Message    string
}

// New constructor for Response
func New(StatusCode int, Status bool, Data interface{}, Message string) *Response {
	return &Response{
		StatusCode,
		Status,
		Data,
		Message,
	}
}

// ConvertMap Converting data before sending to front
func (r *Response) ConvertMap() map[string]any {
	data := make(map[string]any)

	data["status"] = r.Status
	data["data"] = r.Data
	data["message"] = r.Message

	return data
}

// SendResponse method to send the data
func (r *Response) SendResponse(ctx *gin.Context) {
	ctx.JSON(r.StatusCode, r.ConvertMap())
}
