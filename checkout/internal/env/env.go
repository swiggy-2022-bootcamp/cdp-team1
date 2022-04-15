package env

import "os"

//GetEnv ..
func GetEnv(env, defaultValue string) string {
	environment := os.Getenv(env)
	if environment == "" {
		return defaultValue
	}

	return environment
}
