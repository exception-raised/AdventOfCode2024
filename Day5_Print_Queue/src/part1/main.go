package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Pair struct {
	x int
	y int
}

func calculateMiddlePage(rules []string, updates [][]string) int {
	parsedRules := []Pair{}

	for _, rule := range rules {
		splitRule := strings.Split(rule, "|")
		x, err := strconv.Atoi(splitRule[0])
		if err != nil {
			log.Panic(err)
			return 0
		}
		y, err := strconv.Atoi(splitRule[1])
		if err != nil {
			log.Panic(err)
			return 0
		}
		myPair := Pair{x: x, y: y}
		fmt.Println(myPair.x, myPair.y)
		parsedRules = append(parsedRules, myPair)
	}

	middleSum := 0

	for _, update := range updates {
		if isUpdateValid(update, parsedRules) {
			middlePage := findMiddlePage(update)
			middleP, _ := strconv.Atoi(middlePage)
			middleSum += middleP
		}
	}

	return middleSum
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

func findMiddlePage(update []string) string {
	middleIndex := len(update) / 2

	m := update[middleIndex]

	return m
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
		return
	}

	result := calculateMiddlePage(rules, updates)
	fmt.Println("Sum of middle pages from valid updates:", result)
}