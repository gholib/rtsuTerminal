package models

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
)

type Response struct {
	ContentType     string
	StatusCode      int
	ContentLanguage string
	ErrorMessage    string
	ReqID           string
	Payload         interface{}
}

func (res *Response) Send(w http.ResponseWriter, r *http.Request) {

	if res.Payload == nil {
		res.Payload = struct {
			XMLName     xml.Name `xml:"response" json:"-"`
			Error       bool     `xml:"error,tag" json:"error,omitempty"`
			Code        int      `xml:"code,tag" json:"code,omitempty"`
			Description string   `xml:"description,tag" json:"description,omitempty"`
		}{
			Error:       true,
			Code:        res.StatusCode,
			Description: res.ErrorMessage,
		}
	}

	switch res.ContentType {
	case "application/xml":
		w.Header().Set("Content-Type", "application/xml; charset=utf-8")
		w.WriteHeader(res.StatusCode)
		fmt.Fprintf(w, "<?xml version=\"1.0\" encoding=\"UTF-8\"?>")
		xml.NewEncoder(w).Encode(&res.Payload)
	default:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(res.StatusCode)
		json.NewEncoder(w).Encode(&res.Payload)
	}

	log.Println(res.ReqID, "Response:", r.RemoteAddr, r.Method, r.RequestURI)
	log.Println(res.ReqID, "Response Body:", r.RemoteAddr, r.Method, res.Payload)
}
