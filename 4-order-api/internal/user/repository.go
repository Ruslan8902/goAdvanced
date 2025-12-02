package user

import (
	"orderApiStart/pkg/db"

	"gorm.io/gorm/clause"
)

type UserRepository struct {
	Database *db.Db
}

func NewUserRepository(database *db.Db) *UserRepository {
	return &UserRepository{
		Database: database,
	}
}

func (repo *UserRepository) Create(user *User) (*User, error) {
	result := repo.Database.DB.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (repo *UserRepository) Update(user *User) (*User, error) {
	result := repo.Database.DB.Clauses(clause.Returning{}).Updates(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (repo *UserRepository) Delete(id uint) error {
	result := repo.Database.DB.Delete(&User{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *UserRepository) GetById(id uint) (*User, error) {
	var user User
	result := repo.Database.DB.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (repo *UserRepository) GetByPhone(phone string) (*User, error) {
	var user User
	result := repo.Database.DB.First(&user, "Phone = ?", phone)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
