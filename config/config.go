package config

import "os"

var (
	FilePath = getEnvOrDefault("PROGNOS_FILE_PATH", "/tmp/prognos_data")
	UseSQL   = getEnvOrDefaultBool("PROGNOS_USE_SQL", false)
	DBHost   = getEnvOrDefault("PROGNOS_DB_HOST", "localhost")
	DBDB     = getEnvOrDefault("PROGNOS_DB_DB", "")
	DBUser   = getEnvOrDefault("PROGNOS_DB_USER", "")
	DBPass   = getEnvOrDefault("PROGNOS_DB_PASS", "")
)

func getEnvOrDefault(envVar string, defEnvVar string) (newEnvVar string) {
	if newEnvVar = os.Getenv(envVar); len(newEnvVar) == 0 {
		return defEnvVar
	} else {
		return newEnvVar
	}
}

func getEnvOrDefaultBool(envVar string, defEnvVar bool) (newEnvVar bool) {
	newEnvVarStr := os.Getenv(envVar)
	if len(newEnvVarStr) == 0 {
		return defEnvVar
	}
	return newEnvVarStr == "true"
}
