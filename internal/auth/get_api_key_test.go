package auth

import (
	"net/http"
	"reflect"
	"testing"
)

func TestGetApiKey(t *testing.T) {
	header := make(http.Header)
	header.Add("Authorization", "ApiKey supersecret")

	got, _ := GetAPIKey(header)
	want := "supersecret"

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}
