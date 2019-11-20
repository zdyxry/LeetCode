package main 

import (
	"strconv"
	"fmt"
	"strings"
)

func splitIntoFibonacci(S string) []int {
	if len(S) < 3 {
		return []int{}
	}
	res, isComplete := []int{}, false
	for firstEnd := 0; firstEnd < len(S)/2; firstEnd++ {
		if S[0] == '0' && firstEnd > 0 {
			break
		}
		first, _ := strconv.Atoi(S[:firstEnd+1])
		if first >= 1 << 31 {
			break
		}
		for secondEnd := firstEnd +1; max(firstEnd, secondEnd-firstEnd) <= len(S)-secondEnd;secondEnd++ {
			if S[firstEnd+1] == '0' && secondEnd-firstEnd >1 {
				break
			}
			second, _ := strconv.Atoi(S[firstEnd+1:secondEnd+1])
			if second >= 1 << 31 {
				break
			}
			findRecursiveCheck(S, first, second, secondEnd+1, &res, &isComplete)
		}
	}

	return res
}

func findRecursiveCheck(S string, x1 int, x2 int, left int, res *[]int, isComplete *bool ) {
	if x1 >= 1 << 31 || x2 >= 1<<31 {
		return 
	}
	if left == len(S) {
		if !*isComplete {
			*isComplete = true
			*res = append(*res, x1)
			*res = append(*res, x2)
		}
		return
	}
	if strings.HasPrefix(S[left:], strconv.Itoa(x1+x2)) && !*isComplete {
		*res = append(*res, x1)
		findRecursiveCheck(S, x2, x1+x2, left+len(strconv.Itoa(x1+x2)), res, isComplete)
		return
	}

	if len(*res) > 0 && !*isComplete {
		*res = (*res)[:len(*res)-1]
	}
	return
}

func max(a,b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	S := "123456579"
	res := splitIntoFibonacci(S)
	fmt.Println(res)
}