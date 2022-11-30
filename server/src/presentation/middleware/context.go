package middleware

import (
	"net/http"

	"github.com/Doer-org/hack-camp_vol9_2022-1/core/context"
)

func Context(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := context.NewContext(r)
		h.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}
