package main

import (
	"fmt"
	"github.com/gomarkdown/markdown"
	"io/ioutil"
)

type Article struct {
	Filename string
	Body     []byte
}

func loadFile(title string) *Article {
	filename := title + ".md"
	body, _ := ioutil.ReadFile(filename)
	return &Article{Filename: filename, Body: body}
}

func main() {
	File := loadFile("DESIGN")
	md := File.Body
	html := markdown.ToHTML(md, nil, nil)
	fmt.Printf("%s\n", html)
}
