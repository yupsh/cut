package command

type Delimiter string
type Fields []int
type Chars []int
type Bytes []int

type OnlyDelimitedFlag bool

const (
	OnlyDelimited    OnlyDelimitedFlag = true
	NotOnlyDelimited OnlyDelimitedFlag = false
)

type flags struct {
	Delimiter     Delimiter
	Fields        Fields
	Chars         Chars
	Bytes         Bytes
	OnlyDelimited OnlyDelimitedFlag
}

func (d Delimiter) Configure(flags *flags)         { flags.Delimiter = d }
func (f Fields) Configure(flags *flags)            { flags.Fields = f }
func (c Chars) Configure(flags *flags)             { flags.Chars = c }
func (b Bytes) Configure(flags *flags)             { flags.Bytes = b }
func (f OnlyDelimitedFlag) Configure(flags *flags) { flags.OnlyDelimited = f }
