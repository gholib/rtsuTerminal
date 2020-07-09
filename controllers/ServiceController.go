package controllers

import (
	"encoding/json"
	"github.com/gholib/rtsuTerminal/models"
	"github.com/gholib/rtsuTerminal/utils"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func CreateService(w http.ResponseWriter, r *http.Request, _ httprouter.Params)  {
	body := utils.BodyParser(r)
	var service models.Service
	err := json.Unmarshal(body, &service)
	if err != nil {
		utils.ToJson(w, err.Error(), http.StatusUnprocessableEntity)
	}
	err = models.NewService(service)
	if err != nil {
		utils.ToJson(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	utils.ToJson(w, "Услуга успешно добавлена", http.StatusCreated)


}


func GetServices(w http.ResponseWriter, r *http.Request, _ httprouter.Params)  {
	service := models.GetAll(models.Services)
	utils.ToJson(w, service, http.StatusOK)

}