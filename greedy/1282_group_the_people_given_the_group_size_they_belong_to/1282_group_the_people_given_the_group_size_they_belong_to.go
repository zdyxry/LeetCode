package main

import "fmt"

func groupThePeople(groupSizes []int) [][]int {
	var groups [][]int
	groupMemberCount := make(map[int][]int)

	for i := 0; i < len(groupSizes); i++ {
		groupMemberCount[groupSizes[i]] = append(groupMemberCount[groupSizes[i]], i)
	}

	for groupSize, groupMembers := range groupMemberCount {
		subGroup := []int{}
		for i := 0; i < len(groupMembers); i++ {
			if i%groupSize == groupSize-1 {
				subGroup = append(subGroup, groupMembers[i])
				groups = append(groups, subGroup)
				subGroup = []int{}
				continue
			}
			subGroup = append(subGroup, groupMembers[i])
		}
	}
	return groups
}

func main() {
	groupSizes := []int{3, 3, 3, 3, 3, 1, 3}
	res := groupThePeople(groupSizes)
	fmt.Println(res)
}
