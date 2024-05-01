package services

import (
	"errors"
	"fmt"
	m "share-Gutenberg/models"

	"github.com/jmoiron/sqlx"
)

// COLLECTION MANAGER INTERFACE
type CMI interface {
	GetCollection(collectionName string) (*m.Collection, error)
	CreateCollection(*m.CollectionInfo) (*m.Collection, error)
	UpdateCollection(uint, *m.CollectionInfo) error
	DeleteCollection(uint) error
}

type CMT struct {
	DB *sqlx.DB
}

func (cm *CMT) GetCollection(cn string) (*m.Collection, error) {
	var collection m.Collection
	if err := cm.DB.Get(
		&collection,
		"SELECT * FROM collections WHERE collection_name = $1",
		cn); err != nil {
		return nil, err
	}
	return &collection, nil
}
func (cm *CMT) GetCollectionById(id uint) (*m.Collection, error) {
	var collection m.Collection
	if err := cm.DB.Get(
		&collection,
		"SELECT * FROM collections WHERE id = $1",
		id); err != nil {
		return nil, err
	}
	return &collection, nil
}

// TODO
func (cm *CMT) CreateCollection(ci *m.CollectionInfo) (*m.Collection, error) {
	var collection m.Collection
	if err := cm.DB.Get(
		&collection,
		`INSERT INTO collections (collection_name, description, documents, owner_id, category, public, created_at, updated_at)
		VALUES (:collection_name, :description, :documents, :owner_id, :category, :public, :created_at, :updated_at) RETURNING *`,
		&ci,
	); err != nil {
		return nil, err
	}
	return &collection, nil
}
func (cm *CMT) UpdateCollection(id uint, ci *m.CollectionInfo) error {
	if len(*ci.Iter()) == 0 {
		return errors.New("Error: there is no field to be update")
	}
	if _, err := cm.GetCollectionById(id); err != nil {
		return errors.New("Error: the collection couldn't be found")
	}
	updates := "UPDATE Users SET "
	for key, value := range *ci.Iter() {
		updates += fmt.Sprintf("%v = '%v' ,", key, value)
	}
	if _, err2 := cm.DB.Exec(
		fmt.Sprintf("%s WHERE id = $1;", updates[:len(updates)-1]), id); err2 != nil {
		return err2
	}
	return nil
}

func (cm *CMT) DeleteCollection(id uint) error {
	if _, err := cm.DB.Exec("DELETE FROM collections WHERE id = $1", id); err != nil {
		return err
	}
	return nil
}
