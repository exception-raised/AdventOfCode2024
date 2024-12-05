package shared

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Pair struct {
	X int
	Y int
}

func ParseRules(rules []string) []Pair {
	parsedRules := []Pair{}

	for _, rule := range rules {
		splitRule := strings.Split(rule, "|")
		x, _ := strconv.Atoi(splitRule[0])
		y, _ := strconv.Atoi(splitRule[1])
		parsedRules = append(parsedRules, Pair{X: x, Y: y})
	}

	return parsedRules
}

func ParseFile(filePath string) ([]string, [][]string, error) {
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

func FindMiddlePage(update []string) string {
	middleIndex := len(update) / 2
	return update[middleIndex]
}

func IsUpdateValid(update []string, parsedRules []Pair) bool {
	indexMap := map[int]int{}
	for idx, pageStr := range update {
		page, _ := strconv.Atoi(pageStr)
		indexMap[page] = idx
	}

	for _, rule := range parsedRules {
		idxX, existsX := indexMap[rule.X]
		idxY, existsY := indexMap[rule.Y]

		if existsX && existsY {
			if idxX > idxY {
				return false
			}
		}
	}

	return true
}