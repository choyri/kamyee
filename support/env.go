package support

import (
	"os"
	"strconv"
	"strings"
)

func GetBoolEnv(key string, defaultValue ...bool) bool {
	disposeDefault := func() bool {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return false
	}

	value, found := os.LookupEnv(key)
	if !found {
		return disposeDefault()
	}

	ret, err := strconv.ParseBool(strings.TrimSpace(value))
	if err != nil {
		return disposeDefault()
	}

	return ret
}

func GetInt64Env(key string, defaultValue ...int64) int64 {
	disposeDefault := func() int64 {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return 0
	}

	value, found := os.LookupEnv(key)
	if !found {
		return disposeDefault()
	}

	ret, err := strconv.ParseInt(strings.TrimSpace(value), 10, 64)
	if err != nil {
		return disposeDefault()
	}

	return ret
}

func GetStringEnv(key string, defaultValue ...string) string {
	disposeDefault := func() string {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return ""
	}

	ret, found := os.LookupEnv(key)
	if !found {
		return disposeDefault()
	}

	return strings.TrimSpace(ret)
}
