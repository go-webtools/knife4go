package knife4go
import (
	"os"
)

func init() {

	f, err := FS.OpenFile(CTX, "/chunk-845d989c.20e6d994.css", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		panic(err)
	}

	_, err = f.Write([] byte(base64Decoding("LmFwaS10aXRsZVtkYXRhLXYtNWRmNzgwMjBde21hcmdpbi10b3A6MTBweDttYXJnaW4tYm90dG9tOjVweDtmb250LXNpemU6MTZweDtmb250LXdlaWdodDo2MDA7aGVpZ2h0OjMwcHg7bGluZS1oZWlnaHQ6MzBweDtib3JkZXItbGVmdDo0cHggc29saWQgIzAwYWI2ZDt0ZXh0LWluZGVudDo4cHh9")))
	if err != nil {
		panic(err)
	}

	err = f.Close()
	if err != nil {
		panic(err)
	}
}