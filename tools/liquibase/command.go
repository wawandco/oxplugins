package liquibase

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"regexp"

	"github.com/wawandco/oxpecker/plugins"
)

var _ plugins.Command = (*Command)(nil)

type Command struct {
	connectionURL string
}

func (lb Command) Name() string {
	return "migrate"
}

func (lb Command) ParentName() string {
	return ""
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
	r := regexp.MustCompile(`postgres:\/\/(?P<username>.*):(?P<password>.*)@(?P<host>.*):(?P<port>.*)\/(?P<database>.*)\?(?P<extras>.*)`)
	match := r.FindStringSubmatch(lb.connectionURL)
	if match == nil {
		return []string{}, fmt.Errorf("could not convert %v url into liquibase", environment)
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

func NewPlugin(connURL string) *Command {
	return &Command{
		connectionURL: connURL,
	}
}
