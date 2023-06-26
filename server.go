package echo

import (
	"echo/db"
	"echo/handle"
	"github.com/gin-gonic/gin"
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"os"
)

func Run() error {
	initDb()

	engine := initEngine()
	return engine.Run(":10023")
}

func initDb() {
	err := db.Init(os.Getenv("connStr"))
	if err != nil {
		log.Fatal(err)
	}
}

func initEngine() *gin.Engine {
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
	return engine
}
