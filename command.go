package command

import (
	"strings"

	yup "github.com/gloo-foo/framework"
)

type command yup.Inputs[yup.File, flags]

func Cut(parameters ...any) yup.Command {
	cmd := command(yup.Initialize[yup.File, flags](parameters...))
	if cmd.Flags.Delimiter == "" {
		cmd.Flags.Delimiter = "\t"
	}
	return cmd
}

func (p command) Executor() yup.CommandExecutor {
	return yup.Inputs[yup.File, flags](p).Wrap(
		yup.LineTransform(func(line string) (string, bool) {
		// Handle fields mode
		if len(p.Flags.Fields) > 0 {
			delim := string(p.Flags.Delimiter)

			// Check if line contains delimiter
			if bool(p.Flags.OnlyDelimited) && !strings.Contains(line, delim) {
				return "", false
			}

			fields := strings.Split(line, delim)
			var selected []string

			for _, fieldNum := range p.Flags.Fields {
				idx := fieldNum - 1 // Fields are 1-indexed
				if idx >= 0 && idx < len(fields) {
					selected = append(selected, fields[idx])
				}
			}

			return strings.Join(selected, delim), true
		}

		// Handle characters mode
		if len(p.Flags.Chars) > 0 {
			runes := []rune(line)
			var selected []rune

			for _, charNum := range p.Flags.Chars {
				idx := charNum - 1 // Chars are 1-indexed
				if idx >= 0 && idx < len(runes) {
					selected = append(selected, runes[idx])
				}
			}

			return string(selected), true
		}

		// Handle bytes mode
		if len(p.Flags.Bytes) > 0 {
			bytes := []byte(line)
			var selected []byte

			for _, byteNum := range p.Flags.Bytes {
				idx := byteNum - 1 // Bytes are 1-indexed
				if idx >= 0 && idx < len(bytes) {
					selected = append(selected, bytes[idx])
				}
			}

			return string(selected), true
		}

		return line, true
	}).Executor(),
	)
}
