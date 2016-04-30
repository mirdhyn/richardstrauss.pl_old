package main

import (
	"html/template"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type websiteContent struct {
	Title    string
	Home     home
	About    about
	News     news
	Projects projects
	Links    links
}

type home struct {
	Title string
}

type about struct {
	Title   string
	Content []string
}

type news struct {
	Title    string
	Articles []subsection
}

type projects struct {
	Title    string
	Projects []subsection
}

type links struct {
	Title    string
	Sections []subsection
}

type subsection struct {
	Title   string
	Date    string
	links   []link
	Picture string
	Content []string
}

type link struct {
	Desc string
	link string
}

func main() {
	data := websiteContent{}
	yamlData, _ := ioutil.ReadFile("data/pl.yaml")
	_ = yaml.Unmarshal(yamlData, &data)

	f, _ := os.Create("index.html")
	t, _ := template.ParseFiles("index.tmpl")
	t.Execute(f, data)
}
