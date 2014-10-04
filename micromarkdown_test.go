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
	fmt.Println("\n\nStarting tests...\nplease wait...\n")
}

func Test_Bold(t *testing.T) {
	if Micromarkdown("**BOLD**") != "<b>BOLD</b>" {
		t.Error(Micromarkdown("**BOLD**"))
	}
}

func Test_Headline(t *testing.T) {
	if Micromarkdown("###Head") != "<h3>Head</h3>" {
		t.Error(Micromarkdown("###Head"))
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

func Test_Table(t *testing.T) {
	if nltrim(Micromarkdown("\nthis | left | center   | right\n-----|--------|----------|-------\nwith | sample | content  | for\nlorem| ipsum  | dolor    | sit\nsit  | amet   | sed      | do\ndo   | eiusom | tempor   | with\n")) != "<table><tr><th>this </th><th> left </th><th> center   </th><th> right\n</th></tr><tr><td>with </td><td> sample </td><td> content  </td><td> for</td></tr>\n<tr><td>lorem</td><td> ipsum  </td><td> dolor    </td><td> sit</td></tr>\n<tr><td>sit  </td><td> amet   </td><td> sed      </td><td> do</td></tr>\n<tr><td>do   </td><td> eiusom </td><td> tempor   </td><td> with</td></tr>\n</table>" {
		t.Error(Micromarkdown("\nthis | left | center   | right\n-----|--------|----------|-------\nwith | sample | content  | for\nlorem| ipsum  | dolor    | sit\nsit  | amet   | sed      | do\ndo   | eiusom | tempor   | with\n"))
	}
}
