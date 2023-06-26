package echo

import "embed"

//go:embed static
var staticFS embed.FS

//go:embed views
var viewsFS embed.FS
