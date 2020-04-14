1410. HTML Entity Parser


Medium


HTML entity parser is the parser that takes HTML code as input and replace all the entities of the special characters by the characters itself.

The special characters and their entities for HTML are:

```
Quotation Mark: the entity is &quot; and symbol character is ".
Single Quote Mark: the entity is &apos; and symbol character is '.
Ampersand: the entity is &amp; and symbol character is &.
Greater Than Sign: the entity is &gt; and symbol character is >.
Less Than Sign: the entity is &lt; and symbol character is <.
Slash: the entity is &frasl; and symbol character is /.
Given the input text string to the HTML parser, you have to implement the entity parser.
```

Return the text after replacing the entities by the special characters.

 

Example 1:

```
Input: text = "&amp; is an HTML entity but &ambassador; is not."
Output: "& is an HTML entity but &ambassador; is not."
Explanation: The parser will replace the &amp; entity by &
```

Example 2:
```
Input: text = "and I quote: &quot;...&quot;"
Output: "and I quote: \"...\""
```

Example 3:

```
Input: text = "Stay home! Practice on Leetcode :)"
Output: "Stay home! Practice on Leetcode :)"
```

Example 4:

```
Input: text = "x &gt; y &amp;&amp; x &lt; y is always false"
Output: "x > y && x < y is always false"
```

Example 5:

```
Input: text = "leetcode.com&frasl;problemset&frasl;all"
Output: "leetcode.com/problemset/all"
```
 

Constraints:

1 <= text.length <= 10^5  
The string may contain any possible characters out of all the 256 ASCII characters.

## 方法


```go
func entityParser(text string) string {
    textList := []byte(text)
    var res bytes.Buffer
    temp := make([]byte, 0, 7)
    for _, v := range textList {
        switch {
            case v == '&':
                if len(temp) == 0 {
                    temp = append(temp, v)
                }else{
                    res.Write(temp)
                    temp = temp[:1]
                }
            case v == ';':{
                temp = append(temp, v)
                switch string(temp) {
                    case "&quot;":
                        res.WriteByte('"')
                    case "&apos;":
                        res.WriteByte('\'')
                    case "&amp;":
                        res.WriteByte('&')
                    case "&gt;":
                        res.WriteByte('>')
                    case "&lt;":
                        res.WriteByte('<')
                    case "&frasl;":
                        res.WriteByte('/')
                    default:
                        res.Write(temp)
                }
                temp = temp[:0]
            }               
            case len(temp) != 0 :
                temp = append(temp, v)
            default:
                res.WriteByte(v)
        }
    }
    if len(temp) != 0 {
        res.Write(temp)
    } 
    return res.String()
}
```


```python
class Solution:
    def entityParser(self, text: str) -> str:
        dat={"&quot;":"\"",
            "&apos;":"'",
            "&amp;":"&",
            "&gt;":">",
            "&lt;":"<",
            "&frasl;":"/",
             }
        txt=''
        amp_idx,sem_idx=None,None
        for i,e in enumerate(text):
	    	# if & we update the amp_idx 
            if e=="&":
                amp_idx=i
			# if ; we update sem idx	
            if e==";":
                sem_idx=i
		   # if we don't have any amp_idx yet to be tracked just add the curr char from text
            if amp_idx==None:
                txt+=e
			# if we have amp_idx and sem_idx, means we have a contiguous block to compare in dat dictonary
            if amp_idx!=None  and sem_idx!=None:
                key = text[amp_idx:sem_idx+1] # we get that block to compare from text
				# if key in dat then we add the replacement in txt. e.g: &gt replace with >
                if key in dat.keys():
                    txt+=dat[key]
				# but what if we don't have that key in dat? e.g: &ambassador;. so we just add the full string aka key	
                else:
                    txt+=key
				# assign the idx tracker to None to track next blocks
                amp_idx,sem_idx=None,None

        return txt
```