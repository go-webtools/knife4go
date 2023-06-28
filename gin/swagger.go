package knife4go

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"

	"golang.org/x/net/webdav"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/swag"
)

type swaggerConfig struct {
	URL                      string
	DocExpansion             string
	Title                    string
	Oauth2RedirectURL        template.JS
	DefaultModelsExpandDepth int
	DeepLinking              bool
	PersistAuthorization     bool
	Oauth2DefaultClientID    string
}

// Config stores ginSwagger configuration variables.
type Config struct {
	// The url pointing to API definition (normally swagger.json or swagger.yaml). Default is `doc.json`.
	URL                      string
	DocExpansion             string
	InstanceName             string
	Title                    string
	DefaultModelsExpandDepth int
	DeepLinking              bool
	PersistAuthorization     bool
	Oauth2DefaultClientID    string
}

func (config Config) toSwaggerConfig() swaggerConfig {
	return swaggerConfig{
		URL:                      config.URL,
		DeepLinking:              config.DeepLinking,
		DocExpansion:             config.DocExpansion,
		DefaultModelsExpandDepth: config.DefaultModelsExpandDepth,
		Oauth2RedirectURL: "`${window.location.protocol}//${window.location.host}$" +
			"{window.location.pathname.split('/').slice(0, window.location.pathname.split('/').length - 1).join('/')}" +
			"/oauth2-redirect.html`",
		Title:                 config.Title,
		PersistAuthorization:  config.PersistAuthorization,
		Oauth2DefaultClientID: config.Oauth2DefaultClientID,
	}
}

// URL presents the url pointing to API definition (normally swagger.json or swagger.yaml).
func URL(url string) func(*Config) {
	return func(c *Config) {
		c.URL = url
	}
}

// DocExpansion list, full, none.
func DocExpansion(docExpansion string) func(*Config) {
	return func(c *Config) {
		c.DocExpansion = docExpansion
	}
}

// DeepLinking set the swagger deep linking configuration.
func DeepLinking(deepLinking bool) func(*Config) {
	return func(c *Config) {
		c.DeepLinking = deepLinking
	}
}

// DefaultModelsExpandDepth set the default expansion depth for models
// (set to -1 completely hide the models).
func DefaultModelsExpandDepth(depth int) func(*Config) {
	return func(c *Config) {
		c.DefaultModelsExpandDepth = depth
	}
}

// InstanceName set the instance name that was used to generate the swagger documents
// Defaults to swag.Name ("swagger").
func InstanceName(name string) func(*Config) {
	return func(c *Config) {
		c.InstanceName = name
	}
}

// PersistAuthorization Persist authorization information over browser close/refresh.
// Defaults to false.
func PersistAuthorization(persistAuthorization bool) func(*Config) {
	return func(c *Config) {
		c.PersistAuthorization = persistAuthorization
	}
}

// Oauth2DefaultClientID set the default client ID used for OAuth2
func Oauth2DefaultClientID(oauth2DefaultClientID string) func(*Config) {
	return func(c *Config) {
		c.Oauth2DefaultClientID = oauth2DefaultClientID
	}
}

// WrapHandler wraps `http.Handler` into `gin.HandlerFunc`.
func WrapHandler(handler *webdav.Handler, options ...func(*Config)) gin.HandlerFunc {
	var config = Config{
		URL:                      "doc.json",
		DocExpansion:             "list",
		InstanceName:             swag.Name,
		Title:                    "Swagger UI",
		DefaultModelsExpandDepth: 1,
		DeepLinking:              true,
		PersistAuthorization:     false,
		Oauth2DefaultClientID:    "",
	}

	for _, c := range options {
		c(&config)
	}

	return CustomWrapHandler(&config, handler)
}

