package main

import (
	"image/png"
	"net/http"

	"github.com/afocus/captcha"
)

var cap *captcha.Captcha

func main() {

	cap = captcha.New()
	// 设置字体
	if err := cap.SetFont("comic.ttf");
		err != nil {
		panic(err.Error())
	}


	http.HandleFunc("/r", func(w http.ResponseWriter, r *http.Request) {
		// 返回验证码图像对象以及验证码字符串 后期可以对字符串进行对比 判断验证
		img, str := cap.Create(6, captcha.ALL)
		png.Encode(w, img)
		println(str)
	})

	http.HandleFunc("/c", func(w http.ResponseWriter, r *http.Request) {
		str := r.URL.RawQuery
		img := cap.CreateCustom(str)
		png.Encode(w, img)
	})

	http.ListenAndServe(":8085", nil)

}
