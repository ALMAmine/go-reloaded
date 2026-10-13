package main

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRunFiles(t *testing.T) {
	cases := []struct {
		name string
		in   string
		want string
	}{
		{"multiline", "it (cap) was the worst of times (up) , it was\nSimply add 42 (hex) and 10 (bin)\nbearing a untold story ,what do you think ?\n", "It was the worst of TIMES, it was\nSimply add 66 and 2\nbearing an untold story, what do you think?\n"},
		{"blank lines", "hello(up)\n\nworld(cap)\n", "HELLO\n\nWorld\n"},
		{"no final newline", "hello(up)\nworld", "HELLO\nworld"},
		{"windows line endings", "hello(up)\r\n\r\nworld(cap)\r\n", "HELLO\r\n\r\nWorld\r\n"},
		{"marker does not cross lines", "hello\n(up) world", "hello\nworld"},
		{"empty file", "", ""},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			dir := t.TempDir()
			input := filepath.Join(dir, "input.txt")
			output := filepath.Join(dir, "output.txt")
			if err := os.WriteFile(input, []byte(c.in), 0o600); err != nil {
				t.Fatal(err)
			}
			var stderr bytes.Buffer
			if code := run([]string{input, output}, &stderr); code != 0 || stderr.Len() != 0 {
				t.Fatalf("code = %d, stderr = %q", code, stderr.String())
			}
			got, err := os.ReadFile(output)
			if err != nil {
				t.Fatal(err)
			}
			if string(got) != c.want {
				t.Errorf("got %q, want %q", got, c.want)
			}
		})
	}
}

func TestRunArgumentErrors(t *testing.T) {
	for _, args := range [][]string{nil, {"input.txt"}, {"a", "b", "c"}} {
		var stderr bytes.Buffer
		if code := run(args, &stderr); code != 1 || !strings.Contains(stderr.String(), "Usage") {
			t.Errorf("args = %v, code = %d, stderr = %q", args, code, stderr.String())
		}
	}
}

func TestRunFileErrors(t *testing.T) {
	dir := t.TempDir()
	input := filepath.Join(dir, "input.txt")
	output := filepath.Join(dir, "output.txt")
	if err := os.WriteFile(input, []byte("hello(up)"), 0o600); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(output, []byte("existing output"), 0o600); err != nil {
		t.Fatal(err)
	}
	cases := []struct {
		name string
		args []string
		want string
	}{
		{"missing input", []string{filepath.Join(dir, "missing.txt"), output}, "lecture"},
		{"input is a directory", []string{dir, output}, "lecture"},
		{"missing output directory", []string{input, filepath.Join(dir, "missing", "output.txt")}, "ecriture"},
		{"output is a directory", []string{input, dir}, "ecriture"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var stderr bytes.Buffer
			if code := run(c.args, &stderr); code != 1 || !strings.Contains(stderr.String(), c.want) {
				t.Errorf("code = %d, stderr = %q", code, stderr.String())
			}
		})
	}
	got, err := os.ReadFile(output)
	if err != nil {
		t.Fatal(err)
	}
	if string(got) != "existing output" {
		t.Errorf("unreadable input changed the output to %q", got)
	}
}

func TestRunSameFile(t *testing.T) {
	path := filepath.Join(t.TempDir(), "text.txt")
	if err := os.WriteFile(path, []byte("hello(up)"), 0o600); err != nil {
		t.Fatal(err)
	}
	var stderr bytes.Buffer
	if code := run([]string{path, path}, &stderr); code != 0 {
		t.Fatalf("code = %d, stderr = %q", code, stderr.String())
	}
	got, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if string(got) != "HELLO" {
		t.Errorf("got %q, want HELLO", got)
	}
}
