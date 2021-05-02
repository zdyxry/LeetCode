package main 
import "fmt"
func sumBase(n int, k int) int {
    res := 0
    for  n > 0{
        res += n % k
        n = n / k
    }
    return res
    
}

func main() {
	n := 34
	k := 6
	res := sumBase(n, k)
	fmt.Println(res)
}