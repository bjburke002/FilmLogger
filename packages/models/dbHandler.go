package models

import (
	"database/sql"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

//Add a film to the database
func AddFilm(movie Film) (Film, error) {
	sqlConfig := mysql.Config{
		User:                 os.Getenv("DBUSER"),
		Passwd:               os.Getenv("DBPASS"),
		Net:                  "tcp",
		Addr:                 "192.168.1.18:3306",
		DBName:               "FilmLogger",
		AllowNativePasswords: true,
	}
	db, err := sql.Open("mysql", sqlConfig.FormatDSN()) //Convert sqlConfig to a DSN connection string
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	//Insert record into DB
	filmInsert, err := db.Query(
		"INSERT INTO films (title, actors, director, genre, year, rating, review) VALUES (?,?,?,?,?,?,?)",
		movie.Title, movie.Actors, movie.Director, movie.Genre, movie.Year, movie.Rating, movie.Review)
	if err != nil {
		panic(err.Error())
	}
	defer filmInsert.Close()

	//Verify that the film ahs been added to the database
	var film Film
	row := db.QueryRow("SELECT * FROM films WHERE title = ?", movie.Title)
	if err := row.Scan(&film.ID, &film.Title, &film.Actors, &film.Director, &film.Genre, &film.Year, &film.Rating, &film.Review); err != nil {
		log.Fatal(err)
	}
	return film, nil

}

func AddRating(title string, rating int64) (Film, error) {
	sqlConfig := mysql.Config{
		User:                 os.Getenv("DBUSER"),
		Passwd:               os.Getenv("DBPASS"),
		Net:                  "tcp",
		Addr:                 "192.168.1.18:3306",
		DBName:               "FilmLogger",
		AllowNativePasswords: true,
	}
	db, err := sql.Open("mysql", sqlConfig.FormatDSN()) //Convert sqlConfig to a DSN connection string
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	ratingInsert, err := db.Query("UPDATE films SET rating = ? WHERE title = ?", rating, title)
	if err != nil {
		log.Fatal(err)
	}
	defer ratingInsert.Close()

	var film Film
	row := db.QueryRow("SELECT id, title, rating FROM films WHERE title = ?", title)
	if err := row.Scan(&film.ID, &film.Title, &film.Rating); err != nil {
		log.Fatal(err)
	}
	return film, nil

}

//Returns all films in the database and all information about them
func GetFilms() ([]Film, error) {
	//Used to set up SQL connection string
	sqlConfig := mysql.Config{
		User:                 os.Getenv("DBUSER"),
		Passwd:               os.Getenv("DBPASS"),
		Net:                  "tcp",
		Addr:                 "192.168.1.18:3306",
		DBName:               "FilmLogger",
		AllowNativePasswords: true,
	}
	db, err := sql.Open("mysql", sqlConfig.FormatDSN()) //Convert sqlConfig to a DSN connection string
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM films")
	if err != nil {
		log.Fatal(err)
	}

	var films []Film
	for rows.Next() {
		var film Film
		//For each row returned, scan film information into film. If error present, throw error
		if err := rows.Scan(&film.ID, &film.Actors, &film.Title, &film.Director, &film.Genre, &film.Year, &film.Rating, &film.Review); err != nil {
			log.Fatal(err)
		}
		films = append(films, film)
	}
	return films, nil
}

//Return one film based on search query
func GetFilm(title string) (Film, error) {
	sqlConfig := mysql.Config{
		User:                 os.Getenv("DBUSER"),
		Passwd:               os.Getenv("DBPASS"),
		Net:                  "tcp",
		Addr:                 "192.168.1.18:3306",
		DBName:               "FilmLogger",
		AllowNativePasswords: true,
	}
	db, err := sql.Open("mysql", sqlConfig.FormatDSN()) //Convert sqlConfig to a DSN connection string
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var film Film
	row := db.QueryRow("SELECT id, title, director, genre, year, rating FROM films WHERE title = ?", title)

	if err := row.Scan(&film.ID, &film.Title, &film.Director, &film.Genre, &film.Year, &film.Rating); err != nil {
		log.Fatal(err)
	}
	return film, nil
}

//
func GetRatings(rating int) ([]Film, error) {
	sqlConfig := mysql.Config{
		User:                 os.Getenv("DBUSER"),
		Passwd:               os.Getenv("DBPASS"),
		Net:                  "tcp",
		Addr:                 "192.168.1.18:3306",
		DBName:               "FilmLogger",
		AllowNativePasswords: true,
	}
	db, err := sql.Open("mysql", sqlConfig.FormatDSN()) //Convert sqlConfig to a DSN connection string
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var films []Film
	rows, err := db.Query("SELECT id, title, director, rating FROM films WHERE rating = ?", rating)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var film Film
		if err := rows.Scan(&film.ID, &film.Title, &film.Director, &film.Rating); err != nil {
			log.Fatal(err)
		}
		films = append(films, film)
	}

	return films, nil
}

