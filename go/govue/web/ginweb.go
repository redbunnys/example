package web

import (
	"embed"
	"github.com/gin-gonic/gin"
	"io/fs"
	"net/http"
	"strings"
)

//go:embed govueweb/dist
var staticFs embed.FS
var stafs fs.FS

func init() {
	stafs, _ = fs.Sub(staticFs, "govueweb/dist")
}

func RunWeb() {
	r := gin.New()
	r.Use(HandlerStaticFiles(stafs))
	r.Run(":9091")
}

func HandlerStaticFiles(files fs.FS) gin.HandlerFunc {
	fileServer := http.FileServer(http.FS(files))
	return func(c *gin.Context) {
		staticFile := isStaticFile(http.FS(files), c.Request.URL.Path, true)
		if staticFile {
			fileServer.ServeHTTP(c.Writer, c.Request)
			c.Abort()
			return
		}
		c.Next()
	}
}

func isStaticFile(fs http.FileSystem, name string, redirect bool) (isFile bool) {
	const indexPage = "/index.html"
	if strings.HasSuffix(name, indexPage) {
		return true
	}
	f, err := fs.Open(name)
	if err != nil {
		return false
	}
	defer f.Close()
	_, err = f.Stat()
	return err == nil
}
