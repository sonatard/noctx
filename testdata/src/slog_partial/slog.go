package log

import (
	"context"
	"log/slog"
)

func _() {
	ctx := context.Background()

	// Only Debug methods are excluded
	slog.Debug("debug message", slog.String("key", "value")) // no warning - excluded
	slog.Warn("warn message", slog.String("key", "value"))   // want `log/slog.Warn must not be called`
	slog.Error("error message", slog.String("key", "value")) // want `log/slog.Error must not be called`
	slog.Info("info message", slog.String("key", "value"))   // want `log/slog.Info must not be called`

	// Context versions are always fine
	slog.DebugContext(ctx, "debug message", slog.String("key", "value"))
	slog.WarnContext(ctx, "warn message", slog.String("key", "value"))
	slog.ErrorContext(ctx, "error message", slog.String("key", "value"))
	slog.InfoContext(ctx, "info message", slog.String("key", "value"))

	// Logger methods - only Debug is excluded
	l := slog.New(slog.NewTextHandler(nil, nil))
	l.Debug("debug message", slog.String("key", "value")) // no warning - excluded
	l.Warn("warn message", slog.String("key", "value"))   // want `\(\*log/slog.Logger\).Warn must not be called`
	l.Error("error message", slog.String("key", "value")) // want `\(\*log/slog.Logger\).Error must not be called`
	l.Info("info message", slog.String("key", "value"))   // want `\(\*log/slog.Logger\).Info must not be called`
}