# Avatar

An [m3o.com](https://m3o.com) API. For example usage see [m3o.com/avatar/api](https://m3o.com/avatar/api).

Endpoints:

## Generate




[https://m3o.com/avatar/api#Generate](https://m3o.com/avatar/api#Generate)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/avatar"
)

// 
func GenerateAvatarAndReturnBase64stringOfTheAvatar() {
	avatarService := avatar.NewAvatarService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := avatarService.Generate(&avatar.GenerateRequest{
		Format: "png",
Gender: "female",
Upload: true,
Username: "",

	})
	fmt.Println(rsp, err)
	
}
```
## Generate




[https://m3o.com/avatar/api#Generate](https://m3o.com/avatar/api#Generate)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/avatar"
)

// 
func GenerateAnAvatarAndUploadTheAvatarToMicrosCdn() {
	avatarService := avatar.NewAvatarService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := avatarService.Generate(&avatar.GenerateRequest{
		Format: "jpeg",
Gender: "female",
Upload: false,
Username: "",

	})
	fmt.Println(rsp, err)
	
}
```
