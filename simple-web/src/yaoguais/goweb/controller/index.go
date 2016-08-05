package controller

import (
	"fmt"
	"net/http"
)

func Index(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "welcome to goweb tiny framework")
}
