package models

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
)

// Response
type Response struct {
	ContentType     string `xml:"-" json:"-"`
	StatusCode      int    `xml:"-" json:"-"`
	ContentLanguage string `xml:"-" json:"-"`
	ErrorMessage    string `xml:"-" json:"-"`
	ReqID           string `xml:"-" json:"-"`

	// XMLName     xml.Name    `xml:"response" json:"-"`
	// APIVersion  string      `xml:"api_version,tag" json:"api_version,omitempty"`
	// Error       bool        `xml:"error,tag,omitempty" json:"error,omitempty"`
	// Code        int         `xml:"code,tag,omitempty" json:"code,omitempty"`
	// Description string      `xml:"description,tag,omitempty" json:"description,omitempty"`
	Payload interface{} //`xml:"-,omitempty" json:"-,omitempty"`
}

func (res *Response) Send(w http.ResponseWriter, r *http.Request) {

	if res.Payload == nil {
		res.Payload = struct {
			XMLName xml.Name `xml:"response" json:"-"`
			Error   bool     `xml:"error,tag" json:"error,omitempty"`
			Message string   `xml:"message,tag" json:"message,omitempty"`
		}{
			Error:   true,
			Message: res.ErrorMessage,
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
