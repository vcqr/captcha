package main

import (
	"fmt"
	"github.com/vcqr/captcha"
	"image/png"
	"log"
	"net/http"
)

func generateCaptchaHandler(w http.ResponseWriter, r *http.Request) {
	cp := captcha.NewCaptcha(120, 40, 4)
	cp.SetFontPath("bin/")
	cp.SetFontName("free")
	cp.SetMode(1) // 设置为数学公式
	code, img := cp.OutPut()
	//备注：code 可以根据情况存储到session，并在使用时取出验证

	fmt.Println(code)

	w.Header().Set("Content-Type", "image/png; charset=utf-8")

	png.Encode(w, img)
}

func main() {
	http.HandleFunc("/captcha", generateCaptchaHandler)

	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		log.Fatal(err)
	}
}
