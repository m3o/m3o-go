package sunnah

import (
	"go.m3o.com/client"
)

type Sunnah interface {
	Books(*BooksRequest) (*BooksResponse, error)
	Chapters(*ChaptersRequest) (*ChaptersResponse, error)
	Collections(*CollectionsRequest) (*CollectionsResponse, error)
	Hadiths(*HadithsRequest) (*HadithsResponse, error)
}

func NewSunnahService(token string) *SunnahService {
	return &SunnahService{
		client: client.NewClient(&client.Options{
			Token: token,
		}),
	}
}

type SunnahService struct {
	client *client.Client
}

// Get a list of books from within a collection. A book can contain many chapters
// each with its own hadiths.
func (t *SunnahService) Books(request *BooksRequest) (*BooksResponse, error) {

	rsp := &BooksResponse{}
	return rsp, t.client.Call("sunnah", "Books", request, rsp)

}

// Get all the chapters of a given book within a collection.
func (t *SunnahService) Chapters(request *ChaptersRequest) (*ChaptersResponse, error) {

	rsp := &ChaptersResponse{}
	return rsp, t.client.Call("sunnah", "Chapters", request, rsp)

}

// Get a list of available collections. A collection is
// a compilation of hadiths collected and written by an author.
func (t *SunnahService) Collections(request *CollectionsRequest) (*CollectionsResponse, error) {

	rsp := &CollectionsResponse{}
	return rsp, t.client.Call("sunnah", "Collections", request, rsp)

}

// Hadiths returns a list of hadiths and their corresponding text for a
// given book within a collection.
func (t *SunnahService) Hadiths(request *HadithsRequest) (*HadithsResponse, error) {

	rsp := &HadithsResponse{}
	return rsp, t.client.Call("sunnah", "Hadiths", request, rsp)

}

type Book struct {
	// arabic name of the book
	ArabicName string `json:"arabic_name,omitempty"`
	// number of hadiths in the book
	Hadiths int32 `json:"hadiths,omitempty"`
	// number of the book e.g 1
	Id int32 `json:"id,omitempty"`
	// name of the book
	Name string `json:"name,omitempty"`
}

type BooksRequest struct {
	// Name of the collection
	Collection string `json:"collection,omitempty"`
	// Limit the number of books returned
	Limit int32 `json:"limit,omitempty"`
	// The page in the pagination
	Page int32 `json:"page,omitempty"`
}

type BooksResponse struct {
	// A list of books
	Books []Book `json:"books,omitempty"`
	// Name of the collection
	Collection string `json:"collection,omitempty"`
	// The limit specified
	Limit int32 `json:"limit,omitempty"`
	// The page requested
	Page int32 `json:"page,omitempty"`
	// The total overall books
	Total int32 `json:"total,omitempty"`
}

type Chapter struct {
	// arabic title
	ArabicTitle string `json:"arabic_title,omitempty"`
	// the book number
	Book int32 `json:"book,omitempty"`
	// the chapter id e.g 1
	Id int32 `json:"id,omitempty"`
	// the chapter key e.g 1.00
	Key string `json:"key,omitempty"`
	// title of the chapter
	Title string `json:"title,omitempty"`
}

type ChaptersRequest struct {
	// number of the book
	Book int32 `json:"book,omitempty"`
	// name of the collection
	Collection string `json:"collection,omitempty"`
	// Limit the number of chapters returned
	Limit int32 `json:"limit,omitempty"`
	// The page in the pagination
	Page int32 `json:"page,omitempty"`
}

type ChaptersResponse struct {
	// number of the book
	Book int32 `json:"book,omitempty"`
	// The chapters of the book
	Chapters []Chapter `json:"chapters,omitempty"`
	// name of the collection
	Collection string `json:"collection,omitempty"`
	// Limit the number of chapters returned
	Limit int32 `json:"limit,omitempty"`
	// The page in the pagination
	Page int32 `json:"page,omitempty"`
	// Total chapters in the book
	Total int32 `json:"total,omitempty"`
}

type Collection struct {
	// Arabic title if available
	ArabicTitle string `json:"arabic_title,omitempty"`
	// Total hadiths in the collection
	Hadiths int32 `json:"hadiths,omitempty"`
	// Name of the collection e.g bukhari
	Name string `json:"name,omitempty"`
	// An introduction explaining the collection
	Summary string `json:"summary,omitempty"`
	// Title of the collection e.g Sahih al-Bukhari
	Title string `json:"title,omitempty"`
}

type CollectionsRequest struct {
	// Number of collections to limit to
	Limit int32 `json:"limit,omitempty"`
	// The page in the pagination
	Page int32 `json:"page,omitempty"`
}

type CollectionsResponse struct {
	Collections []Collection `json:"collections,omitempty"`
}

type Hadith struct {
	// the arabic chapter title
	ArabicChapterTitle string `json:"arabic_chapter_title,omitempty"`
	// the arabic text
	ArabicText string `json:"arabic_text,omitempty"`
	// the chapter id
	Chapter int32 `json:"chapter,omitempty"`
	// the chapter key
	ChapterKey string `json:"chapter_key,omitempty"`
	// the chapter title
	ChapterTitle string `json:"chapter_title,omitempty"`
	// hadith id
	Id int32 `json:"id,omitempty"`
	// hadith text
	Text string `json:"text,omitempty"`
}

type HadithsRequest struct {
	// number of the book
	Book int32 `json:"book,omitempty"`
	// name of the collection
	Collection string `json:"collection,omitempty"`
	// Limit the number of hadiths
	Limit int32 `json:"limit,omitempty"`
	// The page in the pagination
	Page int32 `json:"page,omitempty"`
}

type HadithsResponse struct {
	// number of the book
	Book int32 `json:"book,omitempty"`
	// name of the collection
	Collection string `json:"collection,omitempty"`
	// The hadiths of the book
	Hadiths []Hadith `json:"hadiths,omitempty"`
	// Limit the number of hadiths returned
	Limit int32 `json:"limit,omitempty"`
	// The page in the pagination
	Page int32 `json:"page,omitempty"`
	// Total hadiths in the  book
	Total int32 `json:"total,omitempty"`
}
