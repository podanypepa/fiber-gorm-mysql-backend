package todo

import (
	"time"

	"gorm.io/gorm"
)

// TODO db model
type TODO struct {
	gorm.Model
	ID        string    `json:""`
	CreatedAt time.Time `json:""`
	Done      bool      `json:""`
	Subject   string    `json:""`
	Note      string    `json:""`
}

// Create new TODO in DB
func Create(db *gorm.DB, t *TODO) error {
	return db.Create(t).Error
}

// Read one TODO from DB by ID
func Read(db *gorm.DB, t *TODO, id string) error {
	return db.Where("id = ?", id).First(t).Error
}

// ReadAll TODOs from DB
func ReadAll(db *gorm.DB, t *[]TODO) error {
	return db.Find(t).Error
}

// Update TODO in DB
func Update(db *gorm.DB, t *TODO) error {
	return db.Save(t).Error
}

// Delete TODO from DB
func Delete(db *gorm.DB, t *TODO) error {
	return nil
}

// DeleteByID one TODO by ID
func DeleteByID(db *gorm.DB, id string) error {
	todo := &TODO{}
	if err := Read(db, todo, id); err != nil {
		return err
	}
	return db.Where("id = ?", id).Delete(todo).Error

}
