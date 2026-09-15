package stats

import (
	"fmt"
	"sort"
)

type keyValue struct {
	Key   string
	Value int
}

func PrintStats(levelCount, messageCount map[string]int) {
	fmt.Println("Statistic:")

	for key, value := range levelCount {
		fmt.Printf("%s: %d\n", key, value)
	}

	var messages []keyValue
	for key, value := range messageCount {
		messages = append(messages, keyValue{key, value})
	}
	sort.Slice(messages, func(i, j int) bool {
		return messages[i].Value > messages[j].Value
	})

	for i, item := range messages {
		if i >= 10 {
			break
		}
		fmt.Printf("%d. %s:  %d\n", i+1, item.Key, item.Value)
	}
}
