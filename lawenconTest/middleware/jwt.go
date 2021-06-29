package middleware

import (
	"net/http"
	"strings"
	"lawenconTest/response"
	"lawenconTest/utils"

)

//JWTValidation ...
func JWTValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		rAuth := strings.Split(header, " ")

		if len(rAuth) != 2 {
			response.Errors(http.StatusUnauthorized, "Invalid Bearer").Default(w)
			return
		}

		claim, err := utils.ParseJWT(rAuth[1])
		if err == utils.ErrInvalidToken {
			response.Errors(http.StatusUnauthorized, "Unathorized").Default(w)
			return
		}
		if err != nil {
			response.Errors(http.StatusUnauthorized, "Invalid Tokens").Default(w)
			return
		}
		ctx := claim.SetValueToContext(r.Context())
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
