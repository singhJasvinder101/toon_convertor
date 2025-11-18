# TOON Converter

A Go package for converting JSON to TOON (Token-Oriented Object Notation) format, optimized for Large Language Models.

## Features

- **Token Efficient**: 30-60% fewer tokens compared to JSON
- **LLM Optimized**: Designed for AI comprehension
- **Modular Design**: Clean package structure with separation of concerns
- **Edge Case Handling**: Robust handling of various JSON structures
- **Customizable**: Configurable indent size and other options



## Installation

### As a Library (for Go projects)
```bash
go get github.com/singhJasvinder101/toon_cli/converter
```

### As a CLI Tool
```bash
go install github.com/singhJasvinder101/toon_cli
```

## Usage (as a library)

```go
package main

import (
    "fmt"
    "github.com/singhJasvinder101/toon_cli/converter"
)

func main() {
    conv := converter.NewToonConverter()
    
    json := `{"users": [{"id": 1, "name": "Alice"}, {"id": 2, "name": "Bob"}]}`
    toon, err := conv.Convert(json)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }
    
    fmt.Println(toon)
    // Output: users[2]{id,name}:
    //           1,Alice
    //           2,Bob
}
```

### Custom Options

```go
opts := converter.ConversionOptions{
    IndentSize: 4,
    SortKeys:   true,
}
conv := converter.NewToonConverterWithOptions(opts)
```

### Multiple Input Types

```go
conv := converter.NewToonConverter()

// from JSON string
toon1, err := conv.Convert(jsonString)

// from bytes
toon2, err := conv.ConvertFromBytes(jsonBytes)

// from Go interface{}
toon3, err := conv.ConvertFromInterface(data)
```

## Usage (as a cli)
```bash
echo '{"name": "Alice", "age": 30}' | toon_cli

toon_cli -input data.json

# convert from file to file
toon_cli -input data.json -output data.toon

toon_cli -input data.json -output data.toon -indent 4
```

## JSON vs TOON Comparison

### JSON (51 tokens)
```json
{
  "users": [
    {"id": 1, "name": "Alice", "role": "admin"},
    {"id": 2, "name": "Bob", "role": "user"}
  ]
}
```

### TOON (24 tokens - 53% reduction)
```
users[2]{id,name,role}:
  1,Alice,admin
  2,Bob,user
```

## Running Examples

```bash
go run main.go
```

This will run various examples demonstrating:
- Simple objects
- Array of objects (tabular format)
- Complex nested structures
- Edge cases (empty objects/arrays, special characters)
- Custom options

## API Reference

### Core Types

- `ToonConverter`: Main converter struct
- `ConversionOptions`: Configuration options

### Main Methods

- `NewToonConverter()`: Create converter with default options
- `NewToonConverterWithOptions(opts)`: Create converter with custom options
- `Convert(jsonStr string)`: Convert JSON string to TOON
- `ConvertFromBytes(data []byte)`: Convert JSON bytes to TOON
- `ConvertFromInterface(data interface{})`: Convert Go data to TOON

### Features Handled

- Objects and arrays
- Nested structures
- Primitive types (strings, numbers, booleans, null)
- Empty objects and arrays
- Special characters in strings
- Mixed-type arrays
- Consistent key ordering

## Testing

The package handles various edge cases including:
- Empty JSON objects `{}`
- Empty arrays `[]`
- Mixed primitive arrays `[1, "hello", true, null]`
- Strings with special characters and newlines
- Deeply nested structures
- Large datasets with consistent object arrays