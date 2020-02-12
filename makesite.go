package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"context"

	"cloud.google.com/go/translate"
	"golang.org/x/text/language"
)

// func main() {
// 	fmt.Println(translateText("es", "Translate this string from english"))
// }

func main() {

	filePtr := flag.String("dir", "txt", " Help text.")
	flag.Parse()
	txtFiles := traverseFiles(*filePtr)
	fmt.Println(txtFiles)
	for _, file := range txtFiles {
		textToTemplate(file)
		fmt.Println("file opened:", file)
	}

	http.Handle("/", http.FileServer(http.Dir("./static/views")))
	fmt.Println("Preparing to listen on port 8081...")
	log.Fatal(http.ListenAndServe(":8081", nil))

}

func readFile(filepath string) string {
	fileContents, err := ioutil.ReadFile(filepath)
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
			output = append(output, dir+"/"+file.Name())
			fmt.Println(file.Name())
		}
	}
	return output
}

func extractFileName(filename string) string {
	fileExt := "html"
	outfile := strings.Split(filename, "/")
	outfile = (strings.Split(outfile[len(outfile)-1], "."))
	return "./static/views/" + outfile[0] + "." + fileExt
}

func textToTemplate(filename string) {
	fileContents := readFile(filename)
	translatedFileContents, err := translateText("es", fileContents)
	pathOut := extractFileName(filename)
	tpl, err := template.ParseFiles("./static/templates/template.tmpl")
	check(err)

	type Content struct {
		Contents string
	}
	content := Content{
		Contents: string(translatedFileContents),
	}
	f, err := os.Create(pathOut)
	check(err)
	err = tpl.Execute(f, content)
	check(err)
}

func translateText(targetLanguage, text string) (string, error) {
	// text := "The Go Gopher is cute"
	ctx := context.Background()

	lang, err := language.Parse(targetLanguage)
	if err != nil {
		return "", fmt.Errorf("language.Parse: %v", err)
	}

	client, err := translate.NewClient(ctx)
	if err != nil {
		return "", err
	}
	defer client.Close()

	resp, err := client.Translate(ctx, []string{text}, lang, nil)
	if err != nil {
		return "", fmt.Errorf("Translate: %v", err)
	}
	if len(resp) == 0 {
		return "", fmt.Errorf("Translate returned empty response to text: %s", text)
	}
	return resp[0].Text, nil
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
