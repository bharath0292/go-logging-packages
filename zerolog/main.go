package main

import (
	"os"

	"github.com/rs/zerolog"
)

func main() {
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	logger.Info().
		Str("service", "billing").
		Int("active_users", 3214).
		Msg("report generated")
}
