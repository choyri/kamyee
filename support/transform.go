package support

import (
	"strconv"
	"strings"
)

func Str2Int64(value string) int64 {
	ret, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		ret = 0
	}

	return ret
}

func Str2Uint64(value string) uint64 {
	ret, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		ret = 0
	}

	return ret
}

func Str2Int(value string) int {
	if len(value) == 0 {
		return 0
	}

	ret, err := strconv.Atoi(value)
	if err != nil {
		return 0
	}

	return ret
}

func Str2Uint(value string) uint {
	ret, err := strconv.ParseUint(value, 10, 32)
	if err != nil {
		ret = 0
	}

	return uint(ret)
}

func Str2Bool(value string) bool {
	switch strings.ToLower(strings.TrimSpace(value)) {
	case "true", "yes", "ok", "y", "1":
		return true
	default:
		return false
	}
}

func Int642Str(value int64) string {
	return strconv.FormatInt(value, 10)
}

func Uint642Str(value uint64) string {
	return strconv.FormatUint(value, 10)
}

func Int2Str(value int) string {
	return strconv.Itoa(value)
}

func Uint2Str(value uint) string {
	return strconv.FormatUint(uint64(value), 10)
}
