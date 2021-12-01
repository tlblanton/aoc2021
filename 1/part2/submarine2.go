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

	val1, _ := strconv.ParseInt(list[0], 10, 64)
	val2, _ := strconv.ParseInt(list[1], 10, 64)
	val3, _ := strconv.ParseInt(list[2], 10, 64)

	lastWindow := val1 + val2 + val3
	windowIncreases := 0

	for i, _ := range list[2:] {
		val1, _ := strconv.ParseInt(list[i], 10, 64)
		val2, _ := strconv.ParseInt(list[i+1], 10, 64)
		val3, _ := strconv.ParseInt(list[i+2], 10, 64)
		window := val1 + val2 + val3

		if window > lastWindow {
			windowIncreases++
		}
		lastWindow = window
	}
	fmt.Println(windowIncreases)
}
