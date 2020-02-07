package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

func main() {

	filePtr := flag.String("example", "defaultValue", " Help text.")
	flag.Parse()
	fmt.Println("file opened:", *filePtr)
	textToTemplate(*filePtr)
}

func readFile(filename string) string {
	fileContents, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	return string(fileContents)
}

func traverseFiles() {
	directory := "."
	files, err := ioutil.ReadDir(directory)

	// var txtfiles []string
	output := []string{}
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".txt" {
			output = append(output, file.Name())
		}
	}
}

func extractFileName(filename string) string {

	var newFileName string
	newFileName = strings.SplitAfter(filename, ".")[0]
	fmt.Println(newFileName)
	fileExt := "html"
	fmt.Println(newFileName + fileExt)
	return "./" + newFileName + fileExt
}

func textToTemplate(filename string) {
	fileContents := readFile(filename)
	fileOut := extractFileName(filename)
	tpl, err := template.ParseFiles("template.tmpl")
	check(err)
	type Content struct {
		Contents string
	}
	content := Content{
		Contents: string(fileContents),
	}
	fmt.Printf(content.Contents)
	// err = tpl.Execute(os.Stdout, content)

	// bytesToWrite := []byte(fileContents)
	err = tpl.Execute(os.Stdout, content)
	if err != nil {
		panic(err)
	}
	err1 := ioutil.WriteFile(fileOut, []byte(content.Contents), 0644)
	check(err1)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
