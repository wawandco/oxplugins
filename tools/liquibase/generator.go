package liquibase

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"time"

	"github.com/gobuffalo/flect"
	"github.com/spf13/pflag"
)

var (
	ErrNameArgMissing = errors.New("name arg missing")
	ErrInvalidName    = errors.New("invalid migration name")
	ErrInvalidPath    = errors.New("invalid path")
)

// Generator for liquibase SQL migrations, it generates xml liquibase
// for SQL in the root + basedir folder. It uses the argument passed
// to determine both the name of the migration and the destination.
// Some examples are:
// - "ox generate migration name" generates [timestamp]-name.xml
// - "ox generate migration folder/name" generates folder/[timestamp]-name.xml
// - "ox generate migration name --base migrations" generates migrations/[timestamp]-name.xml
type Generator struct {
	// mockTimestamp is used for testing purposes, it would replace the
	// timestamp at the beggining of the migration name.
	mockTimestamp string

	// Basefolder for the migrations, if a path is passed, then we will append that
	// path to the baseFolder when generating the migration.
	baseFolder string

	flags *pflag.FlagSet
}

// Name is the name used to identify the generator and also
// the plugin
func (g Generator) Name() string {
	return "migration"
}

func (g Generator) Generate(ctx context.Context, root string, args []string) error {
	if len(args) < 3 {
		return ErrNameArgMissing
	}

	timestamp := time.Now().UTC().Format("20060102150405")
	if g.mockTimestamp != "" {
		timestamp = g.mockTimestamp
	}

	filename, err := g.composeFilename(args[2], timestamp)
	if err != nil {
		return err
	}

	path := g.baseFolder
	if dir := filepath.Dir(args[2]); dir != "." {
		path = filepath.Join(g.baseFolder, dir)
	}

	path = filepath.Join(path, filename)
	_, err = os.Stat(path)
	if err == nil {
		fmt.Printf("[info] %v already exists\n", path)
		return nil
	}

	if !os.IsNotExist(err) {
		return err
	}

	// Creating the folder
	err = os.MkdirAll(filepath.Dir(path), 0755)
	if err != nil {
		return (err)
	}

	tmpl, err := template.New("migration-template").Parse(migrationTemplate)
	if err != nil {
		return err
	}

	var tpl bytes.Buffer
	err = tmpl.Execute(&tpl, strings.ReplaceAll(filename, ".xml", ""))
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(path, tpl.Bytes(), 0655)
	if err != nil {
		return err
	}

	fmt.Printf("[info] migration generated in %v\n", path)
	return nil
}

func (g Generator) composeFilename(passed, timestamp string) (string, error) {
	name := filepath.Base(passed)
	//Should we check the name here ?
	if name == "." || name == "/" {
		return "", ErrInvalidName
	}

	underscoreName := flect.Underscore(name)
	result := timestamp + "-" + underscoreName + ".xml"

	return result, nil
}

func (g *Generator) ParseFlags(args []string) {
	g.flags = pflag.NewFlagSet(g.Name(), pflag.ContinueOnError)
	g.flags.StringVarP(&g.baseFolder, "base", "b", "", "base folder for the migrations")
	g.flags.Parse(args) //nolint:errcheck,we don't care hence the flag
}

func (g *Generator) Flags() *pflag.FlagSet {
	return g.flags
}
