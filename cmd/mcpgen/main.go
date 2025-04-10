package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"

	"github.com/chslink/mcp-tools/gen"
)

var rootDir string

func main() {
	flag.StringVar(&rootDir, "root", "tools", "root directory")
	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		return gen.GenTool(path)
	})
	if err != nil {
		log.Fatal(err)
	}
	//gen.GenTool(fileRoot)
}
