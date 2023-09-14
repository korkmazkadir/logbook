package controller

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/korkmazkadir/logbook"
	"github.com/korkmazkadir/logbook/cryptopuzzle"
)

type APIController struct {
	db logbook.Database
}

func NewAPIController(db logbook.Database) APIController {
	return APIController{db: db}
}

func (cont APIController) GetGinRouter() *gin.Engine {

	router := gin.Default()

	// client side needs to be implemented
	router.POST("/api/book", func(c *gin.Context) {
		var puzzle cryptopuzzle.Puzzle
		err := c.ShouldBindJSON(&puzzle)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		result, err := cryptopuzzle.VerifyPuzzleSolution(puzzle)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if !result {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "could not validated the puzzle"})
			return
		}

		//bookID is encoded to make it URL safe
		bookID := base64.RawURLEncoding.EncodeToString(puzzle.RandomBytes)

		//a book is created using the bookID
		err = cont.db.CreateBook(bookID)
		if err != nil {
			fmt.Printf("error: %s\n", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"url": fmt.Sprintf("/app?book_id=%s", bookID)})
	})

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
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, logs)
	})

	router.GET("/api/puzzle", func(c *gin.Context) {

		puzzle, err := cryptopuzzle.CreatePuzzle(23, time.Minute*2)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, puzzle)
	})

	router.POST("/api/puzzle", func(c *gin.Context) {

		var puzzle cryptopuzzle.Puzzle
		err := c.ShouldBindJSON(&puzzle)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		result, err := cryptopuzzle.VerifyPuzzleSolution(puzzle)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, result)
	})

	return router
}
