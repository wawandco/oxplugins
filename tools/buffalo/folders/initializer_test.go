package folders

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestInitializerRun(t *testing.T) {

	t.Run("valid args", func(t *testing.T) {
		root := t.TempDir()
		err := os.Chdir(root)
		if err != nil {
			t.Fatal("could not move to temp dir")
		}

		i := Initializer{}
		err = i.Initialize(context.Background(), root, []string{"new", "app"})
		if err != nil {
			t.Errorf("err should be nil but got %v", err)
		}

		for _, v := range folders {
			v = strings.ReplaceAll(v, "[name]", "app")
			v = filepath.Join(root, v)

			if _, err := os.Stat(v); os.IsNotExist(err) {
				t.Errorf("should have created %v", v)
			}
		}
	})

	t.Run("invalid args", func(t *testing.T) {
		i := Initializer{}
		err := i.Initialize(context.Background(), "", []string{"new"})
		if err != ErrNameNeeded {
			t.Errorf("err should ne ErrNameNeeded but got %v", err)
		}
	})

	t.Run("valid args", func(t *testing.T) {
		root := t.TempDir()
		err := os.Chdir(root)
		if err != nil {
			t.Fatal("could not move to temp dir")
		}

		i := Initializer{}
		err = i.Initialize(context.Background(), root, []string{"new", "github.com/paganotoni/app"})
		if err != nil {
			t.Errorf("err should be nil but got %v", err)
		}

		v := filepath.Join(root, "app")
		if _, err := os.Stat(v); os.IsNotExist(err) {
			t.Errorf("should have created %v", v)
		}
	})

	t.Run("force disabled", func(t *testing.T) {
		root := t.TempDir()
		err := os.Chdir(root)
		if err != nil {
			t.Fatal("could not move to temp dir")
		}

		err = os.MkdirAll(filepath.Join(root, "app"), 0777)
		if err != nil {
			t.Fatal("could not create dir")
		}

		i := Initializer{}
		err = i.Initialize(context.Background(), root, []string{"new", "app"})
		if err != ErrFolderExists {
			t.Errorf("err should be ErrFolderExists but got %v", err)
		}
	})

	t.Run("force enabled", func(t *testing.T) {
		root := t.TempDir()
		err := os.Chdir(root)
		if err != nil {
			t.Fatal("could not move to temp dir")
		}

		err = os.MkdirAll(filepath.Join(root, "app"), 0777)
		if err != nil {
			t.Fatal("could not create dir")
		}

		i := Initializer{force: true}
		err = i.Initialize(context.Background(), root, []string{"new", "app"})
		if err != nil {
			t.Errorf("err should be nil but got %v", err)
		}
	})

}
