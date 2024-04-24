package services

import (
	c "share-Gutenberg/config"
	"testing"

	"github.com/joho/godotenv"
)

func TestUM(t *testing.T) {
	errEnv := godotenv.Load("../.env")
	if errEnv != nil {
		t.Fatal(errEnv)
	}
	db, err := c.ConnectUsersDB()
	if err != nil {
		t.Fatal("error conecting with database")
	}
	defer db.Close()
	t.Log("test UM intances")

	um := UMT{DB: db}
	test, err2 := um.GetUser(1)
	if err2 != nil {
		t.Error(err2)
	}

	t.Logf("result GetUser() id = 1: %v", test)

}
