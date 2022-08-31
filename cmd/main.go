package main

import "ecobake/cmd/app"

func main() {
	app.StartApplication()
}

//package main

//import (
//	"context"
//	"database/sql"
//	"ecobake/ent"
//	"log"
//
//	"entgo.io/ent/dialect"
//	entsql "entgo.io/ent/dialect/sql"
//	_ "github.com/jackc/pgx/v4/stdlib"
//)
//
//// Open new connection
//func Open(databaseUrl string) *ent.Client {
//	db, err := sql.Open("pgx", databaseUrl)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// Create an ent.Driver from `db`.
//	drv := entsql.OpenDB(dialect.Postgres, db)
//	return ent.NewClient(ent.Driver(drv))
//}
//
//func main() {
//	client := Open("postgresql://rvtn:rvtn@127.0.0.1/etn")
//
//	// Your code. For example:
//	ctx := context.Background()
//	if err := client.Schema.Create(ctx); err != nil {
//		log.Fatal(err)
//	}
//	users, err := client.User.Query().All(ctx)
//	if err != nil {
//		log.Fatal(err)
//	}
//	log.Println(users)
//}
