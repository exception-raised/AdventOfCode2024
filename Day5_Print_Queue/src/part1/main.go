package main

import (
	"Day5_Print_Queue/src/shared"
	"fmt"
	"strconv"
)

func calculateMiddlePage(rules []string, updates [][]string) int {
	parsedRules := shared.ParseRules(rules)

	middleSum := 0

	for _, update := range updates {
		if shared.IsUpdateValid(update, parsedRules) {
			middlePage := shared.FindMiddlePage(update)
			middleP, _ := strconv.Atoi(middlePage)
			middleSum += middleP
		}
	}

	return middleSum
}


func main() {
	filePath := "../../input.txt"

	rules, updates, err := shared.ParseFile(filePath)

	if err != nil {
		return
	}

	result := calculateMiddlePage(rules, updates)
	fmt.Println("Sum of middle pages from valid updates:", result)
}