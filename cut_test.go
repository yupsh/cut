package cut_test

import (
	"context"
	"os"
	"strings"

	"github.com/yupsh/cut"
	"github.com/yupsh/cut/opt"
)

func ExampleCut_fields() {
	ctx := context.Background()
	input := strings.NewReader("one,two,three\nfour,five,six\n")

	cmd := cut.Cut(opt.Delimiter(","), opt.Fields([]int{1, 3}))
	cmd.Execute(ctx, input, os.Stdout, os.Stderr)
	// Output: one,three
	// four,six
}

func ExampleCut_chars() {
	ctx := context.Background()
	input := strings.NewReader("hello\nworld\n")

	cmd := cut.Cut(opt.Chars([]int{1, 3, 5}))
	cmd.Execute(ctx, input, os.Stdout, os.Stderr)
	// Output: hlo
	// wrd
}
