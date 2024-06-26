package models

import (
	"database/sql/driver"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const (
	Admin = iota
	Regular
)

type Err struct {
	Error   error  `json:"error"`
	Message string `json:"message"`
	Status  int    `json:"status"`
}

type UserInfo struct {
	Name      string    `json:"name" db:"name"`
	Email     string    `json:"email" db:"email"`
	Password  string    `json:"password" db:"password"`
	Role      string    `json:"role" db:"role"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

func (ui UserInfo) Iter() *map[string]string {
	res := make(map[string]string)
	if ui.Name != "" {
		res["name"] = ui.Name
	}
	if ui.Email != "" {
		res["email"] = ui.Email
	}
	if ui.Password != "" {
		res["password"] = ui.Password
	}
	if ui.Role != "" {
		res["role"] = ui.Role
	}
	return &res
}

type User struct {
	Id int `json:"id"`
	//Collections []int `json:"collections"`
	UserInfo
}

func (u *User) AddCollection(collection Collection) {

} //todo
func (u *User) DeleteCollection(collection Collection) {

}

// custom type of []int to implement scanner interface
type Docs []int

// manual implementation of Scanner interface from pq to properly read documents []interger from collection table to []int
func (t *Docs) Scan(v interface{}) error {
	if v == nil {
		*t = Docs{}
		return nil
	}
	bytes, ok := v.([]byte)
	if !ok {
		return fmt.Errorf("Scan is expected to receive a []byte from database, but got [%+v]", v)
	}
	str := string(bytes)
	str = strings.TrimPrefix(str, "{")
	str = strings.TrimSuffix(str, "}")
	if len(str) == 0 {
		return nil
	}
	numStr := strings.Split(str, ",")
	for _, num := range numStr {
		formated, err := strconv.Atoi(num)
		if err != nil {
			return err
		}
		*t = append(*t, formated)
	}
	return nil
}

// scanner interface needed for insert into table Docs type into documents properly
func (t *Docs) Value() (driver.Value, error) {
	s := ""
	for _, num := range *t {
		s += fmt.Sprintf("%v,", num)
	}
	if len(s) > 0 {
		s = s[:len(s)-1] // Elimina el último carácter del string
	}
	return fmt.Sprintf("{%v}", s), nil
}

type CollectionInfo struct {
	CollectionName string    `json:"name" db:"collection_name"`
	Description    string    `json:"description" db:"description"`
	Documents      Docs      `json:"documents" db:"documents"`
	Owner          int       `json:"owner_id" db:"owner_id"` //only one usser can be owner
	Category       string    `json:"category" db:"category"`
	Public         bool      `json:"public" db:"public"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
}

// custom method to be used inside CMI
func (ci *CollectionInfo) Iter() *map[string]string {
	res := make(map[string]string)
	if ci.CollectionName != "" {
		res["collection_name"] = ci.CollectionName
	}
	if ci.Description != "" {
		res["description"] = ci.Description
	}
	if ci.Owner != 0 {
		res["owner_id"] = strconv.Itoa(ci.Owner)
	}
	if ci.Category != "" {
		res["category"] = ci.Category
	}
	return &res
}

type Collection struct {
	Id int `json:"id" db:"id"`
	CollectionInfo
}
type UserCollections struct {
	Owned []Collection `json:"owned"`
	Added []Collection `json:"added"`
}

func (c *Collection) AddBook(bookId int) {
} //todo
func (c *Collection) DeleteBook(bookId int) {
} //todo
type BookAuthor struct {
	Name      string `json:"name"`
	BirthYear int    `json:"birth_year"`
	DeathYear int    `json:"death_year"`
}
type Book struct {
	Id            int               `json:"id"`
	Title         string            `json:"title"`
	Authors       []BookAuthor      `json:"authors"`
	Subjects      []string          `json:"subjects"`
	BookShelves   []string          `json:"bookshelves"`
	Languages     []string          `json:"languages"`
	Copyright     bool              `json:"copyright"`
	MediaType     string            `json:"media_type"`
	Formats       map[string]string `json:"formats"`
	DownloadCount int               `json:"download_count"`
}
type Gutendex struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []Book `json:"results"`
}
type Params = url.Values
type BookFileInfo struct{ Name, Ext, Dir, Format string }

/*

//just reference for now
// type BookFormat struct {
//     HTML     string `json:"text/html"`
//     EPUB     string `json:"application/epub+zip"`
//     Mobi     string `json:"application/x-mobipocket-ebook"`
//     RDF      string `json:"application/rdf+xml"`
//     JPEG     string `json:"image/jpeg"`
//     Octet    string `json:"application/octet-stream"`
//     Plain    string `json:"text/plain; charset=us-ascii"`
// }
response from gutendex
{
	"count": 1,
	"next": null,
  "previous": null,
  "results": [
    {
      "id": 1,
      "title": "The Declaration of Independence of the United States of America",
      "authors": [
        {
          "name": "Jefferson, Thomas",
          "birth_year": 1743,
          "death_year": 1826
        }
      ],
      "translators": [],
      "subjects": [
        "United States -- History -- Revolution, 1775-1783 -- Sources",
        "United States. Declaration of Independence"
      ],
      "bookshelves": [
        "American Revolutionary War",
        "Politics",
        "United States Law"
      ],
      "languages": [
        "en"
      ],
      "copyright": false,
      "media_type": "Text",
      "formats": {
        "text/html": "https://www.gutenberg.org/ebooks/1.html.images",
        "application/epub+zip": "https://www.gutenberg.org/ebooks/1.epub3.images",
        "application/x-mobipocket-ebook": "https://www.gutenberg.org/ebooks/1.kf8.images",
        "application/rdf+xml": "https://www.gutenberg.org/ebooks/1.rdf",
        "image/jpeg": "https://www.gutenberg.org/cache/epub/1/pg1.cover.medium.jpg",
        "application/octet-stream": "https://www.gutenberg.org/cache/epub/1/pg1-h.zip",
        "text/plain; charset=us-ascii": "https://www.gutenberg.org/ebooks/1.txt.utf-8"
      },
      "download_count": 2272
    }
  ]
}

// Deserializar el JSON en una estructura Response
    var response Response
    err := json.Unmarshal([]byte(jsonString), &response)
    if err != nil {
        fmt.Println("Error al decodificar JSON:", err)
        return
    }

    // Imprimir el primer libro de los resultados
    fmt.Printf("%+v\n", response.Results[0])
}
*/
//User
