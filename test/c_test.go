package test

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func Test_c(t *testing.T) {
	s := "123456"
	sub := fmt.Sprintf("%c", s[0])
	fmt.Println(s[0])
	fmt.Println(sub)
}

func Test_x_parsetint(t *testing.T) {
	amt, er := strconv.ParseInt("5000", 10, 0)
	fmt.Println(amt)
	fmt.Println(er)

}

func Test_map(t *testing.T) {
	m := make(map[int]string)
	m[1] = "a"
	m[2] = "b"
	m[3] = "c"

	for k, _ := range m {
		fmt.Println(k)
	}

}

func Test_aaa(t *testing.T) {
	s := "z中华axy人back民c共d和e国"
	b := []rune(s)
	fmt.Println(s,"  len:",len(s))
	fmt.Println("中", strings.Index(s, "中"))
	fmt.Println("华", strings.Index(s, "华"))
	fmt.Println("x", strings.Index(s, "x"))

	fmt.Println("-----------------")
	fmt.Println(string(b[3:]))
	fmt.Println(string(b[:2]))
	fmt.Println(string(b[:3]))
	fmt.Println("-----------------")
	fmt.Println(s)
	fmt.Println("index(s,a)", string(b[strings.Index(s, "a"):]))
	fmt.Println("index(s,b)", string(b[strings.Index(s, "b"):]))
	fmt.Println("index(s,c)", string(b[strings.Index(s, "c"):]))
}

func removeHeadNumber(s string) string {
	i := strings.Index(s, "||")
	if i < 0 {
		return s
	}
	return s[i+len("||"):]
}

func Test_remove(t *testing.T) {
	s := "z中华axy人back民c共d和e国"
	for i,v := range s{
		fmt.Println(i,":",string(v),"->",s[i+len(string(v)):])
	}

}
