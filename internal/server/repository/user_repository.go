package repository

import (
	"github.com/titoyudha/task_manager_api/internal/server/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	Register(user models.User) models.User
	GetByID(id int64) models.User
	GetByEmail(email string) models.User
	Update(user models.User) models.User
	Delete(id int64) error
}

type dbConn struct {
	connection *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &dbConn{
		connection: db,
	}
}

func (db *dbConn) Register(user models.User) models.User {
	err := db.connection.Save(&user)
	if err != nil {
		panic(err)
	}

}

func (db *dbConn) GetByID(id int64) models.User {

}

func (db *dbConn) GetByEmail(email string) models.User {

}

func (db *dbConn) Update(user models.User) models.User {

}

func (db *dbConn) Delete(id int64) error {

}
