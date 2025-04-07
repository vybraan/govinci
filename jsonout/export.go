package jsonout

import (
	"encoding/json"
	"govinci/core"
)

func Export(node *core.Node) string {
	bytes, err := json.MarshalIndent(node, "", "  ")
	if err != nil {
		return "{}"
	}
	return string(bytes)
}
