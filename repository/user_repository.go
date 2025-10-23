package repository

import (
	"github.com/qobilovvv/patterns/service_repo/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindAll() ([]models.User, error)
	FindByID(ID uint) (*models.User, error)
	Save(user *models.User) error
}

type userRepositoryDB struct {
	db *gorm.DB
}

func NewUserRepositoryDB(db *gorm.DB) UserRepository {
	return &userRepositoryDB{db: db}
}

func (r *userRepositoryDB) FindAll() ([]models.User, error) {
	var users []models.User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *userRepositoryDB) FindByID(ID uint) (*models.User, error) {
	var user models.User
	err := r.db.Find(&user, ID).Error
	if err != nil {
		return nil, err
	}
	return &user, err
}

func (r *userRepositoryDB) Save(user *models.User) error {
	return r.db.Create(&user).Error
}
