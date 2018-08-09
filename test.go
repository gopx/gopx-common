package main

import (
	"fmt"
	"log"

	yaml "gopkg.in/yaml.v2"
)

var data = `
bdd:
  c: 2
  d: [3, 4]
`

type T struct {
	A *string `json:"a" yaml:"a"`
	B struct {
		RenamedC *int  `json:"c" yaml:"c"`
		D        []int `json:"c" yaml:"d"`
	} `json:"b" yaml:"b"`
}

func main() {
	t := T{}

	err := yaml.Unmarshal([]byte(data), &t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("%v", t)

	// file, _ := os.Open("/Users/rousan/Random/electron-quick-start 2/gows.tar.gz")

	// fmt.Println(fs.ReadEntryTarGz(file, os.Stdout, []string{"sdfds", "gopx.yml"}))

}
