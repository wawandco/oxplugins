package liquibase

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/gobuffalo/flect"
)

type Generator struct{}

var ErrName = errors.New("not valid path or name")

func (g Generator) Name() string {
	return "liquibase"
}

func (g Generator) Generate(ctx context.Context, root string, args []string) error {
	var rootFile string
	t := time.Now()
	fecha := fmt.Sprintf("%d%02d%02d%02d%02d%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())

	name := filepath.Base(args[3])

	if name == "." || name == "/" {
		return ErrName
	}

	dir := filepath.Dir(args[3])

	if name == "." && dir == "." {
		return ErrName
	}

	underscoreName := flect.Underscore(name)
	fullName := fecha + "-" + underscoreName + ".xml"

	if dir == "." {
		rootFile = filepath.Join(root, "migrations", fullName)
	} else {
		rootFile = filepath.Join(root, "migrations", dir, fullName)
	}

	_, err := os.Stat(rootFile)
	if err == nil {

		fmt.Println("file/directory already exist ")
		return nil
	}
	if os.IsNotExist(err) {
		err := os.MkdirAll(filepath.Dir(rootFile), 0755)
		if err != nil {
			return (err)
		}

		tmpl, err := template.New("[timestamp]-[name-underscore].xml").Parse(mainTemplate)
		if err != nil {
			return err
		}

		data := struct {
			Name string
			Time string
		}{
			Name: name,
			Time: fecha,
		}
		var tpl bytes.Buffer
		if err := tmpl.Execute(&tpl, data); err != nil {
			return err
		}

		err = ioutil.WriteFile(rootFile, tpl.Bytes(), 0655)

		if err != nil {
			return err
		}
	}
	return nil
}
