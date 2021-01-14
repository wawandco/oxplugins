package liquibase

import (
	"bytes"
	"context"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

type Generator struct{}

func (g Generator) Name() string {
	return "liquibase"
}

func (g Generator) Generate(ctx context.Context, root string, args []string) error {

	t := time.Now()
	fecha := fmt.Sprintf("%d%02d%02d%02d%02d%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())

	name := args[3]
	var underscoreName string
	re := regexp.MustCompile(`[a-zA-Z][^A-Z]*`)
	submatchall := re.FindAllString(name, -1)

	for index, element := range submatchall {
		if index == 0 {
			underscoreName = strings.ToLower(element)
			continue
		}
		underscoreName = underscoreName + "_" + strings.ToLower(element)
	}
	fullName := fecha + "-" + underscoreName + ".xml"

	fmt.Println(fullName)

	rootFile := filepath.Join(root, "migrations", fullName)
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
