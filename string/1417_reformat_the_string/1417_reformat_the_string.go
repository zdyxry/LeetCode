package main

import "fmt"

func reformat(s string) string {
	cm := make([]byte, 0)
	im := make([]byte, 0)
	for _, v := range []byte(s) {
		if v >= '0' && v <= '9' {
			im = append(im, v)
		} else {
			cm = append(cm, v)
		}
	}

	cmLen := len(cm)
	imLen := len(im)

	if !(cmLen-imLen == 0 || cmLen-imLen == 1 || cmLen-imLen == -1) {
		return ""
	}

	minLen := 0
	flag := 0
	result := ""
	if cmLen == imLen {
		minLen = imLen
	} else if cmLen > imLen {
		minLen = imLen
		flag = 1
	} else {
		minLen = cmLen
		flag = 2
	}

	for i := 0; i < minLen; i++ {
		if flag == 1 {
			result += string(cm[i]) + string(im[i])
		} else {
			result += string(im[i]) + string(cm[i])
		}
	}

	if flag == 1 {
		result += string(cm[cmLen-1])
	} else if flag == 2 {
		result += string(im[imLen-1])
	}
	return result

}

func main() {
	s := "a0b1c2"
	res := reformat(s)
	fmt.Println(res)
}
