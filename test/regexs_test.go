package test

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"regexp"
	"strings"
	"testing"
)

func Test_reg_1(t *testing.T) {
	//
	buf := "11.2.3.4.789-hello world"
	str := cutLeadDigit(buf)

	fmt.Println("-src--", buf)
	fmt.Println("-result--", str)
}

func cutLeadDigit(buf string) string {
	//解析正则表达式，如果成功返回解释器
	reg := regexp.MustCompile("^(\\d{1,3}\\.)")
	if reg == nil {
		fmt.Println("regexp err")
		return buf
	}

	//根据规则提取关键信息
	for {
		l := reg.FindAllString(buf, -1)
		if len(l) == 0 {
			break
		}
		//----------------------------------------------
		buf = buf[len(l[0]):]
	}

	return buf
}
func Test_o_disp(t *testing.T) {
	//
	fmt.Println("0360->o:", 0360)

}
func Test_split(t *testing.T) {
	//
	s := "good"
	l := strings.Split(s, ":")
	spew.Dump(l)
}
