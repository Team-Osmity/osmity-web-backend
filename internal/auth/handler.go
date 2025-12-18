package auth

import (
	"encoding/json"
	"net/http"
)

type Handler struct {
	Service *Service
}

type registerRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var req registerRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	if err := h.Service.Register(
		r.Context(),
		req.Email,
		req.Password,
	); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var req loginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	session, err := h.Service.Login(
		r.Context(),
		req.Email,
		req.Password,
	)
	if err != nil {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "session_id",
		Value:    session.ID,
		Path:     "/",
		HttpOnly: true,
		Secure:   true, // prod
		SameSite: http.SameSiteLaxMode,
		Expires:  session.ExpiresAt,
	})

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_id")
	if err != nil {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	if err := h.Service.Logout(
		r.Context(),
		cookie.Value,
	); err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	// Cookie を即時無効化
	http.SetCookie(w, &http.Cookie{
		Name:     "session_id",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Expires:  time.Unix(0, 0),
	})

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) Me(w http.ResponseWriter, r *http.Request) {
	accountID, ok := r.Context().Value(accountIDKey).(string)
	if !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"account_id": accountID,
	})
}

type verifyEmailRequest struct {
	Code string `json:"code"`
}

func (h *Handler) VerifyEmail(w http.ResponseWriter, r *http.Request) {
	var req verifyEmailRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.Code == "" {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	err := h.Service.VerifyEmail(r.Context(), req.Code)
	switch err {
	case nil:
		w.WriteHeader(http.StatusOK)
	case ErrInvalidCode, ErrExpiredCode:
		http.Error(w, "invalid or expired code", http.StatusBadRequest)
	case ErrAlreadyUsed:
		// 冪等性：既に完了なら成功扱いでも良い
		w.WriteHeader(http.StatusOK)
	default:
		http.Error(w, "internal error", http.StatusInternalServerError)
	}
}

