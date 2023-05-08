package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Используйте: ./main <директория>")
		return
	}

	root := os.Args[1]

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.Mode().IsRegular() {
			return nil
		}
		size := info.Size()
		unit := "B"
		if size > 1024 {
			size /= 1024
			unit = "KB"
		}
		if size > 1024 {
			size /= 1024
			unit = "MB"
		}
		if size > 1024 {
			size /= 1024
			unit = "GB"
		}
		if info.IsDir() {
			return nil
		}
		file, err := os.Stat(path)
		if err != nil {
			return err
		}
		fmt.Printf("%s - %d%s - %s\n", path, size, unit, file.ModTime().Format("2006-01-02 15:04:05"))
		return nil
	})

	if err != nil {
		fmt.Println(err)
	}
}
