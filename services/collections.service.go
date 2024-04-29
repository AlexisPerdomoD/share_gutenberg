package services

import (
	m "share-Gutenberg/models"

	"github.com/jmoiron/sqlx"
)

// COLLECTION MANAGER INTERFACE
type CMI interface {
	GetCollection(collectionName string) (*m.Collection, error)
	CreateCollection(*m.CollectionInfo) error
	UpdateCollection(uint, *m.CollectionInfo) error
	DeleteCollection(uint) error
}

type CMT struct {
	DB *sqlx.DB
}

// TODO
func (cm *CMT) GetCollection(cn string) (m.Collection, error)
func (cm *CMT) CreateCollection(ci *m.CollectionInfo)
func (cm *CMT) UpdateCollection(id uint, ci *m.CollectionInfo)
func (cm *CMT) DeleteCollection(id uint)
