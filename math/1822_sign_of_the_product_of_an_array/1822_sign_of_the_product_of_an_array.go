package main 

import "fmt"

func arraySign(nums []int) int {
    var res = 1
    for _,v := range nums {
        if v < 0 {
            res = -res
        }else if v == 0{
            return 0
        }
    }
    return res
}

func main() {
	nums := []int{-1,-2,-3,-4,3,2,1}
	res := arraySign(nums)
	fmt.Println(res)
}