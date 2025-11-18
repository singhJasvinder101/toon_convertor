package converter

import (
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func NewToonConverter() *ToonConverter {
	return &ToonConverter{
		indentLevel: 0,
		indentSize:  DefaultIndentSize,
	}
}

func NewToonConverterWithOptions(opts ConversionOptions) *ToonConverter {
	return &ToonConverter{
		indentLevel: 0,
		indentSize:  opts.IndentSize,
	}
}

func (tc *ToonConverter) Convert(jsonStr string) (string, error) {
	if strings.TrimSpace(jsonStr) == "" {
		return "", fmt.Errorf("empty JSON input")
	}

	var data interface{}
	if err := json.Unmarshal([]byte(jsonStr), &data); err != nil {
		return "", fmt.Errorf("invalid JSON %v", err)
	}

	tc.indentLevel = 0
	result := tc.convertValue(data, "")
	return strings.TrimSpace(result), nil
}

func (tc *ToonConverter) convertValue(value interface{}, key string) string {
	switch v := value.(type) {
	case map[string]interface{}:
		return tc.convertObject(v, key)
	case []interface{}:
		return tc.convertArray(v, key)
	case string:
		return escapeString(v)
	case float64:
		if v == float64(int64(v)) {
			return strconv.FormatInt(int64(v), 10)
		}
		return strconv.FormatFloat(v, 'f', -1, 64)
	case bool:
		return strconv.FormatBool(v)
	case nil:
		return "null"
	default:
		return fmt.Sprintf("%v", v)
	}
}


func (tc *ToonConverter) convertObject(obj map[string]interface{}, parentKey string) string {
	if len(obj) == 0 {
		if parentKey != "" {
			return parentKey + ": {}"
		}
		return "{}"
	}

	var result strings.Builder
	indent := strings.Repeat(" ", tc.indentLevel*tc.indentSize)

	keys := make([]string, 0, len(obj))
	for k := range obj {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for i, key := range keys {
		value := obj[key]

		if i > 0 {
			result.WriteString("\n")
		}

		result.WriteString(indent)

		switch v := value.(type) {
		case map[string]interface{}:
			result.WriteString(key + ":")
			if len(v) == 0 {
				result.WriteString(" {}")
			} else {
				result.WriteString("\n")
				tc.indentLevel++
				result.WriteString(tc.convertValue(v, ""))
				tc.indentLevel--
			}
		case []interface{}:
			result.WriteString(tc.convertArray(v, key))
		default:
			result.WriteString(key + ": " + tc.convertValue(v, ""))
		}
	}

	return result.String()
}

func (tc *ToonConverter) convertArray(arr []interface{}, key string) string {
	if len(arr) == 0 {
		if key != "" {
			return key + "[0]: []"
		}
		return "[]"
	}

	if tc.isObjectArray(arr) {
		return tc.convertObjectArray(arr, key)
	}

	return tc.convertPrimitiveArray(arr, key)
}


func (tc *ToonConverter) convertObjectArray(arr []interface{}, key string) string {
	if len(arr) == 0 {
		return ""
	}

	var result strings.Builder

	firstObj := arr[0].(map[string]interface{})
	keys := make([]string, 0, len(firstObj))
	for k := range firstObj {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	if key != "" {
		result.WriteString(fmt.Sprintf("%s[%d]{%s}:", key, len(arr), strings.Join(keys, ",")))
	} else {
		result.WriteString(fmt.Sprintf("[%d]{%s}:", len(arr), strings.Join(keys, ",")))
	}

	tc.indentLevel++
	rowIndent := strings.Repeat(" ", tc.indentLevel*tc.indentSize)

	for _, item := range arr {
		obj := item.(map[string]interface{})
		result.WriteString("\n" + rowIndent)

		values := make([]string, len(keys))
		for i, k := range keys {
			value := obj[k]
			values[i] = tc.convertValue(value, "")
		}

		result.WriteString(strings.Join(values, ","))
	}

	tc.indentLevel--

	return result.String()
}


func (tc *ToonConverter) convertPrimitiveArray(arr []interface{}, key string) string {
	var result strings.Builder

	if key != "" {
		result.WriteString(fmt.Sprintf("%s[%d]: ", key, len(arr)))
	} else {
		result.WriteString(fmt.Sprintf("[%d]: ", len(arr)))
	}

	values := make([]string, len(arr))
	for i, item := range arr {
		values[i] = tc.convertValue(item, "")
	}

	result.WriteString("[" + strings.Join(values, ",") + "]")

	return result.String()
}





