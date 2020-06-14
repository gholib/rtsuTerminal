package controllers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Ping(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var (
		resp = models.Response{
			StatusCode: 200,
		}
	)
	resp.Send(w, r)
}
