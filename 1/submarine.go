package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
		return
	}

	list := strings.Split(string(content), "\n")

	lastElementInt, _ := strconv.ParseInt(list[0], 10, 64)
	increases := 0

	for _, element := range list[1:] {
		elementInt, _ := strconv.ParseInt(element, 10, 64)

		if elementInt > lastElementInt {
			fmt.Println(elementInt, "(increased)")
			increases++
		} else {
			fmt.Println(elementInt, "(decreased)")
		}
		lastElementInt = elementInt
	}

	fmt.Println(increases)
}
