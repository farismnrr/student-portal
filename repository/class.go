/**
 * Package repository provides the data access layer for managing class entities.
 * It interacts with the database layer to perform operations such as fetching all classes.
 *
 * The ClassRepository interface defines the methods available for the class repository:
 * - FetchAll: Retrieves all classes from the database.
 *
 * The classRepoImpl struct implements the ClassRepository interface and contains:
 * - db: A reference to the gorm.DB object for database operations.
 *
 * Functions:
 * - NewClassRepo: Constructor for creating a new instance of classRepoImpl.
 * - FetchAll: Implementation to fetch all classes.
 */

package repository

import (
	"a21hc3NpZ25tZW50/model"

	"gorm.io/gorm"
)

type ClassRepository interface {
	FetchAll() ([]model.Class, error)
}

type classRepoImpl struct {
	db *gorm.DB
}

func NewClassRepo(db *gorm.DB) *classRepoImpl {
	return &classRepoImpl{db}
}

func (s *classRepoImpl) FetchAll() ([]model.Class, error) {
	var classes []model.Class
	return classes, s.db.Find(&classes).Error
}
