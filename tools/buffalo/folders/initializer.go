package folders

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/pflag"
)

var (
	ErrNameNeeded   = errors.New("app name must be specified")
	ErrFolderExists = errors.New("folder exists")

	// Folders that will be created when the initializer runs
	// the [name] parts will be replaced by the name of the app.
	folders = []string{
		filepath.Join("[name]"),
		filepath.Join("[name]", "public"),
		filepath.Join("[name]", "migrations"),
		filepath.Join("[name]", "config"),
		filepath.Join("[name]", "app"),
		filepath.Join("[name]", "app", "models"),
		filepath.Join("[name]", "app", "actions"),
		filepath.Join("[name]", "app", "templates"),
		filepath.Join("[name]", "app", "assets"),
		filepath.Join("[name]", "app", "assets", "js"),
		filepath.Join("[name]", "app", "assets", "css"),
		filepath.Join("[name]", "app", "render"),
		filepath.Join("[name]", "app", "tasks"),
		filepath.Join("[name]", "cmd", "[name]"),
	}
)

// Initializer is in charge of building the bones of the
// Buffalo application. it will use the name argument and take the
// base part of it to build the folders.
// Some examples:
// - `ox new bongo` 					=> creates the new app in the bongo folder
// - `ox new github.com/wawandco/bongo` => creates the new app in the bongo folder
// - `ox new wawandco/bongo` 			=> creates the new app in the bongo folder
type Initializer struct {
	// force folder creation if exists.
	force bool

	flags *pflag.FlagSet
}

// Name of the plugin
func (i Initializer) Name() string {
	return "buffalo/initializer"
}

// Initialize the app by creating the needed folders. It will infer the name of the
// folder from the args passed.
func (i *Initializer) Initialize(ctx context.Context, root string, args []string) error {
	if len(args) < 2 {
		return ErrNameNeeded
	}

	name := filepath.Base(args[1])
	base := filepath.Join(root, name)
	if _, err := os.Stat(base); err == nil && !i.force {
		return ErrFolderExists
	}

	err := os.RemoveAll(base)
	if err != nil {
		return err
	}

	for _, v := range folders {
		v = strings.ReplaceAll(v, "[name]", name)
		v = filepath.Join(root, v)

		err := os.MkdirAll(v, 0777)
		if err != nil {
			return err
		}

		fmt.Printf("[info] Created %v folder\n", v)
	}

	return nil
}

func (d *Initializer) ParseFlags(args []string) {
	d.flags = pflag.NewFlagSet(d.Name(), pflag.ContinueOnError)
	d.flags.BoolVarP(&d.force, "force", "f", false, "force the creation by removing folder if exists")
	d.flags.Parse(args) //nolint:errcheck,we don't care hence the flag
}

func (d *Initializer) Flags() *pflag.FlagSet {
	return d.flags
}
