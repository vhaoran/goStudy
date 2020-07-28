package hashcase

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math"
	"strings"
	"testing"
	"time"
)

type TestHash struct {
	A int    `json:"a,omitempty"`
	B string `json:"x_b,omitempty"`
}

func Test_unmarshal(t *testing.T) {
	pat := `{"a":%d,"b":"test"}`
	t0 := time.Now()
	for i := 0; i < 1000; i++ {
		s := fmt.Sprintf(pat, i)

		b := new(TestHash)
		//
		if err := json.Unmarshal([]byte(s), b); err != nil {
			fmt.Println(err)
		}
	}

	fmt.Println(time.Since(t0))
}

func calcHash(s string) string {
	var h = sha1.New()
	//--------a -----------------------------
	io.WriteString(h, s)
	return fmt.Sprintf("%X", string(h.Sum(nil)))
}

func Test_hash_speed(t *testing.T) {
	sub := `{"a":%d,"b":"test"}`
	pat := ""
	for i := 0; i < 1000; i++ {
		pat += fmt.Sprintf(sub, i)
	}

	a := ""
	t0 := time.Now()
	for i := 0; i < 1000; i++ {
		s := fmt.Sprint(pat, i)
		p := calcHash(s)
		a = p
	}
	fmt.Println(time.Since(t0), a)
}

func Test_hash(t *testing.T) {
	h := sha1.New()
	//--------a -----------------------------
	io.WriteString(h, fmt.Sprint(1))
	h.Reset()
	fmt.Println(fmt.Sprintf("%X", string(h.Sum(nil))))
	log.Println(h.Size())

	fmt.Println("------b-----------")
	h = sha1.New()
	io.WriteString(h, fmt.Sprint(1))
	fmt.Println(fmt.Sprintf("%X", string(h.Sum(nil))))
	log.Println(h.Size())

	fmt.Println("------c-----------")
	h = sha1.New()
	io.WriteString(h, fmt.Sprint(2))
	fmt.Println(fmt.Sprintf("%X", string(h.Sum(nil))))
	log.Println(h.Size())

	fmt.Println("-----------------")
	log.Println(h.Size())
}

func Test_slice_1_n(t *testing.T) {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7}
	b := a[0:(len(a))]
	fmt.Println(b)
}

func Test_slice_revers(t *testing.T) {
	l := []int{0, 1, 2, 3, 4, 5, 6}
	for i := 0; i < len(l)/2; i++ {
		l[i], l[len(l)-1-i] = l[len(l)-1-i], l[i]
	}
	fmt.Println(l)
}
func Test_aaa(t *testing.T) {
	a := [...]int{2: 5, 4: 8}
	fmt.Println("-----------------")
	fmt.Println(a)
}

func Test_split(t *testing.T) {
	s := "中国人民解放军"
	l := strings.Split(s, "")
	for _, v := range l {
		fmt.Println("->", v)
	}

}
func Test_max_int64(t *testing.T) {
	f := math.Pow(2, 64)
	fmt.Println("-----------------")
	fmt.Println(f)
}

func Test_slice_add(t *testing.T) {
	l := []int{1, 2, 3, 4, 5}
	l = append(l, l...)
	fmt.Println("-----------------")
	fmt.Println(l)
}

func Test_slice_vs_map(t *testing.T) {
	h := 5
	l := make([]int, h)
	for i := 0; i < h; i++ {
		l = append(l, i)
	}

	//
	c := 0
	t0 := time.Now()
	for j := 0; j < 10*10000; j++ {
		dst := j % h
		for _, v := range l {
			if v == dst {
				c++
				break
			}
		}
	}
	fmt.Println("-----------------")
	fmt.Println(time.Since(t0).Milliseconds())
}

func Test_slice_vs_map_x(t *testing.T) {
	h := 500
	m := make(map[int]int)
	for i := 0; i < h; i++ {
		m[i] = i
	}

	//
	c := 0
	t0 := time.Now()
	for j := 0; j < 10*10000; j++ {
		dst := j % h
		_, ok := m[dst]
		if ok {
			c++
		}
	}
	fmt.Println("-----------------")
	fmt.Println(time.Since(t0).Milliseconds())
}
