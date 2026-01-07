package appgofig

import "strings"

func ShouldBeMasked(key string) bool {
	uppedKey := strings.ToUpper(key)

	if strings.HasSuffix(uppedKey, "PASSWORD") ||
		strings.HasSuffix(uppedKey, "TOKEN") ||
		strings.HasSuffix(uppedKey, "API_KEY") ||
		strings.HasSuffix(uppedKey, "SECRET") {
		return true
	}

	return false
}
