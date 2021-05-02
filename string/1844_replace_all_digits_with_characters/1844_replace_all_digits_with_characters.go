package main 

import "fmt"
func replaceDigits(s string) string {
    ans := []byte{}
    for i := range s{
        if s[i] <= '9'&&s[i] >= '0'{
            ans = append(ans, ((s[i - 1] - 'a') + s[i] - '0') % 26 + 'a' )
        }else{
            ans = append(ans, s[i])
        }
    }
    return string(ans)
}
func main() {
	s := "a1c1e1"
	res := replaceDigits(s)
	fmt.Println(res)
}