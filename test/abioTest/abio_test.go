package abioTest

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func Test_abio_txt(t *testing.T) {
	//w := new(io.WriteCloser)
	f, err := os.Create("/home/whr/abc.txt")
	if err != nil {
		log.Println(err)
		return
	}
	f.WriteString("Hello world!")
	f.Close()
}

func Test_ioutil_txt(t *testing.T) {
	err := ioutil.WriteFile("/home/whr/abc1.txt", []byte("good me"), 777)
	if err != nil {
		log.Println(err)
	}
}

func Test_buffio_write(t *testing.T) {
	f, err := os.Create("/home/whr/aaa.txt")
	if err != nil {
		fmt.Println("err", err)
		return
	}

	// abc
	w := new(bytes.Buffer)
	w.WriteString("good")
	_, _ = f.Write(w.Bytes())
	_ = f.Close()

}

func Test_bufio_write_2(t *testing.T) {
	f, err := os.Create("/home/whr/aaa.txt")
	if err != nil {
		fmt.Println("err", err)
		return
	}
	defer f.Close()

	// abc
	w := bufio.NewWriter(f)
	for i := 0; i < 100; i++ {
		w.WriteString(fmt.Sprint("good_", i))
	}
	_ = w.Flush()
	//f.Write(w.Bytes())

}
