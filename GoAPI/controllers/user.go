package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aleroawani/mongo-golang/models"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserController struct{
	session		*mgo.Session
}

// NewuserController takes in session 

func NewUserController(s *mgo.Session) *UserController{
	return &UserController{s}
}

//struct method
func (uc UserController) GetUser (w http.ResponseWriter, r *http.Request, p httprouter.Params){
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id){
		// w is the response we are sending back to our postman or frontend
		//Writerheader usally has the status code like 404
		w.WriteHeader(http.StatusNotFound)
	}
	oid := bson.ObjectIdHex(id)

	u := models.User{}

	// if you dont have an existing database, mongo will create one for you 
	// you find the data from the database and store it in u
	if err := uc.session.DB("mongo-golang").C("users").FindId(oid); err != nil {
		w.WriteHeader(404)
		return
	}

	// you convert the data to json by marsahlling it to send to the front end 
	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}
	//if everything works out fine, here we are telling postman that we are sending json data 
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", uj)

}

//http.Request is a pointer to the request that the user has sent from the postman to golang
// since its a post method we dont need any params which is why we have the _ for httprouter.params

func (uc UserController) CreateUser (w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	u := models.User{}

	// decode the json values we get from postman so that golang can work with it 
	// we are decoding the r.Body using our struct
	json.NewDecoder(r.Body).Decode(&u)

	// to automatically create a new id for the new user 
	u.Id = bson.NewObjectId()

	// C stands for collection, so we are inserting something in the users collection
	uc.session.DB("mongo-golang").C("users").Insert(u)

	// after inserting u in the database, you want to send it back as a json object to tell the user that something
	// has been created in the database 

	uj, err := json.Marshal(u)

	if err != nil {
		fmt.Println(err)
	}

	// if theres no errror we will send uj back to the postman 
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params){
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id){
		w.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(id)

	if err := uc.session.DB("mongo-golang").C("users").RemoveId(oid); err != nil {
		w.WriteHeader(404)
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Deleted user", oid, "\n")
}