package internal

import "os"

func EnvOrDefault(env string, def string) string {
	e, ok := os.LookupEnv(env)
	if ok {
		return e
	}
	return def
}
