package refresh

import (
	"bytes"
	"context"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestInitializer(t *testing.T) {
	t.Run("Empty directory", func(t *testing.T) {
		root := t.TempDir()

		err := os.Chdir(root)
		if err != nil {
			t.Error("could not change to temp directory")
		}

		i := Initializer{}

		err = i.Initialize(context.Background(), root, []string{"new", "something/myapp"})
		if err != nil {
			t.Fatalf("error should be nil, got %v", err)
		}

		path := filepath.Join(root, "myapp", ".buffalo.dev.yml")
		_, err = os.Stat(path)
		if os.IsNotExist(err) {
			t.Fatalf("Did not create file in %v", path)
		}

		d, err := ioutil.ReadFile(path)
		if err != nil {
			t.Fatal("could not read the file")
		}

		if !bytes.Contains(d, []byte("myapp")) {
			t.Fatal("did not containt app name")
		}

	})
	t.Run("BuffaloFileExist", func(t *testing.T) {

		root := t.TempDir()
		err := os.Chdir(root)
		if err != nil {
			t.Error("could not change to temp directory")
		}

		rootYml := root + "/.buffalo.dev.yml"
		_, err = os.Create(rootYml)
		if err != nil {
			t.Fatalf("Problem creating file, %v", err)
		}

		i := Initializer{}

		err = i.Initialize(context.Background(), root, []string{"new", "app"})

		if err != nil {
			t.Fatalf("error should be nil, got %v", err)
		}

	})
}
