1104. Path In Zigzag Labelled Binary Tree


Medium

![1](1104.png)

In an infinite binary tree where every node has two children, the nodes are labelled in row order.

In the odd numbered rows (ie., the first, third, fifth,...), the labelling is left to right, while in the even numbered rows (second, fourth, sixth,...), the labelling is right to left.



Given the label of a node in this tree, return the labels in the path from the root of the tree to the node with that label.

 

Example 1:

```
Input: label = 14
Output: [1,3,4,14]
```

Example 2:

```
Input: label = 26
Output: [1,2,6,10,26]
```

Constraints:

1 <= label <= 10^6


## 方法




```go
func pathInZigZagTree(label int) []int {
	res := []int{}
	for label > 0 {
		res = append([]int{label}, res...)
		l2 := log2(label)
		label = l2 - 1 - (label-l2)/2
	}
	return res
}

func log2(x int) int {
	for i := 1; i < 31; i++ {
		if x < 1<<i {
			return 1 << (i - 1)
		}
	}
	return -1
}
```