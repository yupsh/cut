package opt

// Custom types for parameters
type Delimiter string
type Fields []int
type Chars []int
type Bytes []int

// Boolean flag types with constants
type OnlyDelimitedFlag bool

const (
	OnlyDelimited    OnlyDelimitedFlag = true
	NotOnlyDelimited OnlyDelimitedFlag = false
)

// Flags represents the configuration options for the cut command
type Flags struct {
	Delimiter     Delimiter         // Field delimiter
	Fields        Fields            // Field numbers to extract
	Chars         Chars             // Character positions to extract
	Bytes         Bytes             // Byte positions to extract
	OnlyDelimited OnlyDelimitedFlag // Only output lines with delimiter
}

// Configure methods for the opt system
func (d Delimiter) Configure(flags *Flags)         { flags.Delimiter = d }
func (f Fields) Configure(flags *Flags)            { flags.Fields = f }
func (c Chars) Configure(flags *Flags)             { flags.Chars = c }
func (b Bytes) Configure(flags *Flags)             { flags.Bytes = b }
func (f OnlyDelimitedFlag) Configure(flags *Flags) { flags.OnlyDelimited = f }
