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

func (answers GroupAnswers) CountMatchingItems() int {
	var checkedItems string
	var matchingItems string

	for _, answer := range answers {
		for _, item := range answer {
			if strings.Contains(checkedItems, string(item)) {
				continue
			}

			matching := true
			for _, answer := range answers {
				if !strings.Contains(answer, string(item)) {
					matching = false
					break
				}
			}

			if matching {
				matchingItems += string(item)
			}

			checkedItems += string(item)
		}
	}

	return len(matchingItems)
}
