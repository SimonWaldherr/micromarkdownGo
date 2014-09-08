package micromarkdown

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
)

func check(e error, t *testing.T) {
	if e != nil {
		t.Error(e)
	}
}

func nltrim(s string) string {
	return strings.Trim(s, " \n\r\t")
}

func Test_Main(t *testing.T) {
	mdfile, err1 := ioutil.ReadFile("./demo/markdown.md")
	check(err1, t)
	htmlfile, err2 := ioutil.ReadFile("./demo/html.html")
	check(err2, t)
	md := Micromarkdown(string(mdfile))
	output := []byte(strings.Replace(string(htmlfile), "$INSERT_PARSED_MARKDOWN", md, 1))
	fmt.Println(string(md))
	err3 := ioutil.WriteFile("./demo/index.html", output, 0644)
	check(err3, t)
	t.Log("please check the index.html file")
}

func Test_Bold(t *testing.T) {
	if Micromarkdown("**BOLD**") != "<b>BOLD</b>" {
		t.Error(Micromarkdown("**BOLD**"))
	}
}

func Test_Italic(t *testing.T) {
	if Micromarkdown("*italic*") != "<i>italic</i>" {
		t.Error(Micromarkdown("*italic*"))
	}
}

func Test_BoldItalic(t *testing.T) {
	if Micromarkdown("*italic and **BOLD**!*") != "<i>italic and <b>BOLD</b>!</i>" {
		t.Error(Micromarkdown("*italic and **BOLD**.*"))
	}
}

func Test_List(t *testing.T) {
	if nltrim(Micromarkdown("\n* this\n* is a\n* list\n")) != "<ul><li>this</li>\n<li>is a</li>\n<li>list</li>\n</ul>" {
		t.Error(Micromarkdown("\n* this\n* is a\n* list\n"))
	}
}

func Test_Link(t *testing.T) {
	if nltrim(Micromarkdown("[SimonWaldherr](http://simon.waldherr.eu/)")) != "<a href=\"http://simon.waldherr.eu/\">SimonWaldherr</a>" {
		t.Error(Micromarkdown("[SimonWaldherr](http://simon.waldherr.eu/)"))
	}
}

func Test_HR(t *testing.T) {
	if nltrim(Micromarkdown("---")) != "<hr/>" {
		t.Error(Micromarkdown("---"))
	}
}
