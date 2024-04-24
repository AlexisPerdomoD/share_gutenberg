package services

import (
	m "share-Gutenberg/models"

	"github.com/jmoiron/sqlx"
)

// USER MANAGER INTERFACE
type UMI interface {
	GetUser(uint) (*m.User, *m.Err)
	DeleteUser(uint) *m.Err
	UpdateUser(uint, *m.UserInfo) (*m.User, *m.Err)
	CreateUser(*m.UserInfo)
}

// USER MANAGER TYPE
type UMT struct {
	DB *sqlx.DB
}

func (um *UMT) GetUser(id uint) (*m.User, *m.Err) {
	var user m.User
	if err := um.DB.Get(&user, "SELECT * FROM Users WHERE id = $1;", id); err != nil {
		return nil, &m.Err{
			Error:   err,
			Status:  1,
			Message: "",
		}
	}
	return &user, nil
}
