package main 
import "fmt"
func numDifferentIntegers(word string) int {
    all := make(map[string]bool)
    
    for i := 0; i < len(word); i++ {
        c := word[i]
        // ignore leading zeros
        if c == '0' && i + 1 < len(word) && isDigit(word[i + 1]) {
            continue
        }
        if !isDigit(c) {
            continue
        }
        // when first digit encountered, let's collect the slice to the map
        start := i
        i++
        for i < len(word) && isDigit(word[i]) {
            i++
        }
        all[word[start:i]] = true
    }
    
    return len(all)
}

func isDigit(c byte) bool {    
    return '0' <= c && c <= '9';
}

func main() {
	word := "a123bc34d8ef34"
	res := numDifferentIntegers(word)
	fmt.Println(res)
}