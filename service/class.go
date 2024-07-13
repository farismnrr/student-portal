/**
 * Package service provides the business logic for managing classes.
 * It interacts with the repository layer to perform operations such as fetching all classes.
 *
 * The ClassService interface defines the methods available for the class service:
 * - FetchAll: Retrieves all classes from the repository.
 *
 * The classService struct implements the ClassService interface and contains:
 * - classRepository: A repository layer object to interact with class data.
 *
 * Functions:
 * - NewClassService: Constructor for creating a new instance of classService.
 * - FetchAll: Implementation to fetch all classes.
 */

package service

import (
	"a21hc3NpZ25tZW50/model"
	"a21hc3NpZ25tZW50/repository"
)

type ClassService interface {
	FetchAll() ([]model.Class, error)
}

type classService struct {
	classRepository repository.ClassRepository
}

func NewClassService(classRepository repository.ClassRepository) ClassService {
	return &classService{classRepository}
}

func (s *classService) FetchAll() ([]model.Class, error) {
	return s.classRepository.FetchAll()
}
