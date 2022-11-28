package internal

import (
	"context"
	"database/sql"
	"ecobake/cmd/config"
	"ecobake/ent"
	"fmt"
	"log"
	"net/url"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"

	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func EntConn(cfg config.PostgresConn) (*ent.Client, error) {
	dsn := url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(cfg.DBUser, cfg.DBPass),
		Host:   fmt.Sprintf("%s:%s", cfg.URL, cfg.Port),
		Path:   cfg.DBName,
	}

	q := dsn.Query()
	dsn.RawQuery = q.Encode()
	pool, err := pgxpool.Connect(context.Background(), dsn.String())
	if err != nil {
		return nil, err
	}
	if err := pool.Ping(context.Background()); err != nil {
		return nil, err
	}

	db, err := sql.Open("pgx", dsn.String())
	if err != nil {
		log.Fatal(err)
	}
	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB(dialect.Postgres, db)
	client := ent.NewClient(ent.Driver(drv))

	ctx := context.Background()
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatal(err)
	}
	return client, nil
}
