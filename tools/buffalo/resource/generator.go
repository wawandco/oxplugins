package resource

import (
	"context"
	"path/filepath"

	"github.com/pkg/errors"

	"github.com/wawandco/oxplugins/tools/buffalo/model"
	"github.com/wawandco/oxplugins/tools/pop/migration/creator"
)

// Generator allows to identify resource as a plugin
type Generator struct{}

// Name returns the name of the generator plugin
func (g Generator) Name() string {
	return "resource"
}

// Generate generates the actions, model, templates and migration files from the given attrs
// app/actions/[name].go
// app/actions/[name]_test.go
// app/models/[name].go
// app/models/[name]_test.go
// app/templates/[name]/index.plush.html
// app/templates/[name]/new.plush.html
// app/templates/[name]/edit.plush.html
// app/templates/[name]/show.plush.html
// migration/[name].up.fizz
// migration/[name].down.fizz
func (g Generator) Generate(ctx context.Context, root string, args []string) error {
	if len(args) < 3 {
		return errors.Errorf("no name specified, please use `ox generate resource [name]`")
	}

	name := args[2]
	attrs := args[3:]

	// Generating Model
	modelsPath := filepath.Join(root, "app", "models")
	model := model.New(modelsPath, name, attrs)
	if err := model.Create(); err != nil {
		return errors.Wrap(err, "error creating model")
	}

	// Generating Migration
	migrationPath := filepath.Join(root, "migrations")
	creator, err := creator.CreateMigrationFor("fizz")
	if err != nil {
		return errors.Wrap(err, "error looking for migration creator")
	}

	if err = creator.Create(migrationPath, args[2:]); err != nil {
		return errors.Wrap(err, "failed creating migrations")
	}

	return nil
}
