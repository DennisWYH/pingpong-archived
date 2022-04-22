package database

import (
	"context"
	"log"
	"testing"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"

	"pingpong2/ent"
	"pingpong2/ent/user"
)

func queryUserByName(username string) (*ent.User, error){
	client, err := ent.Open("postgres", "host=localhost port=5432 user=postgres dbname=testdb sslmode=disable")
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()
	ctx := context.Background()

	user, err := client.User.Query().Where(user.Name(username)).Only(ctx)
	if err != nil {
		return nil, err
	}
	return user, nil
}


func TestAddUser(t *testing.T){
	err := AddUser("mariTestUser", "mari")
	assert.NoError(t, err)
	err = AddUser("yunhaoTestUser", "yunhao")
	assert.NoError(t, err)
	userMari, err := queryUserByName("mariTestUser")
	assert.NoError(t, err)
	userYunhao, err := queryUserByName("yunhaoTestUser")
	assert.NoError(t, err)
	assert.Equal(t, "mari", userMari.Password)
	assert.Equal(t, "yunhao", userYunhao.Password)
	DeleteUserByName("mariTestUser")
	DeleteUserByName("yunhaoTestUser")
}