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
	router.LoadHTMLGlob("../../web/static/index.html")
	router.Static("/static", "../../web/static")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	log.Println("server is running on 9000")
	router.Run(":9000")
}
