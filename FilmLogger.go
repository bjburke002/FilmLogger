package main

import (
	"FilmLogger/packages/models"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/films", getFilms)
	router.GET("/films/:title", getFilm)
	router.GET("/films/ratings/:rating", getRatings)
	router.GET("/films/genres/:genre", filmsByGenre)
	router.GET("/films/directors/:director", filmsByDirector)
	router.GET("/films/actors/:actor", filmsByActor)
	router.GET("/films/years/:year", filmsByYear)
	router.POST("/films", addFilm)
	router.POST("/films/:id", addRatings)
	router.Run("localhost:8080")
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
	} else {
		c.IndentedJSON(http.StatusOK, films)
	}
}

//Grab one film, search by name
func getFilm(c *gin.Context) {
	title := c.Param("title")

	movie, err := models.GetFilm(title)
	if err != nil {
		fmt.Printf("Film not found: %s", title)
		c.AbortWithStatus(http.StatusNotFound)
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
		fmt.Printf("No films found with %d star rating", intRating)
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, films)
	}

}

func filmsByGenre(c *gin.Context) {
	genre := c.Param("genre")
	films, err := models.FilmsByGenre(genre)
	if err != nil {
		fmt.Printf("No films found in %s genre.", genre)
		c.AbortWithStatus(http.StatusNotFound)
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
		fmt.Printf("No films found for: %s", director)
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, films)
	}
}

func filmsByActor(c *gin.Context) {
	actor := c.Param("actor")
	films, err := models.FilmsByActor(actor)
	if err != nil {
		fmt.Printf("No films found for: %s", actor)
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, films)
	}
}

func filmsByYear(c *gin.Context) {
	year := c.Param("year")
	yearInt, err := strconv.Atoi(year)
	films, err := models.FilmsByYear(yearInt)
	if err != nil {
		fmt.Println(err)
		fmt.Printf("No films found for the year %d", year)
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, films)
	}
}
