# knife4go
simply Assembled knife4j + gin-swagger, it means an enhanced version of gin-swagger, Leading  to a beautiful UI.

# go get
`go get -u github.com/go-webtools/knife4go`

## usage
Based on the usage of gin-swagger, add sth as follows:
```go
import (
	knife4goGin "github.com/go-webtools/knife4go/gin"
	knife4goFiles "github.com/go-webtools/knife4go"
    ...
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func InitRouters() *gin.Engine {
    router := gin.Default()
	// swagger base router
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// knife4go: beautify swagger-ui
	router.GET("/knife4go/*any", knife4goGin.WrapHandler(knife4goFiles.Handler))
	// knife4go: config openapi info, modify this port 
	router.GET("/openapi.json", func(c *gin.Context) {
		c.Data(200, "application/json", []byte(`[
			{
				"url": "/swagger/doc.json",
				"swaggerVersion": "2.0"
			}
		]`))
	})
    ...
}
```

## Disclaimer
Public welfare projects.
The disclaimer asserts that the company won't be held responsible for any.

# Links
- gin-swagger: https://github.com/swaggo/gin-swagger
- knife4j: https://github.com/xiaoymin/knife4j
- swag: https://github.com/swaggo/swag/cmd/swag
