package handlers

import (
	"Practice7/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"sync"
)

var books = []models.Book{
	{Id: 1, Title: "Gods history", AuthorId: 1, CategoryId: 1, Price: 10000},
	{Id: 2, Title: "Garry Potter", AuthorId: 2, CategoryId: 2, Price: 8500},
}
var mu sync.Mutex

func GetAllBooks(c *gin.Context) {
	c.JSON(http.StatusOK, books)
}
func AddBooks(c *gin.Context) {
	var newBooks []models.Book
	if err := c.ShouldBindJSON(&newBooks); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	mu.Lock()
	defer mu.Unlock()
	startId := len(books) + 1
	for i := range newBooks {
		newBooks[i].Id = startId + i
		books = append(books, newBooks[i])
	}
	c.JSON(http.StatusCreated, newBooks)
}

func GetIdBook(c *gin.Context) {
	idParam := c.Param("id")
	idB, err := strconv.Atoi(idParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid Book id"})
		return
	}
	var oneBook models.Book
	found := false
	mu.Lock()
	for _, book := range books {
		if book.Id == idB {
			oneBook = book
			found = true
			break
		}
	}
	mu.Unlock()

	if !found {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Book not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"This book": oneBook})
}

func UpdateBooks(c *gin.Context) {
	idParam := c.Param("id")
	idB, err := strconv.Atoi(idParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Book id"})
		return
	}
	var oldBook models.Book
	found := -1
	mu.Lock()
	for i, book := range books {
		if book.Id == idB {
			oldBook = book
			found = i
			break
		}
	}
	mu.Unlock()
	if found == -1 {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Book not found"})
		return
	}
	var updateBook models.Book

	if err := c.ShouldBindJSON(&updateBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	mu.Lock()
	updateBook.Id = books[found].Id
	books[found] = updateBook
	mu.Unlock()
	c.JSON(http.StatusOK, gin.H{"message": "Book updated Successfully!",
		"OldBook":    oldBook,
		"UpdateBook": books[found]})
}

func DeleteBook(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid Book id"})
	}
	var newBooks []models.Book
	found := false
	mu.Lock()
	defer mu.Unlock()
	for _, book := range books {
		if book.Id == id {
			found = true
			continue
		}
		newBooks = append(newBooks, book)
	}
	if !found {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Book not found"})
		return
	}
	books = newBooks

	c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}
