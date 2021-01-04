package cli

import (
	"context"
	"io/ioutil"
	"os"
	"testing"
)

func TestShouldBuild(t *testing.T) {
	t.Run("File Exists", func(t *testing.T) {
		err := os.Chdir(t.TempDir())
		if err != nil {
			t.Fail()
		}

		err = os.MkdirAll("cmd/ox", 0777)
		if err != nil {
			t.Fatal("error creating dir")
		}

		err = ioutil.WriteFile("cmd/ox/main.go", []byte("package main"), 0777)
		if err != nil {
			t.Fatal("error creating file")
		}
		b := &Builder{}
		if !b.shouldBuild() {
			t.Fatal("Should build returned false, should return true")
		}
	})

	t.Run("File Does not exist", func(t *testing.T) {
		err := os.Chdir(t.TempDir())
		if err != nil {
			t.Fail()
		}

		b := &Builder{}
		if b.shouldBuild() {
			t.Fatal("ShouldBuild returned true, should return false")
		}
	})

	t.Run("Folder instead of file", func(t *testing.T) {
		err := os.Chdir(t.TempDir())
		if err != nil {
			t.Fail()
		}

		err = os.MkdirAll("cmd/ox/main.go", 0777)
		if err != nil {
			t.Fatal("error creating dir")
		}

		b := &Builder{}
		if b.shouldBuild() {
			t.Fatal("ShouldBuild returned true, should return false")
		}
	})

}

func TestBuild(t *testing.T) {
	root := t.TempDir()
	err := os.Chdir(root)
	if err != nil {
		t.Fail()
	}

	err = os.MkdirAll("cmd/ox", 0777)
	if err != nil {
		t.Fatal("error creating dir")
	}

	err = ioutil.WriteFile("go.mod", []byte("module app"), 0777)
	if err != nil {
		t.Fatal("error creating go.mod")
	}
	content := `
	package main

	func main() {
		
	}
	`
	err = ioutil.WriteFile("cmd/ox/main.go", []byte(content), 0777)
	if err != nil {
		t.Fatal("error creating file")
	}

	b := &Builder{}
	err = b.Run(context.Background(), root, []string{})
	if err != nil {
		t.Errorf("error building: %v", err)
	}
}