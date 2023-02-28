package config

import "os"

var (
	FilePath = getEnvOrDefault("PROGNOS_FILE_PATH", "/tmp/prognos_data")
)

func getEnvOrDefault(envVar string, defEnvVar string) (newEnvVar string) {
	if newEnvVar = os.Getenv(envVar); len(newEnvVar) == 0 {
		return defEnvVar
	} else {
		return newEnvVar
	}
}
