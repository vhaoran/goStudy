package imageValidate

import (
	"fmt"
	"log"
	"os"
	"testing"

	cap "github.com/dchest/captcha"
)

func Test_t222(t *testing.T) {

}

func Test_write_image_image_validate(t *testing.T) {
	for i := 0; i < 10; i++ {
		//bt := make([]byte, 5000)
		//buffer := bytes.NewBuffer(bt)
		d := cap.RandomDigits(cap.DefaultLen)
		fmt.Println("------", string(d), "-----------")
		id := fmt.Sprint("www", i)
		w := cap.NewImage(id, d, cap.StdWidth, cap.StdHeight)
		//n, err := w.WriteTo(buffer)
		//if err != nil {
		//	fmt.Println("------", "writeImage", "-----------")
		//	log.Println(err)
		//	return
		//}

		f, err := os.Create(fmt.Sprint("/Users/whr/Documents/abc", id, ".png"))

		if err != nil {
			log.Println(err)
			return
		}
		_ = f.Chmod(777)

		m, err := w.WriteTo(f)
		if err != nil {
			fmt.Println("------", "", "-----------")
			log.Println(err)
		}
		log.Println(m)
		_ = f.Sync()
		_ = f.Close()
	}

}

func Test_captcha_New(t *testing.T) {
	store := NewMyStore()
	cap.SetCustomStore(store)

	id := cap.New()
	log.Println("----------", "id", id, "------------")

	buffer := store.Get(id,false)


	for i := 0; i < 10; i++ {
		a := cap.NewLen(4)
		log.Println("----------", "a", a, "------------")
	}
	cap.NewImage()
}
