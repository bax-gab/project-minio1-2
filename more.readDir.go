package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	localfiles, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range localfiles {
		fmt.Println(f.Size(), "Directory?:", f.IsDir(), f.Mode, f.ModTime())
	}
}
