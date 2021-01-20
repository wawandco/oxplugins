package model

import (
	"bytes"
	"context"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/gobuffalo/flect"
	"github.com/pkg/errors"
)

// Generator allows to identify model as a plugin
type Generator struct {
	name string
	dir  string
}

// Name returns the name of the generator plugin
func (g Generator) Name() string {
	return "model"
}

// Generate generates an empty [name].plush.html file
func (g Generator) Generate(ctx context.Context, root string, args []string) error {
	if len(args) < 3 {
		return errors.Errorf("no name specified, please use `ox generate model [name]`")
	}

	dirPath := filepath.Join(root, "app", "models")
	if !g.exists(dirPath) {
		return errors.Errorf("folder '%s' do not exists on your buffalo app, please ensure the folder exists in order to proceed", dirPath)
	}

	g.name = args[2]
	g.dir = dirPath

	if g.exists(filepath.Join(g.dir, g.name+".go")) {
		return errors.Errorf("model already exists")
	}

	if err := g.generateModelFiles(args[3:]); err != nil {
		return err
	}

	return nil
}

func (g Generator) generateModelFiles(args []string) error {
	if err := g.createModelFile(args); err != nil {
		return errors.Wrap(err, "creating model file")
	}

	if err := g.createModelTestFile(); err != nil {
		return errors.Wrap(err, "creating model test file")
	}

	return nil
}

func (g Generator) createModelFile(args []string) error {
	filename := flect.Singularize(g.name) + ".go"
	path := filepath.Join(g.dir, filename)
	attrs := buildAttrs(args)
	data := opts{
		Name:    g.name,
		Attrs:   attrs,
		Imports: buildImports(attrs),
	}

	tmpl, err := template.New(filename).Funcs(templateFuncs).Parse(modelTemplate)
	if err != nil {
		return errors.Wrap(err, "parsing new template error")
	}

	var tpl bytes.Buffer
	if err := tmpl.Execute(&tpl, data); err != nil {
		return errors.Wrap(err, "executing new template error")
	}

	err = ioutil.WriteFile(path, tpl.Bytes(), 0655)
	if err != nil {
		return errors.Wrap(err, "writing new template error")
	}

	return nil
}

func (g Generator) createModelTestFile() error {
	filename := flect.Singularize(g.name) + "_test.go"
	path := filepath.Join(g.dir, filename)
	data := opts{
		Name: g.name,
	}

	tmpl, err := template.New(filename).Funcs(templateFuncs).Parse(modelTestTemplate)
	if err != nil {
		return errors.Wrap(err, "parsing new template error")
	}

	var tpl bytes.Buffer
	if err := tmpl.Execute(&tpl, data); err != nil {
		return errors.Wrap(err, "executing new template error")
	}

	err = ioutil.WriteFile(path, tpl.Bytes(), 0655)
	if err != nil {
		return errors.Wrap(err, "writing new template error")
	}

	return nil
}

func (g Generator) exists(path string) bool {
	_, err := os.Stat(path)

	return !os.IsNotExist(err)
}
