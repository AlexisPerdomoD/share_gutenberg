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
	CreateUser(*m.UserInfo) (*m.User, *m.Err)
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
			Status:  500,
			Message: err.Error(),
		}
	}
	return &user, nil
}

// Named queries can use structs, so if you have an existing struct (i.e. person := &Person{}) that you have populated, you can pass it in as &person
func (um UMT) CreateUser(uf *m.UserInfo) error {

	_, err := um.DB.NamedExec("INSERT INTO Users(name, email, password, role)values(:name, :email, :password, :role)", uf)
	if err != nil {
		return err
	}
	return nil
}
