package home

import (
	"fmt"
	"net/http"
)

func HomePageHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "Hello World!")
}