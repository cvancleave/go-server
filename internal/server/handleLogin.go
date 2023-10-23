package server

import (
	"go-server/internal/utils"
	"net/http"

	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
)

// handleLogin - handles a login request
func (s *server) handleLogin(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	utils.SetCorsHeaders(w, r)

	log.Infof("login attempt by %s", r.RemoteAddr)

	email := r.FormValue("email")
	pass := r.FormValue("password")

	if email == "" || pass == "" {
		log.Errorf("error with login form: missing values")
		utils.RespondError(w, 400, "error with login form")
		return
	}

	// get user from database
	dbUser, err := s.database.GetUserByEmail(email)
	if err != nil {
		log.Errorf("error getting user from database: %s", err)
		utils.RespondError(w, 500, "error getting user")
		return
	}

	// validate password
	if pass != dbUser.Password {
		log.Errorf("error validating user password: %s", err)
		utils.RespondError(w, 403, "error validating user")
		return
	}

	// respond with jwt
	token, err := utils.NewToken(dbUser.Email, dbUser.Id, s.config.TokenKey, s.config.TokenIssuer, s.config.TokenTimeout)
	if err != nil || token == "" {
		log.Errorf("error creating token: %s", err)
		utils.RespondError(w, 500, "error creating token")
		return
	}

	utils.RespondJson(w, 200, token)
}
