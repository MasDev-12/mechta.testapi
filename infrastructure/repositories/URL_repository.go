package repositories

import (
	"fmt"
	"github.com/MasDev-12/mechta.testapi/domain/entities"
	"github.com/MasDev-12/mechta.testapi/infrastructure/db_context"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type URLRepository struct {
	dbContext *db_context.DbContext
}

func NewURLRepository(dbContext *db_context.DbContext) *URLRepository {
	return &URLRepository{dbContext: dbContext}
}

func (r *URLRepository) GetAll() ([]entities.URL, error) {
	var urls []entities.URL
	result := r.dbContext.Db.Find(&urls)
	if result.Error != nil {
		if result.RowsAffected == 0 {
			return nil, fmt.Errorf("no urls found")
		}
		return nil, result.Error
	}
	return urls, nil
}

func (r *URLRepository) GetById(id uuid.UUID) (*entities.URL, error) {
	var url entities.URL
	result := r.dbContext.Db.First(&url, id)
	if result.Error != nil {
		if result.RowsAffected == 0 {
			return nil, fmt.Errorf("user not found")
		}
		return &entities.URL{}, result.Error
	}
	return &url, nil
}

func (r *URLRepository) Add(URL entities.URL) (entities.URL, error) {
	result := r.dbContext.Db.Create(&URL)
	if result.Error != nil {
		return entities.URL{}, result.Error
	}
	return URL, nil
}

func (r *URLRepository) Update(URL entities.URL) (bool, error) {
	result := r.dbContext.Db.Save(&URL)
	if result.Error != nil {
		return false, result.Error
	}
	return result.RowsAffected > 0, nil
}

func (r *URLRepository) Delete(id uuid.UUID) (bool, error) {
	result := r.dbContext.Db.Where("id = ?", id).Delete(&entities.URL{})
	if result.Error != nil {
		return false, result.Error
	}
	return result.RowsAffected > 0, nil
}

func (r *URLRepository) GetUserUrls(userId uuid.UUID) ([]entities.URL, error) {
	var urls []entities.URL
	result := r.dbContext.Db.Where("user_id = ?", userId).Find(&urls)
	if result.Error != nil {
		if result.RowsAffected == 0 {
			return nil, fmt.Errorf("url not found")
		}
		return urls, result.Error
	}
	return urls, nil
}

func (r *URLRepository) GetUrlByShortName(shortName string) (*entities.URL, error) {
	var url entities.URL

	result := r.dbContext.Db.Where("short_url = ?", shortName).First(&url)
	if result.Error != nil {
		if result.RowsAffected == 0 {
			return nil, fmt.Errorf("url not found")
		}
		return nil, result.Error
	}

	updateResult := r.dbContext.Db.Model(&url).UpdateColumn("click_count", gorm.Expr("click_count + ?", 1))
	if updateResult.Error != nil {
		return nil, updateResult.Error
	}

	// Вернем обновленный объект URL
	return &url, nil
}

func (r *URLRepository) DeleteUrlByShortName(shortName string) (bool, error) {
	result := r.dbContext.Db.Where("short_url = ?", shortName).Delete(&entities.URL{})
	if result.Error != nil {
		return false, result.Error
	}
	return result.RowsAffected > 0, nil
}

func (r *URLRepository) GetUrlByOriginalName(originalName string) (*entities.URL, error) {
	var url entities.URL

	result := r.dbContext.Db.Where("origin_url = ?", originalName).First(&url)
	if result.Error != nil {
		if result.RowsAffected == 0 {
			return nil, fmt.Errorf("url not found")
		}
		return nil, result.Error
	}
	return &url, nil
}
