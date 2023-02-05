package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title  string
	Author string
}

type BookSchema struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

func main() {
	db, err := gorm.Open(sqlite.Open("db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Book{})

	r := gin.Default()

	r.GET("/books", func(c *gin.Context) {
		var books []Book
		db.Find(&books)
		c.JSON(200, books)
	})

	r.POST("/books", func(c *gin.Context) {
		var bookSchema BookSchema
		c.BindJSON(&bookSchema)
		book := Book{Title: bookSchema.Title, Author: bookSchema.Author}
		db.Create(&book)
		c.JSON(200, book)
	})

	r.POST("/books/:id", func(c *gin.Context) {
		var bookSchema BookSchema
		c.BindJSON(&bookSchema)
		var book Book
		db.First(&book, c.Param("id"))
		book.Title = bookSchema.Title
		book.Author = bookSchema.Author
		db.Save(&book)
		c.JSON(200, book)
	})

	r.DELETE("/books/:id", func(c *gin.Context) {
		var book Book
		db.First(&book, c.Param("id"))
		db.Delete(&book)
		c.JSON(200, book)
	})

	r.Run()
}
