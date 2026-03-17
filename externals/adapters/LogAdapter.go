package adapters

import (
	"io"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/storybuilder/storybuilder/app/config"
)

// NewLogAdapter creates a new slog.Logger instance.
func NewLogAdapter(cfg config.LogConfig) (*slog.Logger, error) {
	var writers []io.Writer

	if cfg.Console {
		writers = append(writers, os.Stdout)
	}

	if cfg.File {
		err := os.MkdirAll(cfg.Directory, 0755)
		if err != nil {
			return nil, err
		}
		f, err := os.OpenFile(filepath.Join(cfg.Directory, "out.log"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			return nil, err
		}
		writers = append(writers, f)
	}

	if len(writers) == 0 {
		writers = append(writers, io.Discard)
	}

	multiWriter := io.MultiWriter(writers...)

	var level slog.Level
	switch cfg.Level {
	case "DEBUG":
		level = slog.LevelDebug
	case "INFO":
		level = slog.LevelInfo
	case "WARN":
		level = slog.LevelWarn
	case "ERROR":
		level = slog.LevelError
	default:
		level = slog.LevelInfo
	}

	opts := &slog.HandlerOptions{Level: level}

	var handler slog.Handler
	if cfg.Colors {
		handler = slog.NewTextHandler(multiWriter, opts)
	} else {
		handler = slog.NewJSONHandler(multiWriter, opts)
	}

	return slog.New(handler), nil
}
