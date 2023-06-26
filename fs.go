package echo

import "embed"

//go:embed static
var staticFS embed.FS

//go:embed template
var templateFS embed.FS
