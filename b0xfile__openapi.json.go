package knife4go

import (
	"os"
)

func init() {

	f, err := FS.OpenFile(CTX, "/openapi.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		panic(err)
	}

	_, err = f.Write([]byte(base64Decoding("WwogICAgewogICAgICAgICJuYW1lIjogIktuaWZlNGrnpLrkvosiLAogICAgICAgICJ1cmwiOiAiaHR0cDovL2xvY2FsaG9zdDo4MDgwL3N3YWdnZXIvZG9jLmpzb24iLAogICAgICAgICJzd2FnZ2VyVmVyc2lvbiI6ICIyLjAiCiAgICB9Cl0=")))
	if err != nil {
		panic(err)
	}

	err = f.Close()
	if err != nil {
		panic(err)
	}
}
