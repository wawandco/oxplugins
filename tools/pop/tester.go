package pop

import (
	"context"
	"fmt"

	"github.com/gobuffalo/pop/v5"
	"github.com/wawandco/oxpecker/plugins"
	"github.com/wawandco/oxplugins/tools/pop/migrate"
)

type Tester struct {
	migrator migrate.Migrator
}

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
		fmt.Printf("could not drop `%v` database, continuing\n", db.Dialect)
	}

	err = db.Dialect.CreateDB()
	if err != nil {
		return err
	}

	if p.migrator == nil {
		return nil
	}

	// Running migrations
	fmt.Println(">>> Running migrations")
	p.migrator.SetConn("test")
	return p.migrator.Run(ctx, root, args)
}

func (p *Tester) Receive(pls []plugins.Plugin) {
	for _, plugin := range pls {
		c, ok := plugin.(migrate.Migrator)
		if !ok || c.Direction() != "up" {
			continue
		}

		p.migrator = c
		break
	}
}
