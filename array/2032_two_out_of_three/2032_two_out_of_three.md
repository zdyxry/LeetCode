2032. Two Out of Three


Easy



Given three integer arrays nums1, nums2, and nums3, return a distinct array containing all the values that are present in at least two out of the three arrays. You may return the values in any order.
 

Example 1:

```
Input: nums1 = [1,1,3,2], nums2 = [2,3], nums3 = [3]
Output: [3,2]
Explanation: The values that are present in at least two arrays are:
- 3, in all three arrays.
- 2, in nums1 and nums2.
```

Example 2:

```
Input: nums1 = [3,1], nums2 = [2,3], nums3 = [1,2]
Output: [2,3,1]
Explanation: The values that are present in at least two arrays are:
- 2, in nums2 and nums3.
- 3, in nums1 and nums2.
- 1, in nums1 and nums3.
```

Example 3:

```
Input: nums1 = [1,2,2], nums2 = [4,3,3], nums3 = [5]
Output: []
Explanation: No value is present in at least two arrays.
```

Constraints:

```
1 <= nums1.length, nums2.length, nums3.length <= 100
1 <= nums1[i], nums2[j], nums3[k] <= 100
```


## 方法


```
func twoOutOfThree(nums1 []int, nums2 []int, nums3 []int) []int {
    ans := []int{}
    tmp := make([][]int, 3)
    for i := range tmp { tmp[i] = make([]int, 101) }
    for _, v := range nums1 {
        tmp[0][v] = 1
    }
    for _, v := range nums2 {
        tmp[1][v] = 1
    }
    for _, v := range nums3 {
        tmp[2][v] = 1
    }
    for i := 1; i < 101; i++ {
        if tmp[0][i] + tmp[1][i] + tmp[2][i] >= 2 { ans = append(ans, i) }
    }
    return ans
}
```


```
class Solution:
    def twoOutOfThree(self, nums1: List[int], nums2: List[int], nums3: List[int]) -> List[int]:
        a = [0] * 200
        for i in set(nums1):
            a[i] += 1
        for i in set(nums2):
            a[i] += 1
        for i in set(nums3):
            a[i] += 1

        result = []
        for idx, v in enumerate(a):
            if v >= 2:
                result.append(idx)
        return result


```