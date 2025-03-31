package handlers

import (
	"Practice7/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"sync"
)

var authors = []models.Author{
	{Id: 1, Name: "Mukhtar Auezov"},
	{Id: 2, Name: "Nikolay Berdayev"},
}

var mux *sync.Mutex

func GetAuthor(c *gin.Context) {
	c.JSON(http.StatusOK, authors)
}

func PostAuthors(c *gin.Context) {
	var newAuthor models.Author
	if err := c.ShouldBindJSON(&newAuthor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	newAuthor.Id = len(authors) + 1
	authors = append(authors, newAuthor)
	c.JSON(http.StatusCreated, newAuthor)
}

func GetByIdAuthor(c *gin.Context) {
	idParam := c.Param("id")
	idA, err := strconv.Atoi(idParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid Author id"})
		return
	}

	var oneAuthor models.Author
	found := false
	mux.Lock()
	for _, author := range authors {
		if author.Id == idA {
			oneAuthor = author
			found = true
			break
		}
	}
	mux.Unlock()

	if !found {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Author not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"This author": oneAuthor})

}

func UpdateAuthor(c *gin.Context) {
	idParam := c.Param("id")
	idA, err := strconv.Atoi(idParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid Author id"})
		return
	}

	var updateAuthor models.Author
	var OldAuthor models.Author
	found := -1

	if err := c.ShouldBindJSON(&updateAuthor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	mux.Lock()
	for i, author := range authors {
		if author.Id == idA {
			OldAuthor = author
			updateAuthor.Id = authors[i].Id
			authors[i] = updateAuthor
			found = i
			break
		}
	}
	mux.Unlock()

	if found == -1 {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Author not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"OldAuthor": OldAuthor,
		"UpdateAuthor": authors[found],
		"Message":      "Author updated Successfully"})
}

func DeleteAuthor(c *gin.Context) {
	idParam := c.Param("id")
	idA, err := strconv.Atoi(idParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid Author id"})
		return
	}
	var deleteAuthor []models.Author
	found := false

	mux.Lock()
	for _, author := range authors {
		if author.Id == idA {
			found = true
			continue
		}
		deleteAuthor = append(deleteAuthor, author)
	}
	mux.Unlock()

	if !found {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Author not found"})
		return
	}
	authors = deleteAuthor

	c.JSON(http.StatusOK, gin.H{"Message": "Author deleted successfully"})

}
