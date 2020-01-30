package main

import (
	"fmt"
	"github.com/dchest/captcha"
	"io"
	"log"
	"net/http"
)

func showFormHandler(w http.ResponseWriter, r *http.Request) {
	id := captcha.New()

	pat := `
    <html><body>
<hr>
     <img src="./captcha/%s.png" width="200" height="120"/>
<hr>
     <img src="./captcha/%s.png?reload" width="200" height="120"/>
<hr>

<a href="./captcha/%s.png"> view images</a>
<hr>

   </body></html>
    `
	s := fmt.Sprintf(pat, id)
	w.WriteHeader(200)
	w.Write([]byte(s))
}

func processFormHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if !captcha.VerifyString(r.FormValue("captchaId"), r.FormValue("captchaSolution")) {
		io.WriteString(w, "Wrong captcha solution! No robots allowed!\n")
	} else {
		io.WriteString(w, "Great job, human! You solved the captcha.\n")
	}
	io.WriteString(w, "<br><a href='/'>Try another one</a>")
}

func main() {
	http.HandleFunc("/", showFormHandler)
	http.HandleFunc("/process", processFormHandler)
	http.Handle("/captcha/", captcha.Server(captcha.StdWidth, captcha.StdHeight))
	fmt.Println("Server is at localhost:8666")
	if err := http.ListenAndServe("localhost:8666", nil); err != nil {
		log.Fatal(err)
	}
}
