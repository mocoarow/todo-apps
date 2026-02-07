package handler_test

import (
	"testing"

	"github.com/mocoarow/todo-apps/backend-gin-gorm/controller/handler"
)

var config *handler.Config

func TestMain(m *testing.M) {
	config = &handler.Config{
		CORS: &handler.CORSConfig{
			AllowOrigins: "*",
			AllowMethods: "GET,POST,PUT,DELETE",
		},
		Log:   &handler.LogConfig{},
		Debug: &handler.DebugConfig{},
	}
	m.Run()
}
