1365. How Many Numbers Are Smaller Than the Current Number


Easy


Given the array nums, for each nums[i] find out how many numbers in the array are smaller than it. That is, for each nums[i] you have to count the number of valid j's such that j != i and nums[j] < nums[i].

Return the answer in an array.

 

Example 1:

```
Input: nums = [8,1,2,2,3]
Output: [4,0,1,1,3]
Explanation: 
For nums[0]=8 there exist four smaller numbers than it (1, 2, 2 and 3). 
For nums[1]=1 does not exist any smaller number than it.
For nums[2]=2 there exist one smaller number than it (1). 
For nums[3]=2 there exist one smaller number than it (1). 
For nums[4]=3 there exist three smaller numbers than it (1, 2 and 2).
```

Example 2:

```
Input: nums = [6,5,4,8]
Output: [2,1,0,3]
```

Example 3:

```
Input: nums = [7,7,7,7]
Output: [0,0,0,0]
```
 

Constraints:

2 <= nums.length <= 500  

0 <= nums[i] <= 100


## 方法

```go
func smallerNumbersThanCurrent(nums []int) []int {
    var sli = []int{}
	count := 0
	length := len(nums)
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			if i == j {
				continue
			} else if nums[i] > nums[j] {
				count++
			}
		}
		sli = append(sli, count)
		count = 0
	}
	return sli
}
```


```python
class Solution(object):
    def smallerNumbersThanCurrent(self, nums):
        """
        :type nums: List[int]
        :rtype: List[int]
        """
        nums2 = sorted(nums)    # 对nums排序
        mapping = {}            # mapping储存num和其序号的键值对
        for i, num in enumerate(nums2):         # 遍历排序后的数组
            if i > 0 and nums2[i] == nums2[i-1]:# 如果某一位跟前一位的值相等，那么它的序号跟前一位一样（并列）
                mapping[nums2[i]] = mapping[nums2[i-1]]
            else:                               # 如果某一位跟前一位的值不等，那么以index为序号
                mapping[nums2[i]] = i
        res = []
        for num in nums:
            res.append(mapping[num])            # res用来储存返回值
        return res

```


```python
class Solution(object):
    def smallerNumbersThanCurrent(self, nums):
        """
        :type nums: List[int]
        :rtype: List[int]
        """
        from collections import Counter
        num_occ_dict, accumlation = Counter(nums), 0
        print num_occ_dict
        for unique_num in sorted(num_occ_dict):
            
            # update histogram of the number of smaller numbers
            num_occ_dict[unique_num], accumlation = accumlation, num_occ_dict[unique_num]+accumlation
        
        print num_occ_dict
        # output by list comprehension
        return [ num_occ_dict[number] for number in nums]
```