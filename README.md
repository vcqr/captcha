# Go生成图形验证码

Go生成图形验证码，带噪点，干扰线，正弦干扰线；支持生成普通的字符串和简单的数学算术运算公式；注意：字体需要（.ttf）格式， 另需指定字体所在的目录

## Installation

安装前请确认你的电脑上go环境已安装好，可以使用以下命令

```go
go get -u github.com/vcqr/captcha
```

引用包到你的项目中

```go
import "github.com/vcqr/captcha"
```

## Example

```go
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
	cp.SetFontPath("bin/") //指定字体目录
	cp.SetFontName("free") //指定字体名字
	cp.SetMode(1) //1：设置为简单的数学算术运算公式； 其他为普通字符串
	code, img := cp.OutPut()
	//备注：code 可以根据情况存储到session，并在使用时取出验证

	fmt.Println(code)

	w.Header().Set("Content-Type", "image/png; charset=utf-8")

	png.Encode(w, img) //输出为图片
}

func main() {
	http.HandleFunc("/captcha", generateCaptchaHandler)

	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		log.Fatal(err)
	}
}

```

### 输出图片举例

![exp0](https://github.com/vcqr/captcha/blob/master/example/exp0.png)

![exp1](https://github.com/vcqr/captcha/blob/master/example/exp1.png)
