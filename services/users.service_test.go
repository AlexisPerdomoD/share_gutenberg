package services

import (
	c "share-Gutenberg/config"
	m "share-Gutenberg/models"
	"testing"
	"time"

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
		return
	}
	defer db.Close()
	t.Log("test UM intances")

	um := UMT{DB: db}

	t.Log("test UM CreateUser")
	userInfo := m.UserInfo{
		Name:      "testino",
		Email:     "correo@unico2",
		Password:  "HASHEADO",
		Role:      "USER",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err4 := um.CreateUser(&userInfo)
	if err4 != nil {
		t.Error(err4)
	}
	t.Log("test UM CreateUser with same email")
	err5 := um.CreateUser(&userInfo)
	if err5 == nil {
		t.Error("there is not error with used email in register")
	}

	test, err2 := um.GetUserByEmail("correo@unico2")
	if err2 != nil {
		t.Error(err2)
	}
	t.Logf("created user: %v", test)
	t.Logf("result GetUser() id = 1: %v", test)
	t.Log("test not existing user")
	test2, err3 := um.GetUser(0)
	if err3 == nil {
		t.Error("there must be an error when is not found the user and it wasnt")
	}
	if test2 != nil {
		t.Errorf("expected nil, have: %v", test2)
	}

	t.Log("test update")
	userUpdates := m.UserInfo{
		Name:     "nuevo Nombre separado",
		Email:    "nuevo@email",
		Password: "",
		Role:     "",
	}
	if err6 := um.UpdateUser(uint(test.Id), &userUpdates); err6 != nil {
		t.Error(err6)
	}
	test3, err7 := um.GetUser(uint(test.Id))

	if test3.Email != "nuevo@email" && test3.Name != "nuevo Nombre separado" || err7 != nil {
		t.Error("fail updating user fields")
	}
	t.Log("test delete user")
	if err7 := um.DeleteUser(uint(test.Id)); err7 != nil {
		t.Error("error to delete the user")
	}

}
