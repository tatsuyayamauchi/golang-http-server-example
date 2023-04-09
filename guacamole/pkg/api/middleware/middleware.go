package middleware

import (
	"net/http"
)

type Middleware func(http.Handler) http.Handler
type Middlewares []Middleware

// mwsの先頭に設定されたものからチェーンで適用します
func (mws Middlewares) Chain(h http.Handler) http.Handler {
	for i := len(mws) - 1; i >= 0; i-- {
		h = mws[i](h)
	}
	return h
}
