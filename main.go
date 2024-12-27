package main

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

var pathMap map[string]string = make(map[string]string)

func main() {
	r := gin.Default()
	r.SetHTMLTemplate(template.Must(template.ParseGlob("layout.html")))

	r.GET("/*path", func(ctx *gin.Context) {
		p := ctx.Param("path")
		rawStr := pathMap[p]

		ctx.HTML(http.StatusOK, "layout.html", gin.H{
			"Data": rawStr,
		})
	})

	r.POST("/*path", func(ctx *gin.Context) {
		p := ctx.Param("path")
		raw, _ := ctx.GetRawData()
		pathMap[p] = string(raw)
	})

	r.Run(":9521")
}
