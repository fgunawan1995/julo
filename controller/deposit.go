package controller

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/fgunawan1995/julo/model"
	"github.com/fgunawan1995/julo/usecase"
)

func DepositBalance(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var p model.DepositRequest
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(model.BuildAPIResponseError(err))
		return
	}
	userID, ok := ctx.Value(model.ContextUserIDKey).(string)
	if !ok {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(model.BuildAPIResponseError(errors.New("type assertion failed")))
		return
	}
	result, err := usecase.DepositBalance(p, userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.BuildAPIResponseError(err))
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model.BuildAPIResponseSuccess(result))
}
