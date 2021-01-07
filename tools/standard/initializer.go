package standard

import (
	"context"
)

type Initializer struct{}

func (i Initializer) Name() string {
	return "standard/initializer"
}

// - Initializes module based on app name
// - Creates cmd/name/main.go
func (i *Initializer) Initialize(ctx context.Context, root string, args []string) error {
	return nil
}

func (i *Initializer) ParseFlags(flags []string) {}
