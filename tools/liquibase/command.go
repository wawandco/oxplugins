package liquibase

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"regexp"

	pop4 "github.com/gobuffalo/pop"
	pop5 "github.com/gobuffalo/pop/v5"
	"github.com/wawandco/oxpecker/plugins"
)

var _ plugins.Command = (*Command)(nil)
var _ plugins.HelpTexter = (*Command)(nil)

type Command struct {
	connections map[string]URLProvider
}

func (lb Command) Name() string {
	return "migrate"
}

func (lb Command) ParentName() string {
	return "db"
}

func (lb Command) HelpText() string {
	return "runs Liquibase command to update database with current GO_ENV"
}

func (lb *Command) Run(ctx context.Context, root string, args []string) error {
	currentEnv := os.Getenv("GO_ENV")
	if currentEnv == "" {
		currentEnv = "development"
	}

	return lb.update(currentEnv)
}

func (lb *Command) RunBeforeTest(ctx context.Context, root string, args []string) error {
	return lb.update("test")
}

func (lb Command) update(env string) error {
	runArgs, err := lb.buildRunArgsFor(env)
	if err != nil {
		return err
	}

	runArgs = append(runArgs, []string{
		"--changeLogFile=./migrations/changelog.xml",
		"update",
	}...)

	c := exec.Command("liquibase", runArgs...)
	c.Stdin = os.Stdin
	c.Stderr = os.Stderr
	c.Stdout = os.Stdout

	return c.Run()
}

func (lb Command) buildRunArgsFor(environment string) ([]string, error) {
	conn := lb.connections[environment]
	if conn == nil {
		return []string{}, errors.New("connection not found")
	}

	r := regexp.MustCompile(`postgres:\/\/(?P<username>.*):(?P<password>.*)@(?P<host>.*):(?P<port>.*)\/(?P<database>.*)\?(?P<extras>.*)`)
	match := r.FindStringSubmatch(conn.URL())
	if match == nil {
		return []string{}, fmt.Errorf("could not convert `%v` url into Liquibase", environment)
	}

	URL := fmt.Sprintf("jdbc:postgresql://%v:%v/%v?%v", match[3], match[4], match[5], match[6])
	runArgs := []string{
		"--driver=org.postgresql.Driver",
		"--url=" + URL,
		"--logLevel=info",
		"--username=" + match[1],
		"--password=" + match[2],
	}
	return runArgs, nil
}

func NewPlugin(conns interface{}) *Command {
	switch v := conns.(type) {
	case map[string]*pop4.Connection:
		result := map[string]URLProvider{}
		for k, conn := range v {
			result[k] = conn
		}

		return &Command{
			connections: result,
		}
	case map[string]*pop5.Connection:
		result := map[string]URLProvider{}
		for k, conn := range v {
			result[k] = conn
		}

		return &Command{
			connections: result,
		}

	default:
		fmt.Println("DB plugin ONLY receives pop v4 and v5 connections")
	}

	return nil
}
