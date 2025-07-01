package log

import (
	"context"
	"log/slog"
)

func _() {
	ctx := context.Background()

	// default logger
	slog.Debug("debug message", slog.String("key", "value")) // want `log/slog.Debug must not be called. use log/slog.DebugContext`
	slog.Warn("warn message", slog.String("key", "value"))   // want `log/slog.Warn must not be called. use log/slog.WarnContext`
	slog.Error("error message", slog.String("key", "value")) // want `log/slog.Error must not be called. use log/slog.ErrorContext`
	slog.Info("info message", slog.String("key", "value"))   // want `log/slog.Info must not be called. use log/slog.InfoContext`

	slog.DebugContext(ctx, "debug message", slog.String("key", "value"))
	slog.WarnContext(ctx, "warn message", slog.String("key", "value"))
	slog.ErrorContext(ctx, "error message", slog.String("key", "value"))
	slog.InfoContext(ctx, "info message", slog.String("key", "value"))

	slog.Default().Debug("debug message", slog.String("key", "value")) // want `\(\*log/slog.Logger\).Debug must not be called. use \(\*log/slog.Logger\).DebugContext`
	slog.Default().Warn("warn message", slog.String("key", "value"))   // want `\(\*log/slog.Logger\).Warn must not be called. use \(\*log/slog.Logger\).WarnContext`
	slog.Default().Error("error message", slog.String("key", "value")) // want `\(\*log/slog.Logger\).Error must not be called. use \(\*log/slog.Logger\).ErrorContext`
	slog.Default().Info("info message", slog.String("key", "value"))   // want `\(\*log/slog.Logger\).Info must not be called. use \(\*log/slog.Logger\).InfoContext`

	slog.Default().DebugContext(ctx, "debug message", slog.String("key", "value"))
	slog.Default().WarnContext(ctx, "warn message", slog.String("key", "value"))
	slog.Default().ErrorContext(ctx, "error message", slog.String("key", "value"))
	slog.Default().InfoContext(ctx, "info message", slog.String("key", "value"))

	// Logger struct
	l := slog.New(slog.NewTextHandler(nil, nil))
	l.Debug("debug message", slog.String("key", "value")) // want `\(\*log/slog.Logger\).Debug must not be called. use \(\*log/slog.Logger\).DebugContext`
	l.Warn("warn message", slog.String("key", "value"))   // want `\(\*log/slog.Logger\).Warn must not be called. use \(\*log/slog.Logger\).WarnContext`
	l.Error("error message", slog.String("key", "value")) // want `\(\*log/slog.Logger\).Error must not be called. use \(\*log/slog.Logger\).ErrorContext`
	l.Info("info message", slog.String("key", "value"))   // want `\(\*log/slog.Logger\).Info must not be called. use \(\*log/slog.Logger\).InfoContext`

	l.DebugContext(ctx, "debug message", slog.String("key", "value"))
	l.WarnContext(ctx, "warn message", slog.String("key", "value"))
	l.ErrorContext(ctx, "error message", slog.String("key", "value"))
	l.InfoContext(ctx, "info message", slog.String("key", "value"))
}
