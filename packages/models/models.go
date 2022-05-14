package models

type Film struct {
	ID       int     `json:"id"`
	Actors   string  `json:"actors"`
	Title    string  `json:"title"`
	Director string  `json:"director"`
	Genre    string  `json:"genre"`
	Year     int     `json:"year"`
	Rating   float64 `json:"rating"`
	Review   string  `json:"review"`
}

type Director struct {
	ID          int      `json:"id"`
	Name        string   `json:"fullName"`
	Filmography []string `json:"filmography"`
}

type Actor struct {
	ID          int      `json:"id"`
	Name        string   `json:"fullName"`
	Filmography []string `json:"filmography"`
}

type errorString struct {
	err string
}
