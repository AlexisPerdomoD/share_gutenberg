package services

import (
	"errors"
	"fmt"
	m "share-Gutenberg/models"

	"github.com/jmoiron/sqlx"
)

// USER MANAGER INTERFACE
type UMI interface {
	GetUser(uint) (*m.User, *m.Err)
	GetUserByEmail(string) (*m.User, *m.Err)
	DeleteUser(uint) *m.Err
	UpdateUser(int, *m.UserInfo) (*m.User, *m.Err)
	CreateUser(*m.UserInfo) error
	//TODO
	AddCollection(uint, uint) error
	RemoveCollection(uint, uint) error
	GetUserCollections(uint) (*m.UserCollections, *m.Err)
}

// USER MANAGER TYPE
type UMT struct {
	DB *sqlx.DB
}

// same as GetUser but reciving email as argument
func (um *UMT) GetUserByEmail(e string) (*m.User, *m.Err) {
	var user m.User
	if err := um.DB.Get(
		&user,
		"SELECT * FROM Users WHERE email = $1;",
		e); err != nil {
		return nil, &m.Err{
			Error:   err,
			Status:  000,
			Message: err.Error(),
		}
	}
	return &user, nil
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
	_, err := um.DB.Exec(`
		INSERT INTO Users(name, email, password, role, created_at, updated_at) 
		VALUES($1,$2, $3, $4, $5, $6);`,
		uf.Name, uf.Email, uf.Password, uf.Role, uf.CreatedAt, uf.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

// recives as argument an userinfo type with the changes and the email that belongs the row wanted to be changed
func (um UMT) UpdateUser(id uint, uf *m.UserInfo) error {
	if len(*uf.Iter()) == 0 {
		return errors.New("there is no field to be update")
	}
	if _, err := um.GetUser(id); err != nil {
		return err.Error
	}

	updates := "UPDATE Users SET "
	for key, value := range *uf.Iter() {
		updates += fmt.Sprintf("%v = '%v' ,", key, value)
	}
	if _, err2 := um.DB.Exec(
		fmt.Sprintf("%s WHERE id = $1;", updates[:len(updates)-1]), id); err2 != nil {
		return err2
	}

	return nil
}

func (um *UMT) DeleteUser(id uint) error {
	if _, err := um.DB.Exec("DELETE FROM users WHERE id = $1", id); err != nil {
		return err
	}
	return nil
}
func (um UMT) AddCollection(user uint, collection uint) error {
	if _, err := um.DB.Exec(`INSERT INTO user_collection (user_id, collection_id)
	VALUES($1, $2)`, user, collection); err != nil {
		return err
	}
	return nil
}
func (um UMT) RemoveCollection(user uint, collection uint) error {
	if _, err := um.DB.Exec(`DELETE FROM user_collection
	WHERE user_id = $1 AND collection_id = $2`,
		user, collection); err != nil {
		return err
	}
	return nil
}
func (um UMT) GetUserCollections(user uint) (*m.UserCollections, error) {
	var userCollections m.UserCollections
	if err := um.DB.Select(&userCollections.Added,
		`SELECT c.* FROM collections as c
		JOIN user_collection as uc ON
		c.id = uc.collection_id
		WHERE uc.user_id = $1`, user); err != nil {
		return nil, err
	}
	if err := um.DB.Select(&userCollections.Owned,
		`SELECT * FROM collections WHERE owner_id = $1`, user,
	); err != nil {
		return nil, err
	}
	return &userCollections, nil
}
