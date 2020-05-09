package support

import (
	"bytes"
	"encoding/json"
)

func JSONIndent(data interface{}) []byte {
	tmp, _ := json.Marshal(data)

	var ret bytes.Buffer
	_ = json.Indent(&ret, tmp, "", "    ")

	return ret.Bytes()
}
