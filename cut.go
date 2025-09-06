package cut

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"strings"

	yup "github.com/yupsh/framework"
	"github.com/yupsh/framework/opt"

	localopt "github.com/yupsh/cut/opt"
)

// Flags represents the configuration options for the cut command
type Flags = localopt.Flags

// Command implementation
type command opt.Inputs[string, Flags]

// Cut creates a new cut command with the given parameters
func Cut(parameters ...any) yup.Command {
	cmd := command(opt.Args[string, Flags](parameters...))
	// Set default delimiter to tab
	if cmd.Flags.Delimiter == "" {
		cmd.Flags.Delimiter = "\t"
	}
	return cmd
}

func (c command) Execute(ctx context.Context, stdin io.Reader, stdout, stderr io.Writer) error {
	return yup.ProcessFilesWithContext(
		ctx, c.Positional, stdin, stdout, stderr,
		yup.FileProcessorOptions{
			CommandName:     "cut",
			ContinueOnError: true,
		},
		func(ctx context.Context, source yup.InputSource, output io.Writer) error {
			return c.processReader(ctx, source.Reader, output)
		},
	)
}

func (c command) processReader(ctx context.Context, reader io.Reader, output io.Writer) error {
	scanner := bufio.NewScanner(reader)

	for yup.ScanWithContext(ctx, scanner) {
		line := scanner.Text()
		result := c.extractFields(ctx, line)

		// Skip lines without delimiter if OnlyDelimited is set
		if bool(c.Flags.OnlyDelimited) && !strings.Contains(line, string(c.Flags.Delimiter)) {
			continue
		}

		fmt.Fprintln(output, result)
	}

	// Check if context was cancelled
	if err := yup.CheckContextCancellation(ctx); err != nil {
		return err
	}

	return scanner.Err()
}

func (c command) extractFields(ctx context.Context, line string) string {
	if len(c.Flags.Fields) > 0 {
		return c.extractFieldsByNumber(ctx, line)
	}

	if len(c.Flags.Chars) > 0 {
		return c.extractCharsByPosition(ctx, line)
	}

	if len(c.Flags.Bytes) > 0 {
		return c.extractBytesByPosition(ctx, line)
	}

	// Default: return whole line
	return line
}

func (c command) extractFieldsByNumber(ctx context.Context, line string) string {
	delimiter := string(c.Flags.Delimiter)
	fields := strings.Split(line, delimiter)

	var result []string
	for i, fieldNum := range c.Flags.Fields {
		// Check for cancellation periodically when processing many fields
		if i%100 == 0 {
			if err := yup.CheckContextCancellation(ctx); err != nil {
				return "" // Return empty on cancellation
			}
		}

		if fieldNum > 0 && fieldNum <= len(fields) {
			result = append(result, fields[fieldNum-1]) // 1-based indexing
		}
	}

	return strings.Join(result, delimiter)
}

func (c command) extractCharsByPosition(ctx context.Context, line string) string {
	runes := []rune(line)
	var result []rune

	for i, pos := range c.Flags.Chars {
		// Check for cancellation periodically when processing many character positions
		if i%100 == 0 {
			if err := yup.CheckContextCancellation(ctx); err != nil {
				return "" // Return empty on cancellation
			}
		}

		if pos > 0 && pos <= len(runes) {
			result = append(result, runes[pos-1]) // 1-based indexing
		}
	}

	return string(result)
}

func (c command) extractBytesByPosition(ctx context.Context, line string) string {
	bytes := []byte(line)
	var result []byte

	for i, pos := range c.Flags.Bytes {
		// Check for cancellation periodically when processing many byte positions
		if i%100 == 0 {
			if err := yup.CheckContextCancellation(ctx); err != nil {
				return "" // Return empty on cancellation
			}
		}

		if pos > 0 && pos <= len(bytes) {
			result = append(result, bytes[pos-1]) // 1-based indexing
		}
	}

	return string(result)
}
