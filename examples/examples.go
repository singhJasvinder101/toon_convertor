package examples

import (
	"fmt"

	"github.com/singhJasvinder101/toon_convertor/converter"
)

func RunAllExamples() {
	conv := converter.NewToonConverter()

	fmt.Println("=== example 1: Simple Object ===")
	jsonStr1 := `{"name": "Alice", "age": 30, "active": true}`
	toon1, err := conv.Convert(jsonStr1)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Println(toon1)
	}
	fmt.Println()

	fmt.Println("=== example 2: Array of Objects ===")
	jsonStr2 := `{
		"users": [
			{"id": 1, "name": "Alice", "role": "admin", "salary": 75000},
			{"id": 2, "name": "Bob", "role": "user", "salary": 65000},
			{"id": 3, "name": "Charlie", "role": "user", "salary": 70000}
		]
	}`
	toon2, err := conv.Convert(jsonStr2)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Println(toon2)
	}
	fmt.Println()

	fmt.Println("=== example 3: Complex Nested Structure ===")
	jsonStr3 := `{
		"company": "TechCorp",
		"departments": [
			{"name": "Engineering", "budget": 1000000, "headCount": 50},
			{"name": "Marketing", "budget": 500000, "headCount": 20}
		],
		"metadata": {
			"version": "1.0",
			"tags": ["tech", "startup"],
			"founded": 2020
		}
	}`
	toon3, err := conv.Convert(jsonStr3)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Println(toon3)
	}
	fmt.Println()

	fmt.Println("=== Example 4: Edge Cases ===")

	fmt.Println("Empty object:")
	toon4a, _ := conv.Convert(`{}`)
	fmt.Println(toon4a)

	fmt.Println("\nEmpty array:")
	toon4b, _ := conv.Convert(`{"items": []}`)
	fmt.Println(toon4b)

	fmt.Println("\nMixed array:")
	toon4c, _ := conv.Convert(`{"mixed": [1, "hello", true, null]}`)
	fmt.Println(toon4c)

	fmt.Println("\nSpecial characters:")
	toon4d, _ := conv.Convert(`{"message": "Hello, \"world\"!\nNew line here."}`)
	fmt.Println(toon4d)

	fmt.Println("\n=== Example 5: Custom Options ===")
	opts := converter.ConversionOptions{IndentSize: 4, SortKeys: true}
	customConv := converter.NewToonConverterWithOptions(opts)

	jsonStr5 := `{"z": "last", "a": "first", "m": "middle"}`
	toon5, _ := customConv.Convert(jsonStr5)
	fmt.Println("Custom indent (4 spaces):")
	fmt.Println(toon5)
}

func RunSimpleExample() {
	converter := converter.NewToonConverter()

	json := `{"users": [{"id": 1, "name": "Alice"}, {"id": 2, "name": "Bob"}]}`
	toon, err := converter.Convert(json)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Println("JSON:")
	fmt.Println(json)
	fmt.Println("\nTOON:")
	fmt.Println(toon)
}
