// Package web chứa trang admin backoffice (HTML tĩnh), được embed vào binary backoffice.
package web

import "embed"

//go:embed index.html
var FS embed.FS