// CustomWrapHandler wraps `http.Handler` into `gin.HandlerFunc`.
func CustomWrapHandler(config *Config, handler *webdav.Handler) gin.HandlerFunc {
	var once sync.Once

	if config.InstanceName == "" {
		config.InstanceName = swag.Name
	}

	if config.Title == "" {
		config.Title = "Knife4go UI"
	}

	// create a template with name
	index, _ := template.New("swagger_index.html").Parse(swaggerIndexTpl)

	// var matcher = regexp.MustCompile(`(.*)(index\.html|doc\.json|favicon-16x16\.png|favicon-32x32\.png|/oauth2-redirect\.html|swagger-ui\.css|swagger-ui\.css\.map|swagger-ui\.js|swagger-ui\.js\.map|swagger-ui-bundle\.js|swagger-ui-bundle\.js\.map|swagger-ui-standalone-preset\.js|swagger-ui-standalone-preset\.js\.map)[?|.]*`)
	var matcher = regexp.MustCompile(`(.*)(index\.html|doc\.json|\.png|\.gif|\.js|\.css|\.svg)[?|.]*`)

	return func(ctx *gin.Context) {
		if ctx.Request.Method != http.MethodGet {
			ctx.AbortWithStatus(http.StatusMethodNotAllowed)

			return
		}

		// handle doc.json
		if strings.Contains(ctx.Request.RequestURI, "doc.json") {
			doc, err := swag.ReadDoc(config.InstanceName)
			if err != nil {
				ctx.AbortWithStatus(http.StatusInternalServerError)
			}

			ctx.String(http.StatusOK, doc)
			return
		}
		matches := matcher.FindStringSubmatch(ctx.Request.RequestURI)

		if len(matches) != 3 {
			ctx.String(http.StatusNotFound, http.StatusText(http.StatusNotFound))

			return
		}

		path := matches[2]
		once.Do(func() {
			handler.Prefix = matches[1]
		})

		switch filepath.Ext(path) {
		case ".html":
			ctx.Header("Content-Type", "text/html; charset=utf-8")
		case ".css":
			ctx.Header("Content-Type", "text/css; charset=utf-8")
		case ".js":
			ctx.Header("Content-Type", "application/javascript")
		case ".png":
			ctx.Header("Content-Type", "image/png")
		case ".json":
			ctx.Header("Content-Type", "application/json; charset=utf-8")
		}

		log.Println("path: ", path, matches)

		switch path {
		case "index.html":
			_ = index.Execute(ctx.Writer, config.toSwaggerConfig())
		default:
			handler.ServeHTTP(ctx.Writer, ctx.Request)
		}
	}
}

// DisablingWrapHandler turn handler off
// if specified environment variable passed.
func DisablingWrapHandler(handler *webdav.Handler, envName string) gin.HandlerFunc {
	if os.Getenv(envName) != "" {
		return func(c *gin.Context) {
			// Simulate behavior when route unspecified and
			// return 404 HTTP code
			c.String(http.StatusNotFound, "")
		}
	}

	return WrapHandler(handler)
}

// DisablingCustomWrapHandler turn handler off
// if specified environment variable passed.
func DisablingCustomWrapHandler(config *Config, handler *webdav.Handler, envName string) gin.HandlerFunc {
	if os.Getenv(envName) != "" {
		return func(c *gin.Context) {
			// Simulate behavior when route unspecified and
			// return 404 HTTP code
			c.String(http.StatusNotFound, "")
		}
	}

	return CustomWrapHandler(config, handler)
}

const swaggerIndexTpl = `<!-- HTML for static distribution bundle build -->
<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width,initial-scale=1">
    <link rel="icon" href="favicon.ico">
    <title></title>
    <link href="./chunk-0b510f4b.a9ffbfcb.css" rel="prefetch">
    <link href="./chunk-845d989c.20e6d994.css" rel="prefetch">
    <link href="./chunk-0b510f4b.d676e2ca.js" rel="prefetch">
    <link href="./chunk-0d102d5a.e684de54.js" rel="prefetch">
    <link href="./chunk-0fd67716.7c05a478.js" rel="prefetch">
    <link href="./chunk-129cbef4.a2947293.js" rel="prefetch">
    <link href="./chunk-214218f0.e57955ea.js" rel="prefetch">
    <link href="./chunk-2d0af44e.83f5fff8.js" rel="prefetch">
    <link href="./chunk-2d0bd799.f81fc7d8.js" rel="prefetch">
    <link href="./chunk-2d0da532.360126eb.js" rel="prefetch">
    <link href="./chunk-3b888a65.4da69ee0.js" rel="prefetch">
    <link href="./chunk-589faee0.d56965ec.js" rel="prefetch">
    <link href="./chunk-735c675c.d8905482.js" rel="prefetch">
    <link href="./chunk-845d989c.7d4dfe1a.js" rel="prefetch">
    <link href="./chunk-9aa77aca.49424edb.js" rel="prefetch">
    <link href="./chunk-adb9e944.a3f9d0d8.js" rel="prefetch">
    <link href="./chunk-vendors.f24a310a.css" rel="preload" as="style">
    <link href="./main.a93a5ea8.css" rel="preload" as="style">
    <link href="./chunk-vendors.4e708c1b.js" rel="preload" as="script">
    <link href="main.js" rel="preload" as="script">
    <link href="./chunk-vendors.f24a310a.css" rel="stylesheet">
    <link href="./main.a93a5ea8.css" rel="stylesheet">
</head>

<body><noscript><strong>We're sorry but knife4j-vue doesn't work properly without JavaScript enabled. Please enable it
            to continue.</strong></noscript>
    <div id="app"></div>
    <script src="./chunk-vendors.4e708c1b.js"></script>
    <script src="main.js"></script>
</body>

</html>
`
