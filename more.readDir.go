package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func readthenwrite() {

}

func main() {
	var filecounter int
	localfiles, err := ioutil.ReadDir("/Users/gb/go/junkfiles")
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range localfiles {
		fmt.Println(f.Name(), f.Size(), "Directory?:", f.IsDir(), f.Mode, f.ModTime())
		filecounter++
	}
	fmt.Println("\n", "Number of files processed: ", filecounter, len(localfiles))
}
