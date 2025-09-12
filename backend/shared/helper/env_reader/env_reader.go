package envreader

import (
	"os"
)

func GetStringOrDefault(key, def string) string {

	if val, exists := os.LookupEnv(key); exists {
		return val
	}

	return def
}
