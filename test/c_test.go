package test

import (
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"

	"goStudy/lib/g"
)

func Test_c(t *testing.T) {
	s := "123456"
	t0 := time.Now()
	for i := 0; i < 1000; i++ {
		fmt.Println(s[0])
	}

	fmt.Println(time.Since(t0))
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
	fmt.Println(s, "  len:", len(s))
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
	for i, v := range s {
		fmt.Println(i, ":", string(v), "->", s[i+len(string(v)):])
	}

}

func Test_1e9(t *testing.T) {
	fmt.Println("1e9", 1e9/100000)
}

func Test_NilStr(t *testing.T) {
	err := errors.New("ssssssdeeer")

	fmt.Println(g.NilStr(err, "err is nil"))
}
func Test_net_addr(t *testing.T) {
	a, err := net.InterfaceAddrs()
	fmt.Println("-----------------")
	fmt.Println(err)
	for _, address := range a {
		// 检查ip地址判断是否回环地址
		if ip, ok := address.(*net.IPNet); ok && !ip.IP.IsLoopback() {
			if ip.IP.To4() != nil {
				fmt.Println(ip.IP.String())
			}
		}
	}
}

func Test_join(t *testing.T) {
	var l []string
	fmt.Println("------aa-----------")
	fmt.Println(strings.Join(l, ","))
	a := strings.Join(l, ",")
	fmt.Println("-----------------")
	fmt.Println(len(a))

	//

}

type ABC struct {
	A int
	B int
	C string
	L []string
	P *ABC
}

func Test_g_dump(t *testing.T) {
	bean := &ABC{
		A: 1,
		B: 2,
		C: "c",
	}
	fmt.Println("------aaa-----------")
	fmt.Println(g.Dump(bean))
	fmt.Println("------aaa-----------")
	//fmt.Println(g.Dump(*bean))

}

func Test_dump_array(t *testing.T) {
	fmt.Println("------L-----------")
	l := make([]int64, 0)
	l = append(l, 50)
	l = append(l, 51)
	l = append(l, 99)
	fmt.Println(g.Dump(l))
}

func Test_dump_array_struct_ptr(t *testing.T) {
	fmt.Println("------L-----------")
	l := make([]*ABC, 0)
	for i := 0; i < 10; i++ {
		bean := &ABC{
			A: 1,
			B: 2,
			C: "3",
			L: []string{"_", "_", fmt.Sprint(i) + "__"},
			P: &ABC{
				A: i * 10,
				B: i * 20,
				C: "",
				P: nil,
			},
		}
		l = append(l, bean)
	}
	fmt.Println(g.Dump(l))
}
func Test_dump_array_ptr(t *testing.T) {
	fmt.Println("------L-----------")
	l := make([]*int64, 0)
	a := int64(64)
	l = append(l, &a)
	b := int64(65)
	l = append(l, &b)

	fmt.Println(g.Dump(l))
}

func Test_dump_array_str(t *testing.T) {
	fmt.Println("------L-----------")
	l := make([]string, 0)
	l = append(l, "aaa")
	l = append(l, "bbb")

	fmt.Println(g.Dump("l:", l))
}

func Test_dump_nil(t *testing.T) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("HERE")
			fmt.Println(err)
			fmt.Println(0)
		}
	}()

	fmt.Println(2 / 3)
}

func Test_div_zero_recover(t *testing.T) {
	divideByZero()
}

func divideByZero() {
	// Use this deferred function to handle errors.
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("HERE")
			fmt.Println(err)
			fmt.Println(0)
		}
	}()
	// Cause an error.
	// ... Go will run the defer func above.
	cat := 0
	dog := 10 / cat
	fmt.Println(dog)
}

func Test_slice_new(t *testing.T) {
	type A struct {
		ID int64
	}
	l := make([]*A, 0)
	for i := int64(1); i < 5; i++ {
		a := &A{
			ID: i,
		}

		l = append(l, a)
	}

	l1 := l

	l = make([]*A, 0)
	for i := int64(5); i < 10; i++ {
		a := &A{
			ID: i,
		}

		l = append(l, a)
	}

	for _, v := range l {
		fmt.Println(*v)
	}

	fmt.Println("-----------")
	for _, v := range l1 {
		fmt.Println(*v)
	}

}

func Test_byte_n(t *testing.T) {
	a := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	fmt.Println(a[:1])
	fmt.Println(a[:2])
	fmt.Println(a[:3])

}

func Test_env_display(t *testing.T) {
	for _, v := range os.Environ() {
		log.Println(v)
	}
}
