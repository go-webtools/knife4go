package knife4go
import (
	"os"
)

func init() {

	f, err := FS.OpenFile(CTX, "/chunk-vendors.4e708c1b.js", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		panic(err)
	}

	if err != nil {
		panic(err)
	}

	err = f.Close()
	if err != nil {
		panic(err)
	}
}