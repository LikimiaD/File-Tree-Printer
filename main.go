package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
)

type FileInfo struct {
	name string
	size int64
	tab  int
}

func printTree(out io.Writer, ans map[string][]FileInfo, path string, prefix string, size bool) error {
	files := ans[path]
	sort.Slice(files, func(i, j int) bool { return files[i].name < files[j].name })

	for i, file := range files {
		isLast := i == len(files)-1

		linePrefix := "├───"
		if isLast {
			linePrefix = "└───"
		}

		if size {
			sizeStr := ""
			if file.size > 0 {
				sizeStr = fmt.Sprintf(" (%db)", file.size)
			} else if file.size == 0 {
				sizeStr = " (empty)"
			}
			_, err := fmt.Fprintf(out, "%s%s%s%s\n", prefix, linePrefix, file.name, sizeStr)
			if err != nil {
				return err
			}
		} else {
			_, err := fmt.Fprintf(out, "%s%s%s\n", prefix, linePrefix, file.name)
			if err != nil {
				return err
			}
		}

		if file.size == -1 {
			nextPrefix := prefix
			if isLast {
				nextPrefix += "\t"
			} else {
				nextPrefix += "│\t"
			}
			err := printTree(out, ans, filepath.Join(path, file.name), nextPrefix, size)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func readFiles(path string, filesBool bool, tab int) (map[string][]FileInfo, error) {
	ans := make(map[string][]FileInfo)
	files, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	sort.Slice(files, func(i, j int) bool { return files[i].Name() < files[j].Name() })

	var tmpSlice []FileInfo
	for _, file := range files {
		if file.IsDir() {
			name := file.Name()
			tmpMap, err := readFiles(filepath.Join(path, name), filesBool, tab+1)
			if err != nil {
				return nil, err
			}
			tmpSlice = append(tmpSlice, FileInfo{name, -1, tab})
			for key, value := range tmpMap {
				ans[key] = value
			}
		} else if filesBool {
			fileInfo, err := file.Info()
			if err != nil {
				return nil, err
			}
			tmpSlice = append(tmpSlice, FileInfo{fileInfo.Name(), fileInfo.Size(), tab})
		}
	}

	if len(tmpSlice) > 0 || tab == 0 {
		ans[path] = tmpSlice
	}
	return ans, nil
}

func dirTree(out io.Writer, path string, printFiles bool) error {
	ansMap, err := readFiles(path, printFiles, 0)
	if err != nil {
		return err
	}

	err = printTree(out, ansMap, path, "", printFiles)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage: go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"

	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}
