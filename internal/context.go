package internal

import (
	"context"
	"database/sql"
	"net/http"
)

type CustomContext struct {
	DB *sql.DB
}

var customContextKey string = "CUSTOM_CONTEXT"

func CreateContext(args *CustomContext, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		customContext := &CustomContext{
			DB: args.DB,
		}
		requestWithCtx := r.WithContext(context.WithValue(r.Context(), customContextKey, customContext))
		next.ServeHTTP(w, requestWithCtx)
	})
}

func GetContext(ctx context.Context) *CustomContext {
	customContext, ok := ctx.Value(customContextKey).(*CustomContext)
	if !ok {
		return nil
	}
	return customContext
}
