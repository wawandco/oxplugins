package pop

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/gobuffalo/pop/v5"
)

type Tester struct{}

func (p *Tester) Name() string {
	return "pop/tester"
}

func (p *Tester) RunBeforeTest(ctx context.Context, root string, args []string) error {
	db, err := pop.Connect("test")
	if err != nil {
		return err
	}

	fmt.Println(">>> Resetting Database")
	err = db.Dialect.DropDB()
	if err != nil {
		fmt.Printf("could not drop `%v` database, continuing.", db.Dialect.Name())
	}

	err = db.Dialect.CreateDB()
	if err != nil {
		return err
	}

	// Running migrations
	fmt.Println(">>> Running migrations")
	ms := filepath.Join(root, "migrations")
	fm, err := pop.NewFileMigrator(ms, db)
	if err != nil {
		return err
	}

	return fm.Up()
}
