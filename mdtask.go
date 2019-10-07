package main

import (
	"fmt"
	"io/ioutil"
	"github.com/gomarkdown/markdown"
)

type Article struct {
	Filename string
	Body []byte
}

func loadFile(title string) *Article {
	filename := title + ".md"
	body, _ := ioutil.ReadFile(filename)
	return &Article{Filename: filename, Body: body}
}

func main() {
	File := loadFile("test")
	md := File.Body
	html := markdown.ToHTML(md, nil, nil)
	fmt.Printf("%s\n", html)
}

