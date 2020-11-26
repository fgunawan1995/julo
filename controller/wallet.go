package controller

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/fgunawan1995/julo/model"
	"github.com/fgunawan1995/julo/usecase"
)

func InitWallet(w http.ResponseWriter, r *http.Request) {
	var p model.InitWalletRequest
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(model.BuildAPIResponseError(err))
		return
	}
	result, err := usecase.InitWallet(p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.BuildAPIResponseError(err))
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model.BuildAPIResponseSuccess(result))
}

func GetWallet(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID, ok := ctx.Value(model.ContextUserIDKey).(string)
	if !ok {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(model.BuildAPIResponseError(errors.New("type assertion failed")))
		return
	}
	result, err := usecase.GetWallet(userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.BuildAPIResponseError(err))
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model.BuildAPIResponseSuccess(result))
}

func EnableWallet(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID, ok := ctx.Value(model.ContextUserIDKey).(string)
	if !ok {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(model.BuildAPIResponseError(errors.New("type assertion failed")))
		return
	}
	result, err := usecase.EnableWallet(userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.BuildAPIResponseError(err))
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model.BuildAPIResponseSuccess(result))
}

func DisableWallet(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID, ok := ctx.Value(model.ContextUserIDKey).(string)
	if !ok {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(model.BuildAPIResponseError(errors.New("type assertion failed")))
		return
	}
	result, err := usecase.DisableWallet(userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.BuildAPIResponseError(err))
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model.BuildAPIResponseSuccess(result))
}
