package controller

import (
	"fmt"
	"net/http"
	"yaoguais/goweb/lib/template"
)

func CreateUser(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(response, "create user success")

	template.Assign("name", "yaogaui")
	fmt.Fprintln(response, template.Fetch("create_user"))
}

func UpdateUser(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "update user success")
}

func DeleteUser(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "delete user success")
}

func FindUser(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "find user success")
}
