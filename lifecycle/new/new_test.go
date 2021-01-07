package new_test

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/wawandco/oxpecker/plugins"
	"github.com/wawandco/oxplugins/lifecycle/new"
)

func TestInitializerRun(t *testing.T) {
	root := t.TempDir()
	os.Chdir(root)

	pl := &new.Command{}
	tinit := &Tinit{}
	pl.Receive([]plugins.Plugin{tinit})

	err := pl.Run(context.Background(), root, []string{})
	if err == nil {
		t.Error("should return an error")
	}

	err = pl.Run(context.Background(), root, []string{"app"})
	if err != nil {
		t.Errorf("should not return and error, got: %v", err)
	}

	//Should create the folder
	fi, err := os.Stat(filepath.Join(root, "app"))
	if err != nil {
		t.Errorf("should not return and error, got: %v", err)
	}

	if !fi.IsDir() {
		t.Errorf("should be a folder, got a file")
	}

	if !tinit.called {
		t.Errorf("should have called initializer")
	}

	if !tinit.afterCalled {
		t.Errorf("should have called afterinitialize")
	}
}
