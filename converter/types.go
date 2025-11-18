package converter

type ToonConverter struct {
	indentLevel int
	indentSize  int
}

const (
	DefaultIndentSize = 2
)

type ConversionOptions struct {
	IndentSize int
	SortKeys   bool
}

func DefaultOptions() ConversionOptions {
	return ConversionOptions{
		IndentSize: DefaultIndentSize,
		SortKeys:   true,
	}
}
