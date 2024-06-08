package web

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:embed govueweb/dist
var HttpstaticFs embed.FS

func RunHttpWeb() {
	// 使用Sub方法获取govueweb/dist目录的子文件系统
	govueWebFs, err := fs.Sub(staticFs, "govueweb/dist")
	if err != nil {
		// 处理错误
		panic(err)
	}

	// 创建一个http.FileServer，它使用govueWebFs作为文件系统
	fs := http.FileServer(http.FS(govueWebFs))

	// 定义一个http.Handler，它将所有请求重定向到根目录
	http.Handle("/", fs)
	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("api"))
	})

	// 启动HTTP服务器
	http.ListenAndServe(":9091", nil)
}
