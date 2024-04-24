package models

import "net/url"

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
	Name        string `json:"name"`
	Email       string `json:"email"`
	Collections []int  `json:"collections"`
	Password    string `json:"password"`
	Role        string `json:"role"`
}
type User struct {
	Id int `json:"id"`
	UserInfo
}

func (u *User) AddCollection(collection Collection) {

} //todo
func (u *User) DeleteCollection(collection Collection) {

} //todo

type Collection struct {
	Id             int    `json:"id"`
	CollectionName string `json:"name"`
	Description    string `json:"description"`
	Documents      []int  `json:"documents"`
	Owner          int    `json:"owner_id"` //only one usser can be owner
	Category       string `json:"category"`
	Public         bool   `json:"public"`
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
