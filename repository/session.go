/**
 * Package repository provides the data access layer for managing user sessions.
 * It interacts with the database layer to perform operations such as adding, updating, and deleting sessions.
 *
 * The SessionsRepository interface defines the methods available for the session repository:
 * - AddSessions: Adds a new session.
 * - DeleteSession: Deletes a session by session token.
 * - UpdateSessions: Updates an existing session.
 * - SessionAvailName: Checks the availability of a session by username.
 * - SessionAvailToken: Retrieves a session by session token.
 *
 * The sessionsRepoImpl struct implements the SessionsRepository interface and contains:
 * - db: A reference to the gorm.DB object for database operations.
 *
 * Functions:
 * - NewSessionRepo: Constructor for creating a new instance of sessionsRepoImpl.
 * - AddSessions: Implementation to add a new session.
 * - DeleteSession: Implementation to delete a session by session token.
 * - UpdateSessions: Implementation to update an existing session.
 * - SessionAvailName: Implementation to check the availability of a session by username.
 * - SessionAvailToken: Implementation to retrieve a session by session token.
 */

package repository

import (
	"a21hc3NpZ25tZW50/model"

	"gorm.io/gorm"
)

type SessionsRepository interface {
	AddSessions(session model.Session) error
	DeleteSession(token string) error
	UpdateSessions(session model.Session) error
	SessionAvailName(name string) error
	SessionAvailToken(token string) (model.Session, error)
}

type sessionsRepoImpl struct {
	db *gorm.DB
}

func NewSessionRepo(db *gorm.DB) *sessionsRepoImpl {
	return &sessionsRepoImpl{db}
}

func (s *sessionsRepoImpl) AddSessions(session model.Session) error {
	return s.db.Create(&session).Error
}

func (s *sessionsRepoImpl) DeleteSession(token string) error {
	return s.db.Where("token = ?", token).Delete(&model.Session{}).Error
}

func (s *sessionsRepoImpl) UpdateSessions(session model.Session) error {
	return s.db.Model(&model.Session{}).Where("username = ?", session.Username).Updates(session).Error
}

func (s *sessionsRepoImpl) SessionAvailName(name string) error {
	session := model.Session{}
	return s.db.Where("username = ?", name).First(&session).Error
}

func (s *sessionsRepoImpl) SessionAvailToken(token string) (model.Session, error) {
	session := model.Session{}
	return session, s.db.Where("token = ?", token).First(&session).Error
}
