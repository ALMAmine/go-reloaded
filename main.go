package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	os.Exit(run(os.Args[1:], os.Stderr))
}

func run(args []string, stderr io.Writer) int {
	if len(args) != 2 {
		fmt.Fprintln(stderr, "Usage : go run . <fichier-entree> <fichier-sortie>")
		return 1
	}

	content, err := os.ReadFile(args[0])
	if err != nil {
		fmt.Fprintf(stderr, "Erreur de lecture de %q : %v\n", args[0], err)
		return 1
	}

	lines := strings.Split(string(content), "\n")
	for i, line := range lines {
		ending := ""
		if strings.HasSuffix(line, "\r") {
			line = strings.TrimSuffix(line, "\r")
			ending = "\r"
		}
		lines[i] = Process(line) + ending
	}

	if err := os.WriteFile(args[1], []byte(strings.Join(lines, "\n")), 0o644); err != nil {
		fmt.Fprintf(stderr, "Erreur d'ecriture de %q : %v\n", args[1], err)
		return 1
	}
	return 0
}
