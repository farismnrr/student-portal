/**
 * This file contains API handlers for managing student data.
 * It interacts with the model package for student data structures and error responses.
 * The FetchAllStudent handler fetches all students from the database.
 * The FetchStudentByID handler fetches a student by ID from the database.
 * The StoreStudent handler stores a new student in the database.
 * The UpdateStudent handler updates an existing student in the database.
 * The DeleteStudent handler deletes a student by ID from the database.
 * The FetchStudentWithClass handler fetches students with their associated classes.
 *
 * Routes/Endpoints:
 * - GET /students: Fetches all students.
 * - GET /student?id={id}: Fetches a student by ID.
 * - POST /student: Stores a new student.
 * - PUT /student?id={id}: Updates an existing student.
 * - DELETE /student?id={id}: Deletes a student by ID.
 * - GET /students-with-class: Fetches students with their associated classes.
 */

package api

import (
	"a21hc3NpZ25tZW50/model"
	"encoding/json"
	"net/http"
	"strconv"
)

func (api *API) FetchAllStudent(w http.ResponseWriter, r *http.Request) {
	student, err := api.studentService.FetchAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(student)
}

func (api *API) FetchStudentByID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	student, err := api.studentService.FetchByID(idInt)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(student)
}

func (api *API) Storestudent(w http.ResponseWriter, r *http.Request) {
	var student model.Student

	err := json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: err.Error()})
		return
	}

	err = api.studentService.Store(&student)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(student)
}

func (api *API) Updatestudent(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var student model.Student
	err = json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: err.Error()})
		return
	}

	err = api.studentService.Update(idInt, &student)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(student)
}

func (api *API) Deletestudent(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = api.studentService.Delete(idInt)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model.SuccessResponse{Message: "student berhasil dihapus"})
}

func (api *API) FetchStudentWithClass(w http.ResponseWriter, r *http.Request) {
	studentClasses, err := api.studentService.FetchWithClass()
	if err != nil {
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(studentClasses)
}
