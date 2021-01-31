package refresh

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"text/template"
)

var (
	// the filename we will use for the generated yml.
	filename = `.buffalo.dev.yml`

	ErrNameRequired = errors.New("name argument is required")
)

type Initializer struct{}

func (i Initializer) Name() string {
	return "refresh/initializer"
}

func (i *Initializer) Initialize(ctx context.Context, root string, args []string) error {
	if len(args) < 2 {
		return ErrNameRequired
	}

	name := filepath.Base(args[1])
	folder := filepath.Join(root, name)
	rootYML := filepath.Join(folder, filename)

	err := os.MkdirAll(folder, 0777)
	if err != nil {
		return err
	}

	_, err = os.Stat(rootYML)
	if err == nil {
		fmt.Printf("%v already exist\n", filename)
		return nil
	}

	if !os.IsNotExist(err) {
		return err
	}

	t, err := template.New("refresh").Parse(templateFile)
	if err != nil {
		return err
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, name); err != nil {
		return err
	}

	err = ioutil.WriteFile(rootYML, tpl.Bytes(), 0777)
	if err != nil {
		return err
	}

	return nil
}
