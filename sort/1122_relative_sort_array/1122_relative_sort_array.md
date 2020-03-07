1122. Relative Sort Array


Easy


Given two arrays arr1 and arr2, the elements of arr2 are distinct, and all elements in arr2 are also in arr1.

Sort the elements of arr1 such that the relative ordering of items in arr1 are the same as in arr2.  Elements that don't appear in arr2 should be placed at the end of arr1 in ascending order.

 

Example 1:

```
Input: arr1 = [2,3,1,3,2,4,6,7,9,2,19], arr2 = [2,1,4,3,9,6]
Output: [2,2,2,1,4,3,3,9,6,7,19]
```

Constraints:

arr1.length, arr2.length <= 1000

0 <= arr1[i], arr2[i] <= 1000

Each arr2[i] is distinct.

Each arr2[i] is in arr1.


## 方法

```go

func relativeSortArray(arr1 []int, arr2 []int) []int {
    counter := map[int]int{}
	others := []int{}
	dest := make([]int, 0, len(arr2))
	for _, v := range arr2 {
		counter[v] = 0
	}
	for _, v := range arr1 {
		if _, exist := counter[v]; exist {
			counter[v]++
		} else {
			others = append(others, v)
		}
	}
	for _, v := range arr2 {
		for i := 0; i < counter[v]; i++ {
			dest = append(dest, v)
		}
	}
	sort.Ints(others)
	dest = append(dest, others...)
	return dest
}
```

```python
class Solution(object):
    def relativeSortArray(self, arr1, arr2):
        """
        :type arr1: List[int]
        :type arr2: List[int]
        :rtype: List[int]
        """
        c = collections.Counter(arr1)
        res = []       
        for i in arr2:
            res += [i]*c.pop(i)  
            print res
        return res + sorted(c.elements())
```



```python
class Solution(object):
    def relativeSortArray(self, arr1, arr2):
        """
        :type arr1: List[int]
        :type arr2: List[int]
        :rtype: List[int]
        """
        # List all elements from arr1 equal to elements from arr2 checked one by one
        a = [x for i in range(len(arr2)) for x in arr1 if x == arr2[i]]
		
		# List all elements in arr1 not equal to elements from arr2 checked one by one
        b = sorted([x for x in arr1 if x not in arr2])
		
		# Return combined list
        return a + b
```



```python
class Solution(object):
    def relativeSortArray(self, arr1, arr2):
        """
        :type arr1: List[int]
        :type arr2: List[int]
        :rtype: List[int]
        """
        from collections import Counter
        count = Counter(arr1)
        ans = []
        for num in arr2:
            
            cnt = count[num]
            ans += [num]*cnt
            count[num] = 0
        
        ans1 = []
        
        for k, v in count.items():
            ans1 += [k]*v
        
        ans1.sort()
        return ans + ans1
```