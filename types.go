package omdb

//Movie is an IMDB Movie responded by omdb api
type Movie struct {
	Title      string
	Year       string
	Rated      string
	Released   string
	Runtime    string
	Genre      string
	Director   string
	Writer     string
	Actors     string
	Plot       string
	Language   string
	Country    string
	Awards     string
	Poster     string
	Ratings    []Rating
	Metascore  string
	IMDBRating string `json:"imdbRating"`
	IMDBVotes  string `json:"imdbVotes"`
	IMDBID     string `json:"imdbID"`
	Type       string
	DVD        string
	BoxOffice  string
	Production string
	Website    string
}

//Rating rating value for from one review site
type Rating struct {
	Source string
	Value  string
}
