package liquibase

import (
	"context"
	"os"
	"path/filepath"
	"testing"
)

func TestLiquibaseGenerator(t *testing.T) {
	t.Run("FileDoesNotExist", func(t *testing.T) {
		root := t.TempDir()
		err := os.Chdir(root)
		if err != nil {
			t.Error("could not change to temp directory")
		}

		args := []string{"ox", "generate", "liquibase", "addDevices"}

		g := Generator{
			testPrefix: "testfile001",
		}

		err = g.Generate(context.Background(), root, args)
		if err != nil {
			t.Fatalf("Error should be nil, got %v", err)
		}

		_, err = os.Stat(filepath.Join(root, "migrations", "testfile001-add_devices.xml"))
		if err != nil {
			t.Fatalf("Error should be nil, got %v", err)
		}

	})
	t.Run("FileDoesNotExist/name", func(t *testing.T) {
		root := t.TempDir()
		err := os.Chdir(root)
		if err != nil {
			t.Error("could not change to temp directory")
		}

		args := []string{"ox", "generate", "liquibase", "location/addDevices"}

		g := Generator{}

		err = g.Generate(context.Background(), root, args)

		if err != nil {
			t.Fatalf("Error should be nil, got %v", err)
		}

	})

	t.Run("FileDoesNotExist/.", func(t *testing.T) {
		root := t.TempDir()
		err := os.Chdir(root)
		if err != nil {
			t.Error("could not change to temp directory")
		}

		args := []string{"ox", "generate", "liquibase", "."}

		g := Generator{}

		err = g.Generate(context.Background(), root, args)

		if err != ErrName {
			t.Fatalf("Error should be type ErrName, got %v", err)
		}

	})

}
