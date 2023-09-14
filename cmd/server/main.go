package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/korkmazkadir/logbook"
	"github.com/korkmazkadir/logbook/controller"
)

func main() {

	db, err := logbook.NewDBbbolt("./logbook.db")

	if err != nil {
		panic(err)
	}

	apiController := controller.NewAPIController(db)
	router := apiController.GetGinRouter()
	router.LoadHTMLGlob("../../web/static/*.html")
	router.Static("/static", "../../web/static")

	router.GET("/create-new-logbook", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})

	router.GET("/app", func(c *gin.Context) {
		c.HTML(http.StatusOK, "editor.html", nil)
	})

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})

	log.Println("server is running on 9000")
	router.Run(":9000")
}
