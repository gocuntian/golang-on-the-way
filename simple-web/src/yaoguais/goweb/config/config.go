package config

import (
	"net/http"
	"yaoguais/goweb/controller"
)

type callback func(http.ResponseWriter, *http.Request)

var Routers = map[string]callback{
	"/user/create": controller.CreateUser,
	"/user/update": controller.UpdateUser,
	"/user/delete": controller.DeleteUser,
	"/user/find":   controller.FindUser,
	"/":            controller.Index,
}
