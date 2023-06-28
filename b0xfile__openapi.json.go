package knife4go
import (
	"os"
)

func init() {

	f, err := FS.OpenFile(CTX, "/openapi.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		panic(err)
	}

	_, err = f.Write([] byte(base64Decoding("WwogICAgewogICAgICAgICJuYW1lIjogIktuaWZlNGdvIiwKICAgICAgICAidXJsIjogImtuaWZlNGdvL2RvYy5qc29uIiwKICAgICAgICAic3dhZ2dlclZlcnNpb24iOiAiMi4wIgogICAgfQpd")))
	if err != nil {
		panic(err)
	}

	err = f.Close()
	if err != nil {
		panic(err)
	}
}