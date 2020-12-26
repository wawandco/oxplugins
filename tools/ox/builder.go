package ox

// ox building builds a cli binary that will
// contain the tooling defined in ox/main.go
type Builder struct {}

// Builds cli binary
func (g *Builder) Build(ctx context.Context, root string, args []string) error {
	// Skip if there is no ox/main.go
	buildArgs := []string{
		"build", 
		"--ldflags",
		`'-linkmode external -extldflags "-static"'` ,
		`-o`, 
		`/bin/cli`,
		"ox/main.go",
	}

	cmd := exec.CommandContext(ctx, "go", buildArgs...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin

	return cmd.Run()
}