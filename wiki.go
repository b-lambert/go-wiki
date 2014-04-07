package main

import (
	"fmt"
	"io/ioutil"
)

type Page struct {
	Title string
	Body []byte
}

// Input: pointer to a Page
// Output: Returns an "error" indicating result of writing the page to file.
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func main() {
	p1 := &Page{Title: "Index", Body: []byte("Welcome to the Wiki!")}
	p1.save()
	p2, _ := loadPage("Index")
	fmt.Println(string(p2.Body))
}
