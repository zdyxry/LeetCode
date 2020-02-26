package main

import (
	"fmt"
	"sort"
	"strings"
)

func largestMultipleOfThree(digits []int) string {
	var (
		i, mod, sum int
		skipMap     = map[int]bool{}
		final       = []byte{}
		length      = len(digits)
		modIndex    = [][]int{
			nil,
			[]int{-1, -1},
			[]int{-1, -1}} // 0, 1, 2
	)

	sort.Ints(digits)
	for i = 0; i < length; i++ {
		sum += digits[length-1-i]
		mod = digits[length-1-i] % 3
		switch mod {
		case 1:
			modIndex[1][1] = modIndex[1][0]
			modIndex[1][0] = length - 1 - i
		case 2:
			modIndex[2][1] = modIndex[2][0]
			modIndex[2][0] = length - 1 - i
		}
	}
	// fmt.Println(digits, modIndex)
	switch sum % 3 {
	case 0:
		break
	case 1:
		if modIndex[1][0] == -1 && modIndex[2][1] == -1 {
			return ""
		}
		if modIndex[1][0] != -1 {
			skipMap[modIndex[1][0]] = true
		} else {
			skipMap[modIndex[2][0]] = true
			skipMap[modIndex[2][1]] = true
		}
	case 2:
		if modIndex[1][1] == -1 && modIndex[2][0] == -1 {
			return ""
		}
		if modIndex[2][0] != -1 {
			skipMap[modIndex[2][0]] = true
		} else {
			skipMap[modIndex[1][0]] = true
			skipMap[modIndex[1][1]] = true
		}
	}
	for i = 0; i < length; i++ {
		if skipMap[length-1-i] {
			continue
		}
		final = append(final, '0'+uint8(digits[length-1-i]))
	}
	rtn := strings.TrimLeft(string(final), "0")
	if len(rtn) == 0 && len(final) != 0 {
		return "0"
	}
	return rtn
}

func main() {
	digits := []int{8, 1, 9}
	res := largestMultipleOfThree(digits)
	fmt.Println(res)
}
