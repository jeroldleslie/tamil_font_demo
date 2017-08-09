package server

import (
	"tamil_font_demo/commons"
	"tamil_font_demo/controllers"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
)

func StartServer(port string) error {
	apiPrefix := "/api/v1"

	router := mux.NewRouter()
	apiRoutes := mux.NewRouter()

	session := commons.GetDBSession()
	defer session.Close()

	
	//Example requests
	// ex := controllers.NewExampleController(session)
	// router.HandleFunc(apiPrefix + "/ping", ex.Ping).Methods("GET")

	auth := controllers.NewAuthController(session)
	router.HandleFunc(apiPrefix+"/login", auth.LoginHandler).Methods("POST")


	qc := controllers.NewQuestionController(session)
	router.HandleFunc(apiPrefix+"/question", qc.CreateQuestion).Methods("POST")
	router.HandleFunc(apiPrefix+"/question", qc.GetAllQuestion).Methods("GET")
	router.HandleFunc(apiPrefix+"/question/{id}", qc.GetQuestionById).Methods("GET")
	router.HandleFunc(apiPrefix+"/question/{id}", qc.RemoveQuestion).Methods("DELETE")
	router.HandleFunc(apiPrefix+"/question/{id}", qc.UpdateQuestion).Methods("PUT")

	router.PathPrefix(apiPrefix).Handler(negroni.New(
		negroni.HandlerFunc(auth.ValidateTokenMiddleware),
		negroni.Wrap(apiRoutes),
	))

	n := negroni.Classic() // Includes some default middlewares
	n.UseHandler(router)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"HEAD", "GET", "POST", "PUT", "OPTIONS", "DELETE"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})

	// Insert the middleware
	handler := c.Handler(n)

	log.Printf("Listening on port: %s", port)
	return http.ListenAndServe(":"+port, handler)
}
