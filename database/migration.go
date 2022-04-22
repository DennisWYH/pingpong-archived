package database

import (
	"context"
	"log"

	"pingpong2/ent"
	"pingpong2/ent/migrate"
)

func MigrateTablesWithDrop() {
	client, err := ent.Open("postgres", "host=localhost port=5432 user=postgres dbname=testdb sslmode=disable")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	ctx := context.Background()
	err = client.Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)
}