692. Top K Frequent Words


Medium


Given a non-empty list of words, return the k most frequent elements.

Your answer should be sorted by frequency from highest to lowest. If two words have the same frequency, then the word with the lower alphabetical order comes first.

Example 1:
```
Input: ["i", "love", "leetcode", "i", "love", "coding"], k = 2
Output: ["i", "love"]
Explanation: "i" and "love" are the two most frequent words.
    Note that "i" comes before "love" due to a lower alphabetical order.
```

Example 2:
```
Input: ["the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"], k = 4
Output: ["the", "is", "sunny", "day"]
Explanation: "the", "is", "sunny" and "day" are the four most frequent words,
    with the number of occurrence being 4, 3, 2 and 1 respectively.
```

Note:
You may assume k is always valid, 1 ≤ k ≤ number of unique elements.

Input words contain only lowercase letters.

Follow up:

Try to solve it in O(n log k) time and O(n) extra space.

## 方法



```go
func topKFrequent(words []string, k int) []string {
    count := make(map[string]int, len(words))
	for _, w := range words {
		count[w]++
	}

	fw := make(freWords, 0, len(count))
	for w, c := range count {
		fw = append(fw, &entry{
			word:      w,
			frequence: c,
		})
	}

	sort.Sort(fw)

	res := make([]string, k)
	for i := 0; i < k; i++ {
		res[i] = fw[i].word
	}

	return res
}

// entry 是 priorityQueue 中的元素
type entry struct {
	word      string
	frequence int
}

// PQ implements heap.Interface and holds entries.
type freWords []*entry

func (f freWords) Len() int { return len(f) }

func (f freWords) Less(i, j int) bool {
	if f[i].frequence == f[j].frequence {
		return f[i].word < f[j].word
	}
	return f[i].frequence > f[j].frequence
}

func (f freWords) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}
```


```python
class Solution(object):
    def topKFrequent(self, words, k):
        """
        :type words: List[str]
        :type k: int
        :rtype: List[str]
        """
        u={}
        for i in range(len(words)):
            if words[i] not in u:
                u[words[i]]=1
            else:
                u[words[i]]+=1
        u_count={}
        heap=set()
        for key,val in u.items():
            if val not in u_count:
                u_count[val]=[]
            u_count[val].append(key)
            heap.add(val*-1)
        heap=list(heap)
        
        heapq.heapify(heap)
        res=[]
        idx=0
        while idx<k:
            vals=u_count[heapq.heappop(heap)*-1]
            heapq.heapify(heap)
            vals.sort()
            for i in range(len(vals)):
                if len(res)==k:
                    return res
                res.append(vals[i])
                idx+=1
        return res
```