package database

import (
	"context"

	_ "github.com/lib/pq"

	"pingpong2/ent"
)

func AddUser(username string, password string) error {
	client, err := ent.Open("postgres", "host=localhost port=5432 user=postgres dbname=yunhaiwang sslmode=disable")
	if err != nil {
		return err
	}
	defer client.Close()
	ctx := context.Background()
	client.User.Create().SetName(username).SetPassword(password).Save(ctx)
	return nil
}

