package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"text/template"
)

// type Content struct {
// 	content string
// }

// type Todo struct {
// 	Name        string
// 	Description string
// }

func main() {

	// fileContents := readFile("./first-post.txt")
	// traverseFiles()
	// filePtr := flag.String("example", "./first-post.txt", "enter a file name flag")
	filePtr := flag.String("example", "defaultValue", " Help text.")
	flag.Parse()
	fmt.Println("file:", *filePtr)
	// fmt.Println("file:", *filePtr)
	writeOut(*filePtr)
	fmt.Print(string("here \n"))
	// fmt.Print(string(files))
}

func readFile(filename string) string {
	fileContents, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	return string(fileContents)
}

func traverseFiles() {
	// Return .txt files
	// var i := 0 uint64
	directory := "."
	files, err := ioutil.ReadDir(directory)

	// var txtfiles []string
	output := []string{}
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".txt" {
			// output = append(txtfiles, file.Name())
			output = append(output, file.Name())
		}
	}
	// fmt.Println(strings.Join(output, ", "))
	// fmt.Printf(txtfiles)
	// return output
}

// func main() {
// 	// Files are provided as a slice of strings.
// 	paths := []string{traverseFiles()}

// 	// t := template.Must(template.New("html-tmpl").ParseFiles(paths...))
// 	// err = t.Execute(os.Stdout, todos)
// 	// if err != nil {
// 	//   panic(err)
// 	// }
// }

// func txtToTemplate(fileContents string) {
// 	paths := []string{
// 		"todo.tmpl",
// 	}

// 	t := template.Must(template.New("html-tmpl").ParseFiles(paths...))
// 	err = t.Execute(os.Stdout, todos)
// 	if err != nil {
// 		panic(err)
// 	}
// }

func writeOut(filename string) {
	// paths := []string{
	// 	"template.tmpl",
	// }
	fileContents := readFile(filename)
	tpl, err := template.ParseFiles("template.tmpl")
	if err != nil {
		log.Fatalln(err)
	}
	type Content struct {
		Contents string
	}
	content := Content{
		Contents: string(fileContents),
	}
	// fmt.Printf(content.Contents)
	// err = tpl.Execute(os.Stdout, content)

	// bytesToWrite := []byte(fileContents)
	err = tpl.Execute(os.Stdout, content)
	if err != nil {
		panic(err)
	}

}
