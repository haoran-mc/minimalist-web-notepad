package main

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/haoran-mc/minimalist-web-notepad/logging"
	"github.com/haoran-mc/minimalist-web-notepad/util"
)

// 是否有过这个路径
var pathMap map[string]bool = make(map[string]bool)

func main() {
	r := gin.Default()
	r.SetHTMLTemplate(template.Must(template.ParseGlob("layout.html")))

	r.GET("/*path", func(ctx *gin.Context) {
		p := ctx.Param("path")

		if p == "/" {
			redirect(ctx)
		}

		var rawStr []byte
		if pathMap[p] {
			rawStr = util.ReadFile(p)
		}
		ctx.HTML(http.StatusOK, "layout.html", gin.H{
			"Data": string(rawStr),
		})
	})
	r.POST("/*path", func(ctx *gin.Context) {
		p := ctx.Param("path")
		raw, _ := ctx.GetRawData()
		if len(raw) > 0 {
			util.WriteFile(p, raw)
			pathMap[p] = true
		} else {
			util.DeleteFile(p)
		}
	})

	r.Run(":9521")
}

func redirect(ctx *gin.Context) {
	var randPath string
	for {
		randPath = util.RandStr(4)
		if !pathMap[randPath] {
			break
		}
	}
	ctx.Redirect(http.StatusFound, randPath)
}
