package micromarkdown

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Stack struct {
	nodes []string
	count int
}

func Lifo() *Stack {
	return &Stack{}
}

func (s *Stack) Push(n string) {
	s.nodes = append(s.nodes[:s.count], n)
	s.count++
}

func (s *Stack) Pop() string {
	if s.count == 0 {
		return ""
	}
	s.count--
	return s.nodes[s.count]
}

func Micromarkdown(str string) string {
	var stra [][]string
	var indent int = -1
	var repstr string
	var count, casca, status, nstatus int
	var line, reparr, helper, helparr1, helparr2 []string
	helper1 := Lifo()

	htmlEncode := strings.NewReplacer("<", "&lt;", ">", "&gt;", "&", "&#38;", "\"", "&#34;", "'", "&#39;", " ", "&nbsp;", "*", "&#8727;", "\t", "&nbsp;&nbsp;", "\n", "\n<br>")
	headline, _ := regexp.Compile("(?m)^(\\#{1,6})([^\\#\\n]+)$")
	code, _ := regexp.Compile("(?m)\\s```\\n?([^`]+)```")
	hr, _ := regexp.Compile("(?m)^([\\*\\-_] ?){3,}$")
	lists, _ := regexp.Compile("(?m)^((\\s*((\\*|\\-)|\\d(\\.|\\))) [^\\n]+)\\n)+")
	liner, _ := regexp.Compile("(?m)^((\\s*)((\\*|\\-)|\\d(\\.|\\))) ([^\\n]+))")
	bolditalics, _ := regexp.Compile("(?m)([\\*_~]{1,3})([^\\*_~\\n]+[^\\*_~\\s])([\\*_~]{1,3})")
	links, _ := regexp.Compile("!?\\[([^\\]<>]+)\\]\\(([^ \\)<>]+)( \"[^\\(\\)\"]+\")?\\)")
	mail, _ := regexp.Compile("<(([a-z0-9_\\-\\.])+\\@([a-z0-9_\\-\\.])+\\.([a-z]{2,7}))>")
	tables, _ := regexp.Compile("(?m)\\n(([^|\\n]+ *\\| *)+([^|\\n]+\\n))(\\-+\\|)+(\\-+\\n)((([^|\\n]+ *\\| *)+([^|\\n]+)\\n)+)")

	/* code */
	stra = code.FindAllStringSubmatch(str, -1)
	for i := 0; i < len(stra); i++ {
		str = strings.Replace(str, stra[i][0], "<code>"+htmlEncode.Replace(strings.TrimSpace(stra[i][1]))+"</code>", 1)
	}

	/* headlines */
	stra = headline.FindAllStringSubmatch(str, -1)
	for i := 0; i < len(stra); i++ {
		count = len(stra[i][1])
		str = strings.Replace(str, stra[i][0], "<h"+strconv.Itoa(count)+">"+strings.TrimSpace(stra[i][2])+"</h"+strconv.Itoa(count)+">", 1)
	}

	/* lists */
	stra = lists.FindAllStringSubmatch(str, -1)
	for i := 0; i < len(stra); i++ {
		casca = 0
		if strings.TrimSpace(stra[i][0])[0:1] == "*" {
			repstr = "<ul>"
		} else {
			repstr = "<ol>"
		}
		helper = strings.Split(strings.TrimSpace(stra[i][0]), "\n")
		status = 0
		indent = -1
		for j := 0; j < len(helper); j++ {
			line = liner.FindStringSubmatch(helper[j])
			if len(line[2]) == 0 {
				nstatus = 0
			} else {
				if indent == -1 {
					indent = len(strings.Replace(line[2], "\t", "    ", -1))
				}
				nstatus = int(len(strings.Replace(line[2], "\t", "    ", -1)) / indent)
			}
			for status > nstatus {
				repstr += helper1.Pop()
				status--
				casca--
			}
			for status < nstatus {
				if strings.TrimSpace(stra[i][0])[0:1] == "*" {
					repstr += "<ul>"
					helper1.Push("</ul>")
				} else {
					repstr += "<ol>"
					helper1.Push("</ol>")
				}
				status++
				casca++
			}
			repstr += "<li>" + line[6] + "</li>\n"
		}
		for casca > 0 {
			repstr += "</ul>"
			casca--
		}
		if strings.TrimSpace(stra[i][0])[0:1] == "*" {
			repstr += "</ul>"
		} else {
			repstr += "</ol>"
		}
		str = strings.Replace(str, stra[i][0], repstr+"\n", 1)
	}

	/* tables */
	stra = tables.FindAllStringSubmatch(str, -1)
	for i := 0; i < len(stra); i++ {
		repstr = "<table><tr>"
		helper = strings.Split(stra[i][1], "|")
		for j := 0; j < len(helper); j++ {
			repstr += "<th>" + helper[j] + "</th>"
		}
		repstr += "</tr>"
		helparr1 = strings.Split(stra[i][6], "\n")
		for j := 0; j < len(helparr1); j++ {
			helparr2 = strings.Split(helparr1[j], "|")
			if len(helparr2[0]) != 0 {
				repstr += "<tr>"
				for k := 0; k < len(helparr2); k++ {
					repstr += "<td>" + helparr2[k] + "</td>"
				}
				repstr += "</tr>\n"
			}
		}
		repstr += "</table>"
		fmt.Println(repstr)
		str = strings.Replace(str, stra[i][0], repstr, 1)
	}

	/* bold and italics */
	for i := 0; i < 3; i++ {
		stra = bolditalics.FindAllStringSubmatch(str, -1)
		for j := 0; j < len(stra); j++ {
			if stra[j][1] == "~~" {
				str = strings.Replace(str, stra[j][0], "<del>"+stra[j][2]+"</del>", 1)
			} else {
				switch len(stra[j][1]) {
				case 1:
					reparr = []string{"<i>", "</i>"}
				case 2:
					reparr = []string{"<b>", "</b>"}
				case 3:
					reparr = []string{"<i><b>", "</b></i>"}
				}
				str = strings.Replace(str, stra[j][0], reparr[0]+stra[j][2]+reparr[1], 1)
			}
		}
	}

	/* links */
	stra = links.FindAllStringSubmatch(str, -1)
	for i := 0; i < len(stra); i++ {
		if strings.TrimSpace(stra[i][0])[0:1] == "!" {
			str = strings.Replace(str, stra[i][0], "<img src=\""+stra[i][2]+"\" alt=\""+stra[i][1]+"\" title=\""+stra[i][1]+"\" />\n", -1)
		} else {
			str = strings.Replace(str, stra[i][0], "<a href=\""+stra[i][2]+"\">"+stra[i][1]+"</a>\n", -1)
		}
	}
	stra = mail.FindAllStringSubmatch(str, -1)
	for i := 0; i < len(stra); i++ {
		str = strings.Replace(str, stra[i][0], "<a href=\"mailto:"+stra[i][1]+"\">"+stra[i][1]+"</a>\n", -1)
	}

	/* horizontal line */
	stra = hr.FindAllStringSubmatch(str, -1)
	for i := 0; i < len(stra); i++ {
		str = strings.Replace(str, stra[i][0], "\n<hr/>\n", -1)
	}

	return str
}
