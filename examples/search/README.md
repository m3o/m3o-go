# Search

An [m3o.com](https://m3o.com) API. For example usage see [m3o.com/search/api](https://m3o.com/search/api).

Endpoints:

## CreateIndex

Create an index by name


[https://m3o.com/search/api#CreateIndex](https://m3o.com/search/api#CreateIndex)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/search"
)

// Create an index by name
func CreateAnIndex() {
	searchService := search.NewSearchService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := searchService.CreateIndex(&search.CreateIndexRequest{
		Index: "customers",

	})
	fmt.Println(rsp, err)
	
}
```
## DeleteIndex

Delete an index by name


[https://m3o.com/search/api#DeleteIndex](https://m3o.com/search/api#DeleteIndex)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/search"
)

// Delete an index by name
func DeleteAnIndex() {
	searchService := search.NewSearchService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := searchService.DeleteIndex(&search.DeleteIndexRequest{
		Index: "customers",

	})
	fmt.Println(rsp, err)
	
}
```
## Index

Index a record i.e. insert a document to search for.


[https://m3o.com/search/api#Index](https://m3o.com/search/api#Index)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/search"
)

// Index a record i.e. insert a document to search for.
func IndexArecord() {
	searchService := search.NewSearchService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := searchService.Index(&search.IndexRequest{
		Data: map[string]interface{}{
	"age": 37,
	"starsign": "Leo",
	"name": "John Doe",
},
Index: "customers",

	})
	fmt.Println(rsp, err)
	
}
```
## Search

Search for records in a given in index


[https://m3o.com/search/api#Search](https://m3o.com/search/api#Search)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/search"
)

// Search for records in a given in index
func SearchForArecord() {
	searchService := search.NewSearchService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := searchService.Search(&search.SearchRequest{
		Index: "customers",
Query: "name == 'John'",

	})
	fmt.Println(rsp, err)
	
}
```
## Search

Search for records in a given in index


[https://m3o.com/search/api#Search](https://m3o.com/search/api#Search)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/search"
)

// Search for records in a given in index
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

Search for records in a given in index


[https://m3o.com/search/api#Search](https://m3o.com/search/api#Search)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/search"
)

// Search for records in a given in index
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

Delete a record given its ID


[https://m3o.com/search/api#Delete](https://m3o.com/search/api#Delete)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/search"
)

// Delete a record given its ID
func DeleteArecord() {
	searchService := search.NewSearchService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := searchService.Delete(&search.DeleteRequest{
		Id: "1234",
Index: "customers",

	})
	fmt.Println(rsp, err)
	
}
```
