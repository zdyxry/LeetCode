779 K-th Symbol in Grammar

Medium

On the first row, we write a 0. Now in every subsequent row, we look at the previous row and replace each occurrence of 0 with 01, and each occurrence of 1 with 10.

Given row N and index K, return the K-th indexed symbol in row N. (The values of K are 1-indexed.) (1 indexed).

Examples:

```
Input: N = 1, K = 1
Output: 0

Input: N = 2, K = 1
Output: 0

Input: N = 2, K = 2
Output: 1

Input: N = 4, K = 5
Output: 1

Explanation:
row 1: 0
row 2: 01
row 3: 0110
row 4: 01101001
```


Note:

    N will be an integer in the range [1, 30].  
    K will be an integer in the range [1, 2^(N-1)].



## 方法


N==1 时row 只有一个数值 0。0 转换为 01， 1 转换为 10：如果 K 是奇数，那么与 N-1 行的 (K+1)/2 相同；如果 K 是偶数，那么与 N-1 行的 (K+1)/2 数字取反。


```go
func kthGrammar(N int, K int) int {
    if K == 1 {
		return 0
	}
	if K%2 == 0 {
		source := kthGrammar(N-1, K/2)
		if source == 1 {
			return 0
		} else {
			return 1
		}
	} else {
		source := kthGrammar(N-1, (K+1)/2)
		if source == 1 {
			return 1
		} else {
			return 0
		}
	}
}
```




row 1: 0
row 2: 01
row 3: 0110
row 4: 01101001
row 5: 0110100110010110

每一行的前半部分都和上一行一致，后半部分为上一行数值取反，递归求解，比较 K 与 2^(N-2) ，如果 K <=2^(N-2) ，表示 K 在N 行的前半部分，那么此时 N 行 K 值与 N-1 行 K 值相同；如果 K > 2^(N-2) ，表示 K 在 N 行的后半部分，此时 N 行 K 值与 N-1 行的 K-(1<<(N-2)) 数值取反相同。



```python
class Solution(object):
    def kthGrammar(self, N, K):
        """
        :type N: int
        :type K: int
        :rtype: int
        """
        if K == 1:
            return 0
        if K <= (1 << (N - 2)):
            return self.kthGrammar(N - 1, K)
        else:
            return 1 - self.kthGrammar(N - 1, K - (1 << (N - 2)))
```