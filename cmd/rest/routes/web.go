package routes

import (
	"io/fs"
	"net/http"
)

// WebRoutes phục vụ trang backoffice tĩnh.
func WebRoutes(mux *http.ServeMux, static fs.FS) {
	mux.Handle("/", http.FileServerFS(static))
}
