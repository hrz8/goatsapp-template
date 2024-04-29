package main

import (
	"context"

	"github.com/hrz8/goatsapp/internal/port"
	"github.com/jackc/pgx/v5"
)

func connect(ctx context.Context, cfg port.AppConfigor) *pgx.Conn {
	conn, err := pgx.Connect(ctx, cfg.GetDatabaseURL())
	if err != nil {
		panic(err)
	}

	return conn
}
