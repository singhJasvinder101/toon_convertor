package converter

import (
	"testing"
)

func TestBasicConversion(t *testing.T) {
	conv := NewToonConverter()

	json := `{"name": "Alice", "age": 30}`
	expected := "age: 30\nname: Alice"

	result, err := conv.Convert(json)
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}

	if result != expected {
		t.Errorf("expected:\n%s\ngot:\n%s", expected, result)
	}
}

func TestArrayConversion(t *testing.T) {
	conv := NewToonConverter()

	json := `{"users": [{"id": 1, "name": "Alice"}, {"id": 2, "name": "Bob"}]}`

	result, err := conv.Convert(json)
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}

	if !contains(result, "users[2]{id,name}:") {
		t.Errorf("expected tabular format not found in %s", result)
	}

	if !contains(result, "1,Alice") || !contains(result, "2,Bob") {
		t.Errorf("expected data rows not found in %s", result)
	}
}

func TestEmptyInput(t *testing.T) {
	conv := NewToonConverter()

	_, err := conv.Convert("")
	if err == nil {
		t.Error("expected error for empty input")
	}
}

func TestInvalidJSON(t *testing.T) {
	conv := NewToonConverter()

	_, err := conv.Convert("{invalid json}")
	if err == nil {
		t.Error("expected error for invalid JSON")
	}
}

func TestSpecialCharacters(t *testing.T) {
	conv := NewToonConverter()

	json := `{"message": "Hello, \"world\"!\nNew line"}`

	result, err := conv.Convert(json)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if !contains(result, "\\\"") || !contains(result, "\\n") {
		t.Errorf("special characters not properly escaped %s", result)
	}
}

func TestCustomOptions(t *testing.T) {
	opts := ConversionOptions{IndentSize: 4, SortKeys: true}
	conv := NewToonConverterWithOptions(opts)

	json := `{"z": "last", "a": "first"}`

	result, err := conv.Convert(json)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if !contains(result, "a: first") {
		t.Errorf("keys not sorted properly %s", result)
	}
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) &&
		(s == substr ||
			s[:len(substr)] == substr ||
			s[len(s)-len(substr):] == substr ||
			containsAt(s, substr))
}

func containsAt(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
