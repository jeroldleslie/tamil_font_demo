package routes

import (
	"tamil_font_demo/commons"
	"tamil_font_demo/controllers"
	"github.com/gorilla/mux"
)

var ApiPrefix = "/api/v1"

func ApiRouter() *mux.Router {

	session := commons.GetDBSession()

	apiRoutes := mux.NewRouter()

	auth := controllers.NewAuthController(session)
	apiRoutes.HandleFunc(ApiPrefix+"/login", auth.LoginHandler).Methods("POST")

	//User routes
	// Get a UserController instance
	uc := controllers.NewUserController(session)
	apiRoutes.HandleFunc(ApiPrefix+"/users", uc.CreateUser).Methods("POST")
	apiRoutes.HandleFunc(ApiPrefix+"/users", uc.GetAllUser).Methods("GET")

	return apiRoutes
}
