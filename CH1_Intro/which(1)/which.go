package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Пожалуйста укажите аргумент!")
		return
	}

	file := arguments[1]
	path := os.Getenv("PATH")
	pathSplit := filepath.SplitList(path)
	for _, directory := range pathSplit {
		fullPath := filepath.Join(directory, file)

		// Существует ли файл?
		fileInfo, err := os.Stat(fullPath)
		if err != nil {
			continue
		}

		mode := fileInfo.Mode()
		// Это обычный файл?
		if !mode.IsRegular() {
			continue
		}

		// Является ли он исполняемым?
		if mode&0111 != 0 {
			fmt.Println(fullPath)
			return
		}
	}
}
