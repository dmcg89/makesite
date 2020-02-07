package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

func main() {

	filePtr := flag.String("dir", ".", " Help text.")
	flag.Parse()
	txtFiles := traverseFiles(*filePtr)
	fmt.Println(txtFiles)
	for _, file := range txtFiles {
		textToTemplate(file)
		fmt.Println("file opened:", file)
	}

}

func readFile(filename string) string {
	fileContents, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	return string(fileContents)
}

func traverseFiles(dir string) []string {
	files, err := ioutil.ReadDir(dir)

	// var txtfiles []string
	output := []string{}
	check(err)

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".txt" {
			output = append(output, file.Name())
		}
	}
	return output
}

func extractFileName(filename string) string {

	var newFileName string
	newFileName = strings.SplitAfter(filename, ".")[0]
	fileExt := "html"
	fmt.Println(newFileName + fileExt)
	return "./" + newFileName + fileExt
}

func textToTemplate(filename string) {
	fileContents := readFile(filename)
	pathOut := extractFileName(filename)
	tpl, err := template.ParseFiles("template.tmpl")
	check(err)

	type Content struct {
		Contents string
	}
	content := Content{
		Contents: string(fileContents),
	}
	f, err := os.Create(pathOut)
	check(err)
	err = tpl.Execute(f, content)
	check(err)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
