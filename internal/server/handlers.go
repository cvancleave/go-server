package server

import (
	"encoding/json"
	"go-server/internal/models"
	"go-server/internal/utils"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
)

// getUserHandler - gets a user via user ID passed in through the URL query
func (s *server) getUserHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	utils.SetCorsHeaders(w, r)

	userID := r.URL.Query().Get("id")
	if userID == "" {
		log.Error("getUserHandler() - error getting id from query")
		utils.RespondError(w, 400, "error getting id from query")
		return
	}

	userIDint, err := strconv.Atoi(userID)
	if err != nil {
		log.Errorf("getUserHandler() - error parsing id: %s", err.Error())
		utils.RespondError(w, 400, "error parsing id")
		return
	}

	log.Debugf("getUserHandler() - getting user by id: %d", userIDint)

	// TODO - get user from database here

	// return dummy user
	user := models.User{
		ID:   1,
		Name: "gotten user",
	}

	utils.RespondJson(w, 200, user)
}

// postUserHandler - creates a user from info in the request body
func (s *server) postUserHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	utils.SetCorsHeaders(w, r)

	user := models.User{}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Errorf("postUserHandler() - error reading request body: %s", err.Error())
		utils.RespondError(w, 400, "error reading request body")
		return
	}

	log.Debugf("postUserHandler() - creating user with name: %s", user.Name)

	// TODO - insert user into database here

	// return dummy user
	user = models.User{
		ID:   1,
		Name: "created user",
	}

	utils.RespondJson(w, 200, user)
}
