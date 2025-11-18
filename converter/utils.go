package converter

import "strings"

func escapeString(s string) string {
	s = strings.ReplaceAll(s, "\\", "\\\\")
	s = strings.ReplaceAll(s, "\"", "\\\"")
	s = strings.ReplaceAll(s, "\n", "\\n")
	s = strings.ReplaceAll(s, "\r", "\\r")
	s = strings.ReplaceAll(s, "\t", "\\t")

	if strings.ContainsAny(s, ",\n\r\t\"\\") || strings.TrimSpace(s) != s {
		return "\"" + s + "\""
	}

	return s
}

func slicesEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func (tc *ToonConverter) isObjectArray(arr []interface{}) bool {
	if len(arr) == 0 {
		return false
	}

	var firstKeys []string
	for i, item := range arr {
		obj, ok := item.(map[string]interface{})
		if !ok {
			return false
		}

		if i == 0 {
			for k := range obj {
				firstKeys = append(firstKeys, k)
			}
			for j := 0; j < len(firstKeys); j++ {
				for k := j + 1; k < len(firstKeys); k++ {
					if firstKeys[j] > firstKeys[k] {
						firstKeys[j], firstKeys[k] = firstKeys[k], firstKeys[j]
					}
				}
			}
		} else {
			currentKeys := make([]string, 0, len(obj))
			for k := range obj {
				currentKeys = append(currentKeys, k)
			}
			for j := 0; j < len(currentKeys); j++ {
				for k := j + 1; k < len(currentKeys); k++ {
					if currentKeys[j] > currentKeys[k] {
						currentKeys[j], currentKeys[k] = currentKeys[k], currentKeys[j]
					}
				}
			}

			if !slicesEqual(firstKeys, currentKeys) {
				return false
			}
		}
	}

	return len(firstKeys) > 0
}
