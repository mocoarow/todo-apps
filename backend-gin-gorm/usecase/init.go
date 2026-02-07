// Package usecase implements application use cases following the CQRS pattern.
// Commands handle state mutations (create, update, delete) and queries handle reads.
// Each use case is composed of fine-grained command/query objects injected via constructors.
package usecase

import (
	"go.opentelemetry.io/otel"
)

var (
	tracer = otel.Tracer("github.com/mocoarow/todo-apps/backend-gin-gorm/usecase")
)
