// NOTE: Not well tested, just an illustration of what's possible
package main

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"log/slog"
	"os"
	"runtime"
	"time"

	"github.com/fatih/color"
)

type PrettyHandlerOptions struct {
	SlogOpts slog.HandlerOptions
}

type PrettyHandler struct {
	slog.Handler
	l *log.Logger
}

func (h *PrettyHandler) Handle(ctx context.Context, r slog.Record) error {
	level := r.Level.String() + ":"

	switch r.Level {
	case slog.LevelDebug:
		level = color.MagentaString(level)
	case slog.LevelInfo:
		level = color.BlueString(level)
	case slog.LevelWarn:
		level = color.YellowString(level)
	case slog.LevelError:
		level = color.RedString(level)
	}

	fields := make(map[string]interface{}, r.NumAttrs())
	r.Attrs(func(a slog.Attr) bool {
		fields[a.Key] = a.Value.Any()

		return true
	})

	b, err := json.MarshalIndent(fields, "", "  ")
	if err != nil {
		return err
	}

	timeStr := r.Time.Format(time.DateTime)
	msg := color.CyanString(r.Message)

	// Извлечение источника (имя файла, строка, функция)
	source := ""
	if r.PC != 0 {
		fs := runtime.CallersFrames([]uintptr{r.PC})
		frame, _ := fs.Next()
		source = color.GreenString("%s:%d (%s)", frame.File, frame.Line, frame.Function)
	}

	h.l.Println(timeStr, level, msg, source, string(b))

	return nil
}

func NewPrettyHandler(
	out io.Writer,
	opts PrettyHandlerOptions,
) *PrettyHandler {
	h := &PrettyHandler{
		Handler: slog.NewJSONHandler(out, &opts.SlogOpts),
		l:       log.New(out, "", 0),
	}

	return h
}
func main() {

	var err error
	opts := PrettyHandlerOptions{
		SlogOpts: slog.HandlerOptions{
			AddSource: true,
			Level:     slog.LevelDebug,
		},
	}
	handler := NewPrettyHandler(os.Stdout, opts)
	logger := slog.New(handler)

	logger.Debug(
		"executing database query",
		slog.String("query", "SELECT * FROM users"),
	)
	logger.Info("image upload successful", slog.String("image_id", "39ud88"))
	logger.Warn(
		"storage is 90% full",
		slog.String("available_space", "900.1 MB"),
	)
	logger.Error(
		"An error occurred while processing the request",
		slog.String("url", "https://example.com"),
	)

	logger.Error("err", "errkey", err)
}
