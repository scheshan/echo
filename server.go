package echo

import (
	"echo/handle"
	"github.com/gin-gonic/gin"
	"html/template"
	"io/fs"
	"log"
	"net/http"
)

func Run() error {
	engine := gin.Default()

	templateDir, err := fs.Sub(templateFS, "template")
	if err != nil {
		log.Fatal("Cannot find template dir")
	}
	engine.SetHTMLTemplate(template.Must(template.New("test").ParseFS(templateDir, "**")))

	staticDir, err := fs.Sub(staticFS, "static")
	if err != nil {
		log.Fatal("Cannot find static dir")
	}
	engine.StaticFS("static", http.FS(staticDir))

	for _, handler := range handle.Handlers {
		engine.GET(handler.Path(), handler.Handle)
	}

	return engine.Run(":10023")
}
