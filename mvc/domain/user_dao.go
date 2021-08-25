package domain

import (
	"fmt"
	"golang-microservices/mvc/utils"
	"log"
	"net/http"
)

var (
	users = map[int64]*User{
		123: {Id: 123, FirstName: "Alexey", LastName: "Gaskov", Email: "alexey.gaskov@gmail.com"},
		124: {Id: 124, FirstName: "Paul", LastName: "McCartney", Email: "beatles_mc@gmail.com"},
		125: {Id: 125, FirstName: "Demis", LastName: "Rousses", Email: "the_greek@gmail.com"},
		126: {Id: 126, FirstName: "Joe", LastName: "Dassen", Email: "joe.gaskov@gmail.com"},
	}

	UserDao userDaoInterface
)

func init() {
	UserDao = &userDao{}
}

type userDaoInterface interface {
	GetUser(int64) (*User, *utils.ApplicationError)
}

type userDao struct{}

func (u *userDao) GetUser(userId int64) (*User, *utils.ApplicationError) {
	log.Println("we're accessing the database")
	if user := users[userId]; user != nil {
		return user, nil
	}

	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("user %v does not exists", userId),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}
}
