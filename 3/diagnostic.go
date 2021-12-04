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
	gammaRateBin, epsilonRateBin := getGammaAndEpsilonBinaries(list)

	// Oxygen calculations
	oxygenGeneratorContenders := list
	for i := 0; i < len(gammaRateBin); i++ {
		if len(oxygenGeneratorContenders) < 2 {
			break
		}
		oxygenGeneratorContenders = keepElementsWithXinPositionY(oxygenGeneratorContenders, string(gammaRateBin[i]), i)
		gammaRateBin, _ = getGammaAndEpsilonBinaries(oxygenGeneratorContenders) // must recalculate with new list
	}

	// CO2 calculations
	co2ScrubberContenders := strings.Split(string(content), "\n")
	for i := 0; i < len(epsilonRateBin); i++ {
		if len(co2ScrubberContenders) < 2 {
			break
		}
		co2ScrubberContenders = keepElementsWithXinPositionY(co2ScrubberContenders, string(epsilonRateBin[i]), i)
		_, epsilonRateBin = getGammaAndEpsilonBinaries(co2ScrubberContenders) // must recalculate with new list
	}

	// Converting bin to decimal
	oxy, err := strconv.ParseInt(oxygenGeneratorContenders[0], 2, 64)
	if err != nil {
		panic(err)
	}
	co2, err := strconv.ParseInt(co2ScrubberContenders[0], 2, 64)
	if err != nil {
		panic(err)
	}

	fmt.Println(oxygenGeneratorContenders, co2ScrubberContenders)
	fmt.Println(oxy, co2)
}

func keepElementsWithXinPositionY(strList []string, findChar string, pos int) []string {

	returnList := strList
	justDeleted := false

	for i := 0; i < len(strList); i++ {
		// Because we are slicing lists, we need this to reset our i when we delete a value,
		// so we check the new nth position after slicing a new value over into it
		if justDeleted {
			i -= 1
			justDeleted = false
		}

		if string(strList[i][pos]) != findChar {
			// if the strList element does not have the proper binary bit in the given position, get it outta here
			for j := 0; j < len(returnList); j++ {
				if returnList[j] == strList[i] {
					returnList = RemoveIndex(returnList, j)
					justDeleted = true
				}
			}
		}
	}
	return returnList
}

func RemoveIndex(s []string, index int) []string {
	if index >= len(s) || index < 0 {
		return s
	}
	// Slicing things in order to delete
	return append(s[:index], s[index+1:]...)
}

func getGammaAndEpsilonBinaries(binList []string) (string, string) {
	gammaRateBin := ""
	epsilonRateBin := ""

	var mapList []map[string]int // list of maps. There will be one map for each vertical rown, containing 0, and 1 as keys, and whose values with be frequency
	for i := 0; i < len(binList[0]); i++ {
		mapList = append(mapList, map[string]int{})
	}

	// iterating over every binary value in list
	for _, element := range binList[0:] {
		for j, bit := range element {
			mapList[j][string(bit)]++
		}
	}

	// iterating over map to see most common element per row
	for _, element := range mapList {
		if element["0"] == element["1"] {
			gammaRateBin += "1"
			epsilonRateBin += "0"
		} else if element["0"] > element["1"] {
			gammaRateBin += "0"
			epsilonRateBin += "1"
		} else {
			gammaRateBin += "1"
			epsilonRateBin += "0"
		}
	}
	return gammaRateBin, epsilonRateBin
}
