package form

import (
	"encoding/json"
	"fmt"
)

func tojson(ErrMsg string) string {
	if ErrMsg != "" {
		output, _ := json.Marshal(jsonformat{ErrMsg: ErrMsg})
		return tostring(output)
	}
	return ""
}

func tostring(output []byte) string {
	return fmt.Sprintf("%s", output)
}
