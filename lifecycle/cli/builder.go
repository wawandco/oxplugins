package cli

import (
	"context"
	"os"
	"os/exec"
)

// ox building builds a cli binary that will
// contain the tooling defined in ox/main.go
type Builder struct{}

func (b *Builder) Name() string {
	return "cli"
}

func (b *Builder) ParentName() string {
	return ""
}

func (b *Builder) HelpText() string {
	return "Builds a binary from `cmd/ox/main.go` on `bin/cli`"
}

// Run builds cli binary
func (b *Builder) Run(ctx context.Context, root string, args []string) error {
	// Skip if there is no ox/main.go
	if !b.shouldBuild() {
		return nil
	}

	buildArgs := []string{
		"build",

		//Static binary
		"--ldflags",
		"-linkmode external",
		"--ldflags",
		`-extldflags "-static"`,

		//Output
		`-o`,
		`bin/cli`,
		"./cmd/ox/main.go",
	}

	cmd := exec.CommandContext(ctx, "go", buildArgs...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin

	return cmd.Run()
}

//
func (b *Builder) shouldBuild() bool {
	fi, err := os.Stat("cmd/ox/main.go")

	return err == nil && !fi.IsDir()
}
