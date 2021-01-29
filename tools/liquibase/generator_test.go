package liquibase

import (
	"context"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"testing"
)

func TestLiquibaseGenerator(t *testing.T) {
	t.Run("FileDoesNotExist", func(t *testing.T) {
		root := t.TempDir()
		err := os.Chdir(root)
		if err != nil {
			t.Error("could not change to temp directory")
		}

		args := []string{"generate", "migration", "addDevices"}

		g := Generator{
			testPrefix: "testfile001",
		}

		err = g.Generate(context.Background(), root, args)
		if err != nil {
			t.Fatalf("Error should be nil, got %v", err)
		}
		path := filepath.Join(root, "migrations", "testfile001-add_devices.xml")
		_, err = os.Stat(path)
		if err != nil {
			t.Fatalf("Error should be nil, got %v", err)
		}
		content, err := ioutil.ReadFile(path)
		if err != nil {
			log.Fatal(err)
		}
		text := string(content)
		matched, err := regexp.MatchString(`<changeSet id="testfile001-add_devices" author="ox">`, text)

		if !matched {
			t.Fatalf("File's content is not correct, %v", err)
		}

	})
	t.Run("FileDoesNotExist/name", func(t *testing.T) {
		root := t.TempDir()
		err := os.Chdir(root)
		if err != nil {
			t.Error("could not change to temp directory")
		}

		args := []string{"generate", "migration", "location/addDevices"}

		g := Generator{
			testPrefix: "testfile001",
		}

		err = g.Generate(context.Background(), root, args)

		if err != nil {
			t.Fatalf("Error should be nil, got %v", err)
		}
		path := filepath.Join(root, "migrations", "location", "testfile001-add_devices.xml")
		_, err = os.Stat(path)
		if err != nil {
			t.Fatalf("Error should be nil, got %v", err)
		}

		content, err := ioutil.ReadFile(path)
		if err != nil {
			log.Fatal(err)
		}
		text := string(content)
		matched, err := regexp.MatchString(`<changeSet id="testfile001-add_devices" author="ox">`, text)

		if !matched {
			t.Fatalf("File's content is not correct, %v", err)
		}

	})

	t.Run("FileDoesNotExist/.", func(t *testing.T) {
		root := t.TempDir()
		err := os.Chdir(root)
		if err != nil {
			t.Error("could not change to temp directory")
		}

		args := []string{"generate", "migration", "."}

		g := Generator{}

		err = g.Generate(context.Background(), root, args)

		if err != ErrName {
			t.Fatalf("Error should be type ErrName, got %v", err)
		}

	})

	t.Run("genpathTest", func(t *testing.T) {

		root := t.TempDir()
		err := os.Chdir(root)
		if err != nil {
			t.Error("could not change to temp directory")
		}

		args := []string{"generate", "migration", "addDevices"}

		g := Generator{
			testPrefix: "testfile001",
		}
		ret, err := g.genPath(args, root)
		if err != nil {
			t.Fatalf("error should bw nil got, %v", err)
		}
		path := filepath.Join(root, "migrations", "testfile001-add_devices.xml")

		expected := []string{path, "add_devices", "testfile001"}

		match := g.equal(ret, expected)

		if !match {
			t.Fatalf("did not generate the correct path")
		}

	})

}

func TestComposeName(t *testing.T) {
	g := Generator{
		testPrefix: "composename",
	}

	filename, err := g.composeName("addDevices")
	if err != nil {
		t.Errorf("err should be nil, got %v", err)
	}

	expected := "composename-add_devices.xml"
	if filename != expected {
		t.Errorf("filename should be %v, got %v", expected, filename)
	}
}

func TestComposeNameInvalid(t *testing.T) {
	g := Generator{
		testPrefix: "composename",
	}

	_, err := g.composeName(".")
	if err != ErrInvalidName {
		t.Errorf("err should be ErrInvalidName, got %v", err)
	}

	_, err = g.composeName("/")
	if err != ErrInvalidName {
		t.Errorf("err should be ErrInvalidName, got %v", err)
	}

}
