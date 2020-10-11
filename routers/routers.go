package routers

import (
	"net/http"

	"github.com/fabiosebastiano/go-web/web"
	"github.com/gin-gonic/gin"
)

func HandleIndex() gin.HandlerFunc {
	return func(c *gin.Context) {
		html := web.MustAsset("index.html")
		c.Data(200, "text/hmtl; charset=UTF-8", html)
	}
}

func Register(e *gin.Engine) {
	h := gin.WrapH(http.FileServer(web.AssetFile()))

	e.GET("/favicon.ico", h)
	e.GET("/js/*filepath", h)
	e.GET("/css/*filepath", h)
	e.GET("/img/*filepath", h)
	e.GET("/fonts/*filepath", h)
	e.NoRoute(HandleIndex())
}
