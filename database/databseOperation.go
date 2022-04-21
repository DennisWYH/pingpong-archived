package database

import (
	"context"
	_ "github.com/lib/pq"
	"log"

	"pingpong2/ent"
	"pingpong2/ent/user"
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

func DeleteUserByName(username string) error{
	client, err := ent.Open("postgres", "host=localhost port=5432 user=postgres dbname=yunhaiwang sslmode=disable")
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()
	ctx := context.Background()

	user, err := client.User.Query().Where(user.Name(username)).Only(ctx)
	if err != nil {
		return err
	}
	client.User.DeleteOne(user).Exec(ctx)
	return nil
}
