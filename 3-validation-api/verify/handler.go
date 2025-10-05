package verify

import (
	"encoding/json"
	"fmt"
	"goAdvancedAPI/configs"
	"goAdvancedAPI/pkg/req"
	"goAdvancedAPI/pkg/res"
	"math/rand"
	"net/http"

	"net/smtp"

	"github.com/jordan-wright/email"
)

type VerifyHandlerDeps struct {
	EmailConfig *configs.EmailConfig
}

type VerifyHandler struct{}

func NewVerifyHandler(router *http.ServeMux, emailConfig VerifyHandlerDeps) {
	handler := &VerifyHandler{}
	router.HandleFunc("POST /send", handler.Send(&emailConfig))
	router.HandleFunc("GET /verify/{hash}", handler.Verify())

}

var letterRunes = []rune("abcdefghijklmnoprstuvwxyzABCDEFGHIJKLMNOPRSTUVWXYZ")

func (handler *VerifyHandler) Send(ec *VerifyHandlerDeps) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[SendRequestStruct](&w, r)
		if err != nil {
			return
		}

		hash := make([]rune, 10)
		for i := range hash {
			hash[i] = letterRunes[rand.Intn(len(letterRunes))]
		}

		db := NewStorage("data.json")
		emailHashListWithDb := NewBinListWithDb(db)

		s := EmailHash{
			Email: body.Email,
			Hash:  string(hash),
		}

		emailHashListWithDb.EmailHashs = append(emailHashListWithDb.EmailHashs, s)
		content, _ := json.Marshal(emailHashListWithDb.EmailHashList.EmailHashs)
		emailHashListWithDb.Db.WriteStorage(content)

		e := email.NewEmail()
		e.From = fmt.Sprintf("Ruslan Araslanov <%s>", ec.EmailConfig.Email)
		e.To = []string{body.Email}
		e.Subject = "Random Hash"
		e.HTML = []byte(fmt.Sprintf("<a href='http://localhost:8081/verify/%s'>http://localhost:8081/verify/%s</a>", hash, hash))
		err = e.Send("smtp.mail.ru:465", smtp.PlainAuth("", ec.EmailConfig.Email, ec.EmailConfig.Password, ec.EmailConfig.Address))
		if err != nil {
			fmt.Print("Ошибка при отправке письма")
			res.Json(w, "Failed to send verificztion email", 500)
		}
		res.Json(w, "Hash created", 201)

	}
}

func (handler *VerifyHandler) Verify() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hashString := r.PathValue("hash")
		db := NewStorage("data.json")
		emailHashListWithDb := NewBinListWithDb(db)

		var found bool
		var foundIdx int
		for idx, hm := range emailHashListWithDb.EmailHashs {
			if hm.Hash == hashString {
				found = true
				foundIdx = idx
				break
			}
		}

		if !found {
			res.Json(w, found, 404)
		} else {
			emailHashListWithDb.EmailHashs = append(emailHashListWithDb.EmailHashs[:foundIdx], emailHashListWithDb.EmailHashs[foundIdx+1:]...)
			content, _ := json.Marshal(emailHashListWithDb.EmailHashs)
			emailHashListWithDb.Db.WriteStorage(content)
			res.Json(w, found, 200)
		}
	}
}
