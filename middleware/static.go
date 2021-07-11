package middleware

import (
	"embed"
	"net/http"
	"path"
	"strings"
)

//contentType 返回对应资源类型的响应头
func contentType(uPath string) string {
	var ct = "text/html; charset=UTF-8"
	switch path.Ext(uPath) {
	case ".js":
		ct = "text/javascript"
	case ".css":
		ct = "text/css"
	case ".png":
		ct = "image/png"
		// other type ...
	}

	return ct
}

func StaticFileServer(fs embed.FS) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		upath := r.URL.Path
		//设置对应的内容类型
		w.Header().Set("content-type", contentType(upath))

		if strings.HasPrefix(upath, "/") {
			upath = strings.TrimPrefix(upath, "/")
		}
		indexPage := "index.html"
		file, err := fs.Open(upath)
		if err != nil {
			index, _ := fs.ReadFile(indexPage)
			_, _ = w.Write(index)
			return
		}
		_, err = file.Stat()
		if err != nil {
			index, _ := fs.ReadFile(indexPage)
			_, _ = w.Write(index)
			return
		}
		page, _ := fs.ReadFile(upath)
		_, _ = w.Write(page)
	}
}
