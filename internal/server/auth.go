package server

import (
	"go-server/internal/utils"
	"net/http"

	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
)

// WithAuth - middleware to require authentication on various endpoints
func (s *server) WithAuth(handler httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		utils.SetCorsHeaders(w, r)

		// read token from header
		token, err := utils.GetTokenFromRequest(r)
		if err != nil {
			log.Error(err)
			utils.RespondError(w, 400, err.Error())
			return
		}

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
