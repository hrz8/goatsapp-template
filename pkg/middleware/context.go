package middleware

import (
	"context"

	"github.com/labstack/echo/v4"
)

type EchoCtx struct {
	echo.Context
}

func (ec EchoCtx) Get(key string) any {
	val := ec.Context.Get(key)
	if val != nil {
		return val
	}

	return ec.Request().Context().Value(key)
}

func (ec EchoCtx) Set(key string, val any) {
	ec.SetRequest(ec.Request().WithContext(context.WithValue(ec.Request().Context(), key, val)))
}
