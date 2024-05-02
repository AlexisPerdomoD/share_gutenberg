package services

import (
	c "share-Gutenberg/config"
	m "share-Gutenberg/models"
	"testing"
	"time"

	"github.com/joho/godotenv"
)

func TestCM(t *testing.T) {
	errEnv := godotenv.Load("../.env")
	if errEnv != nil {
		t.Fatal("error getting env variables")
	}
	db, err := c.ConnectUsersDB()
	if err != nil {
		t.Error("error connecting with database")
		return
	}
	defer db.Close()
	t.Log("test CM intances")

	cm := CMT{DB: db}
	um := UMT{DB: db}

	usrErr := um.CreateUser(&m.UserInfo{
		Name:     "luis",
		Email:    "email@valido",
		Password: "HASHEADO",
		Role:     "user",
	})
	if usrErr != nil {
		t.Fatal(usrErr)
	}
	user, usrErr2 := um.GetUserByEmail("email@valido")
	if usrErr2 != nil {
		t.Fatal("error getting user by email")
	}

	collectionInfo := m.CollectionInfo{
		CollectionName: "testino_collection",
		Description:    "esta es una coleccion para libros de terror",
		Category:       "horror",
		Public:         true,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
		Owner:          user.Id,
		Documents:      make([]int, 0),
	}

	if err := cm.CreateCollection(&collectionInfo); err != nil {
		t.Error(err)
		t.Fatal("create collection failed")
	}

	test, err := cm.GetCollection("testino_collection")
	if err != nil {
		t.Error(err)
	}
	if test.CollectionName != "testino_collection" ||
		test.Description != "esta es una coleccion para libros de terror" ||
		test.Category != "horror" {
		t.Fatal("Error: the given collection does not have the right field's values'")
	}
	_, err3 := cm.GetCollectionById(uint(test.Id))
	if err3 != nil {
		t.Fatal("Error: problem found getting collection by id")
	}
	collectionUpdates := m.CollectionInfo{
		Description: "nueva descripcion",
		Category:    "terror",
	}
	if err4 := cm.UpdateCollection(uint(test.Id), &collectionUpdates); err4 != nil {
		t.Fatal(err4)
	}
	updatedTest, err5 := cm.GetCollection(test.CollectionName)
	if err5 != nil {
		t.Fatal(err5)
	}

	if updatedTest.Category != "terror" ||
		updatedTest.Description != "nueva descripcion" {
		t.Fatal("the given updated collection does not have the right field's name")
	}

	if err6 := cm.DeleteCollection(uint(updatedTest.Id)); err6 != nil {
		t.Fatal("Error: problem found trying to delete the collection")
	}
	if deleteErr := um.DeleteUser(uint(user.Id)); deleteErr != nil {
		t.Fatal("error deleting user")
	}
}
