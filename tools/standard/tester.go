package standard

import (
	"context"
	"fmt"
	"os"
	"os/exec"
)

type Tester struct{}

func (t Tester) Name() string {
	return "standard/tester"
}

func (b *Tester) RunBeforeTest(ctx context.Context, root string, args []string) error {
	return os.Setenv("GO_ENV", "test")
}

func (p *Tester) Test(ctx context.Context, root string, args []string) error {
	fmt.Println(">>> Running Tests")

	cargs := []string{
		"test",
	}

	if len(args) > 0 {
		cargs = append(cargs, args...)
	} else {
		cargs = append(cargs, "./...", "-p", "1")
	}

	cmd := exec.CommandContext(ctx, "go", cargs...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
