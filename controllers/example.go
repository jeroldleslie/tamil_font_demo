package controllers

import(
	"labix.org/v2/mgo"
	"tamil_font_demo/commons"
	"net/http"
)

// UserController represents the controller for operating on the User resource
type ExampleController struct {
	session *mgo.Session
	response *commons.ResponseController
}

// NewUserController provides a reference to a UserController with provided mongo session
func NewExampleController(s *mgo.Session) *ExampleController {
	u := &ExampleController{
		session : s,
		response : commons.NewResponseController(),
	}
	//ensureIndex(u.session)
	return u
}

func (ec ExampleController) Ping(w http.ResponseWriter, r *http.Request) {
	// A very simple health check.
    w.WriteHeader(http.StatusOK)
    w.Header().Set("Content-Type", "application/json")

    // In the future we could report back on the status of our DB, or our cache 
    // (e.g. Redis) by performing a simple PING, and include them in the response.
	w.Write([]byte(`{"alive": true}`))
    //io.WriteString(w, `{"alive": true}`)
}