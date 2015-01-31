package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"gopkg.in/yaml.v1"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("usage: convert file.yml")
		os.Exit(1)
	}

	filename := os.Args[1]
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	data := make(map[interface{}]interface{})
	err = yaml.Unmarshal(file, data)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// remove top-level locale and faker keys
	for _, v := range data {
		data = v.(map[interface{}]interface{})["faker"].(map[interface{}]interface{})
	}

	name := strings.Split(path.Base(filename), ".")[0]
	parts := strings.Split(name, "-")
	parts[0] = strings.Title(parts[0])
	for i, p := range parts[1:] {
		parts[i+1] = strings.ToUpper(p)
	}
	name = strings.Join(parts, "_")

	fmt.Printf("package locales\n\nvar %v = ", name)
	fmt.Printf("%#v\n", data)
}
