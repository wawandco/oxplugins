package creator

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"path/filepath"
	"time"

	"github.com/pkg/errors"
)

// FizzCreator model struct for fizz generation files
type FizzCreator struct {
	opts opts
}

// Name is the name of the migration type
func (f FizzCreator) Name() string {
	return "fizz"
}

// Create will create 2 .fizz files for the migration
func (f *FizzCreator) Create(dir string, args []string) error {
	timestamp := time.Now().UTC().Format("20060102150405")
	f.opts = newOptions(args)

	fileName := fmt.Sprintf("%s_%s", timestamp, f.opts.TableName)

	if err := f.createUPFile(dir, fileName); err != nil {
		return err
	}

	if err := f.createDownFile(dir, fileName); err != nil {
		return err
	}

	return nil
}

func (f *FizzCreator) createDownFile(dir, name string) error {
	filename := fmt.Sprintf("%s.down.fizz", name)
	path := filepath.Join(dir, filename)

	tmpl, err := template.New(filename).Funcs(templateFuncs).Parse(fizzDownTemplate)
	if err != nil {
		return errors.Wrap(err, "parsing new template error")
	}

	var tpl bytes.Buffer
	if err := tmpl.Execute(&tpl, f.opts); err != nil {
		return errors.Wrap(err, "executing new template error")
	}

	err = ioutil.WriteFile(path, tpl.Bytes(), 0655)
	if err != nil {
		return errors.Wrap(err, "writing new template error")
	}

	return nil
}

func (f *FizzCreator) createUPFile(dir, name string) error {
	filename := fmt.Sprintf("%s.up.fizz", name)
	path := filepath.Join(dir, filename)

	tmpl, err := template.New(filename).Funcs(templateFuncs).Parse(fizzUPTemplate)
	if err != nil {
		return errors.Wrap(err, "parsing new template error")
	}

	var tpl bytes.Buffer
	if err := tmpl.Execute(&tpl, f.opts); err != nil {
		return errors.Wrap(err, "executing new template error")
	}

	err = ioutil.WriteFile(path, tpl.Bytes(), 0655)
	if err != nil {
		return errors.Wrap(err, "writing new template error")
	}

	return nil
}
