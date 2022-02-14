package main

func maxDistance(colors []int) int {
	result := 0
	for i := 0; i < len(colors); i++ {
		for j := i; j < len(colors); j++ {
			if colors[i] != colors[j] {
				if result < j-i {
					result = j - i
				}
			}
		}
	}
	return result

}
