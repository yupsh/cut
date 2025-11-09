package command_test

import (
	"errors"
	"testing"

	"github.com/gloo-foo/testable/assertion"
	"github.com/gloo-foo/testable/run"
	command "github.com/yupsh/cut"
)

func TestCut_Fields(t *testing.T) {
	result := run.Command(command.Cut(command.Fields{1, 3})).
		WithStdinLines("a\tb\tc").Run()
	assertion.NoError(t, result.Err)
	assertion.Count(t, result.Stdout, 1)
}

func TestCut_CustomDelimiter(t *testing.T) {
	result := run.Command(command.Cut(command.Fields{1, 2}, command.Delimiter(","))).
		WithStdinLines("a,b,c").Run()
	assertion.NoError(t, result.Err)
	assertion.Count(t, result.Stdout, 1)
}

func TestCut_Bytes(t *testing.T) {
	result := run.Command(command.Cut(command.Bytes{1, 3})).
		WithStdinLines("abc").Run()
	assertion.NoError(t, result.Err)
	assertion.Count(t, result.Stdout, 1)
}

func TestCut_Characters(t *testing.T) {
	result := run.Command(command.Cut(command.Chars{1, 2})).
		WithStdinLines("abc").Run()
	assertion.NoError(t, result.Err)
	assertion.Count(t, result.Stdout, 1)
}

func TestCut_OnlyDelimited(t *testing.T) {
	result := run.Command(command.Cut(command.Fields{1}, command.OnlyDelimited)).
		WithStdinLines("a\tb", "no-delimiter").Run()
	assertion.NoError(t, result.Err)
	assertion.Count(t, result.Stdout, 1)
}

func TestCut_EmptyInput(t *testing.T) {
	result := run.Quick(command.Cut(command.Fields{1}))
	assertion.NoError(t, result.Err)
	assertion.Empty(t, result.Stdout)
}

func TestCut_InputError(t *testing.T) {
	result := run.Command(command.Cut(command.Fields{1})).
		WithStdinError(errors.New("read failed")).Run()
	assertion.ErrorContains(t, result.Err, "read failed")
}

func TestCut_BytesOutOfRange(t *testing.T) {
	// Test bytes selection beyond line length
	result := run.Command(command.Cut(command.Bytes{10, 20})).
		WithStdinLines("short").Run()
	assertion.NoError(t, result.Err)
	// Should output empty or handle gracefully
}

func TestCut_CharsOutOfRange(t *testing.T) {
	// Test chars selection beyond line length
	result := run.Command(command.Cut(command.Chars{10, 20})).
		WithStdinLines("short").Run()
	assertion.NoError(t, result.Err)
}

