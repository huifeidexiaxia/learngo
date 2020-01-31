package main

import (
	"bufio"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

func main() {
	yamlFile, e := ioutil.ReadFile("src/yaml/test.yaml")
	if e != nil {
		fmt.Print(e)
	}
	fmt.Println("begin")

	fileMap := make(map[string]interface{})
	e = yaml.Unmarshal(yamlFile, fileMap)
	if e != nil {
		fmt.Println(e)
	}

	testMap := make(map[string]interface{})
	testMap["test"] = fileMap
	out, e := yaml.Marshal(&testMap)
	fmt.Println(string(out))

	file, _ := os.OpenFile("test-out.yaml", os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0644)
	buf := bufio.NewWriter(file)
	fmt.Fprintf(buf, string(out))
	buf.Flush()

}
