package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/korkmazkadir/logbook"
)

type APIController struct {
	db logbook.Database
}

func NewAPIController(db logbook.Database) APIController {
	return APIController{db: db}
}

func (cont APIController) GetGinRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/api/book/create/:book_id", func(c *gin.Context) {
		bookID := c.Param("book_id")
		err := cont.db.CreateBook(bookID)
		if err != nil {
			fmt.Printf("error: %s\n", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, struct{}{})
	})

	router.POST("/api/:book_id/logs", func(c *gin.Context) {
		bookID := c.Param("book_id")
		var log logbook.Log
		err := c.ShouldBindJSON(&log)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		log, err = cont.db.AppendLog(bookID, log)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, log)
	})

	router.GET("/api/:book_id/logs", func(c *gin.Context) {
		bookID := c.Param("book_id")
		logs, err := cont.db.GetLogs(bookID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, logs)
	})

	return router
}
