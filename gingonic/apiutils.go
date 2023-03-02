package gingonic

import "github.com/gin-gonic/gin"

// JSONError format
type JSONError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// AbortWithError helper method
func AbortWithError(c *gin.Context, code int, message string) {
	c.AbortWithStatusJSON(code, &JSONError{
		Code:    code,
		Message: message,
	})
}
