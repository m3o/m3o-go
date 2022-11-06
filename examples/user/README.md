# User

An [m3o.com](https://m3o.com) API. For example usage see [m3o.com/user/api](https://m3o.com/user/api).

Endpoints:

## VerifyToken

Check whether the token attached to MagicLink is valid or not.
Ideally, you need to call this endpoint from your http request
handler that handles the endpoint which is specified in the
SendMagicLink request.


[https://m3o.com/user/api#VerifyToken](https://m3o.com/user/api#VerifyToken)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/user"
)

// Check whether the token attached to MagicLink is valid or not.
// Ideally, you need to call this endpoint from your http request
// handler that handles the endpoint which is specified in the
// SendMagicLink request.
func VerifyAtoken() {
	userService := user.NewUserService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := userService.VerifyToken(&user.VerifyTokenRequest{
		Token: "EdsUiidouJJJLldjlloofUiorkojflsWWdld",

	})
	fmt.Println(rsp, err)
	
}
```
## Create

Create a new user account. The email address and username for the account must be unique.


[https://m3o.com/user/api#Create](https://m3o.com/user/api#Create)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/user"
)

// Create a new user account. The email address and username for the account must be unique.
func CreateAnAccount() {
	userService := user.NewUserService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := userService.Create(&user.CreateRequest{
		Email: "joe@example.com",
Id: "user-1",
Password: "Password1",
Username: "joe",

	})
	fmt.Println(rsp, err)
	
}
```
## Update

Update the account username or email


[https://m3o.com/user/api#Update](https://m3o.com/user/api#Update)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/user"
)

// Update the account username or email
func UpdateAnAccount() {
	userService := user.NewUserService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := userService.Update(&user.UpdateRequest{
		Email: "joe+2@example.com",
Id: "user-1",
Username: "joe",

	})
	fmt.Println(rsp, err)
	
}
```
## Read

Read an account by id, username or email. Only one need to be specified.


[https://m3o.com/user/api#Read](https://m3o.com/user/api#Read)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/user"
)

// Read an account by id, username or email. Only one need to be specified.
func ReadAnAccountById() {
	userService := user.NewUserService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := userService.Read(&user.ReadRequest{
		Id: "user-1",

	})
	fmt.Println(rsp, err)
	
}
```
## Read

Read an account by id, username or email. Only one need to be specified.


[https://m3o.com/user/api#Read](https://m3o.com/user/api#Read)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/user"
)

// Read an account by id, username or email. Only one need to be specified.
func ReadAccountByUsernameOrEmail() {
	userService := user.NewUserService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := userService.Read(&user.ReadRequest{
		Username: "joe",

	})
	fmt.Println(rsp, err)
	
}
```
## Read

Read an account by id, username or email. Only one need to be specified.


[https://m3o.com/user/api#Read](https://m3o.com/user/api#Read)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/user"
)

// Read an account by id, username or email. Only one need to be specified.
func ReadAccountByEmail() {
	userService := user.NewUserService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := userService.Read(&user.ReadRequest{
		Email: "joe@example.com",

	})
	fmt.Println(rsp, err)
	
}
```
## SendVerificationEmail

Send a verification email to a user.


[https://m3o.com/user/api#SendVerificationEmail](https://m3o.com/user/api#SendVerificationEmail)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/user"
)

// Send a verification email to a user.
func SendVerificationEmail() {
	userService := user.NewUserService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := userService.SendVerificationEmail(&user.SendVerificationEmailRequest{
		Email: "joe@example.com",
FailureRedirectUrl: "https://m3o.com/verification-failed",
FromName: "Awesome Dot Com",
RedirectUrl: "https://m3o.com",
Subject: "Email verification",
TextContent: `Hi there,

Please verify your email by clicking this link: $micro_verification_link`,

	})
	fmt.Println(rsp, err)
	
}
```
## Login

Login using username or email. The response will return a new session for successful login,
401 in the case of login failure and 500 for any other error


[https://m3o.com/user/api#Login](https://m3o.com/user/api#Login)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/user"
)

// Login using username or email. The response will return a new session for successful login,
// 401 in the case of login failure and 500 for any other error
func LogAuserIn() {
	userService := user.NewUserService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := userService.Login(&user.LoginRequest{
		Email: "joe@example.com",
Password: "Password1",

	})
	fmt.Println(rsp, err)
	
}
```
## Logout

Logout a user account


[https://m3o.com/user/api#Logout](https://m3o.com/user/api#Logout)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/user"
)

// Logout a user account
func LogAuserOut() {
	userService := user.NewUserService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := userService.Logout(&user.LogoutRequest{
		SessionId: "df91a612-5b24-4634-99ff-240220ab8f55",

	})
	fmt.Println(rsp, err)
	
}
```
## List

List all users. Returns a paged list of results


[https://m3o.com/user/api#List](https://m3o.com/user/api#List)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/user"
)

// List all users. Returns a paged list of results
func ListAllUsers() {
	userService := user.NewUserService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := userService.List(&user.ListRequest{
		Limit: 100,
Offset: 0,

	})
	fmt.Println(rsp, err)
	
}
```
## VerifyEmail

Verify the email address of an account from a token sent in an email to the user.


[https://m3o.com/user/api#VerifyEmail](https://m3o.com/user/api#VerifyEmail)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/user"
)

// Verify the email address of an account from a token sent in an email to the user.
func VerifyEmail() {
	userService := user.NewUserService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := userService.VerifyEmail(&user.VerifyEmailRequest{
		Token: "012345",

	})
	fmt.Println(rsp, err)
	
}
```
## Delete

Delete an account by id


[https://m3o.com/user/api#Delete](https://m3o.com/user/api#Delete)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/user"
)

// Delete an account by id
func DeleteUserAccount() {
	userService := user.NewUserService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := userService.Delete(&user.DeleteRequest{
		Id: "8b98acbe-0b6a-4d66-a414-5ffbf666786f",

	})
	fmt.Println(rsp, err)
	
}
```
## ReadSession

Read a session by the session id. In the event it has expired or is not found and error is returned.


[https://m3o.com/user/api#ReadSession](https://m3o.com/user/api#ReadSession)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/user"
)

// Read a session by the session id. In the event it has expired or is not found and error is returned.
func ReadAsessionByTheSessionId() {
	userService := user.NewUserService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := userService.ReadSession(&user.ReadSessionRequest{
		SessionId: "df91a612-5b24-4634-99ff-240220ab8f55",

	})
	fmt.Println(rsp, err)
	
}
```
## SendMagicLink

Login using email only - Passwordless


[https://m3o.com/user/api#SendMagicLink](https://m3o.com/user/api#SendMagicLink)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/user"
)

// Login using email only - Passwordless
func SendAmagicLink() {
	userService := user.NewUserService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := userService.SendMagicLink(&user.SendMagicLinkRequest{
		Address: "www.example.com",
Email: "joe@example.com",
Endpoint: "verifytoken",
FromName: "Awesome Dot Com",
Subject: "MagicLink to access your account",
TextContent: `Hi there,

Click here to access your account $micro_verification_link`,

	})
	fmt.Println(rsp, err)
	
}
```
## LogoutAll

Logout of all user's sessions


[https://m3o.com/user/api#LogoutAll](https://m3o.com/user/api#LogoutAll)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/user"
)

// Logout of all user's sessions
func LogoutAllSessionsForAuser() {
	userService := user.NewUserService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := userService.LogoutAll(&user.LogoutAllRequest{
		
	})
	fmt.Println(rsp, err)
	
}
```
## UpdatePassword

Update the account password


[https://m3o.com/user/api#UpdatePassword](https://m3o.com/user/api#UpdatePassword)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/user"
)

// Update the account password
func UpdateTheAccountPassword() {
	userService := user.NewUserService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := userService.UpdatePassword(&user.UpdatePasswordRequest{
		ConfirmPassword: "Password2",
NewPassword: "Password2",
OldPassword: "Password1",
UserId: "user-1",

	})
	fmt.Println(rsp, err)
	
}
```
## SendPasswordResetEmail

Send an email with a verification code to reset password.
Call "ResetPassword" endpoint once user provides the code.


[https://m3o.com/user/api#SendPasswordResetEmail](https://m3o.com/user/api#SendPasswordResetEmail)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/user"
)

// Send an email with a verification code to reset password.
// Call "ResetPassword" endpoint once user provides the code.
func SendPasswordResetEmail() {
	userService := user.NewUserService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := userService.SendPasswordResetEmail(&user.SendPasswordResetEmailRequest{
		Email: "joe@example.com",
FromName: "Awesome Dot Com",
Subject: "Password reset",
TextContent: `Hi there,
 click here to reset your password: myapp.com/reset/code?=$code`,

	})
	fmt.Println(rsp, err)
	
}
```
## ResetPassword

Reset password with the code sent by the "SendPasswordResetEmail" endpoint.


[https://m3o.com/user/api#ResetPassword](https://m3o.com/user/api#ResetPassword)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/user"
)

// Reset password with the code sent by the "SendPasswordResetEmail" endpoint.
func ResetPassword() {
	userService := user.NewUserService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := userService.ResetPassword(&user.ResetPasswordRequest{
		Code: "012345",
ConfirmPassword: "NewPassword1",
Email: "joe@example.com",
NewPassword: "NewPassword1",

	})
	fmt.Println(rsp, err)
	
}
```
