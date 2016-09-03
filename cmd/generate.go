// +build ignore
package main

import (
	"fmt"
	"go/format"
	"io/ioutil"
	"os"
	"path"
	"regexp"
	"strings"

	"gopkg.in/yaml.v2"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("usage: generate in.yml out.go")
		os.Exit(1)
	}
	infile := os.Args[1]
	outfile := os.Args[2]

	yml, err := ioutil.ReadFile(infile)
	if err != nil {
		panic(err)
	}

	// load yaml locale data
	data := make(map[interface{}]interface{})
	err = yaml.Unmarshal(yml, data)
	if err != nil {
		panic(err)
	}

	// remove top-level locale and faker keys
	for _, v := range data {
		data = v.(map[interface{}]interface{})["faker"].(map[interface{}]interface{})
	}

	// generate locale var name
	name := strings.Split(path.Base(infile), ".")[0]
	parts := strings.Split(name, "-")
	parts[0] = strings.Title(parts[0])
	for i, p := range parts[1:] {
		parts[i+1] = strings.ToUpper(p)
	}
	name = strings.Join(parts, "_")

	// dump go struct
	var out string
	out += fmt.Sprintf("package locales\n\nvar %v = ", name)
	out += fmt.Sprintf("%#v\n", data)

	// reformat and line breaks
	out = regexp.MustCompile(`map\[interface\s*\{}\]`).ReplaceAllString(out, "map[string]")
	out = regexp.MustCompile(`\[\]interface\s*\{}`).ReplaceAllString(out, "[]string")
	out = regexp.MustCompile(`string\{`).ReplaceAllString(out, "string{\n")
	out = regexp.MustCompile(`("\w+":)`).ReplaceAllString(out, "\n${1}\n")
	out = regexp.MustCompile(`\s+$`).ReplaceAllString(out, "")
	out = regexp.MustCompile(`"},`).ReplaceAllString(out, "\",\n},")
	out = regexp.MustCompile(`"}},`).ReplaceAllString(out, "\",\n},\n},")
	out = regexp.MustCompile(`"}}},`).ReplaceAllString(out, "\",\n},\n},\n},")

	// special cases
	out = regexp.MustCompile(`"buzzwords":\s*\[]string`).ReplaceAllString(out, "\"buzzwords\":[][]string")
	out = regexp.MustCompile(`"bs":\s*\[]string`).ReplaceAllString(out, "\"bs\":[][]string")
	out = regexp.MustCompile(`#\{Name`).ReplaceAllString(out, "#{name")
	out = regexp.MustCompile(`#\{Address`).ReplaceAllString(out, "#{address")
	out = regexp.MustCompile(`#\{Company`).ReplaceAllString(out, "#{company")
	out = regexp.MustCompile(`#\{prefix`).ReplaceAllString(out, "#{name.prefix")
	out = regexp.MustCompile(`#\{suffix`).ReplaceAllString(out, "#{name.suffix")
	out = regexp.MustCompile(`#\{first_name`).ReplaceAllString(out, "#{name.first_name")
	out = regexp.MustCompile(`#\{nobility_title_prefix`).ReplaceAllString(out, "#{name.nobility_title_prefix")
	out = regexp.MustCompile(`#\{last_name`).ReplaceAllString(out, "#{name.last_name")
	out = regexp.MustCompile(`#\{city_prefix`).ReplaceAllString(out, "#{address.city_prefix")
	out = regexp.MustCompile(`#\{city_suffix`).ReplaceAllString(out, "#{address.city_suffix")
	out = regexp.MustCompile(`#\{building_number`).ReplaceAllString(out, "#{address.building_number")
	out = regexp.MustCompile(`#\{street_name`).ReplaceAllString(out, "#{address.street_name")
	out = regexp.MustCompile(`#\{street_suffix`).ReplaceAllString(out, "#{address.street_suffix")
	out = regexp.MustCompile(`#\{creature`).ReplaceAllString(out, "#{team.creature")

	outb, err := format.Source([]byte(out))
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(outfile, outb, 0)
	if err != nil {
		panic(err)
	}
}
