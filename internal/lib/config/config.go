package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/EagleLizard/ezd-daemon/internal/lib/constants"
	"github.com/joho/godotenv"
)

type EzdDConfigType struct {
	Host               string
	Port               string
	EzdGhWebhookSecret string
}

var EzdDConfig *EzdDConfigType

func init() {
	baseDir := constants.BaseDir()
	dotenvFilePath := filepath.Join(baseDir, ".env")
	err := godotenv.Load(dotenvFilePath)
	if err != nil {
		log.Fatal(err)
	}
	cfg := EzdDConfigType{
		Host:               getEnvVarOrDefault("EZD_D_HOST", "0.0.0.0"),
		Port:               getEnvVarOrDefault("EZD_D_PORT", "4440"),
		EzdGhWebhookSecret: getEnvVarOrDefault("EZD_GH_WEBHOOK_SECRET", ""),
	}
	EzdDConfig = &cfg
}
func getEnvVarOrDefault(envKey string, defaultVal string) string {
	envVal := os.Getenv(envKey)
	if len(envVal) == 0 {
		return defaultVal
	}
	return envVal
}
