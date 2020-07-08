package controllers

import (
	"encoding/json"
	"github.com/gholib/rtsuTerminal/models"
	"github.com/gholib/rtsuTerminal/utils"
	"github.com/julienschmidt/httprouter"
	"net/http"
)



func CreateTerminal(w http.ResponseWriter, r *http.Request, _ httprouter.Params)  {
	// внутри body лежит слайс байтов
	body := utils.BodyParser(r)
	var terminal models.Terminal
	err :=json.Unmarshal(body, &terminal)
	if err != nil {
		utils.ToJson(w, err.Error(), http.StatusUnprocessableEntity)
	}
	err = models.NewTerminal(terminal)
	if err != nil {
		utils.ToJson(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	utils.ToJson(w, "Терминал успешно добавлен!", http.StatusCreated)

	
}
