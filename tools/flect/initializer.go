package flect

import (
	"context"
	"fmt"
	"os"
)

type Initializer struct{}

func (i Initializer) Name() string {
	return "flect/initializer"
}

func (i *Initializer) Initialize(ctx context.Context, root string, args []string) error {

	root = root + "/inflections.yml"

	content := `
	{
	  "singular": "plural"
	}
	`

	if _, err := os.Stat(root); err == nil {

		fmt.Println("inflections.yml file already exist")
		return nil

	} else if os.IsNotExist(err) {

		// create file if it does not exist
		file, err := os.Create(root)

		if err != nil {
			return (err)
		}

		_, err = os.OpenFile(root, os.O_RDWR, 0644)
		if err != nil {
			return (err)
		}

		_, err = file.WriteString(content)
		if err != nil {
			return (err)
		}

		file.Close()

		return nil

	} else {
		return err

	}
}
