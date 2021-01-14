package liquibase

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"regexp"
)

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
