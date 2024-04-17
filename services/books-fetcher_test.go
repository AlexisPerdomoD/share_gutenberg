package services

import (
	"fmt"
	"net/url"
	"strings"
	"testing"
)

func checkResultValues(expected []string, value string) *[]string {
	for index, expectedValue := range expected {
		if value == expectedValue {
			expected = append(expected[:index], expected[index+1:]...)
		}
	}
	return &expected
}
func TestBooksFetcherGeneral(t *testing.T) {
	t.Log("testing function BooksFetcher")
	defaultParameters := url.Values{}

	_, err := BooksFetcher(defaultParameters)

	if err != nil {
		t.Log("fail no parameters fetching")
		t.Error(err)
	}
	t.Log("done no parameters fetching")

	const ids = "5, 12, 33"
	defaultParameters.Add("ids", ids)
	test, err2 := BooksFetcher(defaultParameters)
	if err2 != nil {
		t.Log("fail fetching with simultaniusly ids")
		t.Error(err2)
	}
	if test.Count != 3 {
		t.Error("the number of results exceeded the number required")
	}

	missing := []int{5, 12, 33}
	for _, book := range test.Results {
		for index, missingNumber := range missing {
			if book.Id == missingNumber {
				missing = append(missing[:index], missing[index+1:]...)
			}

		}
	}

	if len(missing) > 0 {
		t.Error("the ids from the results were not the correct ones")
		return
	}
	defaultParameters.Del("ids")
	t.Log("the function fetch properly addings param ids")

}

func TestSearch(t *testing.T) {
	t.Log("testing function BooksFetcher with search params")
	defaultParameters := url.Values{}
	title, author := "Frankenstein", "Wollstonecraft"
	defaultParameters.Add("search", fmt.Sprintf("%v %v", title, author))

	searchTest, err3 := BooksFetcher(defaultParameters)
	if err3 != nil {
		t.Error("error fetching with search param")
	}
	t.Log("Total books: ", searchTest.Count)
	var c = make(chan bool, searchTest.Count)
	go func(c chan bool) {
		for _, book := range searchTest.Results {
			autorsAndTitle := book.Title
			//collect all authors names
			for _, autor := range book.Authors {
				autorsAndTitle += fmt.Sprintf(" %v ", autor.Name)
			}
			// flag to know wether title or autors names are in the current book
			if strings.Contains(strings.ToLower(autorsAndTitle), strings.ToLower(title)) ||
				strings.Contains(strings.ToLower(autorsAndTitle), strings.ToLower(author)) {
				c <- true
			} else {
				c <- false
			}
		}
		close(c)
	}(c)
	for ok := range c {
		if !ok {
			t.Error("one of the books using title and authors name did not include any of them on those")
			return
		}
	}
	t.Log("test search params ok")
}
