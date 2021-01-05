package grift

import (
	"context"
	"fmt"

	"github.com/markbates/grift/grift"
)

// Grift command is a root command to run tasks
// usage is ox task [name]. If no name is passed this will
// list the tasjs
type Command struct{}

func (c Command) Name() string {
	return "task"
}

func (c Command) ParentName() string {
	return ""
}

func (c Command) HelpText() string {
	return "Runs grifts tasks previously imported in the CLI"
}

func (c *Command) Run(ctx context.Context, root string, args []string) error {
	if len(args) < 2 {
		c.list()
		return nil
	}

	task := args[1]
	gc := grift.NewContext(task)
	if len(args) > 2 {
		gc.Args = args[2:]
	}

	return grift.Run(task, gc)
}

func (c Command) list() {
	list := grift.List()
	fmt.Printf("Available grift tasks:\n\n")
	for _, v := range list {
		fmt.Println(v)
	}

	fmt.Printf("\nrun: ox task [task-name]\n")
}
