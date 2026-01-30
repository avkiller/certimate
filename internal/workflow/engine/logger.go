package engine

import (
	"log/slog"
)

type withLogger interface {
	SetLogger(logger *slog.Logger)
}
