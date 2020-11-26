package middleware

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"context"

	"github.com/fgunawan1995/julo/wallet/dal/cache"
	"github.com/fgunawan1995/julo/wallet/model"
)

func ValidateJWT(next http.HandlerFunc) http.HandlerFunc {
	cacheDAL := cache.New(model.GetCache())
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		authorizationHeader := req.Header.Get(model.HeaderAuth)
		if authorizationHeader == "" && req.FormValue("token") != "" {
			authorizationHeader = "Bearer " + req.FormValue("token")
		}
		if authorizationHeader != "" {
			bearerToken := strings.Split(authorizationHeader, " ")
			if len(bearerToken) == 2 {
				bearerToken := strings.Split(authorizationHeader, " ")
				if len(bearerToken) == 2 {
					ctx := req.Context()
					userID, err := cacheDAL.GetToken(bearerToken[1])
					if err != nil || userID == "0" {
						w.WriteHeader(http.StatusUnauthorized)
						json.NewEncoder(w).Encode(model.BuildAPIResponseError(err))
						return
					}
					ctx = context.WithValue(ctx, model.ContextUserIDKey, userID)
					next(w, req.WithContext(ctx))
				}
			}
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(model.BuildAPIResponseError(errors.New("Authorization header is required")))
		}
	})
}
