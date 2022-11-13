package celestia

import (
	"encoding/base64"
	"encoding/hex"
)

// NamespaceID -
func NamespaceID(base64Value string) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(base64Value)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(decoded), nil
}
