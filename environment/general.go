package environment

import "os"

var (
	RedisHost     = getEnv("REDIS_HOST", "localhost")
	RedisPort     = getEnv("REDIS_PORT", "6379")
	RedisPassword = getEnv("REDIS_PASSWORD", "")
)

func getEnv(key string, fallback string) string {
	value := os.Getenv(key)

	if value != "" {
		return value
	}

	return fallback
}
