package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/singhJasvinder101/toon_convertor/converter"
)

func main() {
	var (
		inputFile  = flag.String("input", "", "Input JSON file (default: stdin)")
		outputFile = flag.String("output", "", "Output TOON file (default: stdout)")
		indentSize = flag.Int("indent", 2, "Indent size for output")
		showHelp   = flag.Bool("help", false, "Show help message")
	)
	flag.Parse()

	if *showHelp {
		showUsage()
		return
	}

	opts := converter.ConversionOptions{
		IndentSize: *indentSize,
		SortKeys:   true,
	}
	conv := converter.NewToonConverterWithOptions(opts)

	var input string
	var err error

	if *inputFile != "" {
		input, err = readFile(*inputFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error when reading input file %v\n", err)
			os.Exit(1)
		}
	} else {
		input, err = readStdin()
		if err != nil {
			fmt.Fprintf(os.Stderr, "error when reading stdin %v\n", err)
			os.Exit(1)
		}
	}

	result, err := conv.Convert(strings.TrimSpace(input))
	if err != nil {
		fmt.Fprintf(os.Stderr, "conversion error %v\n", err)
		os.Exit(1)
	}

	if *outputFile != "" {
		err = writeFile(*outputFile, result)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error when writing output file %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("successfully converted JSON to TOON format in %s\n", *outputFile)
	} else {
		fmt.Println(result)
	}
}

func showUsage() {
	fmt.Println("TOON Converter - Convert JSON to TOON format")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  toon-cli [options]")
	fmt.Println()
	fmt.Println("Options:")
	fmt.Println("  -input string    Input JSON file (default: stdin)")
	fmt.Println("  -output string   Output TOON file (default: stdout)")
	fmt.Println("  -indent int      Indent size for output (default: 2)")
	fmt.Println("  -help           Show this help message")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  echo '{\"name\":\"Alice\"}' | toon-cli")
	fmt.Println("  toon-cli -input data.json -output data.toon")
	fmt.Println("  toon-cli -input data.json -indent 4")
}

func readFile(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func readStdin() (string, error) {
	var lines []string
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return strings.Join(lines, "\n"), nil
}

func writeFile(filename, content string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.WriteString(file, content)
	return err
}
