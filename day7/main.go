package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"regexp"
	"strconv"
)

func check(e error) {
    if e != nil {
			log.Fatal(e)
    }
}

func main() {
	var line, key, trimmedText, color string
	var values []string
	var arr []string
	var firstSpaceIdx, count int
	
	target := "shiny gold"
	associationMap := make(map[string]map[string]int)
	reg, err := regexp.Compile("[^a-zA-Z0-9\\s\\,]+")

	file, err := os.Open("./example2.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line = scanner.Text()
		line = reg.ReplaceAllString(line, "")
		line = strings.ReplaceAll(line, "bags", "")
		line = strings.ReplaceAll(line, "bag", "")
		arr = strings.Split(line, "contain")
		key = strings.TrimSpace(arr[0])
		values = strings.Split(arr[1], ",")

		associationMap[key] = make(map[string]int)
		for _, value := range values {
			trimmedText = strings.TrimSpace(value)
			if trimmedText != "no other" {
				firstSpaceIdx = strings.Index(trimmedText, " ")
				count, err = strconv.Atoi(trimmedText[:firstSpaceIdx])
				check(err)
				color = trimmedText[firstSpaceIdx + 1:]
				// Fill all associations for current entry
				associationMap[key][color] = count

				// If there is entry for any of the values in the map, associate current entry with it's values
				if val, ok := associationMap[color]; ok {
					for k, c := range val {
						if _, ok := associationMap[key][k]; ok {
							associationMap[key][k] += c * count
						} else {
							associationMap[key][k] = c * count
						}
					}
				}
			}
		}

		// If current entry is associated with some other entry, share associations of current entry
		for _, el := range associationMap {
			if count, ok := el[key]; ok {
				for k := range associationMap[key] {
					if _, ok := el[k]; ok {
						el[k] += associationMap[key][k] * count
					} else {
						el[k] = associationMap[key][k] * count
					}
				}
			}
		}
	}

	counter := 0
	for _, entry := range associationMap {
		if _, ok := entry[target]; ok {
			counter++
		}
	}

	counter2 := 0
	for _, value := range associationMap[target] {
		counter2 += value
	}

	fmt.Println("Part 1:", counter)
	fmt.Println("Part 2:", counter2)
	

	check(scanner.Err())
}
