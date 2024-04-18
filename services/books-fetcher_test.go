package services

import (
	"fmt"
	"net/url"
	"share-Gutenberg/models"
	"slices"
	"strconv"
	"strings"
	"testing"
)

func checkResultValues(expected []string, value string) bool {
	return slices.ContainsFunc(expected, func(s string) bool {
		return strings.EqualFold(s, value)
	})
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

	searchTest, err := BooksFetcher(defaultParameters)
	if err != nil {
		t.Error("error fetching with search param")
	}
	t.Log("Total books: ", searchTest.Count)
	var c = make(chan bool, len(searchTest.Results))
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

func TestLanguageQuery(t *testing.T) {
	t.Log("testing function BooksFetcher with languages query")
	defaultParameters := url.Values{}
	languages := "es"
	defaultParameters.Add("languages", languages)
	testLanguages, err := BooksFetcher(defaultParameters)
	if err != nil {
		t.Error("error fetching with languages query")
	}
	if testLanguages.Previous != "" {
		t.Error("error, previous page its diferent from \"\" and it is the first page")
	}
	t.Log("this is first page, testing Previus value is default value \"\" ")

	var c = make(chan string)
	for _, book := range testLanguages.Results {
		go func(c chan string, book models.Book) {
			if slices.Contains(book.Languages, languages) {
				c <- "ok"
			} else {
				t.Log(book.Languages, "run bitch run")
				c <- ""
			}
		}(c, book)
	}
	for range testLanguages.Results {
		ok := <-c
		if ok != "ok" {
			t.Error("there is some books that does not include the languages given as queery")
			return
		}
	}
	t.Log("all books in the first page includes the given language")

	defaultParameters.Add("page", "2")
	defaultParameters.Set("languages", "en")
	languages = "en"
	t.Log(defaultParameters)
	testLanguages2, err := BooksFetcher(defaultParameters)
	if err != nil {
		t.Error("error fetching with languages query on page 2")
	}
	if testLanguages2.Previous == "" {
		t.Error("error, previous page should not be  \"\", it has been used page=2 as query")
	}
	t.Log("this is second page, testing Previus exist ", testLanguages2.Previous)

	var c2 = make(chan bool, len(testLanguages2.Results))
	for _, book2 := range testLanguages2.Results {
		go func(c chan bool, book models.Book) {
			if slices.Contains(book.Languages, languages) {
				c <- true
			} else {
				c <- false
			}
		}(c2, book2)
	}

	for range testLanguages2.Results {
		ok := <-c2
		if !ok {
			t.Error("there is some books that does not include the languages given as queery in the second page")
			return
		}
	}
	t.Log("all books in the second page includes the given language")
}

// as soon as the function detects an error finish and shows it
func TestTopicQuery(t *testing.T) {
	expected := "fiction"
	defaultParameters := url.Values{}
	defaultParameters.Set("topic", expected)

	testTopic, err := BooksFetcher(defaultParameters)
	if err != nil {
		t.Error("error fetching books with topic query")
	}
	t.Log("books with the given topics to be evaluated: ", len(testTopic.Results))
	var finish = make(chan bool, len(testTopic.Results))
	defer close(finish)
	for _, book := range testTopic.Results {

		go func(f chan bool, book models.Book) {
			var testSlice []string
			testSlice = append(testSlice, strings.Split(strings.Join(book.BookShelves, " "), " ")...)
			//format key words from books
			testSlice = append(testSlice, strings.Split(strings.Join(book.Subjects, " "), " ")...)
			testSlice = slices.CompactFunc(testSlice, strings.EqualFold)
			f <- checkResultValues(testSlice, expected)
		}(finish, book)
	}
	for range testTopic.Results {
		ok := <-finish
		if !ok {
			t.Error("there were topic words expected and not found")
			return
		}
	}
}

func TestFetchBook(t *testing.T) {
	const expected int = 84
	test, err := BookFetcher(strconv.Itoa(expected))
	if err != nil {
		t.Error("error fetching book")
	}
	if test.Id != expected {
		t.Error("did not get the book with the given id")
	}
	t.Log("invalid id test")
	const invalid string = "invalid"
	_, err2 := BookFetcher(invalid)
	if err2.Status != 404 {
		t.Error("there were not error with the given invalid parameter")
	}
}

/*topic

Use this to search for a case-insensitive key-phrase in books' bookshelves or subjects. For example, /books?topic=children gives books on the "Children's Literature" bookshelf, with the subject "Sick children -- Fiction", and so on. */
