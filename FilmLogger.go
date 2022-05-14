package main

import (
	"FilmLogger/packages/models"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var NotFoundError = fmt.Errorf("Resource could not be found")

func main() {
	router := gin.Default()
	router.GET("/films", getFilms)
	router.GET("/films/:title", getFilm)
	router.GET("/films/ratings/:rating", getRatings)
	router.GET("/films/genres/:genre", filmsByGenre)
	router.GET("/films/directors/:director", filmsByDirector)
	router.GET("/films/actors/:actor", filmsByActor)
	router.GET("/films/years/:year", filmsByYear)
	//router.GET("/health-check", healthCheck)
	router.POST("/films", addFilm)
	router.POST("/films/:id", addRatings)
	router.Run("localhost:8080")

	//Mapping errors to Status codes
	router.Use(gin.Logger())
}

func addFilm(c *gin.Context) {
	var film models.Film

	if err := c.BindJSON(&film); err != nil {
		fmt.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
	} else {

		models.AddFilm(film)
		c.IndentedJSON(http.StatusCreated, film)
	}
}

func addRatings(c *gin.Context) {
	var title string
	var rating int64

	if err := c.BindJSON(&rating); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		film, err := models.GetFilm(title)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
		} else {
			models.AddRating(title, rating)
			c.IndentedJSON(http.StatusCreated, film)
		}
	}
}

//Grab all films in the database
func getFilms(c *gin.Context) {
	films, err := models.GetFilms()

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else if films == nil {
		c.String(http.StatusOK, "\n\nNo films found in the database. Add some!\n\n")
	} else {
		c.IndentedJSON(http.StatusOK, films)
	}
}

//Grab one film, search by name
func getFilm(c *gin.Context) {
	title := c.Param("title")

	movie, err := models.GetFilm(title)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	} else if movie.Title == "" {
		c.String(http.StatusOK, "\n\nFilm not found.\n\n")
	} else {
		c.IndentedJSON(http.StatusOK, movie)
	}
}

//Get rating for a particular film
func getRatings(c *gin.Context) {
	strRating := c.Param("rating")
	intRating, err := strconv.Atoi(strRating)
	if err != nil {
		log.Fatal(err)
	}
	films, err := models.GetRatings(intRating)

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	} else if films == nil {
		c.String(http.StatusOK, "\n\nNo films with this rating were found.\n\n")
	} else {
		c.IndentedJSON(http.StatusOK, films)
	}

}

func filmsByGenre(c *gin.Context) {
	genre := c.Param("genre")
	films, err := models.FilmsByGenre(genre)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	} else if films == nil {
		c.String(http.StatusOK, "\n\nNo films found in this genre.\n\n")
	} else {
		for _, film := range films {
			c.IndentedJSON(http.StatusOK, film)
		}
	}
}

func filmsByDirector(c *gin.Context) {
	director := c.Param("director")
	films, err := models.FilmsByDirector(director)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	} else if films == nil {
		c.String(http.StatusOK, "\n\n No films found for this director.\n\n")
	} else {
		c.IndentedJSON(http.StatusOK, films)
	}
}

func filmsByActor(c *gin.Context) {
	actor := c.Param("actor")
	films, err := models.FilmsByActor(actor)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	} else if films == nil {
		c.String(http.StatusOK, "\n\nNo films found for this actorn\n\n.")
	} else {
		c.IndentedJSON(http.StatusOK, films)
	}
}

func filmsByYear(c *gin.Context) {
	year := c.Param("year")
	yearInt, err := strconv.Atoi(year)
	films, err := models.FilmsByYear(yearInt)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	} else if films == nil {
		c.String(http.StatusOK, "\n\nNo films found for this year.\n\n")
	} else {
		c.IndentedJSON(http.StatusOK, films)
	}
}

//Recommend a film based on user input (rating, genre, etc)
//func recommendMe(c *gin.Context) {
//	//IMDB/letterboxd API
//}
//
////Built-in health check to wrap into alerting
//func healthCheck(c *gin.Context) {
//	healthy, err := models.HealthCheck()
//	if err != nil {
//		fmt.Println(err)
//		c.AbortWithStatus(http.StatusNotFound)
//	} else {
//		c.IndentedJSON(http.StatusOK, healthy)
//	}
//}
