/**
 * Package model provides the data structures for user, session, student, class, student-class relationship, and credentials.
 *
 * Structs:
 * - User: Represents a user with fields for username and password.
 * - Session: Represents a session with fields for token, username, and expiry time.
 * - Student: Represents a student with fields for name, address, and class ID.
 * - Class: Represents a class with fields for ID, name, professor, and room number.
 * - StudentClass: Represents the relationship between students and classes with fields for name, address, class name, professor, and room number.
 * - Credential: Represents credentials with fields for host, alternative host, username, alternative username, password, alternative password, database name, alternative database name, port, alternative port, schema, and alternative schema.
 */

package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(100);unique"`
	Password string `json:"password"`
}
type Session struct {
	gorm.Model
	Token    string    `json:"token"`
	Username string    `json:"username"`
	Expiry   time.Time `json:"expiry"`
}

type Student struct {
	gorm.Model
	Name    string `json:"name"`
	Address string `json:"address"`
	ClassId int    `json:"class_id"`
}

type Class struct {
	ID         int    `gorm:"primaryKey"`
	Name       string `json:"name"`
	Professor  string `json:"professor"`
	RoomNumber int    `json:"room_number"`
}

type StudentClass struct {
	Name       string `json:"name"`
	Address    string `json:"address"`
	ClassName  string `json:"class_name"`
	Professor  string `json:"professor"`
	RoomNumber int    `json:"room_number"`
}

type Credential struct {
	Host                    string
	HostAlternative         string
	Username                string
	UsernameAlternative     string
	Password                string
	PasswordAlternative     string
	DatabaseName            string
	DatabaseNameAlternative string
	Port                    int
	PortAlternative         int
	Schema                  string
	SchemaAlternative       string
}


type ErrorResponse struct {
	Error string `json:"error"`
}

type SuccessResponse struct {
	Username string `json:"username"`
	Message  string `json:"message"`
}
