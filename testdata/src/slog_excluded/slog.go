package log

import (
	"context"
	"log/slog"
	"net/http"
)

func _() {
	ctx := context.Background()

	// All slog functions are excluded - no warnings expected
	slog.Debug("debug message", slog.String("key", "value"))
	slog.Warn("warn message", slog.String("key", "value"))
	slog.Error("error message", slog.String("key", "value"))
	slog.Info("info message", slog.String("key", "value"))

	// Context versions are always fine
	slog.DebugContext(ctx, "debug message", slog.String("key", "value"))
	slog.WarnContext(ctx, "warn message", slog.String("key", "value"))
	slog.ErrorContext(ctx, "error message", slog.String("key", "value"))
	slog.InfoContext(ctx, "info message", slog.String("key", "value"))

	// Logger methods are also excluded - no warnings expected
	l := slog.New(slog.NewTextHandler(nil, nil))
	l.Debug("debug message", slog.String("key", "value"))
	l.Warn("warn message", slog.String("key", "value"))
	l.Error("error message", slog.String("key", "value"))
	l.Info("info message", slog.String("key", "value"))

	// But other packages should still be checked
	http.Get("https://example.com") // want `net/http.Get must not be called`
}