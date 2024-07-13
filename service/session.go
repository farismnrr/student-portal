/**
 * Package service provides the business logic for managing user sessions.
 * It interacts with the repository layer to perform operations such as adding, updating, and deleting sessions.
 *
 * The SessionService interface defines the methods available for the session service:
 * - AddSession: Adds a new session.
 * - UpdateSession: Updates an existing session.
 * - DeleteSession: Deletes a session by session token.
 * - SessionAvailName: Checks the availability of a session by username.
 * - TokenExpired: Checks if a session token has expired.
 * - TokenValidity: Validates the token and returns the session information.
 *
 * The sessionService struct implements the SessionService interface and contains:
 * - sessionRepository: A repository layer object to interact with session data.
 *
 * Functions:
 * - NewSessionService: Constructor for creating a new instance of sessionService.
 * - AddSession: Implementation to add a new session.
 * - UpdateSession: Implementation to update an existing session.
 * - DeleteSession: Implementation to delete a session by session token.
 * - SessionAvailName: Implementation to check the availability of a session by username.
 * - TokenExpired: Implementation to check if a session token has expired.
 * - TokenValidity: Implementation to validate the token and return session information.
 */

package service

import (
	"a21hc3NpZ25tZW50/model"
	"a21hc3NpZ25tZW50/repository"
	"fmt"
	"time"
)

type SessionService interface {
	AddSession(session model.Session) error
	UpdateSession(session model.Session) error
	DeleteSession(sessionToken string) error
	SessionAvailName(username string) error
	TokenExpired(session model.Session) bool
	TokenValidity(token string) (model.Session, error)
}

type sessionService struct {
	sessionRepository repository.SessionsRepository
}

func NewSessionService(sessionRepository repository.SessionsRepository) SessionService {
	return &sessionService{sessionRepository}
}

func (s *sessionService) SessionAvailName(username string) error {
	return s.sessionRepository.SessionAvailName(username)
}

func (s *sessionService) AddSession(session model.Session) error {
	return s.sessionRepository.AddSessions(session)
}

func (s *sessionService) UpdateSession(session model.Session) error {
	return s.sessionRepository.UpdateSessions(session)
}

func (s *sessionService) DeleteSession(sessionToken string) error {
	return s.sessionRepository.DeleteSession(sessionToken)
}

func (s *sessionService) TokenValidity(token string) (model.Session, error) {
	session, err := s.sessionRepository.SessionAvailToken(token)
	if err != nil {
		return model.Session{}, err
	} else if s.TokenExpired(session) {
		return model.Session{}, fmt.Errorf("%s", "Token is Expired!")
	}
	return session, nil
}

func (s *sessionService) TokenExpired(session model.Session) bool {
	return session.Expiry.Before(time.Now())
}
