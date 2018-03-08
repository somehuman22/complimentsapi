package controllers

import (
	"complimentsapi/controllers"

	"github.com/gorilla/mux"
)

//Router ... the global router
var Router *mux.Router

//InitRouter ... assigns the router and sets up endpoints
func InitRouter() {
	Router = mux.NewRouter()
	Router.HandleFunc("/users/{userename}", controllers.GetProfileHandler)
}
