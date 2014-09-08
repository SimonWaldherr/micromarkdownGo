#micromarkdown.go

[![Build Status](https://travis-ci.org/SimonWaldherr/micromarkdownGo.svg?branch=master)](https://travis-ci.org/SimonWaldherr/micromarkdownGo)
[![GoDoc](https://godoc.org/github.com/SimonWaldherr/micromarkdownGo?status.svg)](https://godoc.org/github.com/SimonWaldherr/micromarkdownGo)
[![Gittip donate button](http://img.shields.io/gittip/bevry.png)](https://www.gittip.com/SimonWaldherr/ "Donate weekly to this project using Gittip")
[![Flattr donate button](https://raw.github.com/balupton/flattr-buttons/master/badge-89x18.gif)](https://flattr.com/submit/auto?user_id=SimonWaldherr&url=http%3A%2F%2Fgithub.com%2FSimonWaldherr%2FmicromarkdownGo "Donate monthly to this project using Flattr")


convert [markdown](http://en.wikipedia.org/wiki/Markdown) to [HTML](http://en.wikipedia.org/wiki/HTML) via golang  
this is the golang version of µmarkdown, take a look at the
[PHP version](https://github.com/SimonWaldherr/micromarkdown.php) and the
[JavaScript version](https://github.com/SimonWaldherr/micromarkdown.js).

##about

License:   MIT  
Version: 0.1.2  
Date:  09.2014  

##howto

###test

```sh
wget https://github.com/SimonWaldherr/micromarkdown.go/archive/master.zip --no-check-certificate
unzip master.zip
cd ./micromarkdown.go-master/demo/
go run example.go
```

###use

```sh
go get github.com/SimonWaldherr/micromarkdown.go
```

```go
import mmd "github.com/SimonWaldherr/micromarkdown.go"
import "fmt"

func main() {
  md := mmd.Micromarkdown("#title\n\nlorem *ipsum* dolor sit\namet pluribus **procrastinatio**\n")
  fmt.Println(string(md))
}
```

##contact

Feel free to contact me via [eMail](mailto:contact@simonwaldherr.de) or on [Twitter](http://twitter.com/simonwaldherr). This software will be continually developed. Suggestions and tips are always welcome.
