package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gholib/rtsuTerminal/models"
	"github.com/julienschmidt/httprouter"
)

//CreateUser : createuser
func CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var (
		request models.UserRequest
	)
	err := json.NewDecoder(r.Body).Decode(&request)

	fmt.Println(err, request.Login)
}
