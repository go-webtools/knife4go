package knife4go
import (
	"os"
)

func init() {

	f, err := FS.OpenFile(CTX, "/chunk-2d0da532.360126eb.js", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		panic(err)
	}

	_, err = f.Write([] byte(base64Decoding("KHdpbmRvdy53ZWJwYWNrSnNvbnA9d2luZG93LndlYnBhY2tKc29ucHx8W10pLnB1c2goW1siY2h1bmstMmQwZGE1MzIiXSx7IjZhYjMiOmZ1bmN0aW9uKG4sdCxlKXsidXNlIHN0cmljdCI7ZS5yKHQpO3ZhciB1PShlKCI1NjA5IiksZSgiYjFjNyIpLHtuYW1lOiJPQXV0aDIiLGNvbXBvbmVudHM6e30sZGF0YTpmdW5jdGlvbigpe3JldHVybnthcGk6bnVsbCxzd2FnZ2VySW5zdGFuY2U6bnVsbCxkZWJ1Z1N1cHBvcnQ6ITF9fSxjb21wdXRlZDp7c3dhZ2dlcjpmdW5jdGlvbigpe3JldHVybiB0aGlzLiRzdG9yZS5zdGF0ZS5nbG9iYWxzLnN3YWdnZXJ9fSxtb3VudGVkOmZ1bmN0aW9uKCl7fSxiZWZvcmVDcmVhdGU6ZnVuY3Rpb24oKXt9LGNyZWF0ZWQ6ZnVuY3Rpb24oKXt9LG1ldGhvZHM6e319KSxhPWUoIjI4NzciKSxvPU9iamVjdChhLmEpKHUsKGZ1bmN0aW9uKCl7cmV0dXJuKDAsdGhpcy5fc2VsZi5fYykoImEtcm93IixbdGhpcy5fdigiIOaIkeaYr09BdXRoMuWbnuiwg+mhtemdoiAiKV0pfSksW10sITEsbnVsbCwiZTVhODRjYTgiLG51bGwpO3QuZGVmYXVsdD1vLmV4cG9ydHN9fV0pOw==")))
	if err != nil {
		panic(err)
	}

	err = f.Close()
	if err != nil {
		panic(err)
	}
}