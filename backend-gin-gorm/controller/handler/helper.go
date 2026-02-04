package handler

import (
	"fmt"
	"math"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/mocoarow/todo-apps/backend-gin-gorm/api"
)

const (
	// ContextFieldUserID is the gin context key for the authenticated user's ID.
	ContextFieldUserID = "userId"
)

// InitRouterGroupFunc is a function that registers routes under a parent router group with optional middleware.
type InitRouterGroupFunc func(parentRouterGroup gin.IRouter, middleware ...gin.HandlerFunc)

// NewErrorResponse creates an ErrorResponse with the given error code and message.
func NewErrorResponse(code string, message string) *api.ErrorResponse {
	return &api.ErrorResponse{
		Code:    code,
		Message: message,
	}
}

// GetIntFromPath extracts an integer value from the URL path parameter with the given name.
func GetIntFromPath(c *gin.Context, param string) (int, error) {
	idS := c.Param(param)
	id, err := strconv.Atoi(idS)
	if err != nil {
		return 0, fmt.Errorf("convert string to int(%s): %w", idS, err)
	}

	return id, nil
}

func safeIntToInt32(v int) (int32, error) {
	if v < math.MinInt32 || v > math.MaxInt32 {
		return 0, fmt.Errorf("value %d overflows int32", v)
	}
	return int32(v), nil
}
