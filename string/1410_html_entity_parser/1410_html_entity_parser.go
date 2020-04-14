package main

import (
	"bytes"
	"fmt"
)

func entityParser(text string) string {
	textList := []byte(text)
	var res bytes.Buffer
	temp := make([]byte, 0, 7)
	for _, v := range textList {
		switch {
		case v == '&':
			if len(temp) == 0 {
				temp = append(temp, v)
			} else {
				res.Write(temp)
				temp = temp[:1]
			}
		case v == ';':
			{
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
		case len(temp) != 0:
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

func main() {
	text := "&amp; is an HTML entity but &ambassador; is not."
	res := entityParser(text)
	fmt.Println(res)
}
