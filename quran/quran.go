package quran

import (
	"go.m3o.com/client"
)

type Quran interface {
	Chapters(*ChaptersRequest) (*ChaptersResponse, error)
	Search(*SearchRequest) (*SearchResponse, error)
	Summary(*SummaryRequest) (*SummaryResponse, error)
	Verses(*VersesRequest) (*VersesResponse, error)
}

func NewQuranService(token string) *QuranService {
	return &QuranService{
		client: client.NewClient(&client.Options{
			Token: token,
		}),
	}
}

type QuranService struct {
	client *client.Client
}

// List the Chapters (surahs) of the Quran
func (t *QuranService) Chapters(request *ChaptersRequest) (*ChaptersResponse, error) {

	rsp := &ChaptersResponse{}
	return rsp, t.client.Call("quran", "Chapters", request, rsp)

}

// Search the Quran for any form of query or questions
func (t *QuranService) Search(request *SearchRequest) (*SearchResponse, error) {

	rsp := &SearchResponse{}
	return rsp, t.client.Call("quran", "Search", request, rsp)

}

// Get a summary for a given chapter (surah)
func (t *QuranService) Summary(request *SummaryRequest) (*SummaryResponse, error) {

	rsp := &SummaryResponse{}
	return rsp, t.client.Call("quran", "Summary", request, rsp)

}

// Lookup the verses (ayahs) for a chapter including
// translation, interpretation and breakdown by individual
// words.
func (t *QuranService) Verses(request *VersesRequest) (*VersesResponse, error) {

	rsp := &VersesResponse{}
	return rsp, t.client.Call("quran", "Verses", request, rsp)

}

type Chapter struct {
	// The arabic name of the chapter
	ArabicName string `json:"arabic_name,omitempty"`
	// The complex name of the chapter
	ComplexName string `json:"complex_name,omitempty"`
	// The id of the chapter as a number e.g 1
	Id int32 `json:"id,omitempty"`
	// The simple name of the chapter
	Name string `json:"name,omitempty"`
	// The pages from and to e.g 1, 1
	Pages []int32 `json:"pages,omitempty"`
	// Should the chapter start with bismillah
	PrefixBismillah bool `json:"prefix_bismillah,omitempty"`
	// The order in which it was revealed
	RevelationOrder int32 `json:"revelation_order,omitempty"`
	// The place of revelation
	RevelationPlace string `json:"revelation_place,omitempty"`
	// The translated name
	TranslatedName string `json:"translated_name,omitempty"`
	// The number of verses in the chapter
	Verses int32 `json:"verses,omitempty"`
}

type ChaptersRequest struct {
	// Specify the language e.g en
	Language string `json:"language,omitempty"`
}

type ChaptersResponse struct {
	Chapters []Chapter `json:"chapters,omitempty"`
}

type Interpretation struct {
	// The unique id of the interpretation
	Id int32 `json:"id,omitempty"`
	// The source of the interpretation
	Source string `json:"source,omitempty"`
	// The translated text
	Text string `json:"text,omitempty"`
}

type Result struct {
	// The associated arabic text
	Text string `json:"text,omitempty"`
	// The related translations to the text
	Translations []Translation `json:"translations,omitempty"`
	// The unique verse id across the Quran
	VerseId int32 `json:"verse_id,omitempty"`
	// The verse key e.g 1:1
	VerseKey string `json:"verse_key,omitempty"`
}

type SearchRequest struct {
	// The language for translation
	Language string `json:"language,omitempty"`
	// The number of results to return
	Limit int32 `json:"limit,omitempty"`
	// The pagination number
	Page int32 `json:"page,omitempty"`
	// The query to ask
	Query string `json:"query,omitempty"`
}

type SearchResponse struct {
	// The current page
	Page int32 `json:"page,omitempty"`
	// The question asked
	Query string `json:"query,omitempty"`
	// The results for the query
	Results []Result `json:"results,omitempty"`
	// The total pages
	TotalPages int32 `json:"total_pages,omitempty"`
	// The total results returned
	TotalResults int32 `json:"total_results,omitempty"`
}

type SummaryRequest struct {
	// The chapter id e.g 1
	Chapter int32 `json:"chapter,omitempty"`
	// Specify the language e.g en
	Language string `json:"language,omitempty"`
}

type SummaryResponse struct {
	// The chapter id
	Chapter int32 `json:"chapter,omitempty"`
	// The source of the summary
	Source string `json:"source,omitempty"`
	// The short summary for the chapter
	Summary string `json:"summary,omitempty"`
	// The full description for the chapter
	Text string `json:"text,omitempty"`
}

type Translation struct {
	// The unique id of the translation
	Id int32 `json:"id,omitempty"`
	// The source of the translation
	Source string `json:"source,omitempty"`
	// The translated text
	Text string `json:"text,omitempty"`
}

type Verse struct {
	// The unique id of the verse in the whole book
	Id int32 `json:"id,omitempty"`
	// The interpretations of the verse
	Interpretations []Interpretation `json:"interpretations,omitempty"`
	// The key of this verse (chapter:verse) e.g 1:1
	Key string `json:"key,omitempty"`
	// The verse number in this chapter
	Number int32 `json:"number,omitempty"`
	// The page of the Quran this verse is on
	Page int32 `json:"page,omitempty"`
	// The arabic text for this verse
	Text string `json:"text,omitempty"`
	// The basic translation of the verse
	TranslatedText string `json:"translated_text,omitempty"`
	// The alternative translations for the verse
	Translations []Translation `json:"translations,omitempty"`
	// The phonetic transliteration from arabic
	Transliteration string `json:"transliteration,omitempty"`
	// The individual words within the verse (Ayah)
	Words []Word `json:"words,omitempty"`
}

type VersesRequest struct {
	// The chapter id to retrieve
	Chapter int32 `json:"chapter,omitempty"`
	// Return the interpretation (tafsir)
	Interpret bool `json:"interpret,omitempty"`
	// The language of translation
	Language string `json:"language,omitempty"`
	// The verses per page
	Limit int32 `json:"limit,omitempty"`
	// The page number to request
	Page int32 `json:"page,omitempty"`
	// Return alternate translations
	Translate bool `json:"translate,omitempty"`
	// Return the individual words with the verses
	Words bool `json:"words,omitempty"`
}

type VersesResponse struct {
	// The chapter requested
	Chapter int32 `json:"chapter,omitempty"`
	// The page requested
	Page int32 `json:"page,omitempty"`
	// The total pages
	TotalPages int32 `json:"total_pages,omitempty"`
	// The verses on the page
	Verses []Verse `json:"verses,omitempty"`
}

type Word struct {
	// The character type e.g word, end
	CharType string `json:"char_type,omitempty"`
	// The QCF v2 font code
	Code string `json:"code,omitempty"`
	// The id of the word within the verse
	Id int32 `json:"id,omitempty"`
	// The line number
	Line int32 `json:"line,omitempty"`
	// The page number
	Page int32 `json:"page,omitempty"`
	// The position of the word
	Position int32 `json:"position,omitempty"`
	// The arabic text for this word
	Text string `json:"text,omitempty"`
	// The translated text
	Translation string `json:"translation,omitempty"`
	// The transliteration text
	Transliteration string `json:"transliteration,omitempty"`
}
