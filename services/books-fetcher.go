package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	m "share-Gutenberg/models"
)

const URL_BASE = "http://gutendex.com/books"

func buildURL(base string, params url.Values) string {
	return fmt.Sprintf("%s?%s", base, params.Encode())
}

func BooksFetcher(params url.Values) (*m.Gutendex, error) {
	// set parameters
	url := URL_BASE

	if len(params) > 0 {
		url = buildURL(url, params)
	}

	//fetching
	response, err := http.Get(url)
	if err != nil {
		return nil, errors.New("something wnet wrong fetching from Gutendex")
	}
	//extracting // using io since response.body only can be read once
	defer response.Body.Close()
	body, err2 := io.ReadAll(response.Body)
	if err2 != nil {
		return nil, errors.New("something went wrong reading gutenbex response")
	}
	//deserialation
	var books m.Gutendex
	err3 := json.Unmarshal(body, &books)

	if err3 != nil {
		return nil, errors.New("something whent wrong deserializating gutenbex response")
	}
	return &books, nil
}

func BookFetcher(id string) (*m.Book, *m.Err) {
	response, err := http.Get(fmt.Sprintf("%v/%v", URL_BASE, id))
	if err != nil {
		return nil, &m.Err{Error: err}
	}
	body, err2 := io.ReadAll(response.Body)
	defer response.Body.Close()
	if err2 != nil {
		return nil, &m.Err{Error: err2}
	}
	var book m.Book
	err3 := json.Unmarshal(body, &book)
	if err3 != nil {
		return nil, &m.Err{Error: err3}
	}
	if book.Id == 0 && book.Title == "" {
		return nil, &m.Err{Error: errors.New("not found"), Message: "there is not result for your request", Status: 404}
	}
	return &book, nil
}
