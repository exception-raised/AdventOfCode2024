package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Pair struct {
	x int
	y int
}


func calculateMiddlePageFromCorrected(rules []string, updates [][]string) int {
	parsedRules := parseRules(rules)
	middleSum := 0

	for _, update := range updates {
		if !isUpdateValid(update, parsedRules) {
			correctedUpdate := correctUpdateOrder(update, parsedRules)
			middlePage := findMiddlePage(correctedUpdate)
			middleP, _ := strconv.Atoi(middlePage)
			middleSum += middleP
		}
	}

	return middleSum
}

func parseRules(rules []string) []Pair {
	parsedRules := []Pair{}

	for _, rule := range rules {
		splitRule := strings.Split(rule, "|")
		x, err := strconv.Atoi(splitRule[0])
		if err != nil {
			log.Panic(err)
		}
		y, err := strconv.Atoi(splitRule[1])
		if err != nil {
			log.Panic(err)
		}
		parsedRules = append(parsedRules, Pair{x: x, y: y})
	}

	return parsedRules
}

func isUpdateValid(update []string, parsedRules []Pair) bool {
	indexMap := map[int]int{}
	for idx, pageStr := range update {
		page, _ := strconv.Atoi(pageStr)
		indexMap[page] = idx
	}

	for _, rule := range parsedRules {
		idxX, existsX := indexMap[rule.x]
		idxY, existsY := indexMap[rule.y]

		if existsX && existsY {
			if idxX > idxY {
				return false
			}
		}
	}

	return true
}

func correctUpdateOrder(update []string, parsedRules []Pair) []string {
	intUpdate := []int{}
	for _, pageStr := range update {
		page, _ := strconv.Atoi(pageStr)
		intUpdate = append(intUpdate, page)
	}

	sort.Slice(intUpdate, func(i, j int) bool {
		for _, rule := range parsedRules {
			if intUpdate[i] == rule.x && intUpdate[j] == rule.y {
				return true 
			} else if intUpdate[i] == rule.y && intUpdate[j] == rule.x {
				return false 
			}
		}
		return intUpdate[i] < intUpdate[j] 
	})

	correctedUpdate := []string{}
	for _, page := range intUpdate {
		correctedUpdate = append(correctedUpdate, strconv.Itoa(page))
	}

	return correctedUpdate
}

func findMiddlePage(update []string) string {
	middleIndex := len(update) / 2
	return update[middleIndex]
}

func parseFile(filePath string) ([]string, [][]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	var rules []string
	var updates [][]string
	ruleSection := true

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			ruleSection = false
			continue
		}

		if ruleSection {
			rules = append(rules, line)
		} else {
			update := strings.Split(line, ",")
			updates = append(updates, update)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return rules, updates, nil
}

func main() {
	filePath := "../../input.txt"

	rules, updates, err := parseFile(filePath)
	if err != nil {
		log.Panic(err)
	}

	result := calculateMiddlePageFromCorrected(rules, updates)
	fmt.Println("Sum of middle pages from corrected updates:", result)
}
