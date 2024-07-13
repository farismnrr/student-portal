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
	Port                    string
	PortAlternative         string
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
