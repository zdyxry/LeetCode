package main 
import "fmt"
func countMatches(items [][]string, ruleKey string, ruleValue string) int {
    keys := make(map[string]int)
    keys["type"] = 0
    keys["color"] = 1
    keys["name"] = 2
    res := 0
    for i := range items {
        if items[i][keys[ruleKey]] == ruleValue {res ++}
    }
    return res
}


func main() {
	items := [][]string{{"phone","blue","pixel"}, {"computer","silver","lenovo"}, {"phone","gold","iphone"}}
	ruleKey := "color"
	ruleValue := "silver"
	res := countMatches(items, ruleKey, ruleValue)
	fmt.Println(res)
}