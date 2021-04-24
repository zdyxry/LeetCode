package main 
import "fmt"
func checkIfPangram(sentence string) bool {
    d := make(map[string]bool)
    for _, i := range sentence {
        d[string(i)] = true
    }
    return len(d) == 26
}

func main () {
	sentence := "thequickbrownfoxjumpsoverthelazydog"
	res := checkIfPangram(sentence)
	fmt.Println(res)
}