//go:generate go-bindata store

package main

import (
	"html/template"
	"log"
	"os"
	"strings"
)

type TemplateConfig struct {
	Name string
}

func main() {
	if len((os.Args)) < 3 {
		log.Fatal("Expected usage: jstore TypeName typefile.go")
	}

	data, err := Asset("store/store.go")
	if err != nil {
		log.Fatal(err)
	}

	text := string(data)
	text = strings.Replace(text, "DataItem", os.Args[1], -1)
	text = strings.Replace(text, "Collection", os.Args[1]+"Collection", -1)
	t, err := template.New("store").Parse(text)

	f, err := os.Create(os.Args[2])
	if err != nil {
		log.Println("create file: ", err)
		return
	}
	defer f.Close()

	t.Execute(f, TemplateConfig{})
}
