package main 

import "fmt"

func busyStudent(startTime []int, endTime []int, queryTime int) int {
    length:=len(startTime)
    res:=0
    for i:=0;i<length;i++{
       if queryTime>=startTime[i] &&  queryTime <=endTime[i] {
           res++
       }
    }
    return res;
}

func main() {
	startTime := []int{1,2,3}
	endTime := []int{3,2,7}
	queryTime := 4
	res := busyStudent(startTime, endTime, queryTime)
	fmt.Println(res)
}