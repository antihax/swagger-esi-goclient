package main

import (
	"flag"
	"log"
	"os/exec"
)

var generatedPath string

func init() {
	flag.StringVar(&generatedPath, "src", "./esi", "path of generated code")
	flag.Parse()
}

func main() {
	if generatedPath == "" {
		log.Fatalf("src must be a valid path\n")
	}
	cmd := exec.Command("go", "fmt", generatedPath)
	if err := cmd.Run(); err != nil {
		log.Fatalf("Failed to run go fmt: %v\n", err)
	}
}
