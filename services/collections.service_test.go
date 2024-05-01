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

	collectionInfo := m.CollectionInfo{
		CollectionName: "testino_collection",
		Description:    "esta es una coleccion para libros de terror",
		Category:       "horror",
		Public:         true,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	if _, err := cm.CreateCollection(&collectionInfo); err != nil {
		t.Fatal("create collection failed")
	}

	test, err := cm.GetCollection("testino_collection")
	if err != nil {
		t.Fatal("error getting collection by name")
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
		Category:    "horror",
	}
	if err4 := cm.UpdateCollection(uint(test.Id), &collectionUpdates); err4 != nil {
		t.Fatal("Error: problem updating the collection")
	}
	updatedTest, err5 := cm.GetCollection(test.CollectionName)
	if err5 != nil {
		t.Fatal("Error: getting collection by name")
	}

	if updatedTest.Category != "horror" ||
		updatedTest.Description != "nueva descripcion" {
		t.Fatal("the given updated collection does not have the right field's name")
	}

	if err6 := cm.DeleteCollection(uint(updatedTest.Id)); err6 != nil {
		t.Fatal("Error: problem found trying to delete the collection")
	}
}
