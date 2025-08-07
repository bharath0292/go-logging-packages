package main

import (
	"log/slog"
	"os"
)

func main() {
	// Json Logging
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger.Info("user login",
		slog.String("username", "gopher"),
		slog.Int("user_id", 123))

	// Text Logging
	logger = slog.New(slog.NewTextHandler(os.Stdout, nil))
	logger.Warn("user login",
		slog.String("username", "gopher"),
		slog.Int("user_id", 123))

}
