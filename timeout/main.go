package main

import (
	"os"
	"context"
	"fmt"
	"time"
	"net/http"
	"net/http/httptest"
)

func main() {
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(time.Hour)
	}))
	defer svr.Close()

	// Default HTTP Client
	//fmt.Println("making request")
	//http.Get(svr.URL)
	//fmt.Println("finished request")

	// Set timeout
	fmt.Println("making request")
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)
	req, _ := http.NewRequest("GET", svr.URL, nil)
	req = req.WithContext(ctx)

	_, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("finished request", err)
		os.Exit(1)
	}
}