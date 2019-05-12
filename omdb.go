package omdb

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	apiEndpoint = "http://www.omdbapi.com/"
)

type movieResponse struct {
	Movie
	Response string
	Error    string
}

type searchResponse struct {
	Search       []Movie
	TotalResults int `json:"totalResults"`
	Response     string
}

//Client is the api client for omdb
type Client struct {
	httpClient *http.Client
	apiKey     string
}

//ByID find IMDB movie by its imdb id
func (c *Client) ByID(id string) (Movie, error) {
	return Movie{}, fmt.Errorf("not implemented")
}

//ByTitle find IMDB movie by matching its title and release year
func (c *Client) ByTitle(title string, year int) (Movie, error) {
	if title == "" {
		return Movie{}, fmt.Errorf("can not find movie with empty title")
	}

	req, err := http.NewRequest("GET", apiEndpoint, nil)
	if err != nil {
		return Movie{}, fmt.Errorf("error when create http request: %v", err)
	}

	req.Header.Add("Accept", "application/json")

	q := req.URL.Query()
	q.Set("apikey", c.apiKey)
	q.Set("t", title)

	if year > 0 {
		q.Set("y", strconv.Itoa(year))
	}

	req.URL.RawQuery = q.Encode()

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Movie{}, fmt.Errorf("failed to send http request: %v", err)
	}

	result := &movieResponse{}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Movie{}, fmt.Errorf("failed to read response body: %v", err)
	}
	err = json.Unmarshal(body, result)
	if err != nil {
		fmt.Printf("%+v\n\n", req)
		fmt.Printf("got response body: %v\n\n", string(body))
		return Movie{}, fmt.Errorf("failed to decode resonse json: %v", err)
	}

	if strings.ToLower(result.Response) == "false" {
		return Movie{}, fmt.Errorf("get error response: %s", result.Error)
	}

	return result.Movie, nil
}

//Search general search for an IMDB movie
func (c *Client) Search(keyword string) ([]Movie, error) {
	return []Movie{}, fmt.Errorf("not implemented")
}

//New create a new api client for omdb
func New(apiKey string) *Client {
	if apiKey == "" {
		panic("can not create omdb client with empty apikey")
	}

	httpClient := &http.Client{Timeout: time.Duration(10) * time.Second}
	return &Client{httpClient, apiKey}
}
