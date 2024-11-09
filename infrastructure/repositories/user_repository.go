package repositories

import (
	"errors"
	"fmt"
	"github.com/MasDev-12/mechta.testapi/domain/entities"
	"github.com/MasDev-12/mechta.testapi/infrastructure/db_context"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository struct {
	dbContext *db_context.DbContext
}

func NewUserRepository(dbContext *db_context.DbContext) *UserRepository {
	return &UserRepository{dbContext: dbContext}
}

func (r *UserRepository) GetAll() ([]entities.User, error) {
	var users []entities.User
	result := r.dbContext.Db.Preload("Accounts").Find(&users)
	if result.Error != nil {
		if result.RowsAffected == 0 {
			return nil, fmt.Errorf("no users found")
		}
		return nil, result.Error
	}
	return users, nil
}

func (r *UserRepository) GetById(id uuid.UUID) (*entities.User, error) {
	var user entities.User
	result := r.dbContext.Db.First(&user, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &user, nil
}

func (r *UserRepository) Add(user entities.User) (*entities.User, error) {
	result := r.dbContext.Db.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *UserRepository) Update(user entities.User) (bool, error) {
	result := r.dbContext.Db.Save(&user)
	if result.Error != nil {
		return false, result.Error
	}
	return result.RowsAffected > 0, nil
}

func (r *UserRepository) Delete(id uuid.UUID) (bool, error) {
	result := r.dbContext.Db.Where("id = ?", id).Delete(&entities.User{})
	if result.Error != nil {
		return false, result.Error
	}
	return result.RowsAffected > 0, nil
}

func (r *UserRepository) GetUserByEmail(email string) (*entities.User, error) {
	var user entities.User
	result := r.dbContext.Db.Where("email = ?", email).Find(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &user, nil
}
