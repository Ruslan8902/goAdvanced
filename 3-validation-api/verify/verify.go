package verify

import (
	"net/http"

	"net/smtp"

	"github.com/jordan-wright/email"
)

type VerifyHandler struct{}

func NewVerifyHandler(router *http.ServeMux) {
	handler := &VerifyHandler{}
	router.HandleFunc("POST /send", handler.Send())
	router.HandleFunc("GET /verify/{hash}", handler.Verify())

}

func (handler *VerifyHandler) Send() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		e := email.NewEmail()
		e.From = "Jordan Wright <test@gmail.com>"
		e.Subject = "Awesome Subject"
		e.Text = []byte("Text Body is, of course, supported!")
		e.HTML = []byte("<h1>Fancy HTML is supported, too!</h1>")
		e.Send("smtp.gmail.com:587", smtp.PlainAuth("", "test@gmail.com", "password123", "smtp.gmail.com"))
	}
}

func (handler *VerifyHandler) Verify() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(200)
	}
}
