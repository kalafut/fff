package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	flag.Parse()
	exclude := map[string]bool{}
	scan(".", flag.Arg(0), exclude)
}

func scan(root, search string, exclude map[string]bool) {
	search = strings.ToLower(search)
	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() && info.Name() == ".git" {
			return filepath.SkipDir
		}
		if !exclude[filepath.Ext(path)] && strings.Contains(strings.ToLower(path), search) {
			fmt.Println(path)
		}
		return nil
	})
}
