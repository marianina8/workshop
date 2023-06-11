package cmd

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/marianina8/example24/models"
	"gopkg.in/yaml.v2"
)

func TestConvert(t *testing.T) {
	t.Run("Test successful convert", func(t *testing.T) {
		b := bytes.NewBufferString("")
		rootCmd.SetOut(b)
		rootCmd.SetArgs([]string{"convert", "-f", "../example.json"})
		err := rootCmd.Execute()
		if err != nil {
			t.Fatal(err)
		}
		actualBytes, err := io.ReadAll(b)
		if err != nil {
			t.Fatal(err)
		}
		filePath := filepath.Join("..", "example.yaml")
		expectedBytes, err := os.ReadFile(filePath)
		if err != nil {
			t.Fatal(err)
		}
		var workshop1, workshop2 models.Workshop
		yaml.Unmarshal(actualBytes, &workshop1)
		yaml.Unmarshal(expectedBytes, &workshop2)
		if !(workshop1.Title == workshop2.Title &&
			workshop1.Instructors[0].Name == workshop2.Instructors[0].Name &&
			workshop1.Instructors[0].Email == workshop2.Instructors[0].Email &&
			workshop1.Instructors[1].Name == workshop2.Instructors[1].Name &&
			workshop1.Instructors[1].Email == workshop2.Instructors[1].Email) {
			t.Fatalf("expected %v, got %v", workshop2, workshop1)
		}
	})

	t.Run("Fail to convert", func(t *testing.T) {
		b := bytes.NewBufferString("")
		rootCmd.SetErr(b)
		rootCmd.SetArgs([]string{"convert", "--random"})
		err := rootCmd.Execute()
		if err == nil {
			t.Fatalf("expected err (unknown flag) but got nil")
		}
		actualBytes, err := io.ReadAll(b)
		if err != nil {
			t.Fatal(err)
		}
		if string(actualBytes) != "Error: unknown flag: --random\n" {
			t.Fatal(err)
		}
	})
}
