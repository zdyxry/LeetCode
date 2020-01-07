71. Simplify Path


Medium


Given an absolute path for a file (Unix-style), simplify it. Or in other words, convert it to the canonical path.

In a UNIX-style file system, a period . refers to the current directory. Furthermore, a double period .. moves the directory up a level. For more information, see: Absolute path vs relative path in Linux/Unix

Note that the returned canonical path must always begin with a slash /, and there must be only a single slash / between two directory names. The last directory name (if it exists) must not end with a trailing /. Also, the canonical path must be the shortest string representing the absolute path.

 

Example 1:

```
Input: "/home/"
Output: "/home"
Explanation: Note that there is no trailing slash after the last directory name.
```

Example 2:

```
Input: "/../"
Output: "/"
Explanation: Going one level up from the root directory is a no-op, as the root level is the highest level you can go.
```

Example 3:

```
Input: "/home//foo/"
Output: "/home/foo"
Explanation: In the canonical path, multiple consecutive slashes are replaced by a single one.
```

Example 4:

```
Input: "/a/./b/../../c/"
Output: "/c"
```

Example 5:
```
Input: "/a/../../b/../c//.//"
Output: "/c"
```

Example 6:

```
Input: "/a//b////c/d//././/.."
Output: "/a/b/c"
```

## 方法

```go
func simplifyPath(path string) string {
    lp := len(path)
	stack := make([]string, 0, lp/2)
	dir := make([]byte, 0, lp)

	for i := 0; i < lp; i++ {
		// 使用前，清空 dir
		// 这个方法比 dir = []byte{} 快了近 8 倍
		dir = dir[:0]
		// 获取 dir
		for i < lp && path[i] != '/' {
			dir = append(dir, path[i])
			i++
		}

		s := string(dir)

		switch s {
		case ".", "":
			// do nothing
		case "..":
			if len(stack) > 0 {
				// pop
				stack = stack[:len(stack)-1]
			}
		default:
			stack = append(stack, s)
		}
	}

	return "/" + strings.Join(stack, "/")
}
```


```python
class Solution(object):
    def simplifyPath(self, path):
        """
        :type path: str
        :rtype: str
        """
        stack, tokens = [], path.split("/")
        for token in tokens:
            if token == ".." and stack:
                stack.pop()
            elif token != ".." and token != "." and token:
                stack.append(token)
        return "/" + "/".join(stack)
```