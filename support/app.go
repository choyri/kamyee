package support

import "github.com/choyri/kamyee"

func GetAppEnv() string {
	return GetStringEnv(kamyee.KeyAppEnv, kamyee.EnvLocal)
}

func GetAppName() string {
	return GetStringEnv(kamyee.KeyAppName)
}
