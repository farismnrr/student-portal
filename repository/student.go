/**
 * Package repository provides the data access layer for managing student entities.
 * It interacts with the database layer to perform operations such as fetching all students, fetching a student by ID, storing a student, updating a student, deleting a student, and fetching students with their classes.
 *
 * The StudentRepository interface defines the methods available for the student repository:
 * - FetchAll: Retrieves all students from the database.
 * - FetchByID: Retrieves a student by their ID.
 * - Store: Stores a new student in the database.
 * - Update: Updates an existing student in the database.
 * - Delete: Deletes a student from the database.
 * - FetchWithClass: Retrieves students with their associated classes.
 *
 * The studentRepoImpl struct implements the StudentRepository interface and contains:
 * - db: A reference to the gorm.DB object for database operations.
 *
 * Functions:
 * - NewStudentRepo: Constructor for creating a new instance of studentRepoImpl.
 * - FetchAll: Implementation to fetch all students.
 * - FetchByID: Implementation to fetch a student by ID.
 * - Store: Implementation to store a new student.
 * - Update: Implementation to update an existing student.
 * - Delete: Implementation to delete a student.
 * - FetchWithClass: Implementation to fetch students with their classes.
 */

package repository

import (
	"a21hc3NpZ25tZW50/model"

	"gorm.io/gorm"
)

type StudentRepository interface {
	FetchAll() ([]model.Student, error)
	FetchByID(id int) (*model.Student, error)
	Store(s *model.Student) error
	Update(id int, s *model.Student) error
	Delete(id int) error
	FetchWithClass() (*[]model.StudentClass, error)
}

type studentRepoImpl struct {
	db *gorm.DB
}

func NewStudentRepo(db *gorm.DB) *studentRepoImpl {
	return &studentRepoImpl{db}
}

func (s *studentRepoImpl) FetchAll() ([]model.Student, error) {
	students := []model.Student{}
	return students, s.db.Find(&students).Error
}

func (s *studentRepoImpl) Store(student *model.Student) error {
	return s.db.Create(student).Error
}

func (s *studentRepoImpl) Update(id int, student *model.Student) error {
	students := model.Student{}
	return s.db.Model(&students).Where("id = ?", id).Updates(student).Error
}

func (s *studentRepoImpl) Delete(id int) error {
	student := model.Student{}
	return s.db.Where("id = ?", id).Delete(&student).Error
}

func (s *studentRepoImpl) FetchByID(id int) (*model.Student, error) {
	student := model.Student{}
	return &student, s.db.First(&student, id).Error
}

func (s *studentRepoImpl) FetchWithClass() (*[]model.StudentClass, error) {
	studentClass := []model.StudentClass{}
	return &studentClass, s.db.Table("students").
		Select("students.name, students.address, classes.name as class_name, classes.professor, classes.room_number").
		Joins("left join classes on students.class_id = classes.id").
		Scan(&studentClass).Error
}
