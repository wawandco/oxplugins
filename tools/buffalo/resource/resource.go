package resource

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
	"text/template"

	"github.com/gobuffalo/flect/name"
	"github.com/pkg/errors"

	"github.com/wawandco/oxplugins/tools/buffalo/model"
	"github.com/wawandco/oxplugins/tools/pop/migration/creator"
)

// Resource model struct
type Resource struct {
	Name  name.Ident
	Model model.Model
	Args  []string

	originalArgs []string
	originalName string
	root         string
}

// New creates a new instance of Resource
func New(root string, args []string) *Resource {
	modelsPath := filepath.Join(root, "app", "models")
	model := model.New(modelsPath, args[2], args[3:])

	return &Resource{
		Args:  args[3:],
		Model: model,
		Name:  name.New(args[2]),

		originalArgs: args[2:],
	}
}

// GenerateModel generates the model for the resource
func (r *Resource) GenerateModel() error {
	if err := r.Model.Create(); err != nil {
		return errors.Wrap(err, "error creating model")
	}

	return nil
}

// GenerateModel generates the migrations for the resource
func (r *Resource) GenerateMigrations() error {
	migrationPath := filepath.Join(r.root, "migrations")
	creator, err := creator.CreateMigrationFor("fizz")
	if err != nil {
		return errors.Wrap(err, "error looking for migration creator")
	}

	if err = creator.Create(migrationPath, r.originalArgs); err != nil {
		return errors.Wrap(err, "failed creating migrations")
	}

	return nil
}

// GenerateModel generates the templates for the resource
func (r *Resource) GenerateTemplates() error {
	templates := map[string]string{
		"index": templateIndexTmpl,
		"new":   templateNewTmpl,
		"edit":  templateEditTmpl,
		"show":  templateShowTmpl,
		"form":  templateFormTmpl,
	}

	dirPath := filepath.Join(r.root, "app", "templates", r.Name.Underscore().String())
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		err = os.Mkdir(dirPath, 0755)
		if err != nil {
			return err
		}
	}

	for name, content := range templates {
		filename := name + ".plush.html"
		path := filepath.Join(dirPath, filename)

		tmpl, err := template.New(filename).Parse(content)
		if err != nil {
			return errors.Wrap(err, "parsing new template error")
		}

		var tpl bytes.Buffer
		if err = tmpl.Execute(&tpl, r); err != nil {
			return errors.Wrap(err, "executing new template error")
		}

		if err = ioutil.WriteFile(path, tpl.Bytes(), 0655); err != nil {
			return errors.Wrap(err, "writing new template error")
		}
	}

	return nil
}
