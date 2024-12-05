package main

import (
	"Day5_Print_Queue/src/shared"
	"fmt"
	"log"
	"sort"
	"strconv"
)

func calculateMiddlePageFromCorrected(rules []string, updates [][]string) int {
	parsedRules := shared.ParseRules(rules)
	middleSum := 0

	for _, update := range updates {
		if !shared.IsUpdateValid(update, parsedRules) {
			correctedUpdate := correctUpdateOrder(update, parsedRules)
			middlePage := shared.FindMiddlePage(correctedUpdate)
			middleP, _ := strconv.Atoi(middlePage)
			middleSum += middleP
		}
	}

	return middleSum
}


func correctUpdateOrder(update []string, parsedRules []shared.Pair) []string {
	intUpdate := []int{}
	for _, pageStr := range update {
		page, _ := strconv.Atoi(pageStr)
		intUpdate = append(intUpdate, page)
	}

	sort.Slice(intUpdate, func(i, j int) bool {
		for _, rule := range parsedRules {
			if intUpdate[i] == rule.X && intUpdate[j] == rule.Y {
				return true 
			} else if intUpdate[i] == rule.Y && intUpdate[j] == rule.X {
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

func main() {
	filePath := "../../input.txt"

	rules, updates, err := shared.ParseFile(filePath)
	if err != nil {
		log.Panic(err)
	}

	result := calculateMiddlePageFromCorrected(rules, updates)
	fmt.Println("Sum of middle pages from corrected updates:", result)
}
