package restcontroller

import (
	"fmt"
	"movieapi/movie"
	"movieapi/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getMovies(c *gin.Context) {
	c.JSON(http.StatusOK, service.FindAll())
}

func getMovie(c *gin.Context) {
	id := c.Param("id")
	int_id, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Error during conversion")
		c.JSON(http.StatusInternalServerError, "id must be a number")
	} else {
		c.JSON(http.StatusOK, service.FindById(int_id))
	}
}

func postMovie(c *gin.Context) {
	var m movie.Movie
	if err := c.BindJSON(&m); err != nil {
		return
	}
	service.PostMovie(m)
}

func InitEndPoints() {
	router := gin.Default()
	router.GET("/movies", getMovies)
	router.GET("/movies/:id", getMovie)
	router.POST("/movies", postMovie)
	router.Run("localhost:8080")

}
