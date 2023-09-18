package server

import (
	"go-server/internal/utils"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
)

// RequireAuth - middleware to require authentication on various endpoints
func (s *server) RequireAuth(handler httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		utils.SetCorsHeaders(w, r)

		// read token from header
		authHeader, ok := r.Header["Authorization"]
		if !ok {
			log.Error("unauthorized: missing authorization header")
			utils.RespondError(w, 400, "unauthorized: missing authorization header")
			return
		}

		// typically an auth header looks like "Bearer <raw jwt token>"
		token := strings.Split(authHeader[0], "Bearer ")[1]

		// validate the token
		if err := utils.ValidateToken(token, s.config.TokenKey, s.config.TokenIssuer); err != nil {
			log.Error(err)
			utils.RespondError(w, 403, err.Error())
			return
		}

		// user is validated, so call the original handler
		handler(w, r, params)
	}
}
