package main 

import "fmt"


func numRollsToTarget(d int, f int, target int) int {
    mod := int((1e9)+7)
    dp := make([]int, target+1)
    dp[0] = 1
    for i := 0; i < d; i++ {
        ndp := make([]int, target+1)
        for j := 1; j <= f; j++ {
            for k := 0; k <= target-j; k++ {
                ndp[k+j] = (ndp[k+j]+dp[k])%mod 
            }
        }
        //fmt.Println(dp)
        dp = ndp
    }
    return dp[target]
}


func main() {
	d := 1
	f := 6
	target := 3
	res := numRollsToTarget(d, f, target)
	fmt.Println(res)
}