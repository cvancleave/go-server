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

// handleGetTimesheets - gets all timesheets for a user
func (s *server) handleGetTimesheets(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	utils.SetCorsHeaders(w, r)

	// get user id url query
	userIdString := r.URL.Query().Get("userId")
	userId, err := strconv.Atoi(userIdString)
	if err != nil {
		log.Errorf("error getting user id from query: %s", err.Error())
		utils.RespondError(w, 400, "error getting user id")
		return
	}

	log.Debugf("getting timesheets for user id: %d", userId)

	// get timesheets from database
	timesheets, err := s.database.GetTimesheetsByUserId(userId)
	if err != nil {
		log.Errorf("error getting timesheets from database: %s", err.Error())
		utils.RespondError(w, 500, "error getting timesheets")
		return
	}

	utils.RespondJson(w, 200, timesheets)
}

// handlePostTimesheet - creates a timesheet for a user
func (s *server) handlePostTimesheet(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	utils.SetCorsHeaders(w, r)

	timesheet := models.Timesheet{}
	if err := json.NewDecoder(r.Body).Decode(&timesheet); err != nil {
		log.Errorf("error reading timesheet from request body: %s", err.Error())
		utils.RespondError(w, 400, "error reading request body")
		return
	}

	log.Debugf("creating timesheet for user id: %d", timesheet.UserID)

	// insert timesheet into database
	if err := s.database.InsertTimesheet(timesheet); err != nil {
		log.Errorf("error inserting timesheet: %s", err.Error())
		utils.RespondError(w, 500, "error inserting timesheet")
		return
	}

	utils.RespondJson(w, 200, true)
}
