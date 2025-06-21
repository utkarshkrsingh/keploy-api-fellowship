package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/schema"
	"github.com/utkarshkrsingh/keploy-api-fellowship/basic-api/initializers"
	"github.com/utkarshkrsingh/keploy-api-fellowship/basic-api/models"
	"gorm.io/gorm"
)

func GetList(c *gin.Context) {
	var watchList []models.WatchList
	var decoder = schema.NewDecoder()
	var body struct {
		Name string `schema:"name"`
	}

	if err := decoder.Decode(&body, c.Request.URL.Query()); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request " + err.Error(),
		})

		return
	}

	query := initializers.DB
	if body.Name != "" {
		query = query.Where("name = ?", body.Name)
	}

	if err := query.Find(&watchList).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Database error " + err.Error(),
		})
		return
	}

	if len(watchList) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "No entry found",
		})
		return
	}

	c.JSON(http.StatusOK, watchList)
}

func InsertAnime(c *gin.Context) {
	var watchList models.WatchList
	if err := c.ShouldBindJSON(&watchList); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request payload: " + err.Error(),
		})

		return
	}

	var existing models.WatchList
	if err := initializers.DB.Where("name = ?", watchList.Name).First(&existing).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{
			"error": "Anime already added",
		})

		return
	} else if err != gorm.ErrRecordNotFound {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Database error " + err.Error(),
		})

		return
	}

	if err := initializers.DB.Create(&watchList).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to insert anime",
		})

		return
	}

	c.JSON(http.StatusOK, watchList)
}

func Update(c *gin.Context) {
	var watchList models.WatchList
	if err := c.ShouldBindJSON(&watchList); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request payload: " + err.Error(),
		})

		return
	}

	name := c.Param("name")
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Anime name is required",
		})

		return
	}

	var existing models.WatchList
	if err := initializers.DB.Where("name = ?", name).Find(&existing).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Anime not found",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Database error: " + err.Error(),
			})
		}
		return
	}

	watchList.Name = existing.Name
	if err := initializers.DB.Model(&existing).Updates(watchList).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Database error " + err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, watchList)

}

func Delete(c *gin.Context) {
	name := c.Param("name")
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Name parameter is required",
		})

		return
	}
	var existing models.WatchList
	if err := initializers.DB.Where("name = ?", name).First(&existing).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "ISBN not found",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Database error: " + err.Error(),
			})
		}
		return
	}

	if err := initializers.DB.Delete(&existing).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Database error " + err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, existing)
}
