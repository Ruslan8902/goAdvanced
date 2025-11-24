package auth

import (
	"net/http"
	"orderApiStart/configs"
	"orderApiStart/internal/user"
	"orderApiStart/pkg/jwt"
	"orderApiStart/pkg/req"
	"orderApiStart/pkg/res"
	"strconv"
)

type AuthHandlerDeps struct {
	Config            *configs.Config
	UserRepository    *user.UserRepository
	SessionRepository *SessionRepository
}

type AuthHandler struct {
	Config            *configs.Config
	UserRepository    *user.UserRepository
	SessionRepository *SessionRepository
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		Config:         deps.Config,
		UserRepository: deps.UserRepository,
	}
	router.HandleFunc("POST /auth/login", handler.LoginRequest())
	router.HandleFunc("POST /auth/confirm-login", handler.LoginConfirm())

}

func (handler *AuthHandler) LoginRequest() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[LoginRequest](&w, r)
		if err != nil {
			return
		}

		userObj, err := handler.UserRepository.GetByPhone(body.Phone)
		if err != nil {
			newUser := user.NewUser(body.Phone)
			userObj, err = handler.UserRepository.Create(newUser)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
		}

		session := NewSession(userObj.ID)
		for {
			existedSessionId, _ := handler.SessionRepository.GetBySessionId(session.SessionID)
			if existedSessionId == nil {
				break
			}
			session.GenerateSessionID()
		}

		createdSession, err := handler.SessionRepository.Create(session)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		data := LoginResponse{
			SessionID: createdSession.SessionID,
		}

		res.Json(w, data, 200)
	}
}

func (handler *AuthHandler) LoginConfirm() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[LoginConfirmRequest](&w, r)
		if err != nil {
			return
		}

		session, err := handler.SessionRepository.GetBySessionId(body.SessionID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if session.ConfirmationCode != body.Code {
			http.Error(w, "Wrong conformation code", http.StatusBadRequest)
			return
		}

		token, err := jwt.NewJWT(handler.Config.Auth.Secret).Create(jwt.JWTData{
			UserID: strconv.FormatUint(uint64(session.UserID), 10),
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		session.Token = token
		_, err = handler.SessionRepository.Update(session)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data := LoginConfirmResponse{
			Token: token,
		}

		res.Json(w, data, 200)
	}
}
