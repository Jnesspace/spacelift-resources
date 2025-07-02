package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
)

func main() {
	docsDir := "docs"

	// Read all files from docs directory
	files, err := ioutil.ReadDir(docsDir)
	if err != nil {
		log.Fatalf("failed to read docs directory: %v", err)
	}

	// Concatenate all markdown files
	for _, file := range files {
		if filepath.Ext(file.Name()) == ".md" {
			content, err := ioutil.ReadFile(filepath.Join(docsDir, file.Name()))
			if err != nil {
				log.Printf("failed to read file %s: %v", file.Name(), err)
				continue
			}

			// Print file name as header
			fmt.Printf("\n# %s\n\n", file.Name())

			// Print file contents
			fmt.Println(string(content))

			// Print separator between files
			fmt.Println("\n---\n")
		}
	}
}
