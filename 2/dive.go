package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	horizontalCount, verticalCount, aim := 0, 0, 0

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		movement, _ := strconv.Atoi(line[1])
		switch line[0] {
		case "up":
			aim -= movement
		case "down":
			aim += movement
		case "forward":
			horizontalCount += movement
			verticalCount += (aim * movement)
		}
	}
	fmt.Println(verticalCount * horizontalCount)
}
