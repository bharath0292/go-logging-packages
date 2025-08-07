package main

import (
	util "logging/zerolog-mysetup/pkg/utils"

	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {

	devMode := os.Getenv("ENV") == "dev"
	logLevel := zerolog.InfoLevel

	if devMode {
		// To set log level
		logLevel = zerolog.DebugLevel

		// To set global timezone early
		os.Setenv("TZ", "UTC")

		// Read env files from current and parent directories
		cwd, err := os.Getwd()
		if err != nil {
			log.Fatal().Msgf("Could not get working directory: %v", err)
		}

		filesToTry := []string{
			filepath.Join(cwd, ".env"),
			filepath.Join(filepath.Dir(cwd), ".env"),
		}

		if err := godotenv.Load(filesToTry...); err != nil {
			log.Fatal().Err(err).Msg("Error loading .env files")
		}
	}

	util.InitLogger(logLevel)
}

func main() {
	log.Info().Str("service", "billing").Int("active_users", 3214).Msg("report generated")
}
