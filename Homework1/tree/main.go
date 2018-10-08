package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
)

func isLastDir(dir os.FileInfo, path string) bool {
	dirEntries, _ := ioutil.ReadDir(path)
	dirs := []os.FileInfo{}
	for _, entry := range dirEntries {
		if entry.IsDir() {
			dirs = append(dirs, entry)
		}
	}
	return dir.Name() == dirs[len(dirs)-1].Name()
}

func recDirTree(outputStream io.Writer, path string, printFiles bool, indent string) (resultErr error) {
	dirEntries, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}

	for i, entry := range dirEntries {
		if entry.IsDir() {
			dirIndent := indent + "├───%s\n"
			levelIndent := "│"
			if (!printFiles && isLastDir(entry, path)) || (printFiles && i == len(dirEntries)-1) {
				dirIndent = indent + "└───%s\n"
				levelIndent = ""
			}
			fmt.Fprintf(outputStream, dirIndent, entry.Name())

			err := recDirTree(outputStream, path+"/"+entry.Name(), printFiles, indent+levelIndent+"\t")
			if err != nil {
				return err
			}
		} else if printFiles {
			fileIndent := indent + "├───%s (%s)\n"
			if i == len(dirEntries)-1 {
				fileIndent = indent + "└───%s (%s)\n"
			}
			size := strconv.FormatInt(entry.Size(), 10) + "b"
			if size == "0b" {
				size = "empty"
			}
			fmt.Fprintf(outputStream, fileIndent, entry.Name(), size)
		}
	}

	return nil
}

func dirTree(outputStream io.Writer, path string, printFiles bool) (resultErr error) {
	return recDirTree(outputStream, path, printFiles, "")
}

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}
