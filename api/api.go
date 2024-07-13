/**
 * This file contains the API handlers for user registration, login, and logout.
 * It interacts with the model package for user data structures and error responses.
 * The Register handler registers a new user, performing validation checks on the provided credentials.
 * The Login handler authenticates a user, creates a session, and sets a cookie for the session.
 * The Logout handler deletes the session and clears the session cookie.
 *
 * Routes/Endpoints:
 * - POST /register: Registers a new user.
 * - POST /login: Authenticates a user and creates a session.
 * - POST /logout: Deletes the session and clears the session cookie.
 */

package api

import (
	"a21hc3NpZ25tZW50/service"
	"fmt"
	"net/http"
)

type API struct {
	userService    service.UserService
	sessionService service.SessionService
	studentService service.StudentService
	classService   service.ClassService
	mux            *http.ServeMux
}

func NewAPI(userService service.UserService, sessionService service.SessionService, studentService service.StudentService, classService service.ClassService) API {
	mux := http.NewServeMux()
	api := API{
		userService,
		sessionService,
		studentService,
		classService,
		mux,
	}

	mux.Handle("/user/register", api.Post(http.HandlerFunc(api.Register)))
	mux.Handle("/user/login", api.Post(http.HandlerFunc(api.Login)))
	mux.Handle("/user/logout", api.Get(api.Auth(http.HandlerFunc(api.Logout))))

	mux.Handle("/student/get-all", api.Get(api.Auth(http.HandlerFunc(api.FetchAllStudent))))
	mux.Handle("/student/get", api.Get(api.Auth(http.HandlerFunc(api.FetchStudentByID))))
	mux.Handle("/student/add", api.Post(api.Auth(http.HandlerFunc(api.Storestudent))))
	mux.Handle("/student/update", api.Put(api.Auth(http.HandlerFunc(api.Updatestudent))))
	mux.Handle("/student/delete", api.Delete(http.HandlerFunc(api.Deletestudent)))
	mux.Handle("/student/get-with-class", api.Get(http.HandlerFunc(api.FetchStudentWithClass)))

	mux.Handle("/class/get-all", api.Get(api.Auth(http.HandlerFunc(api.FetchAllClass))))

	return api
}

func (api *API) Handler() *http.ServeMux {
	return api.mux
}

func (api *API) Start() {
	fmt.Println("starting web server at http://localhost:8080")
	http.ListenAndServe(":8080", api.Handler())
}
