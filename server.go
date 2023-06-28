package echo

import (
	"echo/db"
	"echo/handle"
	"github.com/foolin/goview"
	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path/filepath"
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

	engine.HTMLRender = initGoView()

	staticDir, err := fs.Sub(staticFS, "static")
	if err != nil {
		log.Fatal("Cannot find static dir")
	}
	engine.StaticFS("static", http.FS(staticDir))

	handle.Init(engine)
	return engine
}

func initGoView() *ginview.ViewEngine {
	gv := ginview.Default()
	gv.ViewEngine.SetFileHandler(func(config goview.Config, tplFile string) (content string, err error) {
		path := filepath.Join(config.Root, tplFile)
		bytes, err := viewsFS.ReadFile(path + config.Extension)
		if err != nil {
			return "", err
		}
		content = string(bytes)
		return content, nil
	})

	return gv
}
