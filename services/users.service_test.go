package services

import (
	c "share-Gutenberg/config"
	m "share-Gutenberg/models"
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
	t.Log("test UM GetUser id 1")
	test, err2 := um.GetUser(1)
	if err2 != nil {
		t.Error(err2)
	}

	t.Logf("result GetUser() id = 1: %v", test)
	t.Log("test not existing user")
	test2, err3 := um.GetUser(0)
	if err3 != nil {
		t.Log(err3)
	}
	if test2 != nil {
		t.Errorf("expected nil, have: %v", test2)
	}
	t.Log("test UM CreateUser")
	userInfo := m.UserInfo{
		Name:     "usuario",
		Email:    "false@falseto",
		Password: "HASHEADO",
		Role:     "USER",
	}
	if err4 := um.CreateUser(&userInfo); err4 != nil {
		t.Error(err4)
	}
	t.Log("test UM CreateUser with same email")
	if err5 := um.CreateUser(&userInfo); err5 == nil {
		t.Error("there is not error with used email in register")
	} else {
		t.Log(err5.Error())
	}
}