func FilmsByGenre(genre string) ([]Film, error) {
	sqlConfig := mysql.Config{
		User:                 os.Getenv("DBUSER"),
		Passwd:               os.Getenv("DBPASS"),
		Net:                  "tcp",
		Addr:                 "192.168.1.18:3306",
		DBName:               "FilmLogger",
		AllowNativePasswords: true,
	}
	db, err := sql.Open("mysql", sqlConfig.FormatDSN()) //Convert sqlConfig to a DSN connection string
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var films []Film
	rows, err := db.Query("SELECT * FROM films WHERE genre = ?", genre)
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var film Film
		if err := rows.Scan(&film.ID, &film.Title, &film.Director, &film.Genre, &film.Year, &film.Rating); err != nil {
			log.Fatal(err)
		}
		films = append(films, film)
	}
	return films, nil
}

func FilmsByDirector(director string) ([]Film, error) {
	sqlConfig := mysql.Config{
		User:                 os.Getenv("DBUSER"),
		Passwd:               os.Getenv("DBPASS"),
		Net:                  "tcp",
		Addr:                 "192.168.1.18:3306",
		DBName:               "FilmLogger",
		AllowNativePasswords: true,
	}
	db, err := sql.Open("mysql", sqlConfig.FormatDSN()) //Convert sqlConfig to a DSN connection string
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var films []Film
	rows, err := db.Query("SELECT filmography FROM directors WHERE fullName = ?", director) //Can I do this with director.Filmography instead?
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var film Film
		if err := rows.Scan(&film.ID, &film.Title, &film.Director, &film.Genre, &film.Year, &film.Rating); err != nil {
			log.Fatal(err)
		}
		films = append(films, film)
	}
	return films, nil
}

func FilmsByActor(actor string) ([]Film, error) {
	sqlConfig := mysql.Config{
		User:                 os.Getenv("DBUSER"),
		Passwd:               os.Getenv("DBPASS"),
		Net:                  "tcp",
		Addr:                 "192.168.1.18:3306",
		DBName:               "FilmLogger",
		AllowNativePasswords: true,
	}
	db, err := sql.Open("mysql", sqlConfig.FormatDSN()) //Convert sqlConfig to a DSN connection string
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var films []Film
	rows, err := db.Query("SELECT filmography FROM actors WHERE fullName = ?", actor) //Perhaps use a join on films/actor db? Look into that.
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var film Film
		if err := rows.Scan(&film.ID, &film.Title, &film.Director, &film.Genre, &film.Year, &film.Rating); err != nil {
			log.Fatal(err)
		}
		films = append(films, film)
	}
	return films, nil
}

//List all films for a chosen year
func FilmsByYear(year int) ([]Film, error) {
	sqlConfig := mysql.Config{
		User:                 os.Getenv("DBUSER"),
		Passwd:               os.Getenv("DBPASS"),
		Net:                  "tcp",
		Addr:                 "192.168.1.18:3306",
		DBName:               "FilmLogger",
		AllowNativePasswords: true,
	}
	db, err := sql.Open("mysql", sqlConfig.FormatDSN()) //Convert sqlConfig to a DSN connection string
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var films []Film
	rows, err := db.Query("SELECT * FROM films WHERE year = ?", year)
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var film Film
		if err := rows.Scan(&film.ID, &film.Actors, &film.Title, &film.Director, &film.Genre, &film.Year, &film.Rating, &film.Review); err != nil {
			log.Fatal(err)
		}
		films = append(films, film)
	}
	return films, nil
}
