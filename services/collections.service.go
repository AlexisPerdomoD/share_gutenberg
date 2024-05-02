package services

import (
	"context"
	"errors"
	"fmt"
	m "share-Gutenberg/models"
	u "share-Gutenberg/utils"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

// COLLECTION MANAGER INTERFACE
type CMI interface {
	GetCollection(collectionName string) (*m.Collection, error)
	CreateCollection(*m.CollectionInfo) error
	UpdateCollection(uint, *m.CollectionInfo) error
	DeleteCollection(uint) error
	DeleteBookToCollection(uint, uint) error
	AddBookToCollection(uint, uint) error
}

type CMT struct {
	DB *sqlx.DB
}

func (cm *CMT) GetCollection(cn string) (*m.Collection, error) {
	var collection m.Collection
	var documents interface{}
	//aqui se usa el metodo Scan personalizado
	if err := cm.DB.QueryRowx(
		"SELECT id, collection_name, description, category, documents, public, owner_id, created_at, updated_at FROM collections WHERE collection_name = $1;",
		cn).Scan(&collection.Id, &collection.CollectionName, &collection.Description, &collection.Category, &documents, &collection.Public, &collection.Owner, &collection.CreatedAt, &collection.UpdatedAt); err != nil {
		return nil, err
	}
	if errDoc := collection.ScanDocuments(documents); errDoc != nil {
		return nil, errDoc
	}
	return &collection, nil
}
func (cm *CMT) GetCollectionById(id uint) (*m.Collection, error) {
	var collection m.Collection
	var documents interface{}
	//aqui se usa el metodo Scan personalizado
	if err := cm.DB.QueryRowx(
		"SELECT id, collection_name, description, category, documents, public, owner_id, created_at, updated_at FROM collections WHERE id = $1;",
		id).Scan(&collection.Id, &collection.CollectionName, &collection.Description, &collection.Category, &documents, &collection.Public, &collection.Owner, &collection.CreatedAt, &collection.UpdatedAt); err != nil {
		return nil, err
	}
	if errDoc := collection.ScanDocuments(documents); errDoc != nil {
		return nil, errDoc
	}
	return &collection, nil
}

// TO DO
func (cm *CMT) CreateCollection(ci *m.CollectionInfo) error {
	_, err := cm.DB.Exec(
		`INSERT INTO collections (collection_name, description, documents, owner_id, category, public, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`,
		ci.CollectionName, ci.Description, pq.Array(ci.Documents), ci.Owner, ci.Category, ci.Public, ci.CreatedAt, ci.UpdatedAt,
	)
	if err != nil {
		return err
	}
	return nil
}

func (cm *CMT) UpdateCollection(id uint, ci *m.CollectionInfo) error {
	if len(*ci.Iter()) == 0 {
		return errors.New("Error: there is no field to be update")
	}
	if _, err := cm.GetCollectionById(id); err != nil {
		return errors.New("Error: the collection couldn't be found")
	}
	updates := "UPDATE collections SET "
	for key, value := range *ci.Iter() {
		updates += fmt.Sprintf("%s = '%s' ,", key, value)
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

func (cm *CMT) AddBookToCollection(book int, cID int) error {
	collection, err := cm.GetCollectionById(uint(cID))
	if err != nil {
		return err
	}
	indexCh := make(chan int)
	finding, cancel := context.WithCancel(context.Background())
	for idx, curr := range collection.Documents {
		go u.Find(finding, indexCh, idx, curr, book)
	}
	for {
		idx, ok := <-indexCh
		//the chanel was close due context cancel or there is no more values to check
		if !ok {
			cancel()
			break
		}
		//if there was found a value means the book is already in the documents collection
		if idx != -1 {
			cancel()
			return errors.New("book already in the collection")
		}
	}
	collection.Documents = append(collection.Documents, book)
	if _, errUpdate := cm.DB.Exec(
		`UPDATE collections SET documents = $1`,
		pq.Array(collection.Documents),
	); errUpdate != nil {
		return errUpdate
	}
	return nil
}
func (cm *CMT) DeleteBookToCollection(book int, cID int) error {
	collection, err := cm.GetCollectionById(uint(cID))
	if err != nil {
		return err
	}
	indexCh := make(chan int)
	finding, cancel := context.WithCancel(context.Background())
	for idx, curr := range collection.Documents {
		go u.Find(finding, indexCh, idx, curr, book)
	}
	var idx int
	for {
		val, ok := <-indexCh
		if !ok {
			cancel()
			return errors.New("book not found")
		}
		if idx != -1 {
			cancel()
			idx = val
			break
		}
	}
	collection.Documents = append(collection.Documents[:idx], collection.Documents[idx:]...)
	if _, errUpdate := cm.DB.Exec(
		`UPDATE collections SET documents = $1`,
		pq.Array(collection.Documents),
	); errUpdate != nil {
		return errUpdate
	}
	return nil
}
