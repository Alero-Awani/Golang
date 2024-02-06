package main

import (
	"net/http"

	"github.com/aleroawani/mongo-golang/controllers"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

// The controller will pass the id to mongodb and get the data for the user
func main() {
	r := httprouter.New()
	uc := controllers.NewUserController(getSession())

	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	//this creates the golang server for you
	http.ListenAndServe("localhost:9000", r)

}

// function that returns the mongodb session
func getSession() *mgo.Session {

	//dials the database abd saves the session in s
	s, err := mgo.Dial("mongodb://127.0.0.1:27107/")
	if err != nil {
		panic(err)
	}
	return s
}
