package handlers

import (
	"Practice7/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var categorys = []models.Category{
	{Id: 1, Name: "Poema"},
	{Id: 2, Name: "Drama"},
	{Id: 3, Name: "Povest"},
}

func GetCategories(c *gin.Context) {
	c.JSON(http.StatusOK, categorys)
}

func PostCategories(c *gin.Context) {
	var newCategory models.Category
	if err := c.ShouldBindJSON(&newCategory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	newCategory.Id = len(categorys) + 1
	categorys = append(categorys, newCategory)
	c.JSON(http.StatusCreated, newCategory)
}

func GetByIdCategories(c *gin.Context) {
	idParam := c.Param("id")
	idC, err := strconv.Atoi(idParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid Author id"})
		return
	}

	var oneCategories models.Category
	found := false
	mux.Lock()
	for _, cate := range categorys {
		if cate.Id == idC {
			oneCategories = cate
			found = true
			break
		}
	}
	mux.Unlock()

	if !found {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Category not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"This category": oneCategories})

}

func UpdateCategories(c *gin.Context) {
	idParam := c.Param("id")
	idA, err := strconv.Atoi(idParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid Author id"})
		return
	}

	var updateCategories models.Category
	var OldCategories models.Category
	found := -1

	if err := c.ShouldBindJSON(&updateCategories); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	mux.Lock()
	for i, cate := range categorys {
		if cate.Id == idA {
			OldCategories = cate
			updateCategories.Id = categorys[i].Id
			categorys[i] = updateCategories
			found = i
			break
		}
	}
	mux.Unlock()

	if found == -1 {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Category not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"OldCategory": OldCategories,
		"UpdateCategory": categorys[found],
		"Message":        "Category updated Successfully"})
}

func DeleteCategory(c *gin.Context) {
	idParam := c.Param("id")
	idС, err := strconv.Atoi(idParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid Category id"})
		return
	}
	var deleteCategory []models.Category
	found := false

	mux.Lock()
	for _, cate := range categorys {
		if cate.Id == idС {
			found = true
			continue
		}
		deleteCategory = append(deleteCategory, cate)
	}
	mux.Unlock()

	if !found {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Category not found"})
		return
	}
	categorys = deleteCategory

	c.JSON(http.StatusOK, gin.H{"Message": "Category deleted successfully"})

}
