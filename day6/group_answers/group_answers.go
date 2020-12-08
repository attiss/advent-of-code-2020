package group_answers

import (
	"strings"
)

type GroupAnswers []string

func (answers GroupAnswers) CountUniqueItems() int {
	var uniqueItems string
	for _, answer := range answers {
		for _, item := range answer {
			if !strings.Contains(uniqueItems, string(item)) {
				uniqueItems += string(item)
			}
		}
	}
	return len(uniqueItems)
}
