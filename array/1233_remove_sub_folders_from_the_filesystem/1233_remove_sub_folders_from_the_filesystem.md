1233. Remove Sub-Folders from the Filesystem


Medium


Given a list of folders, remove all sub-folders in those folders and return in any order the folders after removing.

If a folder[i] is located within another folder[j], it is called a sub-folder of it.

The format of a path is one or more concatenated strings of the form: / followed by one or more lowercase English letters. For example, /leetcode and /leetcode/problems are valid paths while an empty string and / are not.

 

Example 1:
```
Input: folder = ["/a","/a/b","/c/d","/c/d/e","/c/f"]
Output: ["/a","/c/d","/c/f"]
Explanation: Folders "/a/b/" is a subfolder of "/a" and "/c/d/e" is inside of folder "/c/d" in our filesystem.
```

Example 2:

```
Input: folder = ["/a","/a/b/c","/a/b/d"]
Output: ["/a"]
Explanation: Folders "/a/b/c" and "/a/b/d/" will be removed because they are subfolders of "/a".
```

Example 3:

```
Input: folder = ["/a/b/c","/a/b/ca","/a/b/d"]
Output: ["/a/b/c","/a/b/ca","/a/b/d"]
```

Constraints:

1 <= folder.length <= 4 * 10^4  
2 <= folder[i].length <= 100  
folder[i] contains only lowercase letters and '/'  
folder[i] always starts with character '/'  
Each folder name is unique.  

## 方法



```go
func removeSubfolders(folder []string) []string {
    sort.Sort(sort.StringSlice(folder))
    out:=[]string{}
    out = append(out,folder[0])
    t:=folder[0]
    for i:=1;i<len(folder);i++{
        if(strings.HasPrefix(folder[i],t+"/")==false){
            out = append(out,folder[i])
            t=folder[i]
        }
    }
    return out
}
```


```python
class Solution:
    def removeSubfolders(self, folder: List[str]) -> List[str]:
        res, t = [], ' '
        for f in sorted(folder):
            print(f)
            if not f.startswith(t):
                res.append(f)
                t = f + '/'
        return res
```