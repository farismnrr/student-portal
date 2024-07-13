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
	"a21hc3NpZ25tZW50/model"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (api *API) Register(c *gin.Context) {
	var creds model.User
	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Internal Server Error"})
		return
	}

	if creds.Username == "" || creds.Password == "" {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Bad Request"})
		return
	}

	if api.userService.CheckPassLength(creds.Password) {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Please provide a password of more than 5 characters"})
		return
	}

	if api.userService.CheckPassAlphabet(creds.Password) {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Please use Password with Contains non Alphabetic Characters"})
		return
	}

	if err := api.userService.Register(creds); err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse{Message: "User Registered"})
}

func (api *API) Login(c *gin.Context) {
	var creds model.User

	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Internal Server Error"})
		return
	}

	if creds.Username == "" || creds.Password == "" {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Bad Request"})
		return
	}

	if api.userService.CheckPassLength(creds.Password) {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Please provide a password of more than 5 characters"})
		return
	}

	if api.userService.CheckPassAlphabet(creds.Password) {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Please use Password with Contains non Alphabetic Characters"})
		return
	}

	if err := api.userService.Login(creds); err != nil {
		c.JSON(http.StatusUnauthorized, model.ErrorResponse{Error: "Wrong User or Password!"})
		return
	}

	sessionToken := uuid.NewString()
	expiresAt := time.Now().Add(5 * time.Hour)
	session := model.Session{Token: sessionToken, Username: creds.Username, Expiry: expiresAt}

	if err := api.sessionService.SessionAvailName(session.Username); err != nil {
		err = api.sessionService.AddSession(session)
	} else {
		err = api.sessionService.UpdateSession(session)
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: "Internal Server Error"})
		return
	}

	c.SetCookie("session_token", sessionToken, 3600*5, "/", "", false, true)

	c.JSON(http.StatusOK, model.SuccessResponse{Message: "Login Success"})
}

func (api *API) Logout(c *gin.Context) {
	sessionToken, err := c.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			c.JSON(http.StatusUnauthorized, model.ErrorResponse{Error: "Internal Server Error"})
			return
		}
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Internal Server Error"})
		return
	}

	api.sessionService.DeleteSession(sessionToken)

	c.SetCookie("session_token", "", -1, "/", "", false, true)

	c.JSON(http.StatusOK, model.SuccessResponse{Message: "Logout Success"})
}
