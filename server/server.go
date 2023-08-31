package server

import (
	"TranscribeHub_HTMX/database"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"os"
)

func Start() error {
	conn, err := pgx.Connect(context.Background(), "postgresql://postgres:test@127.0.0.1:5432/transcribe_hub")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())
	r := newRouter(&database.DaoImpl{Conn: conn})
	return r.Run(":8080")
}
