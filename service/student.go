/**
 * Package service provides the business logic for the student entities.
 * It interacts with the repository layer to perform CRUD operations and other business logic.
 *
 * The StudentService interface defines the methods available for the student service:
 * - FetchAll: Retrieves all students from the repository.
 * - FetchByID: Retrieves a specific student by their ID.
 * - Store: Adds a new student to the repository.
 * - Update: Updates an existing student's information.
 * - Delete: Removes a student from the repository by their ID.
 * - FetchWithClass: Retrieves students along with their associated class information.
 *
 * The studentService struct implements the StudentService interface and contains:
 * - studentRepository: A repository layer object to interact with the student data.
 *
 * Functions:
 * - NewStudentService: Constructor for creating a new instance of studentService.
 * - FetchAll: Implementation to fetch all students.
 * - FetchByID: Implementation to fetch a specific student by ID.
 * - Store: Implementation to add a new student.
 * - Update: Implementation to update an existing student.
 * - Delete: Implementation to delete a student by ID.
 * - FetchWithClass: Implementation to fetch students with their class information.
 */

package service

import (
	"a21hc3NpZ25tZW50/model"
	"a21hc3NpZ25tZW50/repository"
)

type StudentService interface {
	FetchAll() ([]model.Student, error)
	FetchByID(id int) (*model.Student, error)
	Store(s *model.Student) error
	Update(id int, s *model.Student) error
	Delete(id int) error
	FetchWithClass() (*[]model.StudentClass, error)
}

type studentService struct {
	studentRepository repository.StudentRepository
}

func NewStudentService(studentRepository repository.StudentRepository) StudentService {
	return &studentService{studentRepository}
}

func (s *studentService) FetchAll() ([]model.Student, error) {
	return s.studentRepository.FetchAll()
}

func (s *studentService) FetchByID(id int) (*model.Student, error) {
	return s.studentRepository.FetchByID(id)
}

func (s *studentService) Store(student *model.Student) error {
	return s.studentRepository.Store(student)
}

func (s *studentService) Update(id int, student *model.Student) error {
	return s.studentRepository.Update(id, student)
}

func (s *studentService) Delete(id int) error {
	return s.studentRepository.Delete(id)
}

func (s *studentService) FetchWithClass() (*[]model.StudentClass, error) {
	return s.studentRepository.FetchWithClass()
}

