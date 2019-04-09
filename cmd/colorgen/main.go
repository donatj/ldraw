package main

import (
	"flag"
	"fmt"
	"html/template"
	"image/color"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	re = regexp.MustCompile(`COLOUR\s+(\S+)\s+CODE\s+(\d+)\s+VALUE\s+#(\S{6})\s+EDGE\s+#(\S{6})(?:\s+ALPHA\s+(\d+))?`)

	pkg = flag.String("package", "ldraw", "package name to output")
	arg = ""
)

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s [options] <file>:\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()

	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(2)
	}

	arg = flag.Arg(0)
}

const templ = `// Code generated by "colorgen"; DO NOT EDIT
package {{.Package}}

import (
	"image/color"
)

// BrickColor is a brick color number as defined by http://www.ldraw.org/library/official/ldcfgalt.ldr
type BrickColor int

const (
{{range .Colors}}	// {{.Name}} as defined by: 
	//    {{.Descr}} 
	{{.Name}} BrickColor = {{.Value}}
{{end}}
)

// BrickColorMap connects BrickColor's to their coresponding color.Color
var BrickColorMap = map[BrickColor]color.Color{
{{range .Colors}}	{{.Name}}: {{.Color}},
{{end}}
}

`

type templateColor struct {
	Name  string
	Value string
	Descr string
	Color string
}

type templateData struct {
	Package string
	Colors  []templateColor
}

func main() {
	res, err := http.Get("http://www.ldraw.org/library/official/ldcfgalt.ldr")
	if err != nil {
		log.Fatal(err.Error())
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err.Error())
	}

	matches := re.FindAllStringSubmatch(string(b), -1)

	td := templateData{
		Package: *pkg,
		Colors:  []templateColor{},
	}

	for _, match := range matches {
		name := strings.ReplaceAll(match[1], "_", "")

		// spew.Dump(match)
		c := hex(match[3])
		if match[5] != "" {
			i, err := strconv.ParseUint(match[5], 10, 8)
			if err != nil {
				log.Fatal(err)
			}

			c.A = uint8(i)
		}

		td.Colors = append(td.Colors, templateColor{
			name,
			match[2],
			match[0],
			fmt.Sprintf("%#v", c),
		})
	}

	f, err := os.Create(arg)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer f.Close()

	tmpl := template.New("code template")
	tmpl, err = tmpl.Parse(templ)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = tmpl.Execute(f, td)
	if err != nil {
		log.Fatal(err.Error())
	}

}

func hex(s string) (c color.RGBA) {
	var err error

	c.A = 0xff
	switch len(s) {
	case 6:
		_, err = fmt.Sscanf(s, "%02x%02x%02x", &c.R, &c.G, &c.B)
	case 3:
		_, err = fmt.Sscanf(s, "%1x%1x%1x", &c.R, &c.G, &c.B)
		// Double the hex digits:
		c.R *= 17
		c.G *= 17
		c.B *= 17
	default:
		err = fmt.Errorf("invalid length, must be 6 or 3")

	}

	if err != nil {
		log.Fatal(err)
	}

	return
}
