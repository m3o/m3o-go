package client

import (
	"os"
	"testing"
)

func TestBasicCall(t *testing.T) {
	response := map[string]interface{}{}
	if err := NewClient(&Options{
		Token: os.Getenv("TOKEN"),
	}).Call("groups", "list", map[string]interface{}{
		"memberId": "random",
	}, &response); err != nil {
		t.Fatal(err)
	}
	if len(response) > 0 {
		t.Fatal(len(response))
	}
}
