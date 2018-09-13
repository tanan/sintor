package monitor

import (
	"fmt"
	"testing"
)

func TestGet(t *testing.T) {
	r := Requester{}
	result, err := r.get("https://www.google.com/")
	if err != nil {
		t.Fatalf("failed test %v", err)
	}
	if result.Response.Status != 200 {
		t.Fatalf("Status code is not 200, but %v", result.Response.Status)
	}
	if len(result.Response.Contents) > 1024 || len(result.Response.Contents) < 0 {
		t.Fatalf("Contents length is over 1024 bytes or null. Length: %v", len(result.Response.Contents))
	}
	fmt.Println(result.Response.Contents)
}
