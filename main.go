package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Movie struct {
	ID          int    `json: "id"`
	Title       string `json: "title"`
	Description string `json: "descripton"`
}

var movies = []Movie{}

func CreateMovie(c *gin.Context) {
	var newMovie Movie
	if err := c.ShouldBindJSON(&newMovie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
	}

	newMovie.ID = len(movies) + 1
	movies = append(movies, newMovie)
	c.JSON(200, newMovie)
}

func GetMovies(c *gin.Context) {
	c.JSON(http.StatusOK, movies)
}

func main() {
	r := gin.Default()
	r.POST("/movie", CreateMovie)
	r.GET("/movies", GetMovies)
	r.Run(":8070")
}
