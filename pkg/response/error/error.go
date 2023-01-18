package errresp

import "github.com/gin-gonic/gin"

// Response holds the information about the error
type Response struct {
	StatusCode int
	Status     bool
	Message    error
}

// New constructor for ErrorResponse
func New(StatusCode int, Status bool, Message error) *Response {
	return &Response{
		StatusCode,
		Status,
		Message,
	}
}

// ConvertMap Converting data before sending to front
func (r *Response) ConvertMap() map[string]any {
	data := make(map[string]any)

	data["status"] = r.Status
	data["data"] = r.StatusCode
	data["error"] = r.Message.Error()

	return data
}

// SendResponse method to send the data
func (r *Response) SendResponse(ctx *gin.Context) {
	ctx.JSON(r.StatusCode, r.ConvertMap())
}
