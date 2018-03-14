package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_Get(t *testing.T) {
	ts := httptest.NewServer(NewRouter())
	defer ts.Close()

	res, _ := http.Get(ts.URL + "/espresso")

	if res.StatusCode != 200 {
		t.Fatalf("Expected status to == 200, but got %d",res.StatusCode)
	}
}

func Test_Post(t *testing.T) {
	ts := httptest.NewServer(NewRouter())
	defer ts.Close()

	res, _ := http.Post(ts.URL + "/", "", nil)
	
	if res.StatusCode == 200 {
		t.Fatalf("Expected status to == 200, but got %d",res.StatusCode)
	}
}