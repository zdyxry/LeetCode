443. String Compression


Easy


Given an array of characters, compress it in-place.

The length after compression must always be smaller than or equal to the original array.

Every element of the array should be a character (not int) of length 1.

After you are done modifying the input array in-place, return the new length of the array.

 
Follow up:

Could you solve it using only O(1) extra space?

 
Example 1:
```
Input:
["a","a","b","b","c","c","c"]

Output:
Return 6, and the first 6 characters of the input array should be: ["a","2","b","2","c","3"]

Explanation:
"aa" is replaced by "a2". "bb" is replaced by "b2". "ccc" is replaced by "c3".
```


Example 2:

```
Input:
["a"]

Output:
Return 1, and the first 1 characters of the input array should be: ["a"]

Explanation:
Nothing is replaced.
```

Example 3:

```
Input:
["a","b","b","b","b","b","b","b","b","b","b","b","b"]

Output:
Return 4, and the first 4 characters of the input array should be: ["a","b","1","2"].

Explanation:
Since the character "a" does not repeat, it is not compressed. "bbbbbbbbbbbb" is replaced by "b12".
Notice each digit has it's own entry in the array.
```


Note:

All characters have an ASCII value in [35, 126].   
1 <= len(chars) <= 1000.

## 方法

```go
func compress(chars []byte) int {
    end, count := 0, 1

	for i, char := range chars {
		if i+1 < len(chars) && char == chars[i+1] {
			// 统计相同字符的个数
			count++
		} else { // 出现不同的字符
			// 在 end 处填写被压缩的字符
			chars[end] = char
			end++

			// 从 end 处填写被压缩字符的个数
			if count > 1 {
				for _, num := range strconv.Itoa(count) {
					chars[end] = byte(num)
					end++
				}
			}

			// 把 count 重置为 1
			count = 1
		}
	}

	return end
}
```


双指针，使用 read 和 write 分别标记读和写的位置，读写操作均从左到右进行：读入连续的一串字符串，然后将压缩版写到数组中，最终 write 将指向输出答案的结尾。
保留指针 anchor，指向当前读到连续字符串的起始位置，从左到右进行读取，当读到最后一个字符或者下一个字符与当前不同时，则到达连续区块的结尾，就从 write 写入压缩的结果。 chars[anchor] 为字符， read - anchor +1 为长度。

```python
class Solution(object):
    def compress(self, chars):
        """
        :type chars: List[str]
        :rtype: int
        """
        anchor = write = 0
        for read, c in enumerate(chars):
            if read + 1 == len(chars) or chars[read + 1] != c:
                chars[write] = chars[anchor]
                write += 1
                if read > anchor:
                    for digit in str(read - anchor + 1):
                        chars[write] = digit
                        write += 1
                anchor = read + 1
        return write

```