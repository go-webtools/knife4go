# knife4go
simply Assembled knife4j + gin-swagger, it means an enhanced version of gin-swagger with nice UI.  

work well on go1.18+

# go get
`go get github.com/go-webtools/knife4go`

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
    ...
}
```

## Disclaimer
Public welfare projects.  
The disclaimer asserts that the individual won't be held responsible for any.

## MIT LICENSE
[LICENSE](./LICENSE)

# Links
- gin-swagger: https://github.com/swaggo/gin-swagger
- knife4j: https://github.com/xiaoymin/knife4j
- swag: https://github.com/swaggo/swag/cmd/swag
