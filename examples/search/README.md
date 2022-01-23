# Search

An [m3o.com](https://m3o.com) API. For example usage see [m3o.com/search/api](https://m3o.com/search/api).

Endpoints:

## Index

Index a document i.e. insert a document to search for.


[https://m3o.com/search/api#Index](https://m3o.com/search/api#Index)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/search"
)

// Index a document i.e. insert a document to search for.
func IndexAdocument() {
	searchService := search.NewSearchService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := searchService.Index(&search.IndexRequest{
		Document: &search.Document{
	Contents: map[string]interface{}{
		"starsign": "Leo",
		"name": "John Doe",
		"age": 37,
},
	Id: "1234",
},
Index: "customers",

	})
	fmt.Println(rsp, err)
	
}
```
## Search

Search for documents in a given in index


[https://m3o.com/search/api#Search](https://m3o.com/search/api#Search)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/search"
)

// Search for documents in a given in index
func SearchForAdocument() {
	searchService := search.NewSearchService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := searchService.Search(&search.SearchRequest{
		Index: "customers",
Query: "name == 'John'",

	})
	fmt.Println(rsp, err)
	
}
```
## Search

Search for documents in a given in index


[https://m3o.com/search/api#Search](https://m3o.com/search/api#Search)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/search"
)

// Search for documents in a given in index
func SearchOnMultipleFieldsand() {
	searchService := search.NewSearchService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := searchService.Search(&search.SearchRequest{
		Index: "customers",
Query: "name == 'John' AND starsign == 'Leo'",

	})
	fmt.Println(rsp, err)
	
}
```
## Search

Search for documents in a given in index


[https://m3o.com/search/api#Search](https://m3o.com/search/api#Search)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/search"
)

// Search for documents in a given in index
func SearchOnMultipleFieldsor() {
	searchService := search.NewSearchService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := searchService.Search(&search.SearchRequest{
		Index: "customers",
Query: "name == 'John' OR name == 'Jane'",

	})
	fmt.Println(rsp, err)
	
}
```
## Delete

Delete a document given its ID


[https://m3o.com/search/api#Delete](https://m3o.com/search/api#Delete)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/search"
)

// Delete a document given its ID
func DeleteAdocument() {
	searchService := search.NewSearchService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := searchService.Delete(&search.DeleteRequest{
		Id: "1234",
Index: "customers",

	})
	fmt.Println(rsp, err)
	
}
```
## DeleteIndex

Delete an index.


[https://m3o.com/search/api#DeleteIndex](https://m3o.com/search/api#DeleteIndex)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/search"
)

// Delete an index.
func DeleteAnIndex() {
	searchService := search.NewSearchService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := searchService.DeleteIndex(&search.DeleteIndexRequest{
		Index: "customers",

	})
	fmt.Println(rsp, err)
	
}
```
