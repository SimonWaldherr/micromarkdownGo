package main

import (
	"fmt"
	"strings"
	"io/ioutil"
	"./micromarkdown"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	mdfile, err1 := ioutil.ReadFile("./markdown.md")
	check(err1)
	htmlfile, err2 := ioutil.ReadFile("./html.html")
	check(err2)
	md := micromarkdown.Micromarkdown(string(mdfile))
	output := []byte(strings.Replace(string(htmlfile), "$INSERT_PARSED_MARKDOWN", md, 1))
	fmt.Println(string(md))
	err3 := ioutil.WriteFile("./index.html", output, 0644)
	check(err3)
}
