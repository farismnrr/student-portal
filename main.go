/**
 * This Go program is designed to manage an educational system where it handles database connections,
 * migrations, and API server setup for handling various educational entities such as users, sessions,
 * students, and classes.
 *
 * Key components:
 * - `db`: Manages database connections and configurations.
 * - `model`: Defines the data models used in the application.
 * - `repository`: Provides access to the database for CRUD operations on models.
 * - `service`: Contains business logic and interacts with repositories.
 * - `api`: Handles HTTP requests and routes them to the appropriate services.
 *
 * The program starts by loading environment variables, setting up a database connection using credentials
 * that can be configured via environment variables or default values. It then connects to the database
 * and performs auto-migration to ensure the database schema is set up correctly.
 *
 * Default classes are created and stored in the database. Repositories for users, sessions, students,
 * and classes are initialized and used to create corresponding services. These services are then used
 * to set up the API routes.
 *
 * The API routes are as follows:
 * - `/user/register`: Handles user registration.
 * - `/user/login`: Handles user login.
 * - `/user/logout`: Handles user logout.
 * - `/student/get-all`: Fetches all students.
 * - `/student/get`: Fetches a single student by ID.
 * - `/student/add`: Adds a new student.
 * - `/student/update`: Updates an existing student.
 * - `/student/delete`: Deletes a student.
 * - `/student/get-with-class`: Fetches students along with their class information.
 * - `/class/get-all`: Fetches all classes.
 *
 * Finally, the API server is started to listen for incoming HTTP requests on port 8080.
 */

package main

import (
	"a21hc3NpZ25tZW50/api"
	"a21hc3NpZ25tZW50/db"
	"a21hc3NpZ25tZW50/model"
	repo "a21hc3NpZ25tZW50/repository"
	"a21hc3NpZ25tZW50/service"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	db := db.NewDB()
	
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		port = 5432
	}

	dbCredential := model.Credential{
		Host:                    "localhost",
		HostAlternative:         os.Getenv("DB_HOST"),
		Username:                "postgres",
		UsernameAlternative:     os.Getenv("DB_USERNAME"),
		Password:                "postgres",
		PasswordAlternative:     os.Getenv("DB_PASSWORD"),
		DatabaseName:            "kampusmerdeka",
		DatabaseNameAlternative: os.Getenv("DB_NAME"),
		Port:                    5432,
		PortAlternative:         port,
		Schema:                  "public",
	}

	conn, err := db.Connect(&dbCredential)
	if err != nil {
		panic(err)
	}

	conn.AutoMigrate(&model.User{}, &model.Session{}, &model.Student{}, &model.Class{})

	classes := []model.Class{
		{
			Name:       "Mathematics",
			Professor:  "Dr. Smith",
			RoomNumber: 101,
		},
		{
			Name:       "Physics",
			Professor:  "Dr. Johnson",
			RoomNumber: 102,
		},
		{
			Name:       "Chemistry",
			Professor:  "Dr. Lee",
			RoomNumber: 103,
		},
	}

	for _, c := range classes {
		if err := conn.Create(&c).Error; err != nil {
			panic("failed to create default classes")
		}
	}

	userRepo := repo.NewUserRepo(conn)
	sessionRepo := repo.NewSessionRepo(conn)
	studentRepo := repo.NewStudentRepo(conn)
	classRepo := repo.NewClassRepo(conn)

	userService := service.NewUserService(userRepo)
	sessionService := service.NewSessionService(sessionRepo)
	studentService := service.NewStudentService(studentRepo)
	classService := service.NewClassService(classRepo)

	mainAPI := api.NewAPI(userService, sessionService, studentService, classService)
	mainAPI.Start()
}
