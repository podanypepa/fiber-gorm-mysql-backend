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
	return nil
}

// Read one TODO from DB by ID
func Read(db *gorm.DB, t *TODO, id int) error {
	return nil
}

// ReadAll TODOs from DB
func ReadAll(db *gorm.DB, t *TODO) error {
	return nil
}

// Update TODO in DB
func Update(db *gorm.DB, t *TODO) error {
	return nil
}

// Delete TODO from DB
func Delete(db *gorm.DB, t *TODO) error {
	return nil
}
