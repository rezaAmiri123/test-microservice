package http

import (
	"encoding/json"
	"github.com/opentracing/opentracing-go"
	"github.com/rezaAmiri123/test-microservice/user_service/domain"
	"net/http"
	"strings"
)

func (h *HttpServer) GetProfile(w http.ResponseWriter, r *http.Request) {
	token := h.tokenFromHeader(r)
	u, err := h.app.Queries.GetUserToken.Handler(r.Context(), token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *HttpServer) CreateUser(w http.ResponseWriter, r *http.Request) {
	h.metric.CreateUserHttpRequests.Inc()

	span, ctx := opentracing.StartSpanFromContext(r.Context(), "HttpServer.CreateUser")
	defer span.Finish()

	u := &domain.User{}
	if err := json.NewDecoder(r.Body).Decode(u); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.app.Commands.CreateUser.Handle(ctx, u); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *HttpServer) Login(w http.ResponseWriter, r *http.Request) {
	type UserLogin struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	userLogin := &UserLogin{}
	if err := json.NewDecoder(r.Body).Decode(userLogin); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	u, err := h.app.Queries.GetProfile.Handle(r.Context(), userLogin.Username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	token,err := u.GenerateJWTToken()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	resp := map[string]string{
		"token": token,
	}
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *HttpServer) tokenFromHeader(r *http.Request) string {
	headerValue := r.Header.Get("Authorization")

	if len(headerValue) > 7 && strings.ToLower(headerValue[0:6]) == "bearer" {
		return headerValue[7:]
	}

	return ""
}
