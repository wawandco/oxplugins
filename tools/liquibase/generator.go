package liquibase

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/gobuffalo/flect"
)

var (
	ErrNameArgMissing = errors.New("name arg missing")
	ErrInvalidName    = errors.New("invalid migration name")
	ErrInvalidPath    = errors.New("invalid path")
)

type Generator struct {
	// mockTimestamp is used for testing purposes, it would replace the
	// timestamp at the beggining of the migration name.
	mockTimestamp string

	// Basefolder for the migrations, if a path is passed, then we will append that
	// path to the baseFolder when generating the migration.
	baseFolder string
}

// Name is the name used to identify the generator and also
// the plugin
func (g Generator) Name() string {
	return "migration"
}

func (g Generator) Generate(ctx context.Context, root string, args []string) error {
	if len(args) == 0 {
		return
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

	// err = os.MkdirAll(filepath.Dir(path), 0755)
	// if err != nil {
	// 	return (err)
	// }

	// d := data{
	// 	Filename:  filename,
	// 	Timestamp: timestamp,
	// }

	// tmpl, err := template.New("[timestamp]-[name-underscore].xml").Parse(mainTemplate)
	// if err != nil {
	// 	return err
	// }

	// var tpl bytes.Buffer
	// err = tmpl.Execute(&tpl, data)
	// if err != nil {
	// 	return err
	// }

	// err = ioutil.WriteFile(path, tpl.Bytes(), 0655)
	// if err != nil {
	// 	return err
	// }

	fmt.Printf("[info] migration generated in %v\n", path)
	return nil
}

// // Genpath returns the path, the name of the file and the timestamp
// func (g Generator) genPath(args []string, root string) ([]string, error) {
// 	var ret []string
// 	dir := filepath.Dir(args[2])
// 	if name == "." && dir == "." {
// 		return ret, ErrName
// 	}

// 	underscoreName := flect.Underscore(name)
// 	timestamp := time.Now().UTC().Format("20060102150405")
// 	if g.testPrefix != "" {
// 		timestamp = g.testPrefix
// 	}

// 	fullName := timestamp + "-" + underscoreName + ".xml"

// 	path := filepath.Join(root, "migrations", fullName)
// 	if dir != "." {
// 		path = filepath.Join(root, "migrations", dir, fullName)
// 	}
// 	ret = append(ret, path, underscoreName, timestamp)

// 	return ret, nil
// }

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
