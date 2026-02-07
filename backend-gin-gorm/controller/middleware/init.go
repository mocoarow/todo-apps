package middleware

import (
	"go.opentelemetry.io/otel"
)

var (
	tracer = otel.Tracer("github.com/mocoarow/todo-apps/backend-gin-gorm/controller/middleware")
)
