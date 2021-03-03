package main 

func canChoose(groups [][]int, nums []int) bool {
	i:=0
	j:=0
	for i<len(groups)&&j<len(nums){
		k:=0
		t:=j
		for k<len(groups[i])&&j<len(nums)&&groups[i][k]==nums[j]{
			k++
			j++
		}
		if k==len(groups[i]){
			i++
		}else {
			j = t+1
		}
	}
	return i==len(groups)
}
