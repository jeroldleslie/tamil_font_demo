package controllers

import (
	"tamil_font_demo/commons"
	"tamil_font_demo/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"log"
	"net/http"
)

var (
	database   = "tamildb"
	collection = "questions"
)


type QuestionController struct {
	session  *mgo.Session
	response *commons.ResponseController
}


func NewQuestionController(s *mgo.Session) *QuestionController {
	q := &QuestionController{
		session:  s,
		response: commons.NewResponseController(),
	}
	ensureIndex(q.session)
	return q
}

func ensureIndex(s *mgo.Session) {
	session := s.Copy()
	defer session.Close()

	c := session.DB(database).C(collection)

	index := mgo.Index{
		Key:        []string{"id"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
	err := c.EnsureIndex(index)
	if err != nil {
		panic(err)
	}
}

func (uc QuestionController) GetAllQuestion(w http.ResponseWriter, r *http.Request) {
	s := uc.session.Copy()
	defer s.Close()

	c := s.DB(database).C(collection)

	var questions []models.Question
	err := c.Find(bson.M{}).All(&questions)
	if err != nil {
		uc.response.WriteError(w, "Failed get all questions", http.StatusInternalServerError)
		log.Println("Failed get all questions: ", err)
		return
	}

	uc.response.WriteSuccess(w, questions, http.StatusOK)
}

func (uc QuestionController) GetQuestionById(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)["id"]
	var id bson.ObjectId
	if bson.IsObjectIdHex(param) == true {
		id = bson.ObjectIdHex(param)
	} else {
		uc.response.WriteError(w, "Invalid input to ObjectIdHex", http.StatusInternalServerError)
		log.Println("Invalid input to ObjectIdHex: ")
		return
	}
	s := uc.session.Copy()
	defer s.Close()

	c := s.DB(database).C(collection)

	var question models.Question
	err := c.Find(bson.M{"_id": id}).One(&question)

	if err != nil {
		uc.response.WriteError(w, "Record not found", http.StatusNotFound)
		log.Println("Failed get questions: ", err)
		return
	}
	uc.response.WriteSuccess(w, question, http.StatusOK)

}

func (uc QuestionController) RemoveQuestion(w http.ResponseWriter, r *http.Request) {

	// mux.Vars grabs variables from the path
	id := mux.Vars(r)["id"]

	s := uc.session.Copy()
	defer s.Close()

	c := s.DB(database).C(collection)

	err := c.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	if err != nil {
		switch err {
		default:
			uc.response.WriteError(w, "Failed remove question", http.StatusInternalServerError)
			log.Println("Failed delete question: ", err)
			return
		case mgo.ErrNotFound:
			uc.response.WriteError(w, "Question not found", http.StatusNotFound)
			return
		}
	}
	uc.response.WriteSuccess(w, "", http.StatusOK)
}


func (uc QuestionController) CreateQuestion(w http.ResponseWriter, r *http.Request) {
	s := uc.session.Copy()
	defer s.Close()
	question := models.Question{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&question)
	if err != nil {
		uc.response.WriteError(w, "Incorrect body", http.StatusBadRequest)
		return
	}

	question.Id = bson.NewObjectId()

	c := s.DB(database).C(collection)

	err = c.Insert(question)
	if err != nil {
		fmt.Println(err)
		if mgo.IsDup(err) {
			uc.response.WriteError(w, "questions already exists", http.StatusBadRequest)
			return
		}

		uc.response.WriteError(w, "Database error", http.StatusInternalServerError)
		log.Println("Failed insert question: ", err)
		return
	}

	uc.response.WriteSuccess(w, question, http.StatusOK)
}

func (uc QuestionController) UpdateQuestion(w http.ResponseWriter, r *http.Request) {
	// mux.Vars grabs variables from the path
	id := mux.Vars(r)["id"]
	s := uc.session.Copy()
	defer s.Close()

	c := s.DB(database).C(collection)

	question := models.Question{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&question)
	if err != nil {
		uc.response.WriteError(w, "Incorrect body", http.StatusBadRequest)
		return
	}

	err = c.Update(bson.M{"_id": bson.ObjectIdHex(id)}, &question)
	if err != nil {
		switch err {
		default:
			uc.response.WriteError(w, "Database error", http.StatusInternalServerError)
			log.Println("Failed update question: ", err)
			return
		case mgo.ErrNotFound:
			uc.response.WriteError(w, "question not found", http.StatusNotFound)
			return
		}
	}
	uc.response.WriteSuccess(w, "", http.StatusOK)
}
