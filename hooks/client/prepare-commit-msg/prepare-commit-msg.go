#!/usr/bin/env gorun

package main

import (
	"io/ioutil"
	"log"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 1 {
		log.Fatal("Commit message file name is missing!")
	}
	d1 := []byte("# Please include a useful commit message!")
	if err := ioutil.WriteFile(args[1], d1, 0644); err != nil {
		log.Fatal("Write to file err:", err)
	}
}
